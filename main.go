package main

import (
	"github.com/gocolly/colly"
	"fmt"
)

func main(){
	count := 1
	collector := colly.NewCollector(
		colly.AllowedDomains("wionews.com","www.wionews.com"),
	)
	collector.OnHTML(".content-holder h1 a,.content-holder h2 a,.content-holder  h3 a,.content-holder  h4 a,.content-holder  h5 a,.content-holder  h6 a",func (element *colly.HTMLElement){
		fmt.Printf("%v: %v\n",count,element.Text)
		count +=1
	})

	collector.OnRequest(func (request *colly.Request){
		fmt.Println("visiting ",request.URL.String())
	})

	collector.Visit("https://www.wionews.com/")
}