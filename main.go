package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var (
	includeChars string
	fileFilter   string
	verbose      bool
)

func main() {
	flag.StringVar(&includeChars, "i", ".*", "characters to include (regex)")
	flag.StringVar(&fileFilter, "f", ".*", "filter files to read (regex)")
	flag.BoolVar(&verbose, "v", false, "enable verbose output")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("ERROR: no file or directory specified")
		os.Exit(1)
	}
	root := flag.Args()[0]

	include, err := regexp.Compile(includeChars)
	if err != nil {
		fmt.Printf("ERROR: invalid regex for 'include': %s\n", err)
		os.Exit(1)
	}
	filter, err := regexp.Compile(fileFilter)
	if err != nil {
		fmt.Printf("ERROR: invalid regex for 'filter': %s\n", err)
		os.Exit(1)
	}

	reader := NewReader(include)
	count(reader, root, filter)

	NewWriter(os.Stdout).Write(reader.Counts())
}

func count(reader *Reader, root string, filter *regexp.Regexp) error {
	fi, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return countFile(reader, root, filter)
	}

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		return countFile(reader, path, filter)
	})
}

func countFile(reader *Reader, path string, filter *regexp.Regexp) error {
	if !filter.MatchString(path) {
		if verbose {
			log.Printf("ignored: %s\n", path)
		}
		return nil
	}
	if verbose {
		log.Printf("reading: %s\n", path)
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	if err := reader.Read(f); err != nil {
		return fmt.Errorf("reading [%s]: %w", path, err)
	}
	return nil
}
