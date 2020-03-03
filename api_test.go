package recipepuppy

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

type testCase struct {
	Recipes     []string
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
	wrappedSubject := func(nonEmptyArg []string, emptyArg []string, page int) ([]Recipe, error) {
		return FindRecipes(nonEmptyArg, page)
	}

	runEmptyArgsTest(t, wrappedSubject)

	cases := []testCase{
		{[]string{"mashed potatoes"}, []string{}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{[]string{"wagyu steak"}, []string{}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{[]string{"flat bread"}, []string{}, 0, true, httpmock.NewStringResponder(200, `😀`)},
		{[]string{"hamburger"}, []string{}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{[]string{"hotdog"}, []string{}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, wrappedSubject)
}

func TestFindRecipesWithIngredients(t *testing.T) {
	runEmptyArgsTest(t, FindRecipesWithIngredients)

	cases := []testCase{
		{[]string{"bhaji"}, []string{"tumeric"}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{[]string{"sandwich"}, []string{"egg"}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{[]string{"salad"}, []string{"mushroom"}, 0, true, httpmock.NewStringResponder(200, `😺`)},
		{[]string{"nuggets"}, []string{"chicken"}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{[]string{"steak"}, []string{"kangaroo"}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, FindRecipesWithIngredients)
}

func TestFindRecipesByIngredients(t *testing.T) {
	wrappedSubject := func(emptyArg []string, nonEmptyArg []string, page int) ([]Recipe, error) {
		return FindRecipesByIngredients(nonEmptyArg, page)
	}

	runEmptyArgsTest(t, wrappedSubject)

	cases := []testCase{
		{[]string{}, []string{"potato"}, 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{[]string{}, []string{"prawn"}, 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{[]string{}, []string{"mushroom"}, 0, true, httpmock.NewStringResponder(200, `😺`)},
		{[]string{}, []string{"chicken"}, 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{[]string{}, []string{"kangaroo"}, 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest(cases, t, wrappedSubject)
}
