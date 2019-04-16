package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/SrcHndWng/go-serverless-dynamo-report/report/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/signintech/gopdf"
)

type templateData struct {
	Datas []data
}

type data struct {
	ID        int
	ColumnA   string
	ColumnB   string
	CreatedAt string
}

// CreatePdf creates pdf file in local dest path
func CreatePdf(reports model.Reports, ttf, dest string) (err error) {
	datas := createReportData(reports)
	fmt.Println(datas)

	pdf := gopdf.GoPdf{}
	defer pdf.Close()

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err = pdf.AddTTFFont("times", ttf)
	if err != nil {
		return
	}
	err = pdf.SetFont("times", "", 14)
	if err != nil {
		return
	}

	pdf.SetLineWidth(0.5)

	y1 := 10.0
	y2 := 10.0
	for i, data := range datas {
		pdf.Line(10, y1, 585, y2)
		if i == 0 {
			pdf.Br(5)
		} else {
			pdf.Br(20)
		}
		pdf.Cell(nil, fmt.Sprintf("%d, %s, %s, %s", data.ID, data.ColumnA, data.ColumnB, data.CreatedAt))
		y1 += 20
		y2 += 20
	}

	// last line
	pdf.Line(10, y1, 585, y2)

	// horizontal lines
	pdf.Line(10, 10, 10, y2)
	pdf.Line(585, 10, 585, y2)

	pdf.WritePdf(dest)

	return nil
}

// PutToS3 puts src file to s3
func PutToS3(bucket, key, region, src string) (err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	sess := session.New(&aws.Config{Region: aws.String(region)})

	in := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf),
	}

	_, err = s3.New(sess).PutObject(in)
	if err != nil {
		return
	}

	return
}

func createReportData(reports model.Reports) (datas []data) {
	datas = make([]data, 0)
	for _, report := range reports {
		t := time.Unix(report.CreatedAt, 0)
		datas = append(datas, data{ID: report.ID, ColumnA: report.ColumnA, ColumnB: report.ColumnB, CreatedAt: t.Format("2006-1-2 15:04:05")})
	}
	return
}
