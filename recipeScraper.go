package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func recipeScraper(url string, recipeChannel chan<- Recipe) {
	doc := fetchPage(url)
	recipe := Recipe{uri: url}

	fmt.Println(recipe)
	getIngredients(&recipe, doc)
	getMethod(&recipe, doc)

	recipeChannel <- recipe
}

func fetchPage(url string) *html.Node {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}
	return doc
}

func getRecipesFromSearchResult(doc *html.Node) []string {
	var allRecipes []string
	var finda func(*html.Node)
	finda = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.HasPrefix(a.Val, "/food/recipes") {
					allRecipes = append(allRecipes, "http://www.bbc.co.uk"+a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			finda(c)
		}
	}

	var findLi func(*html.Node)
	findLi = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.HasPrefix(a.Val, "article") {
					finda(n)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLi(c)
		}
	}
	findLi(doc)

	return allRecipes
}
func getMethod(r *Recipe, doc *html.Node) {
	action := func(val *html.Node, step int) {
		method := getValue(val)
		r.method = append(r.method, Method{step: step, description: method})
	}
	findAttr(doc, "recipeInstructions", action)
}
func getIngredients(r *Recipe, doc *html.Node) {
	action := func(val *html.Node, step int) {
		ingredient := getValue(val)
		r.ingredients = append(r.ingredients, ingredient)
	}
	findAttr(doc, "ingredients", action)
}

func findAttr(doc *html.Node, name string, action func(doc *html.Node, step int)) {
	var f func(*html.Node)
	i := 0
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, a := range n.Attr {
				if a.Key == "itemprop" && a.Val == name {
					i++
					action(n, i)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func getValue(n *html.Node) string {
	var value string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			value += c.FirstChild.Data
		} else {
			value += c.Data
		}
	}
	return value
}

type Recipe struct {
	uri         string
	ingredients []string
	method      []Method
}

type Method struct {
	step        int
	description string
}
