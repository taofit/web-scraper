package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func createFile() (*csv.Writer, *os.File) {
	fName := "data.csv"
	file, err := os.Create((fName))
	if err != nil {
		log.Fatalf("file cannot be created, err: %q", err)
		return nil, nil
	}

	return csv.NewWriter(file), file
}

func writeToFile(writer *csv.Writer) {
	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var row []string
			el.ForEach("th", func(_ int, el *colly.HTMLElement) {
				row = append(row, el.Text)
			})
			if len(row) > 0 {
				writer.Write(row)
				row = []string{}
			}
			el.ForEach("td", func(_ int, el *colly.HTMLElement) {
				row = append(row, el.Text)
			})
			if len(row) > 0 {
				writer.Write(row)
			}
		})
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
func main() {
	writer, file := createFile()
	defer file.Close()
	defer writer.Flush()
	writeToFile(writer)
	fmt.Println("scraping Complete")
}
