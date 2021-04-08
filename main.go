package main

import (
	"log"
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

func statLeaders() {
    doc, err := goquery.NewDocument("https://www.pro-football-reference.com/years/2020/leaders.htm")
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("table").Each(func(index int, item *goquery.Selection) {
        theadTag := item.Find("thead").Text()
		tableRowTag := item.Find("tr").Text()
        fmt.Println("NFL Leaders:", index, theadTag, tableRowTag)
    })
}

func passingLeaders() {
    doc, err := goquery.NewDocument("https://www.pro-football-reference.com/years/2020/passing.htm")
    if err != nil {
        log.Fatal(err)
    }

    doc.Find("#div_passing.table_container.is_setup").Each(func(index int, item *goquery.Selection) {
        theadTag := item.Find("thead").Text()
		tableRowTag := item.Find("tr").Text()
        fmt.Println("Passing Leaders:", index, theadTag, tableRowTag)
    })
}

func main() {
    statLeaders()
	passingLeaders()
}