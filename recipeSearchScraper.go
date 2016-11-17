package main

import (
	"fmt"
)

func recipeSearchScraper(searchTerm string) []Recipe {

	// url := "http://www.bbc.co.uk/food/recipes/search?dishes[]=" + searchTerm
	searchUrl := "http://www.bbc.co.uk/food/recipes/search?keywords=" + searchTerm

	searchResult := fetchPage(searchUrl)

	done := make(chan []Recipe)

	//recipes[0] = "http://www.bbc.co.uk/food/recipes/apple_pie_with_custard_29696"
	recipes := getRecipesFromSearchResult(searchResult)

	processRecipes(recipes, done)

	result := <-done
	fmt.Println("done done done!!!!!!!!!!!")
	return result
}

func processRecipes(recipes []string, done chan []Recipe) {
	recipeChannel := make(chan Recipe)

	go func() {
		/*for elem := range recipeChannel {
			fmt.Println(elem)
		}*/
		var allRecipes []Recipe
		for {
			rec, more := <-recipeChannel
			if more {
				fmt.Println(rec)
				allRecipes = append(allRecipes, rec)
			} else {
				fmt.Println("received all jobs")
				done <- allRecipes
				return
			}
		}
	}()

	for _, r := range recipes {
		recipeScraper(r, recipeChannel)
	}
	close(recipeChannel)
}
