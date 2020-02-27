package openlibrary

// A Doc represents an item in the Open Library
type Doc struct {
	AuthorAlternativeName []string `json:"author_alternative_name"`
	AuthorKey             []string `json:"author_key"`
	AuthorName            []string `json:"author_name"`
	CoverEditionKey       string   `json:"cover_edition_key"`
	CoverI                int      `json:"cover_i"`
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
