package backup

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
)

// Backups a group using robocopy.
// In case of an error it fails immediately.
func Robocopy(g Group) error {
	var err error
	for _, e := range g.Elements {
		err = writeElement(e, g.Folder)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeElement(e Element, path string) error {
	if !e.IsValid() {
		return errors.New("Invalid file/folder")
	}

	ext := filepath.Ext(e.Path)
	// Source folder
	var folder = e.Path
	// Source file (if existing)
	var file = ""
	if ext != "" {
		folder = filepath.Dir(e.Path)
		file = filepath.Base(e.Path)
	}

	cmd := exec.Command(
		"robocopy",
		"/tee",
		"/MIR",
		"/COPY:DAT",
		folder,
		fmt.Sprintf("%s/%s/", path, e.Name),
		file)

	fmt.Println(folder)
	fmt.Println(fmt.Sprintf("%s/%s/", path, e.Name))
	fmt.Println(file)

	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()

	// TODO: Detect errors
	return nil
}
