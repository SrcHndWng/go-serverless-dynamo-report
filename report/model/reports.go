package model

import "sort"

// Report tableata struct
type Report struct {
	ID        int    `dynamo:"id"`
	ColumnA   string `dynamo:"column_a"`
	ColumnB   string `dynamo:"column_b"`
	CreatedAt int64  `dynamo:"created_at"`
}

// Reports contains report data
type Reports []Report

// GetReports selects all data
func GetReports() (data Reports, err error) {
	table := Table("reports")
	err = table.Scan().All(&data)
	if err != nil {
		return
	}
	sort.Sort(data)
	return
}

func (r Reports) Len() int {
	return len(r)
}

func (r Reports) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Reports) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}
