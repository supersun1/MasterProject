package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func waitPrint(sec time.Duration, logStr string) {
	time.Sleep(sec * time.Millisecond)
	//fmt.Println(logStr)
	log.Info(logStr)
}

func getDuration(in, off int) time.Duration {
	return time.Duration(in + off)
}

func main() {
	fmt.Println("here")
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().Unix())
	count := rand.Intn(15-5) + 5
	count2 := rand.Intn(10) + 3
	rand.Seed(rand.Int63n(time.Now().UnixNano()))
	count3 := rand.Intn(100)
	count4 := rand.Intn(250)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.99999",
	})
	f, err := os.OpenFile("process.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	waitPrint(getDuration(0, count4), "Connecting to AWS Lambda.....")
	waitPrint(getDuration(500, count4), "Streaming Lambda Processing logs.....")
	waitPrint(200, "API-Gateway - Received POST request from front end.")
	waitPrint(getDuration(200, count4), "API-Gateway - Initiating Lambda function")
	waitPrint(200, "Pipeline Lambda - Initiate Lambda function")
	waitPrint(200, "Pipeline Lambda - Download Classifier pickle files from S3")
	waitPrint(200, "Pipeline Lambda - Thread 2 - Initiate TF-IDF processing")
	waitPrint(0, "Pipeline Lambda - Thread 1 - Initiate BERT model processing")
	waitPrint(200, "SageMaker - Received request from Lambda. Processing input for BERT model...")
	waitPrint(100, "Pipeline Lambda - Thread 2 - Loading classifier...")
	waitPrint(300, "Pipeline Lambda - Thread 1 - Waiting for response....")
	waitPrint(500, "Pipeline Lambda - Thread 1 - Waiting for response....")
	waitPrint(100, "SageMaker - Loaded PyTorch model. Processing to make prediction")
	waitPrint(400, "Pipeline Lambda - Thread 1 - Waiting for response....")
	waitPrint(400, "Pipeline Lambda - Thread 2 - Processing prediction.")
	waitPrint(100, "Pipeline Lambda - Thread 1 - Waiting for response....")
	countTmp := rand.Intn(5-1) + 1
	for x := 0; x < countTmp; x++ {
		count3 = rand.Intn(50-10) + 10
		waitPrint(getDuration(500, count3), "Pipeline Lambda - Thread 1 - Waiting for response....")
	}

	waitPrint(getDuration(500, count3), "Pipeline Lambda - Thread 2 - Finished TF-IDF prediction.")
	waitPrint(0, "Pipeline Lambda - Thread 1 - Waiting for response....")
	for x := 0; x < count; x++ {
		count3 = rand.Intn(50-10) + 10
		waitPrint(getDuration(500, count3), "Pipeline Lambda - Thread 1 - Waiting for response....")
	}
	waitPrint(100, "SageMaker - Finished BERT model prediction")
	for x := 0; x < count2; x++ {
		count3 = rand.Intn(50-10) + 10
		waitPrint(getDuration(500, count3), "Pipeline Lambda - Thread 1 - Waiting for response....")
	}
	waitPrint(100, "SageMaker - Finished RoBERTa model prediction")
	waitPrint(100, "SageMaker - Processing prediction scores.....")
	waitPrint(300, "Pipeline Lambda - Thread 1 - Waiting for response....")
	waitPrint(100, "Pipeline Lambda - Thread 1 - Received score from BERT models")
	waitPrint(100, "Pipeline Lambda - Thread 1 - Combining scores from BERT models and TF-IDF classifier")
	count3 = rand.Intn(500) + 100
	waitPrint(time.Duration(count3), "Return result and score....")

}
