package model

import (
	"fmt"
	"testing"
)

func TestGetReports(t *testing.T) {
	dynamoLocalEndpoint = "http://localhost:8001/shell"
	reports, err := GetReports()
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result count = %d, reports = %v\n", len(reports), reports)
}
