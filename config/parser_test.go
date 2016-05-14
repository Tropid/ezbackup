package config

import (
	"bytes"
	"testing"
)

var (
	testConfig = []byte(`[
		{
			"folder": "F:/Temp/b",
			"elements": [
				{
					"name": "A",
					"path": "F:/Temp/a.txt",
					"excludes": []
				}
			]
		}
	]`)
)

func TestLoad(t *testing.T) {
	config, err := loadReader(bytes.NewReader(testConfig))
	if err != nil {
		t.Fatal(err)
	}

	if len(config.Groups) != 1 {
		t.Error("Should be 1 group: ", len(config.Groups))
	}

	if config.Groups[0].Folder != "F:/Temp/b" {
		t.Error("Incorrect output folder: ", config.Groups[0].Folder)
	}

	if len(config.Groups[0].Elements) != 1 {
		t.Error("Should be one element.")
	}

	if config.Groups[0].Elements[0].Name != "A" {
		t.Error("Invalid element name: ", config.Groups[0].Elements[0].Name)
	}
}
