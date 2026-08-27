package storage

import (
	"encoding/csv"
	"os"
	"strconv"
)

func SaveCSV(
	filename string,
	results []CrawlResult,
) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		"url",
		"depth",
	})

	if err != nil {
		return err
	}

	for _, result := range results {

		err := writer.Write([]string{
			result.URL,
			strconv.Itoa(result.Depth),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
