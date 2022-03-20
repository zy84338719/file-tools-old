package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSV struct {
}

func CsvNew() *CSV {
	return &CSV{}
}
func (c *CSV) Read(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()
	w := csv.NewReader(f)
	data, err := w.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}
