package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

//  Documentation
//  =================
//	https://www.zenrows.com/blog/web-scraping-golang#prerequisites
//	https://github.com/gocolly/colly

// 	These functions are executed in the following order:
// 		OnRequest(): Called before performing an HTTP request with Visit().
// 		OnError(): Called if an error occurred during the HTTP request.
// 		OnResponse(): Called after receiving a response from the server.
// 		OnHTML(): Called right after OnResponse() if the received content is HTML.
// 		OnScraped(): Called after all OnHTML() callback executions.

// 	Commnads
//	====================
// 	go mod init worm.go
// 	go get github.com/gocolly/colly/v2

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
