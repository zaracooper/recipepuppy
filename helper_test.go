package recipepuppy

import (
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
)

func runEmptyArgTest(t *testing.T, subject func(string) ([]Recipe, error)) {
	if _, err := subject(""); err == nil {
		t.Error("Expected error, got nil")
	}
}

func runAPICallTest(queryKey string, cases []testCase, t *testing.T, subject func(string) ([]Recipe, error)) {
	query := url.Values{}

	var (
		recipes []Recipe
		err     error
	)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for i := 0; i < len(cases); i++ {
		query.Set(queryKey, cases[i].Recipe)

		httpmock.RegisterResponderWithQuery("GET", APIHREF, query, cases[i].Responder)

		recipes, err = subject(cases[i].Recipe)
		if cases[i].GotError {
			if err == nil {
				t.Errorf("Expected error. Got: %v", err)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error. Got: %v", err)
			}
		}

		if len(recipes) != cases[i].ResultCount {
			t.Errorf("Expected %d recipes. Got: %v", cases[i].ResultCount, recipes)
		}
	}
}
