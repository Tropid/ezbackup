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

	var cmd *exec.Cmd
	if ext != "" {
		// Handle file
		folder := filepath.Dir(e.Path)
		file := filepath.Base(e.Path)

		cmd = exec.Command(
			"robocopy",
			"/COPY:DAT",
			folder,
			fmt.Sprintf("%s/%s/", path, e.Name),
			file)
	} else {
		// Handle folder
		cmd = exec.Command(
			"robocopy",
			"/MIR",
			"/COPY:DAT",
			e.Path,
			fmt.Sprintf("%s/%s/", path, e.Name))
	}

	var outErr bytes.Buffer
	cmd.Stderr = &outErr

	var out bytes.Buffer
	cmd.Stdout = &out

	cmd.Run()

	// TODO: Detect errors
	return nil
}
