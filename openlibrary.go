// Package openlibrary provides an interface to the Open Library API. It is VERY barebones and was created to serve the needs of a handful of developers during a Hackathon. They are interested in improving it once they aren't under a tight time crunch.
package openlibrary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// BaseURL is the hostname of the Open Library API
	BaseURL = "http://openlibrary.org"

	// SearchURL is the search url of the Open Library API
	SearchURL = BaseURL + "/search"

	// WorksURL is the url for works in the Open Library API
	WorksURL = BaseURL + "/works"

	// CoverURL is the hostname for the Open Library covers API
	CoverURL = "http://covers.openlibrary.org"

	// AuthorURL is the url for authors in the Open Library API
	AuthorURL = BaseURL + "/authors"
)

var client *http.Client

// A Search represents the response body from a request.
type Search struct {
	Start int   `json:"start"`
	Found int   `json:"numFound"`
	Docs  []Doc `json:"docs"`
}

// A Work respresents a work response body from a request
type Work struct {
	Description struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"description"`
	Title   string `json:"title"`
	Created struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Photos       []int `json:"photos"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision int    `json:"latest_revision"`
	Key            string `json:"key"`
	BirthDate      string `json:"birth_date"`
	Revision       int    `json:"revision"`
	Type           struct {
		Key string `json:"key"`
	} `json:"type"`
	RemoteIds struct {
		Viaf     string `json:"viaf"`
		Wikidata string `json:"wikidata"`
	} `json:"remote_ids"`
}

// An Author represents a work's author response body from a request
type Author struct {
	Bio   string `json:"bio"`
	Name  string `json:"name"`
	Links []struct {
		URL  string `json:"url"`
		Type struct {
			Key string `json:"key"`
		} `json:"type"`
		Title string `json:"title"`
	} `json:"links"`
	PersonalName string `json:"personal_name"`
	Created      struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Photos       []int `json:"photos"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision int    `json:"latest_revision"`
	Key            string `json:"key"`
	BirthDate      string `json:"birth_date"`
	Revision       int    `json:"revision"`
	Type           struct {
		Key string `json:"key"`
	} `json:"type"`
	RemoteIds struct {
		Viaf     string `json:"viaf"`
		Wikidata string `json:"wikidata"`
	} `json:"remote_ids"`
}

// TitleSearch performs a title search and returns the results.
func TitleSearch(title string) (docs []Doc, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?q=title%%3A%s", SearchURL, url.QueryEscape(title)), nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	c := *getClient()
	resp, err := c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var r Search
	err = json.Unmarshal(body, &r)
	docs = r.Docs
	return
}

// GetWorkByID returns a work given an a Work ID
func GetWorkByID(id string) (work Work, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s.json", WorksURL, id), nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	c := *getClient()
	resp, err := c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &work)
	return
}

// GetAuthorByID returns an author given an Author ID
func GetAuthorByID(id string) (author Author, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s.json", AuthorURL, id), nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	c := *getClient()
	resp, err := c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &author)
	return
}

func getClient() *http.Client {
	if client != nil {
		return client
	}
	client = &http.Client{}
	return client
}
