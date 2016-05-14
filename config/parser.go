package config

import "encoding/json"
import "github.com/tropid/ezbackup/backup"
import "io"
import "io/ioutil"
import "os"

type Configuration struct {
	Groups []backup.Group
}

func Load(path string) (*Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return loadReader(file)
}

func loadReader(r io.Reader) (*Configuration, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(bytes, &config.Groups)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
