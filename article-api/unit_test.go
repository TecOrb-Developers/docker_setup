package main

import (
    "os"
    "log"
    "testing"
    "bytes"
    "net/http"
    "gopkg.in/mgo.v2"
    "net/http/httptest"
    "article-api/store"
    "github.com/gorilla/mux"
)

// SERVER the DB server
const SERVER = "mongodb://test:test@localhost:27017/Articles"

// DBNAME the name of the DB instance
const DBNAME = "Articles"

// COLLECTION is the name of the collection in DB
const COLLECTION = "store"

var router *mux.Router // create routes

func TestMain(m *testing.M) {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    session,err := mgo.Dial(SERVER)
    if err != nil {
	    log.Println("Failed to establish connection to Mongo server:", err)
    }
    session.DB(DBNAME).C(COLLECTION).RemoveAll(nil) //clear all documents from collection - store
    router = store.NewRouter()
    os.Exit(m.Run())
}

func TestGetNonExistentArticle(t *testing.T) {
    req, _ := http.NewRequest("GET", "/articles/nonexistentid", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusNotFound, response.Code)
    if body := response.Body.String(); body != "Article with given id was not found!\n" {
	    t.Errorf("Expected:Article with given id was not found!. Got: %s", body)
    }
}

func TestAddArticle(t *testing.T) {
    var jsonStr = []byte(`{"id": "someuniqueid", "title": "latest science shows that potato chips are better for you than sugar", "date" : "2016-09-22",   "body" : "I love fairfax", "tags" : ["health", "fitness", "media"]}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)
    checkResponseCode(t, http.StatusCreated, response.Code)
    if body := response.Body.String(); body != "" {
	    t.Errorf("Expected an empty response body. Got: %s", body)
    }
}


func TestAddDuplicateIdArticle(t *testing.T) {
    var jsonStr = []byte(`{"id": "someuniqueid", "title": "Australia won Ashes 4-1", "date" : "2019-09-22",   "body" : "I love Australia", "tags" : ["health", "fitness", "sports"]}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)
    checkResponseCode(t, http.StatusConflict, response.Code)
    if body := response.Body.String(); body != "Article with this id already exists!\n" {
	    t.Errorf("Expected:Article with this id already exists!. Got: %s", body)
    }
}


func TestAddInvalidArticle(t *testing.T) {
    var jsonStr = []byte(`{"id": "someotheruniqueid", "heading": "Australia won Ashes 4-1", "date" : "2019-09-22",   "body" : "I love Australia", "tags" : ["health", "fitness", "sports"]}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    response := executeRequest(req)
    checkResponseCode(t, http.StatusBadRequest, response.Code)
    if body := response.Body.String(); body != "json: unknown field \"heading\"\n" {
	    t.Errorf("Expected:json: unknown field \"heading\". Got: %s", body)
    }
}

func TestGetExistentArticle(t *testing.T) {
    req, _ := http.NewRequest("GET", "/articles/someuniqueid", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    if body := response.Body.String(); body == "[]" {
	    t.Errorf("Expected: Body of article. Got: %s", body)
    }
}


func TestGetArticleByTagAndDate(t *testing.T) {

    var jsonStr = []byte(`{"id": "someuniqueid1", "title": "I passed AWS certified solutions architect associate", "date" : "2018-08-13",   "body" : "I am in awe of AWS", "tags" : ["cloud", "technology", "awsome"]}`)
    req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    executeRequest(req)

    jsonStr = []byte(`{"id": "someuniqueid2", "title": "I have 3+ years of Devops related experience", "date" : "2018-08-13",   "body" : "I am in awe of AWS", "tags" : ["cloud", "technology", "awsome"]}`)
    req, _ = http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
    executeRequest(req)

    req, _ = http.NewRequest("GET", "/tags/cloud/20180813", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    if body := response.Body.String(); body == "[]" {
	    t.Errorf("Expected: Tag relevant info. Got: %s", body)
    }
}


func TestGetNonExistentArticleByTagAndDate(t *testing.T) {
    req, _ := http.NewRequest("GET", "/tags/nonexistent/20190922", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusNotFound, response.Code)
    if body := response.Body.String(); body != "Article with given tag and date was not found!\n" {
	    t.Errorf("Expected: Article with given tag and date was not found!. Got: %s", body)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)
    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
