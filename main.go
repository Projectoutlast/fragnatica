package main

import (
	"bufio"
	"log"
	"os"

	scr "fragnatica/scraper"
)

type Images struct {
	Url string
}

func main() {
	target, err := os.Open("target.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	scanner := bufio.NewScanner(target)
	scanner.Split(bufio.ScanLines)

	var urls []string

	for scanner.Scan() {
		line := scanner.Text()
		urls = append(urls, line)
	}

	images := scr.CollyScrape(urls)

	for _, image := range images {
		scr.SaveImage(image.Url, "downloads")
	}
}
