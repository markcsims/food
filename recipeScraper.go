package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func recipeScraper(url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}
	recipe := Recipe{uri: url}

	getIngredients(&recipe, doc)
	getMethod(&recipe, doc)

	fmt.Printf("%s", recipe)
}

func getMethod(r *Recipe, doc *html.Node) {
	var f func(*html.Node)
	step := 0
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, a := range n.Attr {
				if a.Key == "itemprop" && a.Val == "recipeInstructions" {
					var method string
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode {
							method += c.FirstChild.Data
						} else {
							method += c.Data
						}
					}
					step++
					r.method = append(r.method, Method{step: step, description: method})
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
func getIngredients(r *Recipe, doc *html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for _, a := range n.Attr {
				if a.Key == "itemprop" && a.Val == "ingredients" {
					var ingredient string
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode {
							ingredient += c.FirstChild.Data
						} else {
							ingredient += c.Data
						}
					}
					r.ingredients = append(r.ingredients, ingredient)
					fmt.Print("\n")
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

type Recipe struct {
	uri         string
	ingredients []string
	method      []Method
}

type Method struct {
	step        int
	description string
}
