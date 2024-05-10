package main

import (
	"flag"
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

//  Run from command line
//	  go build wcons.go
//    ./wcons.exe -numb 88

func main() {

	// define parameters
	visitPtr := flag.String("visit", "https://www.google.com/", "Requested URL for processing")
	allowPtr := flag.String("allow", "", "Allow domains, white lists, e.g. 'google.com,...'")
	flag.Parse()

	fmt.Println("Run setting:")
	fmt.Println("Visit:", *visitPtr)
	fmt.Println("Allow domain:", *allowPtr)
	fmt.Println("=====================================", *allowPtr)

	c := colly.NewCollector(
		// visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(*allowPtr),
	)

	// on every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// print info about new link
		fmt.Printf("New link : %q -> %s\n", e.Text, link)
		// visit new link found on page (only those links are visited which are in AllowedDomains)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// before making a request print "  Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("  Visiting", r.URL.String())
	})

	// start scraping
	c.Visit(*visitPtr)
}
