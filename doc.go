package openlibrary

import (
	"fmt"
	"strconv"
)

// A Size represents the size of an image.
type Size string

const (
	// SizeSmall is an image suitable for use as a thumbnail on a results page
	SizeSmall Size = "S"

	// SizeMedium is suitable for display on a details page
	SizeMedium Size = "M"

	// SizeLarge is the largest image size
	SizeLarge Size = "L"
)

// A Doc represents an item in the Open Library
type Doc struct {
	AuthorAlternativeName []string `json:"author_alternative_name"`
	AuthorKey             []string `json:"author_key"`
	AuthorName            []string `json:"author_name"`
	CoverEditionKey       string   `json:"cover_edition_key"`
	CoverID               int      `json:"cover_i"`
	EbookCount            int      `json:"ebook_count_i"`
	EditionCount          int      `json:"edition_count"`
	EditionKey            []string `json:"edition_key"`
	FirstPublishYear      int      `json:"first_publish_year"`
	FullText              bool     `json:"has_fulltext"`
	IA                    []string `json:"ia"`
	IABoxID               []string `json:"ia_box_id"`
	IACollectionS         string   `json:"ia_collection_s"`
	IALoadedID            []string `json:"ia_loaded_id"`
	GoodreadsID           []string `json:"id_goodreads"`
	GoogleID              []string `json:"id_google"`
	LibraryThingID        []string `json:"id_librarything"`
	OverdriveID           []string `json:"id_overdrive"`
	ISBN                  []string `json:"isbn"`
	Key                   string   `json:"key"`
	Language              []string `json:"language"`
	LastModifiedID        int      `json:"last_modified_i"`
	LCCN                  []string `json:"lccn"`
	LendingEditionS       string   `json:"lending_edition_s"`
	LendingIDS            string   `json:"lending_identifier_s"`
	OCLC                  []string `json:"oclc"`
	PrintDisabledS        string   `json:"printdisabled_s"`
	PublicScanB           bool     `json:"public_scan_b"`
	PublishDate           []string `json:"publish_date"`
	PublishPlace          []string `json:"publish_place"`
	PublishYear           []int    `json:"publish_year"`
	Publisher             []string `json:"publisher"`
	Seed                  []string `json:"seed"`
	Subject               []string `json:"subject"`
	Text                  []string `json:"text"`
	Title                 string   `json:"title"`
	TitleSuggest          string   `json:"title_suggest"`
	Type                  string   `json:"type"`
}

// CoverURL returns the URL of an image of the specified size for the document.
func (d *Doc) CoverURL(size Size) string {
	var key, value string
	switch {
	case d.Key != "":
		key = "id"
		value = strconv.Itoa(d.CoverID)
	case len(d.ISBN) > 0:
		key = "isbn"
		value = d.ISBN[0]
	case len(d.OCLC) > 0:
		key = "oclc"
		value = d.OCLC[0]
	case len(d.LCCN) > 0:
		key = "lccn"
		value = d.LCCN[0]
	// TODO OLID?

	default:
		return ""
	}
	return fmt.Sprintf("%s/b/%s/%s-%s.jpg", CoverURL, key, value, size)
}
