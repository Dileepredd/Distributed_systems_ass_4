package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main(){
	collector := colly.NewCollector(
		colly.AllowedDomains("wionews.com","www.wionews.com"),
	)
	collector.OnHTML(".content-holder h1 a,.content-holder h2 a,.content-holder  h3 a,.content-holder  h4 a,.content-holder  h5 a,.content-holder  h6 a",func (element *colly.HTMLElement){
		fmt.Printf("%v\n",element.Text)
	})

	collector.OnRequest(func (request *colly.Request){
		fmt.Println("visiting %v",request.URL.String())
	})

	collector.Visit("https://www.wionews.com/")
}