package main // import "github.com/weebagency/goload/cmd/goload"

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Get root folder
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(">>> Observing %s\n", dir)

	// Get the time
	start := time.Now()
	log.Printf(">>> Start time %s", start)
	var chg bool
	for !chg {
		select {
		case <-time.After(600 * time.Millisecond):

			chg, err = hasChanged(dir, start)

			if err != nil {
				log.Fatal(err)
			}

		}
	}

	log.Println("reload...")
}

func hasChanged(path string, start time.Time) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			if start.Sub(file.ModTime()) < 0 {
				log.Printf("Folder %s last modified %s\n", file.Name(), file.ModTime())
				return true, nil
			}
			return hasChanged(path+"/"+file.Name(), start)
		}
		if start.Sub(file.ModTime()) < 0 {
			log.Printf("File %s last modified %s\n", file.Name(), file.ModTime())
			return true, nil
		}
	}

	return false, nil
}
