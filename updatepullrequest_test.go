package stash

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const updatePullRequestResponse string = `
{
    "id": 1,
    "version": 101,
    "title": "a title",
    "description": "a description",
    "state": "OPEN",
    "open": true,
    "closed": false,
    "createdDate": 1435759062673,
    "updatedDate": 1435759062673,
    "fromRef": {
        "id": "refs/heads/feature/file1",
        "displayId": "feature/file1",
        "latestChangeset": "aead30bdfe27e176316bb2e2aedd530052730092",
        "repository": {
            "slug": "test-repo",
            "id": 1419,
            "name": "test-repo",
            "scmId": "git",
            "state": "AVAILABLE",
            "statusMessage": "Available",
            "forkable": true,
            "project": {
                "key": "PLAT",
                "id": 349,
                "name": "plat",
                "description": "plat Bill Pay and Top Up Microservices",
                "public": true,
                "type": "NORMAL",
                "isPersonal": false,
                "link": {
                    "url": "/projects/plat",
                    "rel": "self"
                },
                "links": {
                    "self": [
                        {
                            "href": "http://localhost:7990/projects/plat"
                        }
                    ]
                }
            },
            "public": false,
            "link": {
                "url": "/projects/plat/repos/test-repo/browse",
                "rel": "self"
            },
            "cloneUrl": "http://admin@localhost:7990/scm/stamp/test-repo.git",
            "links": {
                "clone": [
                    {
                        "href": "http://admin@localhost:7990/scm/stamp/test-repo.git",
                        "name": "http"
                    },
                    {
                        "href": "ssh://git@localhost:7999/stamp/test-repo.git",
                        "name": "ssh"
                    }
                ],
                "self": [
                    {
                        "href": "http://localhost:7990/projects/plat/repos/test-repo/browse"
                    }
                ]
            }
        }
    },
    "toRef": {
        "id": "refs/heads/develop",
        "displayId": "develop",
        "latestChangeset": "3558d035edf10cb54e316374b9e8403a686995ac",
        "repository": {
            "slug": "test-repo",
            "id": 1419,
            "name": "test-repo",
            "scmId": "git",
            "state": "AVAILABLE",
            "statusMessage": "Available",
            "forkable": true,
            "project": {
                "key": "PLAT",
                "id": 349,
                "name": "Platform Dev",
                "description": "It's a project",
                "public": true,
                "type": "NORMAL",
                "isPersonal": false,
                "link": {
                    "url": "/projects/plat",
                    "rel": "self"
                },
                "links": {
                    "self": [
                        {
                            "href": "http://localhost:7990/projects/plat"
                        }
                    ]
                }
            },
            "public": false,
            "link": {
                "url": "/projects/plat/repos/test-repo/browse",
                "rel": "self"
            },
            "cloneUrl": "http://admin@localhost:7990/scm/stamp/test-repo.git",
            "links": {
                "clone": [
                    {
                        "href": "http://admin@localhost:7990/scm/stamp/test-repo.git",
                        "name": "http"
                    },
                    {
                        "href": "ssh://git@localhost:7999/stamp/test-repo.git",
                        "name": "ssh"
                    }
                ],
                "self": [
                    {
                        "href": "http://localhost:7990/projects/plat/repos/test-repo/browse"
                    }
                ]
            }
        }
    },
    "author": {
        "user": {
            "name": "mike",
            "emailAddress": "mike@myemail.com",
            "id": 877,
            "displayName": "Mike",
            "active": true,
            "slug": "mike",
            "type": "NORMAL",
            "link": {
                "url": "/users/mike",
                "rel": "self"
            },
            "links": {
                "self": [
                    {
                        "href": "http://localhost:7990/users/mike"
                    }
                ]
            }
        },
        "role": "AUTHOR",
        "approved": false
    },
    "reviewers": [
        {
            "user": {
                "name": "bob",
                "emailAddress": "bob@myemail.com",
                "id": 871,
                "displayName": "Bob",
                "active": true,
                "slug": "bob",
                "type": "NORMAL",
                "link": {
                    "url": "/users/bob",
                    "rel": "self"
                },
                "links": {
                    "self": [
                        {
                            "href": "http://localhost:7990/users/bob"
                        }
                    ]
                }
            },
            "role": "REVIEWER",
            "approved": false
        },
        {
            "user": {
                "name": "bill",
                "emailAddress": "bill@myemail.com",
                "id": 871,
                "displayName": "Bill",
                "active": true,
                "slug": "bill",
                "type": "NORMAL",
                "link": {
                    "url": "/users/bill",
                    "rel": "self"
                },
                "links": {
                    "self": [
                        {
                            "href": "http://localhost:7990/users/bill"
                        }
                    ]
                }
            },
            "role": "REVIEWER",
            "approved": false
        }
    ],
    "participants": [],
    "link": {
        "url": "/projects/plat/repos/test-repo/pull-requests/2",
        "rel": "self"
    },
    "links": {
        "self": [
            {
                "href": "http://localhost:7990/projects/plat/repos/test-repo/pull-requests/2"
            }
        ]
    }
}
`

func TestUpdatePullRequest(t *testing.T) {

	expectedRequestBody := `{"version":100,"title":"a title","description":"a description","toRef":{"id":"develop","repository":{"slug":"bar","project":{"key":"proj"}}},"reviewers":[{"user":{"name":"bob"}},{"user":{"name":"bill"}}]}`

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Fatalf("wanted PUT but found %s\n", r.Method)
		}
		url := *r.URL
		if url.Path != "/rest/api/1.0/projects/proj/repos/bar/pull-requests/1" {
			t.Fatalf("UpdatePullRequest() URL path expected to be /rest/api/1.0/projects/proj/repos/bar/pull-requests/1 but found %s\n", url.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Fatalf("UpdatePullRequest() expected request Accept header to be application/json but found %s\n", r.Header.Get("Accept"))
		}
		if r.Header.Get("Authorization") != "Basic dTpw" {
			t.Fatalf("Want Basic dTpw but found %s\n", r.Header.Get("Authorization"))
		}

		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != string(expectedRequestBody) {
			t.Fatalf("Unexpected request body\n %s\n expected\n %s\n", body, expectedRequestBody)
		}

		w.WriteHeader(200)
		fmt.Fprint(w, updatePullRequestResponse)
	}))
	defer testServer.Close()

	url, _ := url.Parse(testServer.URL)
	stashClient := NewClient("u", "p", url)

	reviewers := []string{"bob", "bill"}
	pullRequest, err := stashClient.UpdatePullRequest("proj", "bar", "1", 100, "a title", "a description", "develop", reviewers)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if pullRequest.ID != 1 {
		t.Fatalf("Want 1 but got %v\n", pullRequest.ID)
	}
	if pullRequest.Title != "a title" {
		t.Fatalf("Want 'a title' but got %v\n", pullRequest.Title)
	}
	if pullRequest.Description != "a description" {
		t.Fatalf("Want 'a description' but got %v\n", pullRequest.Description)
	}
	if pullRequest.Open != true {
		t.Fatalf("Want true but got %v\n", pullRequest.Open)
	}
	if pullRequest.State != "OPEN" {
		t.Fatalf("Want OPEN but got %v\n", pullRequest.State)
	}
	if pullRequest.FromRef.DisplayID != "feature/file1" {
		t.Fatalf("Want feature/file1 but got %v\n", pullRequest.FromRef)
	}
	if pullRequest.ToRef.DisplayID != "develop" {
		t.Fatalf("Want develop but got %v\n", pullRequest.ToRef)
	}
}
