package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>Keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {
	var s SitemapIndex
	var n News

	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _ , location := range s.Locations {

		location = location[:len(location)-1][1:len(location)-1]
		resp, _ := http.Get(location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		fmt.Println(n.Titles)

		for idx, _ := range n.Keywords  {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

	}

	for idx, data := range news_map {

		fmt.Println("\n\n\n",idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}

}
