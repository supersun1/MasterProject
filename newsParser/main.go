package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type RestParameters struct {
	NewsCSV    string
	StoreCSV   string
	LabelValue string
	IsKostya   bool
	IsSplit    bool
	IsClean    bool
	IsSep      bool
	SepCount   int
}

type ArticleInfo struct {
	Date      string `json:"estimatedDate"`
	NewsSorce string `json:"siteName"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"text"`
	URL       string `json:"pageUrl"`
	Label     string `json:"label"`
}

type ScanArtInfo struct {
	Title string `json:"title"`
	Text  string `json:"pageUrl"`
	Label string `json:"label"`
}

var exitOnError = onError

func onError(a ...interface{}) {
	log.Error(a...)
	os.Exit(1)
}

func parseArguments() *RestParameters {
	if len(os.Args) < 2 {
		exitOnError("Please enter arguments... (-h for usage and examples)")
	}
	parameters := &RestParameters{}

	newscsv := flag.String("newscsv", "", "csv to news articles")
	storecsv := flag.String("storecsv", "", "Path to where you want to store your csv")
	label := flag.Int("label", 0, "Path to where you want to store your csv")
	kos := flag.Bool("kos", false, "Path to where you want to store your csv")
	split := flag.Bool("split", false, "Split from a big file into two true and false files")
	clean := flag.Bool("clean", true, "Split from a big file into two true and false files")
	seperateFiles := flag.Bool("sep", false, "seperate")
	sepcount := flag.Int("sepcount", 500, "")

	flag.Parse()

	if *newscsv != "" {
		parameters.NewsCSV = *newscsv
	}
	if *storecsv != "" {
		parameters.StoreCSV = *storecsv
	}
	parameters.LabelValue = strconv.Itoa(*label)
	parameters.IsKostya = *kos
	parameters.IsSplit = *split
	parameters.IsClean = *clean
	parameters.IsSep = *seperateFiles
	parameters.SepCount = *sepcount
	return parameters
}

func openCSV(newsLocation string, label string, isKostya bool, isSplit bool, isClean bool, isSep bool) []ScanArtInfo {
	fmt.Println("================== opening csv ==================")
	var newsArr []ScanArtInfo
	csvfile, err := os.Open(newsLocation)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	isScanned := true

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if isScanned {
			isScanned = false
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if isSep {
			text := record[1]
			sepText := strings.Split(text, " ")
			newArtCount := int(math.Ceil(float64(len(sepText) / 500)))

			for x := 0; x < newArtCount+1; x++ {
				text := ""
				start := x * 500
				end := x*500 + 500
				if x == newArtCount {
					end = len(sepText)
				}

				for i := start; i < end; i++ {
					text += sepText[i] + " "
				}

				newsArr = append(newsArr, ScanArtInfo{
					Title: record[0],
					Text:  text,
					Label: record[2],
				})
			}

		} else if isKostya {
			newsArr = append(newsArr, ScanArtInfo{
				Title: record[0],
				Text:  record[1],
				Label: label,
			})
		} else if isSplit || isClean {
			newsArr = append(newsArr, ScanArtInfo{
				Title: record[0],
				Text:  record[1],
				Label: record[2],
			})
		} else {
			newsArr = append(newsArr, ScanArtInfo{
				Title: record[0],
				Text:  record[1],
				Label: record[6],
			})
		}

	}
	return newsArr
}

func writeCSVDescriptors(csvLoc string) {
	file, err := os.OpenFile(csvLoc, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	w := csv.NewWriter(file)
	t := []string{"Title", "Text", "Label"}
	e := w.Write(t)
	if e != nil {
		fmt.Println(e.Error())
	}
	w.Flush()
}

func writeSplitCSV(csvFolder string, infos []ScanArtInfo) {
	var factual []ScanArtInfo
	var fake []ScanArtInfo
	for _, info := range infos {
		switch info.Label {
		case "0":
			fake = append(fake, info)
		case "1":
			factual = append(factual, info)
		}

	}

	writeCSV(csvFolder+"/fakenews.csv", fake)
	writeCSV(csvFolder+"/realnews.csv", factual)
}

func writeCSV(csvLoc string, infos []ScanArtInfo) {
	fmt.Println("================== write to csv ==================")
	var file *os.File

	if _, err := os.Stat(csvLoc); os.IsNotExist(err) {
		writeCSVDescriptors(csvLoc)
	}

	file, err := os.OpenFile(csvLoc, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	w := csv.NewWriter(file)
	for _, value := range infos {
		var t []string
		if value.Title == "" || value.Text == "" || value.Label == "" {
			continue
		}

		t = append(t, value.Title)
		t = append(t, value.Text)
		t = append(t, value.Label)
		err = w.Write(t)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Flush()
	}
}

func parse(params *RestParameters) {
	newsArr := openCSV(params.NewsCSV, params.LabelValue, params.IsKostya, params.IsSplit, params.IsClean, params.IsSep)
	writeCSV(params.StoreCSV, newsArr)
}

func parseSplit(params *RestParameters) {
	newsArr := openCSV(params.NewsCSV, params.LabelValue, params.IsKostya, params.IsSplit, params.IsClean, params.IsSep)
	writeSplitCSV(params.StoreCSV, newsArr)
}

func main() {
	log.Info("running news parser")
	param := parseArguments()

	if param.IsSplit {
		parseSplit(param)
	} else {
		parse(param)
	}

}
