package main

import (
	"fmt"
)

func recipeSearchScraper(searchTerm string) {

	// url := "http://www.bbc.co.uk/food/recipes/search?dishes[]=" + searchTerm

	// resp, err := http.Get(url)

	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()
	// doc, err := html.Parse(resp.Body)

	// if err != nil {
	// 	panic(err)
	// }
	recipes := make([]string, 3)

	recipes[0] = "http://www.bbc.co.uk/food/recipes/apple_pie_with_custard_29696"
	recipes[1] = "http://www.bbc.co.uk/food/recipes/appleandblueberrypie_85998"
	recipes[2] = "http://www.bbc.co.uk/food/recipes/how_to_make_apple_59768"

	processRecipes(recipes)

	fmt.Println("done done done!!!!!!!!!!!")
}

func processRecipes(recipes []string) {
	recipeChannel := make(chan Recipe)

	for _, r := range recipes {
		fmt.Println(r)
		go recipeScraper(r, recipeChannel)
	}

	go func() {
		for elem := range recipeChannel {
			fmt.Println(elem)
		}
	}()
}
