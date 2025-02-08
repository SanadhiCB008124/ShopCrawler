package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"
    "fmt"
	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

func main() {
	var products []Product
	var visitedUrls sync.Map

	
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		product := Product{
			Url:   e.ChildAttr("a", "href"),
			Image: e.ChildAttr("img", "src"),
			Name:  e.ChildText(".product-name"),
			Price: e.ChildText(".price"),
		}
		products = append(products, product)
	})


	c.OnHTML("a.next",func(e *colly.HTMLElement) {
		nextPage :=e.Attr("href")

		if _,found :=visitedUrls.Load(nextPage); !found{
			fmt.Println("scraping:", nextPage)
			visitedUrls.Store(nextPage,struct{}{})
			e.Request.Visit(nextPage)
		}
	})

	
	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("Products.csv")
		if err != nil {
			log.Fatalln("Failed to create CSV file:", err)
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

		log.Println("Scraping completed. Data saved to products.csv")
	})

	
	err := c.Visit("https://www.scrapingcourse.com/ecommerce")
	if err != nil {
		log.Fatalln("Failed to visit the website:", err)
	}
}
