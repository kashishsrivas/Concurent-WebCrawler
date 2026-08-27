package storage

import (
	"encoding/json"
	"os"
)

func SaveJSON(
	filename string,
	results []CrawlResult,
) error {

	data, err := json.MarshalIndent(
		results,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		filename,
		data,
		0644,
	)
}
