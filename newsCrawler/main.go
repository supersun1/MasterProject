package main

import (
	"encoding/json"
	"fmt"
	"github.com/diffbot/diffbot-go-client"
	"github.com/sirupsen/logrus"
	"os"
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

type News struct {
	Title       string
	NewsContent string
	Date        string
	Author      string
	URL         string
}

//
//type RestParameters struct {
//	Action     string `json:"action,omitempty"`
//	NewsCSV    string `json:"secret-name,omitempty"`
//	StoreCSV   string `json:"secret-value,omitempty"`
//	TokenValue string `json:"-"`
//}
//
//func parseArguments() *RestParameters {
//	if len(os.Args) < 2 {
//		exitOnError("Please enter arguments... (-h for usage and examples)")
//	}
//	parameters := &RestParameters{}
//	parameters.Action = os.Args[1]
//	options := flag.NewFlagSet("action", flag.ExitOnError)
//
//	newscsv := options.String("newscsv", "", "csv to news articles")
//	storecsv := options.String("storecsv", "", "Path to where you want to store your csv")
//	tokenValue := options.String("token", "", "Diff bot token")
//
//	if *newscsv != "" {
//		parameters.NewsCSV = *newscsv
//	}
//	if *storecsv != "" {
//		parameters.StoreCSV = *storecsv
//	}
//	if *tokenValue != "" {
//		parameters.TokenValue = *tokenValue
//	}
//	return parameters
//}

func parseArticle(token, newsURL string) (n News) {
	article, err := diffbot.ParseArticle(token, newsURL, nil)
	if err != nil {
		fmt.Println("failed to parse article. Error: " + err.Error())
	} else {
		n.Title = article.Title
		n.URL = article.Url
		n.Author = article.Author
		n.Date = article.Date
		n.NewsContent = article.Text
	}

	articlez, _ := ParseIt(token, newsURL, nil)
	fmt.Println(article)
	fmt.Println(articlez)

	return
}

//func crawl(params *RestParameters) {
//
//}

func main() {
	//params := parseArguments()
	//crawl(params)

	parseArticle("5c729dda64ab1d0e6fa91e2ee59fdc7c", "https://www.naturalnews.com/2020-11-02-journalist-writers-side-job-destroying-trump-ballots.html")
	//parseArticle("5c729dda64ab1d0e6fa91e2ee59fdc7c", "https://www.cnn.com/2020/11/11/politics/state-department-biden-messages/index.html")
}
