package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/SrcHndWng/go-serverless-dynamo-report/report/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	yaml "gopkg.in/yaml.v2"
)

func TestCreateReportData(t *testing.T) {
	t1 := model.Report{ID: 1, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641307}
	t2 := model.Report{ID: 2, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641326}
	t3 := model.Report{ID: 3, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641337}
	reports := []model.Report{t1, t2, t3}
	datas := createReportData(reports)
	if len(datas) != 3 {
		t.Fatalf("error raise. datas len = %#v", len(datas))
	}
	fmt.Printf("create report data = %v\n", datas)
}

func TestCreatePdf(t *testing.T) {
	const ttf = "../assets/times.ttf"
	const dest = "./sample.pdf"
	t1 := model.Report{ID: 1, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641307}
	t2 := model.Report{ID: 2, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641326}
	t3 := model.Report{ID: 3, ColumnA: "aaa", ColumnB: "bbb", CreatedAt: 1553641337}
	reports := []model.Report{t1, t2, t3}
	err := CreatePdf(reports, ttf, dest)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	_, err = os.Stat(dest)
	if err != nil {
		t.Fatal("dest file is not created")
	}
	// dest file remove. if you want to check dest file, comment out here.
	err = os.Remove(dest)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestPutToS3(t *testing.T) {
	buf, err := ioutil.ReadFile("../../serverless.yml")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}

	bucket := fmt.Sprintf("%v", m["provider"].(map[interface{}]interface{})["environment"].(map[interface{}]interface{})["bucket"])
	region := fmt.Sprintf("%v", m["provider"].(map[interface{}]interface{})["region"])
	const key = "times.ttf"
	const src = "../assets/times.ttf"

	err = PutToS3(bucket, key, region, src)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}

	sess := session.New(&aws.Config{Region: aws.String(region)})
	svc := s3.New(sess)
	listObjInput := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	resp, err := svc.ListObjects(listObjInput)
	for _, item := range resp.Contents {
		fmt.Printf("put file = %s\n", *item.Key)
	}
	deleteObjInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	headObjInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	// file remove. if you want to check dest file, comment out here.
	_, err = svc.DeleteObject(deleteObjInput)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	err = svc.WaitUntilObjectNotExists(headObjInput)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}
