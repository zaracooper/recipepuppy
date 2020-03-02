package recipepuppy

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

type testCase struct {
	Recipe      string
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
	runEmptyArgTest(t, FindRecipes)

	cases := []testCase{
		{"mashed potatoes", 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{"wagyu steak", 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{"flat bread", 0, true, httpmock.NewStringResponder(200, `😀`)},
		{"hamburger", 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{"hotdog", 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest("q", cases, t, FindRecipes)
}

func TestFindRecipesByIngredients(t *testing.T) {
	runEmptyArgTest(t, FindRecipesByIngredient)

	cases := []testCase{
		{"potato", 1, false, httpmock.NewStringResponder(200, `{"results": [{"title": "Mashed Potatoes", "href": "link.to/recipe", "ingredients":"potatoes"}]}`)},
		{"prawn", 0, false, httpmock.NewStringResponder(404, `{"results": []}`)},
		{"mushroom", 0, true, httpmock.NewStringResponder(200, `😺`)},
		{"chicken", 0, true, httpmock.ResponderFromResponse(&http.Response{StatusCode: 501, Body: errReader(0)})},
		{"kangaroo", 0, true, httpmock.NewErrorResponder(errors.New("test error"))},
	}

	runAPICallTest("i", cases, t, FindRecipesByIngredient)
}
