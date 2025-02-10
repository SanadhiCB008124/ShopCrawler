package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

func main() {
	var products []Product
	var visitedUrls sync.Map
	var mutex sync.Mutex

	pagesToScrape := []string{
		"https://www.scrapingcourse.com/ecommerce/page/1/",
        "https://www.scrapingcourse.com/ecommerce/page/2/",
        "https://www.scrapingcourse.com/ecommerce/page/3/",
        "https://www.scrapingcourse.com/ecommerce/page/4/",
        "https://www.scrapingcourse.com/ecommerce/page/5/",
        "https://www.scrapingcourse.com/ecommerce/page/6/",
        "https://www.scrapingcourse.com/ecommerce/page/7/",
        "https://www.scrapingcourse.com/ecommerce/page/8/",
        "https://www.scrapingcourse.com/ecommerce/page/9/",
        "https://www.scrapingcourse.com/ecommerce/page/10/",
        "https://www.scrapingcourse.com/ecommerce/page/11/",
        "https://www.scrapingcourse.com/ecommerce/page/12/",
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
		colly.Async(true),
	)
	c.SetRequestTimeout(100*time.Second)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	c.Limit(&colly.LimitRule{
        Parallelism: 4,
    })

	
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed:", r.Request.URL, "\nStatus Code:", r.StatusCode, "\nError:", err)
	})

	// scraping product data
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img", "src"),
			Name:  e.ChildText(".product-name"), // Verify this selector
			Price: e.ChildText(".price"),        // Verify this selector
		}

		fmt.Println("Scraped Product:", product.Name, product.Price) // Debug log

		// implementing concurrency 
		mutex.Lock()
		products = append(products, product)
		mutex.Unlock()
	})

	//  pagination
	c.OnHTML("a.next", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if _, found := visitedUrls.Load(nextPage); !found {
			fmt.Println("Visiting:", nextPage)
			visitedUrls.Store(nextPage, struct{}{})
			e.Request.Visit(nextPage)
		}
	})

	//scraping
	for _, page := range pagesToScrape {
		c.Visit(page)
	}
	c.Wait()

	// saving to a CSV file
	file, err := os.Create("Products.csv")
	if err != nil {
		log.Fatalln("Failed to create CSV:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()


	headers := []string{"URL", "Image", "Name", "Price"}
	if err := writer.Write(headers); err != nil {
		log.Println("Error writing headers:", err)
	}


	for _, product := range products {
		record := []string{product.Url, product.Image, product.Name, product.Price}
		if err := writer.Write(record); err != nil {
			log.Println("Error writing record:", err)
		}
	}

	if err := writer.Error(); err != nil {
		log.Fatal("CSV write error:", err)
	}

	log.Println("Scraping complete. Data saved to Products.csv")
}
