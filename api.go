package recipepuppy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

// FindRecipes finds recipes that match the search term provided
func FindRecipes(searchTerm string) ([]Recipe, error) {
	query := url.Values{}
	query.Add("q", searchTerm)

	results := &response{}

	err := makeRequest(query, results)
	if err != nil {
		return nil, err
	}

	return results.Recipes, nil
}

// FindRecipesByIngredient finds recipes that use the provided ingredient
func FindRecipesByIngredient(ingredient string) ([]Recipe, error) {
	query := url.Values{}
	query.Add("i", ingredient)

	results := &response{}

	err := makeRequest(query, results)
	if err != nil {
		return nil, err
	}

	return results.Recipes, nil
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
