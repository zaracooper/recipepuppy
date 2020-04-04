package recipepuppy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Recipe is a recipe for a meal.
type Recipe struct {
	Title       string `json:"title,omitempty"`
	Link        string `json:"href,omitempty"`
	Picture     string `json:"thumbnail,omitempty"`
	Ingredients []string
}

type recipeResponse struct {
	Recipe
	IngredientString string `json:"ingredients,omitempty"`
}

type response struct {
	Title   string           `json:"title,omitempty"`
	Version float64          `json:"version,omitempty"`
	Link    string           `json:"href,omitempty"`
	Recipes []recipeResponse `json:"results,omitempty"`
}

// APIHREF is the HREF of the Recipe Puppy API.
const APIHREF = "http://recipepuppy.com/api"

// FindRecipes finds recipes that match the recipe title provided. The API returns
// 10 results at a time. So use page to specify what page of results to fecth.
func FindRecipes(recipeTitle string, page int) ([]Recipe, error) {
	return processQuery(true, false, recipeTitle, []string{}, page, "Recipe title cannot be blank. Page cannot be zero")
}

// FindRecipesWithIngredients finds recipes that match the recipe title and ingredients provided. The API returns
// 10 results at a time. So use page to specify what page of results to fecth.
func FindRecipesWithIngredients(recipeTitle string, ingredients []string, page int) ([]Recipe, error) {
	return processQuery(true, true, recipeTitle, ingredients, page, "Recipe title or ingredients cannot be blank. Page cannot be zero")
}

// FindRecipesByIngredients finds recipes that use the provided ingredient. The API returns
// 10 results at a time. So use page to specify what page of results to fecth.
func FindRecipesByIngredients(ingredients []string, page int) ([]Recipe, error) {
	return processQuery(false, true, "", ingredients, page, "Ingredient cannot be blank. Page cannot be zero")
}

func unpackRecipes(resp response) []Recipe {
	recipeResp := recipeResponse{}
	unpackedRecipes := []Recipe{}

	for i := 0; i < len(resp.Recipes); i++ {
		recipeResp = resp.Recipes[i]
		recipeResp.Title = strings.TrimSpace(recipeResp.Title)
		recipeResp.Ingredients = strings.Split(recipeResp.IngredientString, ", ")
		unpackedRecipes = append(unpackedRecipes, recipeResp.Recipe)
	}

	return unpackedRecipes
}

func processQuery(searchRecipe, searchIngr bool, recipeTitle string, ingredients []string, page int, errString string) ([]Recipe, error) {
	if len(strings.TrimSpace(recipeTitle)) == 0 && searchRecipe || len(strings.TrimSpace(strings.Join(ingredients, ""))) == 0 && searchIngr || page == 0 {
		return []Recipe{}, errors.New(errString)
	}

	var queryVal string

	if searchRecipe {
		queryVal = "q=" + recipeTitle
		if searchIngr {
			queryVal += "&"
		}
	}

	if searchIngr {
		queryVal += ("i=" + strings.Join(ingredients, ","))
	}

	results := response{}

	err := makeRequest(queryVal, &results)
	if err != nil {
		return nil, err
	}

	return unpackRecipes(results), nil
}

func makeRequest(query string, results interface{}) error {
	req, err := http.NewRequest("GET", APIHREF, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = query

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, results)
	if err != nil {
		return err
	}

	return nil
}
