package main

import (
	"fmt"
	"strings"

	color "github.com/fatih/color"
	"github.com/zaracooper/recipepuppy/v2"
)

func main() {
	var (
		page        = 1
		recipeTitle string
		ingredients []string
	)

	recipeTitle = "french toast"
	recipes, err := recipepuppy.FindRecipes(recipeTitle, page)
	printRecipes(recipes, page, []string{recipeTitle}, err)

	recipeTitle = "fried chicken"
	ingredients = []string{"eggs", "garlic"}
	recipes, err = recipepuppy.FindRecipesWithIngredients(recipeTitle, ingredients, page)
	printRecipes(recipes, page, append(ingredients, recipeTitle), err)

	ingredients = []string{"duck", "cheese"}
	recipes, err = recipepuppy.FindRecipesByIngredients([]string{"duck"}, page)
	printRecipes(recipes, page, ingredients, err)
}

func printRecipes(recipes []recipepuppy.Recipe, page int, searchTerms []string, err error) {
	color.Green(fmt.Sprintf("Query: %s [page %v]\n\n", strings.Join(searchTerms, ", "), page))

	if err == nil && len(recipes) != 0 {
		var recipe recipepuppy.Recipe

		for i := 0; i < len(recipes); i++ {
			if recipe = recipes[i]; len(recipe.Picture) == 0 {
				recipe.Picture = "None"
			}

			fmt.Println(color.BlueString("No: "), "#", i+1)
			fmt.Println(color.BlueString("Title: "), color.YellowString(recipe.Title))
			fmt.Println(color.BlueString("Link: "), recipe.Link)
			fmt.Println(color.BlueString("Picture: "), recipe.Picture)
			fmt.Print(color.BlueString("Ingredients: "), strings.Join(recipe.Ingredients, ", "), "\n\n")
		}
	} else {
		color.Red("Sorry. There were no results for your query! 😔\n\n")
	}
}
