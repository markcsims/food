package main

import (
	"fmt"
)

func recipeSearchScraper(searchTerm string) []Recipe {

	// url := "http://www.bbc.co.uk/food/recipes/search?dishes[]=" + searchTerm
	/*searchUrl := "http://www.bbc.co.uk/food/recipes/search?keywords=" + searchTerm

	resp, err := http.Get(searchUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}
	*/
	recipes := make([]string, 3)
	done := make(chan []Recipe)

	recipes[0] = "http://www.bbc.co.uk/food/recipes/apple_pie_with_custard_29696"
	recipes[1] = "http://www.bbc.co.uk/food/recipes/appleandblueberrypie_85998"
	recipes[2] = "http://www.bbc.co.uk/food/recipes/how_to_make_apple_59768"

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
		fmt.Println(r)
		recipeScraper(r, recipeChannel)
	}
	close(recipeChannel)
}
