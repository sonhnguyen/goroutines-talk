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
	fmt.Println(elapsed)

}

func searchVNExpress() []Article {
	var results []Article

	categories, err := crawlVNExpressCategory()
	checkError(err)

	for category, url := range categories {
		resultsEachCategory, err := crawlVNExpress(category, url)
		checkError(err)
		results = append(results, resultsEachCategory...)
	}
	return results
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
