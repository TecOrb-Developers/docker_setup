package store

type TagInfo struct {
        Tag  string  `json:"tag"`
        Count int     `json:"count"`
        Articles []string     `json:"articles"`
        RelatedTags []string     `json:"related_tags"`
}

// Article represents an article item
type Article struct {
	ID     string	     `bson:"_id"`
	Title  string        `json:"title"`
	Date  string        `json:"date"`
	Body   string       `json:"body"`
	Tags  []string        `json:"tags"`
}

// Articles is an array of Article objects
type Articles []Article
