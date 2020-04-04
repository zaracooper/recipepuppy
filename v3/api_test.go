package recipepuppy

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

type testCase struct {
	Recipe      string
	Ingredients []string
	ResultCount int
	GotError    bool
	Responder   httpmock.Responder
}

type errReader int

func (errReader) Read(p []byte) (int, error) {
	return 0, errors.New("test error")
}

func (errReader) Close() error {
	return errors.New("test error")
}

func TestFindRecipes(t *testing.T) {
	wrappedSubject := func(recipe string, ingredients []string, page int) ([]Recipe, error) {
		return FindRecipes(recipe, page)
	}

	runEmptyArgsTest(t, wrappedSubject)

	cases := []testCase{
		{"mashed potatoes", []string{}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{"wagyu steak", []string{}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{"flat bread", []string{}, 0, true, httpmock.NewStringResponder(200, `😀`)},
		{"hamburger", []string{}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{"hotdog", []string{}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, wrappedSubject)
}

func TestFindRecipesWithIngredients(t *testing.T) {
	runEmptyArgsTest(t, FindRecipesWithIngredients)

	cases := []testCase{
		{"bhaji", []string{"tumeric"}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{"sandwich", []string{"egg"}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{"salad", []string{"mushroom"}, 0, true, httpmock.NewStringResponder(200, `😺`)},
		{"nuggets", []string{"chicken"}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{"steak", []string{"kangaroo"}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, FindRecipesWithIngredients)
}

func TestFindRecipesByIngredients(t *testing.T) {
	wrappedSubject := func(recipe string, ingredients []string, page int) ([]Recipe, error) {
		return FindRecipesByIngredients(ingredients, page)
	}

	runEmptyArgsTest(t, wrappedSubject)

	cases := []testCase{
		{"", []string{"potato"}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{"", []string{"prawn"}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{"", []string{"mushroom"}, 0, true, httpmock.NewStringResponder(200, `😺`)},
		{"", []string{"chicken"}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{"", []string{"kangaroo"}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, wrappedSubject)
}
