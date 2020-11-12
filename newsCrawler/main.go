package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	urlPkg "net/url"
	"os"
	"strconv"
	"time"
)

const (
	//DefaultServer = `https://labs.diffbot.com/testdrive/article`
	DefaultServer = `https://api.diffbot.com/v3/article`
)

var exitOnError = onError

func onError(a ...interface{}) {
	logrus.Error(a...)
	os.Exit(1)
}

func exitIfError(err error, a ...interface{}) {
	if err != nil {
		exitOnError(a, err.Error())
	}
}

type Request struct {
	URL string `json:"pageUrl"`
}

type ArticleInfo struct {
	Date      string `json:"estimatedDate"`
	NewsSorce string `json:"siteName"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"text"`
	URL       string `json:"pageUrl"`
}

type TestArticle struct {
	Request     Request       `json:"request"`
	ArticleInfo []ArticleInfo `json:"objects"`
}

type Options struct {
	Fields                 string
	Timeout                time.Duration
	Callback               string
	FrontpageAll           string
	ClassifierMode         string
	ClassifierStats        string
	BulkNotifyEmail        string
	BulkNotifyWebHook      string
	BulkRepeat             string
	BulkMaxRounds          string
	BulkPageProcessPattern string
	CrawlMaxToCrawl        string
	CrawlMaxToProcess      string
	CrawlRestrictDomain    string
	CrawlNotifyEmail       string
	CrawlNotifyWebHook     string
	CrawlDelay             string
	CrawlRepeat            string
	CrawlOnlyProcessIfNew  string
	CrawlMaxRounds         string
	BatchMethod            string
	BatchRelativeUrl       string
	CustomHeader           http.Header
}

//
type RestParameters struct {
	NewsCSV    string
	StoreCSV   string
	TokenValue string
}

func parseArguments() *RestParameters {
	if len(os.Args) < 2 {
		exitOnError("Please enter arguments... (-h for usage and examples)")
	}
	parameters := &RestParameters{}

	newscsv := flag.String("newscsv", "", "csv to news articles")
	storecsv := flag.String("storecsv", "", "Path to where you want to store your csv")
	tokenValue := flag.String("token", "", "Diff bot token")
	flag.Parse()

	if *newscsv != "" {
		parameters.NewsCSV = *newscsv
	}
	if *storecsv != "" {
		parameters.StoreCSV = *storecsv
	}
	if *tokenValue != "" {
		parameters.TokenValue = *tokenValue
	}
	return parameters
}

func parseArticle(token, newsURL string) (n ArticleInfo, e error) {
	article, e := ParseIt(token, newsURL, nil)
	if e != nil {
		return
	}

	return article.ArticleInfo[0], e
}

func ParseIt(token, url string, opt *Options) (*TestArticle, error) {
	body, err := Diffbot("article", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result TestArticle
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ArticleInfo[0].Author == "" {
		result.ArticleInfo[0].Author = "anonymous"
	}
	if result.ArticleInfo[0].Date == "" {
		result.ArticleInfo[0].Date = "none"
	}
	if result.ArticleInfo[0].URL == "" {
		result.ArticleInfo[0].URL = result.Request.URL
	}
	return &result, nil
}

type Error struct {
	ErrCode    int    `json:"errorCode"` // Description of the error
	ErrMessage string `json:"error"`     // Error code per the chart below
	RawString  string `json:"-"`         // Raw json format error string
}

// ParseJson parses the JSON-encoded error data.
func (p *Error) ParseJson(s string) error {
	if err := json.Unmarshal([]byte(s), p); err != nil {
		return err
	}
	p.RawString = s
	return nil
}

func (p *Error) Error() string {
	d, _ := json.Marshal(p)
	return string(d)
}

// Diffbot uses computer vision, natural language processing
// and machine learning to automatically recognize
// and structure specific page-types.
func Diffbot(method, token, url string, opt *Options) (body []byte, err error) {
	return DiffbotServer(DefaultServer, method, token, url, opt)
}

// DiffbotServer like Diffbot function, but support custom server.
func DiffbotServer(server, method, token, url string, opt *Options) (body []byte, err error) {
	requestUrl := makeRequestUrl(server, method, token, url, opt)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	if opt != nil && opt.CustomHeader != nil {
		req.Header = opt.CustomHeader
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if len(body) != 0 {
			var apiError Error
			if err = apiError.ParseJson(string(body)); err != nil {
				err = &Error{
					ErrCode:    resp.StatusCode,
					ErrMessage: string(body),
				}
				return
			} else {
				err = &apiError
				return
			}
		} else {
			err = &Error{
				ErrCode:    resp.StatusCode,
				ErrMessage: resp.Status,
			}
			return
		}
	}
	return
}

func makeRequestUrl(server, method, token, webUrl string, opt *Options) string {
	return fmt.Sprintf("%s?token=%s&url=%s%s",
		server, token, urlPkg.QueryEscape(webUrl), opt.MethodParamString(method),
	)
}

// MethodParamString return string as the url params.
//
// If the Options is not empty, the return string begin with a '&'.
func (p *Options) MethodParamString(method string) string {
	if p == nil || method == "" {
		return ""
	}

	switch method {
	case "article", "image", "product":
		var s []byte
		if p.Fields != "" {
			s = append(s, ("&fields=" + p.Fields)...)
		}
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.Callback != "" {
			s = append(s, ("&callback=" + urlPkg.QueryEscape(p.Callback))...)
		}
		return string(s)

	case "frontpage":
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.FrontpageAll != "" {
			s = append(s, ("&all=" + p.FrontpageAll)...)
		}
		return string(s)

	case "analyze":
		var s []byte
		if p.ClassifierMode != "" {
			s = append(s, ("&mode=" + p.ClassifierMode)...)
		}
		if p.Fields != "" {
			s = append(s, ("&fields=" + p.Fields)...)
		}
		if p.ClassifierStats != "" {
			s = append(s, ("&stats=" + p.ClassifierStats)...)
		}
		return string(s)

	case "bulk":
		var s []byte
		if p.BulkNotifyEmail != "" {
			s = append(s, ("&notifyEmail=" + p.BulkNotifyEmail)...)
		}
		if p.BulkNotifyWebHook != "" {
			s = append(s, ("&notifyWebHook=" + p.BulkNotifyWebHook)...)
		}
		if p.BulkRepeat != "" {
			s = append(s, ("&repeat=" + p.BulkRepeat)...)
		}
		if p.BulkMaxRounds != "" {
			s = append(s, ("&maxRounds=" + p.BulkMaxRounds)...)
		}
		if p.BulkPageProcessPattern != "" {
			s = append(s, ("&pageProcessPattern=" + p.BulkPageProcessPattern)...)
		}
		return string(s)

	case "crawl":
		var s []byte
		if p.CrawlMaxToCrawl != "" {
			s = append(s, ("&maxToCrawl=" + p.CrawlMaxToCrawl)...)
		}
		if p.CrawlMaxToProcess != "" {
			s = append(s, ("&maxToProcess=" + p.CrawlMaxToProcess)...)
		}
		if p.CrawlRestrictDomain != "" {
			s = append(s, ("&restrictDomain=" + p.CrawlRestrictDomain)...)
		}
		if p.CrawlNotifyEmail != "" {
			s = append(s, ("&notifyEmail=" + p.CrawlNotifyEmail)...)
		}
		if p.CrawlNotifyWebHook != "" {
			s = append(s, ("&notifyWebHook=" + p.CrawlNotifyWebHook)...)
		}
		if p.CrawlDelay != "" {
			s = append(s, ("&crawlDelay=" + p.CrawlDelay)...)
		}
		if p.CrawlRepeat != "" {
			s = append(s, ("&repeat=" + p.CrawlRepeat)...)
		}
		if p.CrawlOnlyProcessIfNew != "" {
			s = append(s, ("&onlyProcessIfNew=" + p.CrawlOnlyProcessIfNew)...)
		}
		if p.CrawlMaxRounds != "" {
			s = append(s, ("&maxRounds=" + p.CrawlMaxRounds)...)
		}
		return string(s)

	case "batch":
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.BatchMethod != "" {
			s = append(s, ("&method=" + p.BatchMethod)...)
		}
		if p.BatchRelativeUrl != "" {
			s = append(s, ("&relative_url=" + urlPkg.QueryEscape(p.BatchRelativeUrl))...)
		}
		return string(s)

	default: // Custom APIs
		var s []byte
		if p.Timeout != 0 {
			timeout := strconv.FormatInt(int64(p.Timeout/time.Millisecond), 10)
			s = append(s, ("&timeout=" + timeout)...)
		}
		if p.Callback != "" {
			s = append(s, ("&callback=" + urlPkg.QueryEscape(p.Callback))...)
		}
		return string(s)
	}

	return ""
}

func openCSV(newsLocation string) []string {
	fmt.Println("================== openning csv ==================")
	var newsArr []string
	csvfile, err := os.Open(newsLocation)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("found news site: " + record[0])
		newsArr = append(newsArr, record[0])
	}
	return newsArr
}

func scanForArticle(token string, newsArr []string) (artInfoArr []ArticleInfo) {
	fmt.Println("================== scanning csv ==================")

	for _, news := range newsArr {
		fmt.Println("scanning for article " + news)

		newsObj, e := parseArticle(token, news)
		if e != nil {
			log.Fatal("failed to parse article. Error: " + e.Error())
		} else {
			artInfoArr = append(artInfoArr, newsObj)
		}
	}
	return
}

func writeCSV(csvLoc string, infos []ArticleInfo) {
	fmt.Println("================== write to csv ==================")
	var file *os.File
	file, err := os.OpenFile(csvLoc, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	w := csv.NewWriter(file)
	for _, value := range infos {
		var t []string

		t = append(t, value.Title)
		t = append(t, value.Content)
		t = append(t, value.URL)
		t = append(t, value.Author)
		t = append(t, value.Date)
		t = append(t, value.NewsSorce)
		err = w.Write(t)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Flush()
	}
}

func crawl(params *RestParameters) {
	newsArr := openCSV(params.NewsCSV)
	artInfo := scanForArticle(params.TokenValue, newsArr)
	writeCSV(params.StoreCSV, artInfo)

}

func main() {
	params := parseArguments()
	crawl(params)
}
