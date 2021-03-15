package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type NewsMap struct {
    Keyword string
    Location string
}

type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

type Sitemapindex struct {
    Locations []string `xml:"sitemap>loc"`
}

type News struct {
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsRoutine(c chan News, Location string){
    defer wg.Done()
    var n News

	fmt.Println(Location)

    resp, err := http.Get(Location)


	if err != nil {
		fmt.Println(err)
	}

    bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

    xml.Unmarshal(bytes, &n)



    resp.Body.Close()

    c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

    var s Sitemapindex
    resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")

	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

    xml.Unmarshal(bytes, &s)
    news_map := make(map[string]NewsMap)
    resp.Body.Close()
    queue := make(chan News, 30)

    for _, Location := range s.Locations {
        wg.Add(1)

		Location:=strings.Trim(Location,"\n")
        go newsRoutine(queue, Location)
    }

    wg.Wait()
    close(queue)

    for elem := range queue {
        for idx, value := range elem.Keywords {

			fmt.Println(value)

            news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
        }
    }

    p := NewsAggPage{Title: "Amazing News  Aggregator", News: news_map}

    t, err := template.ParseFiles("aggregatorfinish.html")
	
	if err != nil {
		fmt.Println(err)
	}

    defer ime.Time()

    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", newsAggHandler)
    http.ListenAndServe(":8000", nil) 
}