package store

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"

)

//Controller ...
type Controller struct {
    Repository Repository
}


// AddArticle POST /
func (c *Controller) AddArticle(w http.ResponseWriter, r *http.Request) {
    var article Article

    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields() //unknown fields in json body will be disallowed

    if err := decoder.Decode(&article); err != nil { // unmarshall body contents as a type Candidate
        log.Println(err)
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
    }

    log.Println(article)
    db_article := c.Repository.GetArticleById(article.ID)
    if db_article.ID==article.ID {
	    http.Error(w, "Article with this id already exists!", http.StatusConflict)
	    return
    }
    success := c.Repository.AddArticle(article) // adds the article to the DB
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    return
}

// Get Articles by tag and date GET /
func (c *Controller) GetArticlesByTagAndDate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    queryTag := vars["tagName"] // param tagName
    log.Println("Search tagName - " + queryTag);
    queryDate := vars["date"] // param date
    log.Println("Search date - " + queryDate);

    if len(queryDate)==8{
	    queryDate = queryDate[:4] + "-" + queryDate[4:6] + "-" +  queryDate[6:8]
    }

    articles := c.Repository.GetArticlesByTagAndDate(queryTag, queryDate)

    if len(articles)==0{
	    http.Error(w, "Article with given tag and date was not found!", http.StatusNotFound)
	    return
    }
    var taginfo TagInfo

    taginfo.Tag = queryTag //populating tag
    taginfo.Count = len(articles) //populating count

    // Finding last 10 articles in newest to older order
    last_articles := make([]string, 0)
    for _, v := range articles {
	    last_articles = append(last_articles, v.ID)
    }
    // Limiting list of articles to 10 
    if len(last_articles) > 10{
	    last_articles = last_articles[:10]
    }

    taginfo.Articles = last_articles ////populating last 10 articles or less

    // Trying to find related tags by create a map and using map feature of unique keys to get unique list of related tags
    set := make(map[string]struct{})
    for _, v := range articles { //traversing all articles
        for _, u := range v.Tags {  //traversing individual article tags here
		set[u] = struct{}{} // 'u' is nothing but tag here
	}
    }
    // populating related tags array with keys of map formed above while excluding the queried tag
    related_tags := make([]string, 0)
    for key := range set {
	if key!=queryTag {
		related_tags = append(related_tags, key)
	}
    }

    taginfo.RelatedTags = related_tags
    // taginfo struct should be ready now to be marshalled
    data, _ := json.Marshal(taginfo)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}


// GetArticle GET - Gets a single article by ID /
func (c *Controller) GetArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    id := vars["id"] // param id
    log.Println(id);


    article := c.Repository.GetArticleById(id)
    if article.ID==""{
	    http.Error(w, "Article with given id was not found!", http.StatusNotFound)
	    return
    }
    data, _ := json.Marshal(article)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

