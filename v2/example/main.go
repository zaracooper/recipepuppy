package main

import (
	"fmt"
	"strings"

	"github.com/zaracooper/recipepuppy/v2"
)

func main() {
	var (
		page         = 1
		recipeTitles []string
		ingredients  []string
	)

	recipeTitles = []string{"bhaji"}
	recipes, err := recipepuppy.FindRecipes(recipeTitles, page)
	if err != nil {
		fmt.Println("Failed to get bhaji recipes")
	} else {
		printRecipes(recipes, page, recipeTitles)
	}

	recipeTitles = []string{"steak"}
	ingredients = []string{"eggs"}
	recipes, err = recipepuppy.FindRecipesWithIngredients(recipeTitles, ingredients, page)
	if err != nil {
		fmt.Println("Failed to get steak recipes made with eggs")
	} else {
		printRecipes(recipes, page, append(recipeTitles, ingredients...))
	}

	ingredients = []string{"duck"}
	recipes, err = recipepuppy.FindRecipesByIngredients([]string{"duck"}, page)
	if err != nil {
		fmt.Println("Failed to get recipes with duck as an ingredient")
	} else {
		printRecipes(recipes, page, ingredients)
	}
}

func printRecipes(recipes []recipepuppy.Recipe, page int, searchTerms []string) {
	if len(recipes) != 0 {
		fmt.Printf("===RESULTS FOR %s===\n---PAGE %v---\n\n", strings.Join(searchTerms, ", "), page)
		var recipe recipepuppy.Recipe

		for i := 0; i < len(recipes); i++ {
			if recipe = recipes[i]; len(recipe.Picture) == 0 {
				recipe.Picture = "None"
			}

			fmt.Printf("Recipe %v\n---\nTitle: %s,\nLink: %s,\nPicture: %s,\nIngredients: %s\n~\n\n",
				i+1,
				recipe.Title, recipe.Link, recipe.Picture, recipe.Ingredients)
		}

		fmt.Print("\n===DONE===\n\n")
	} else {
		fmt.Print("Sorry\n---\nThere were no results for your query! 😔\n\n\n")
	}
}
