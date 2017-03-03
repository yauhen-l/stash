package stash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDeclinePullRequest(t *testing.T) {
	var tests = []struct {
		project      string
		slug         string
		id           int
		version      int
		responseCode int
	}{
		{
			project:      "PROJ",
			slug:         "slug",
			id:           777,
			version:      1,
			responseCode: 200,
		},
		{
			project:      "PROJ",
			slug:         "slug",
			id:           777,
			version:      1,
			responseCode: 404,
		},
	}

	for testNumber, test := range tests {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Test %d: wanted POST but found %s\n", testNumber, r.Method)
			}

			url := *r.URL
			wantPath := fmt.Sprintf("/rest/api/1.0/projects/%s/repos/%s/pull-requests/%d/decline", test.project, test.slug, test.id)
			if url.Path != wantPath {
				t.Errorf("Test %d: want %s, got %s\n", testNumber, wantPath, url.Path)
			}

			version := url.Query()["version"][0]
			if version != fmt.Sprintf("%d", test.version) {
				t.Errorf("Test %d: want %d but got %s\n", testNumber, test.version, version)
			}

			noCheckHeader := r.Header.Get("X-Atlassian-Token")
			if noCheckHeader != "no-check" {
				t.Errorf("Test %d: Want X-Atlassian-Token header value no-check, but got %s\n", testNumber, noCheckHeader)
			}

			if r.Header.Get("Authorization") != "Basic dTpw" {
				t.Errorf("Test %d: want Basic dTpw but found %s\n", testNumber, r.Header.Get("Authorization"))
			}

			w.WriteHeader(test.responseCode)
		}))
		defer testServer.Close()

		url, _ := url.Parse(testServer.URL)
		stashClient := NewClient("u", "p", url)
		err := stashClient.DeclinePullRequest(test.project, test.slug, test.id, test.version)

		if test.responseCode != 200 {
			if err == nil {
				t.Fatalf("Test %d: not expecting error for non-200 response code: %v\n", testNumber, err)
			}
		} else if err != nil {
			t.Fatalf("Test %d: not expecting error: %v\n", testNumber, err)
		}
	}
}
