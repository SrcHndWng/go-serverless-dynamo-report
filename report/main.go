package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/SrcHndWng/go-serverless-dynamo-report/report/model"
	"github.com/SrcHndWng/go-serverless-dynamo-report/report/service"
	"github.com/aws/aws-lambda-go/lambda"
)

const timeFormat = "2006-01-02 15:04:05"
const messageFormat = "%s, create report success."

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (string, error) {
	const ttf = "/tmp/times.ttf"

	err := createTtfFromAssets(ttf)
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("bucket")
	reportFileName := os.Getenv("reportFileName")
	pdfPath := os.Getenv("pdfPath")
	region := os.Getenv("region")

	reports, err := model.GetReports()
	if err != nil {
		return "", err
	}

	err = service.CreatePdf(reports, ttf, pdfPath)
	if err != nil {
		return "", err
	}

	err = service.PutToS3(bucket, reportFileName, region, pdfPath)
	if err != nil {
		return "", err
	}

	t := time.Now()
	message := createMessage(t)
	return message, nil
}

func createTtfFromAssets(dest string) (err error) {
	assetFile, err := Assets.Open("/assets/times.ttf")
	if err != nil {
		return
	}
	defer assetFile.Close()

	ttfFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer ttfFile.Close()

	_, err = io.Copy(ttfFile, assetFile)
	if err != nil {
		return
	}

	return
}

func createMessage(t time.Time) string {
	return fmt.Sprintf(messageFormat, t.Format(timeFormat))
}

func main() {
	lambda.Start(Handler)
}
