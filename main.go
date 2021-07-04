package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	input := "input.txt"
	_, err := os.Stat(input); if err != nil {
		log.Fatalf("unable stat input file, %v", err)
	}

	in, err := os.Open(input); if err != nil {
		log.Fatalf("unable to open input file, %v", err)
	}
	defer func(i *os.File) {
		err := i.Close(); if err != nil {
			log.Fatalf("unable to close input file, %v", err)
		}
	}(in)

	scanner := bufio.NewScanner(in)
	out, _ := os.Create("out.md.tmp")
	for scanner.Scan() {
		re := regexp.MustCompile(`^(\d{4})[/-](\d{2})[/-](\d{2})$`)
		if re.MatchString(scanner.Text()) {
			y := re.ReplaceAllString(scanner.Text(), "$1")
			m := re.ReplaceAllString(scanner.Text(), "$2")
			d := re.ReplaceAllString(scanner.Text(), "$3")
			dir := filepath.Join("dist", y, m)
			err = os.MkdirAll(dir, os.ModePerm); if err != nil {
				log.Fatalf("unable to make directories, %v", err)
			}
			f := fmt.Sprintf("%s-%s-%s.md", y, m, d)
			out, err = os.Create(filepath.Join(dir, f)); if err != nil {
				log.Fatalf("unable to create file, %v", err)
			}
			_, err = out.WriteString(fmt.Sprintf("# %s-%s-%s\n", y, m, d)); if err != nil {
				log.Fatalf("unable to write content, %v", err)
			}
		} else {
			_, err = out.WriteString(fmt.Sprintf("%s\n", scanner.Text())); if err != nil {
				log.Fatalf("unable to write content, %v", err)
			}
		}
	}
}
