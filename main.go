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
		e.ForEachWithBreak("tr", func(_ int, el *colly.HTMLElement) bool {
			writer.Write([]string{
				el.ChildText("th:nth-child(1)"),
				el.ChildText("th:nth-child(2)"),
				el.ChildText("th:nth-child(3)"),
			})
			return false
		})
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var firstChildText = el.ChildText("td:nth-child(1)")
			var secondChildText = el.ChildText("td:nth-child(2)")
			var thirdChildText = el.ChildText("td:nth-child(3)")
			if firstChildText != "" && secondChildText != "" && thirdChildText != "" {
				writer.Write([]string{
					firstChildText,
					secondChildText,
					thirdChildText,
				})
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
