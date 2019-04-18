package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := searchVNExpress()
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println("results: ", len(results))
	fmt.Println(elapsed)

}

func searchVNExpress() []Article {
	var results []Article
	c := make(chan []Article)

	categories, err := crawlVNExpressCategory()
	checkError(err)

	for category, url := range categories {
		go func() {
			resultsEachCategory, err := crawlVNExpress(category, url)
			checkError(err)
			c <- resultsEachCategory
		}()
	}

	for i := 0; i < len(categories); i++ {
		result := <-c
		results = append(results, result...)
	}

	return results
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
