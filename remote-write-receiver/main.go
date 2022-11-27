package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var requestCount int
var timeStart time.Time

func main() {
	timeStart = time.Now()
	if err := serve(":9179"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func serve(addr string) error {
	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		t2 := time.Now()
		timeDiff := t2.Sub(timeStart)
		if timeDiff.Seconds() > 60 {
			return
		}
		requestCount = requestCount + 1
		outputDir := os.Getenv("OUTPUT_DIR")

		os.MkdirAll(outputDir, 0755)

		t := time.Now()
		timeStr := t.Format("20060102150405")
		fileName := "remote-write-output-" + timeStr + "-" + strconv.Itoa(requestCount)

		out, err := os.Create(filepath.Join(outputDir, fileName))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("count=%d time_spent=%s file=%s\n", requestCount, timeDiff, fileName)
	})
	return http.ListenAndServe(addr, nil)
}
