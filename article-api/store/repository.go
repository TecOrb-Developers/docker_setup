package store

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "mongodb://test:test@localhost:27017/Articles"

// DBNAME the name of the DB instance
const DBNAME = "Articles"

// COLLECTION is the name of the collection in DB
const COLLECTION = "store"

var articleId = 10;

// GetArticleById returns a unique Article
func (r Repository) GetArticleById(id string) Article {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Article

	fmt.Println("ID in GetArticleById", id);

	if err := c.FindId(id).One(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetArticlesByTagAndDate takes tagName and date as search parameters and returns articles
func (r Repository) GetArticlesByTagAndDate(queryTag string, queryDate string) Articles {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := Articles{}

	and := []bson.M{ bson.M{ "date":queryDate }, bson.M{ "tags": bson.M{ "$in":[]string{queryTag} } } }

	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Sort("-_id").All(&result); err != nil {
		fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddArticle adds a Article in the DB
func (r Repository) AddArticle(article Article) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTION)
        c.Insert(article);
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Article ID- ", article.ID)

	return true
}
