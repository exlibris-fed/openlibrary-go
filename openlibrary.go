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

	// EditionsURL is the url for editions of a work in the Open Library API
	// must provide work ID
	EditionsURL = BaseURL + "/works/%s/editions.json"
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
	Description WorkDescription `json:"description"`
	Title       string          `json:"title"`
	Created     struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Photos       []int `json:"photos"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision int          `json:"latest_revision"`
	Key            string       `json:"key"`
	BirthDate      string       `json:"birth_date"`
	Revision       int          `json:"revision"`
	Type           WorkType     `json:"type"`
	RemoteIds      WorkRemoteID `json:"remote_ids"`
	Authors        []WorkAuthor `json:"authors"`
}

// WorkDescription extracts a work's description field
type WorkDescription string

// UnmarshalJSON will try to unmarshal a complex object, falling back to string before failing
func (d *WorkDescription) UnmarshalJSON(data []byte) error {
	var description WorkDescription
	var obj struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		var value string
		err = json.Unmarshal(data, &value)
		if err != nil {
			return err
		}
		description = WorkDescription(value)
	} else {
		description = WorkDescription(obj.Value)
	}
	*d = description
	return nil
}

// WorkType is the work type
type WorkType struct {
	Key string `json:"key"`
}

// WorkRemoteID is some remote ID
type WorkRemoteID struct {
	Viaf     string `json:"viaf"`
	Wikidata string `json:"wikidata"`
}

// WorkAuthor describes the author of a work
type WorkAuthor struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Author struct {
		Key string `json:"key"`
	} `json:"author"`
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

// EditionsResponse represents the response for editions of a work
type EditionsResponse struct {
	Entries []Edition `json:"entries"`
	Links   struct {
		Self string `json:"self"`
		Work string `json:"work"`
	} `json:"links"`
	Size int `json:"size"`
}

// Edition represents the edition of a work
type Edition struct {
	Publishers     []string `json:"publishers"`
	NumberOfPages  int      `json:"number_of_pages"`
	Subtitle       string   `json:"subtitle"`
	Covers         []int    `json:"covers"`
	LocalID        []string `json:"local_id"`
	PhysicalFormat string   `json:"physical_format"`
	LastModified   struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision  int    `json:"latest_revision"`
	Key             string `json:"key"`
	Classifications struct {
	} `json:"classifications"`
	SourceRecords []string `json:"source_records"`
	Title         string   `json:"title"`
	Identifiers   struct {
		Goodreads []string `json:"goodreads"`
	} `json:"identifiers"`
	Created struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Isbn13      []string `json:"isbn_13"`
	Isbn10      []string `json:"isbn_10"`
	PublishDate string   `json:"publish_date"`
	Works       []struct {
		Key string `json:"key"`
	} `json:"works"`
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Revision int `json:"revision"`
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

// GetEditionsByID returns the editions for a work
func GetEditionsByID(id string) (editions []Edition, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(EditionsURL, id), nil)
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
	var editionsResponse EditionsResponse
	err = json.Unmarshal(body, &editionsResponse)

	editions = editionsResponse.Entries
	return
}

func getClient() *http.Client {
	if client != nil {
		return client
	}
	client = &http.Client{}
	return client
}
