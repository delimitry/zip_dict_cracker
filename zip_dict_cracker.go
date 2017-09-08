package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/yeka/zip"
)

func unzip(filename string, password string) bool {
	// open a zip archive for reading
	r, err := zip.OpenReader(filename)
	if nil != err {
		panic(err)
	}
	defer r.Close()
	// create a buffer to write archive to
	buffer := new(bytes.Buffer)
	// iterate through the files in the archive
	for _, f := range r.File {
		f.SetPassword(password)
		// open a file for reading
		r, err := f.Open()
		if err != nil {
			return false
		}
		defer r.Close()
		// try to copy file data to buffer
		n, err := io.Copy(buffer, r)
		if n == 0 || err != nil {
			return false
		}
		break
	}
	return true
}

func crack(zipFile string, dictFile string) {
	file, err := os.Open(dictFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	step := 100000
	startTime := time.Now()
	for scanner.Scan() {
		if count != 0 && count%step == 0 {
			fmt.Printf("Checked %d passwords. Rate: %.2f passwords per second\n",
				count, float64(step)/time.Since(startTime).Seconds())
			startTime = time.Now()
		}
		password := scanner.Text()
		res := unzip(zipFile, password)
		if res == true {
			fmt.Printf("Password found: %s\n", password)
			return
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// prepare command-line parsing
	zipFilePtr := flag.String("f", "", "Path to zip file")
	dictFilePtr := flag.String("d", "", "Path to dictionay file")
	flag.Parse()
	if *zipFilePtr == "" || *dictFilePtr == "" {
		flag.Usage()
		os.Exit(1)
	}
	zipFile := *zipFilePtr
	dictFile := *dictFilePtr

	// check files existence
	if _, err := os.Stat(zipFile); os.IsNotExist(err) {
		fmt.Printf("File \"%s\" not exists!", zipFile)
		os.Exit(1)
	}
	if _, err := os.Stat(dictFile); os.IsNotExist(err) {
		fmt.Printf("File \"%s\" not exists!", dictFile)
		os.Exit(1)
	}

	crack(zipFile, dictFile)
}
