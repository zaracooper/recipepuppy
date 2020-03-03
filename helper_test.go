package recipepuppy

import (
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func runEmptyArgsTest(t *testing.T, subject func([]string, []string, int) ([]Recipe, error)) {
	if _, err := subject([]string{}, []string{}, 0); err == nil {
		t.Error("Expected error, got nil")
	}
}

func runAPICallTest(cases []testCase, t *testing.T, subject func([]string, []string, int) ([]Recipe, error)) {
	query := url.Values{}

	var (
		recipes  []Recipe
		err      error
		testCase testCase
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for i := 0; i < len(cases); i++ {
		testCase = cases[i]

		if len(testCase.Recipes) > 0 {
			query.Set("q", strings.Join(testCase.Recipes, ","))
		}

		if len(testCase.Ingredients) > 0 {
			query.Set("i", strings.Join(testCase.Ingredients, ","))
		}

		httpmock.RegisterResponderWithQuery("GET", APIHREF, query, testCase.Responder)

		recipes, err = subject(testCase.Recipes, testCase.Ingredients, 1)
		if testCase.GotError {
			if err == nil {
				t.Errorf("Expected error. Got: %v", err)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error. Got: %v", err)
			}
		}

		if len(recipes) != testCase.ResultCount {
			t.Errorf("Expected %d recipes. Got: %v", testCase.ResultCount, recipes)
		}
	}
}
