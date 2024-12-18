package main

import "os"

const tmpFileName = "tmp.txt"

func updateFile() error {
	content := RndName(15)
	err := os.WriteFile(tmpFileName, []byte(content), 0777)
	if err != nil {
		return err
	}

	return nil
}
