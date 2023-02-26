package main

import (
	"fmt"
	"strconv"
	"strings"
	"flag"
	
	"github.com/gocolly/colly"
)

func scrape(url string, verbose bool) float64 {
	c := colly.NewCollector()

	var value float64
	var err error
	c.OnRequest(func(r *colly.Request) {
		if verbose {
			fmt.Println("Getting info from:", r.URL)
		}
	})

	c.OnHTML("span.text-success", func(h *colly.HTMLElement) {
		text := strings.TrimSpace(h.Text)
		value, err = strconv.ParseFloat(text, 64)
		if err != nil {
			err = fmt.Errorf("Error converting string to float64")
			panic(err)
		}
	})
	
	c.OnResponse(func(r *colly.Response) {
		if verbose {
			fmt.Println("Status: ", r.StatusCode)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		if verbose {
			fmt.Println("Request URL: ", r.Request.URL, "failed with response:", err)
		}
	})

	c.Visit(url)
	
	return value
}

func main() {
	c_f := flag.String("from", "usd", "name of the currency you want to convert from")
	c_t := flag.String("to", "bdt", "name of the currency you want to convert to")
	amount := flag.Int64("amount", 1, "amount of currency")
	verbose := flag.Bool("verbose", false, "verbose log")
	flag.Parse()

	url := fmt.Sprintf("https://wise.com/gb/currency-converter/%s-to-%s-rate?amount=%d",
		*c_f, *c_t, *amount)

	converted_value := scrape(url, *verbose)
	fmt.Printf("%.3f\n", converted_value * float64(*amount))
}
