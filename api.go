package recipepuppy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Recipe is a recipe for a meal
type Recipe struct {
	Title       string `json:"title,omitempty"`
	Link        string `json:"href,omitempty"`
	Ingredients string `json:"ingredients,omitempty"`
	Picture     string `json:"thumbnail,omitempty"`
}

type response struct {
	Title   string   `json:"title,omitempty"`
	Version float64  `json:"version,omitempty"`
	Link    string   `json:"href,omitempty"`
	Recipes []Recipe `json:"results,omitempty"`
}

// APIHREF is the HREF of the Recipe Puppy API
const APIHREF = "http://recipepuppy.com/api"

// FindRecipes finds recipes that match the recipe titles provided
func FindRecipes(recipeTitles []string, page int) ([]Recipe, error) {
	if isQueryBlank(recipeTitles) {
		return []Recipe{}, errors.New("Recipe titles cannot be blank")
	}

	results := response{}

	err := makeRequest(url.Values{"q": recipeTitles}, &results)
	if err != nil {
		return nil, err
	}

	return results.Recipes, nil
}

// FindRecipesWithIngredients finds recipes that match the recipe titles and ingredients provided
func FindRecipesWithIngredients(recipeTitles []string, ingredients []string, page int) ([]Recipe, error) {
	if isQueryBlank(recipeTitles) || isQueryBlank(ingredients) {
		return []Recipe{}, errors.New("Recipe titles or ingredients cannot be blank")
	}

	results := response{}

	err := makeRequest(url.Values{"q": recipeTitles}, &results)
	if err != nil {
		return nil, err
	}

	return results.Recipes, nil
}

// FindRecipesByIngredient finds recipes that use the provided ingredient
func FindRecipesByIngredient(ingredients []string, page int) ([]Recipe, error) {
	if isQueryBlank(ingredients) {
		return []Recipe{}, errors.New("Ingredient cannot be blank")
	}

	results := response{}

	err := makeRequest(url.Values{"i": ingredients}, &results)
	if err != nil {
		return nil, err
	}

	return results.Recipes, nil
}

func isQueryBlank(query []string) bool {
	return len(strings.TrimSpace(strings.Join(query, ""))) == 0
}

func makeRequest(query url.Values, results interface{}) error {
	req, err := http.NewRequest("GET", APIHREF, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = query.Encode()

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
