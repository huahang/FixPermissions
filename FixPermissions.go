package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: FixPermissions [path to scan]\n")
		return
	}
	_ = filepath.Walk(args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("[Error] Hit an error! " + err.Error() + "\n")
			return err
		}
		if info.Mode().IsDir() && info.Mode().Perm() != 0755 {
			fmt.Printf("[Fixing] 0755 %s\n", path)
			err = os.Chmod(path, 0755)
		}
		if info.Mode().IsRegular() && info.Mode().Perm() != 0644 {
			fmt.Printf("[Fixing] 0644 %s\n", path)
			err = os.Chmod(path, 0644)
			return nil
		}
		return nil
	})
}
