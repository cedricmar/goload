package looper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Loop does not return, it just blocks until a change happens
func Loop() {

	var mod bool
	start := time.Now()

	// Get root folder
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(">>> Observing %s\n", dir)

	for !mod {
		select {
		case <-time.After(500 * time.Millisecond):
			mod, err = hasChanged(dir, start)

			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

func hasChanged(path string, start time.Time) (bool, error) {

	var chg bool

	err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {

		if file.IsDir() && file.Name() == "tmp" {
			return filepath.SkipDir
		}

		if start.Sub(file.ModTime()) < 0 && !strings.HasPrefix(file.Name(), ".") {
			chg = true
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return chg, nil
}
