package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)


type RestParameters struct {
	NewsCSV    string
	StoreCSV   string
	LabelValue string
	IsKostya bool
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
	Title     string `json:"title"`
	Text       string `json:"pageUrl"`
	Label     string `json:"label"`
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
	flag.Parse()

	if *newscsv != "" {
		parameters.NewsCSV = *newscsv
	}
	if *storecsv != "" {
		parameters.StoreCSV = *storecsv
	}
	parameters.LabelValue = strconv.Itoa(*label)
	parameters.IsKostya = *kos
	return parameters
}

func openCSV(newsLocation string, label string, isKostya bool) []ScanArtInfo {
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
		if isKostya {
			newsArr = append(newsArr, ScanArtInfo{
				Title:   record[0],
				Text: record[1],
				Label: label,
			})
		} else {
			newsArr = append(newsArr, ScanArtInfo{
				Title: record[0],
				Text: record[1],
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
	newsArr := openCSV(params.NewsCSV, params.LabelValue, params.IsKostya)
	writeCSV(params.StoreCSV, newsArr)
}

func main() {
	log.Info("running news parser")
	param := parseArguments()
	parse(param)

}

