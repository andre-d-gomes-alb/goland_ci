package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//---------//
// HELPERS //
//---------//
var mRouter *mux.Router
var mArticles []Article
var updArticles []map[string]interface{}
var badArticles []map[string]interface{}
var mErrors []Error
var badIds []interface{}

func init() {
	// Mock Router
	mRouter = Router()
	// Mock Articles
	mArticles = []Article{
		{
			Id:      "1",
			Title:   "Title test1",
			Desc:    "Description test1",
			Content: "Content test1",
		}, {
			Id:      "2",
			Title:   "Title test1",
			Desc:    "Description test2",
			Content: "Content test2",
		},
	}
	updArticles = []map[string]interface{}{
		{
			"Title": "Title test1 update",
		}, {
			"Desc": "Description test1 update",
		}, {
			"Content": "Content test1 update",
		},
	}
	badArticles = []map[string]interface{}{
		{
			"Id":      3, //Type error
			"Title":   "Title error1",
			"Desc":    "Description error1",
			"Content": "Content error1",
		},
		{
			"Id":      "4",
			"Title":   1, //Type error
			"Desc":    "Description error2",
			"Content": "Content error2",
		},
	}
	// Mock Error
	mErrors = []Error{
		{
			Msg:  "Invalid body",
			Code: 400,
		},
		{
			Msg:  "Invalid id",
			Code: 400,
		},
	}
	// Mock Id´s
	badIds = []interface{}{
		4,       //Not exists
		"badId", //It's not id type
		true,    //It's not id type
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mRouter.ServeHTTP(rr, req)
	return rr
}

//-------//
// TESTS //
//-------//

// Create
func TestCreate(t *testing.T) {
	article, _ := json.Marshal(mArticles[0])
	req, _ := http.NewRequest("POST", "/article", bytes.NewBuffer(article))
	response := executeRequest(req)

	var out Article
	err := json.Unmarshal(response.Body.Bytes(), &out)

	assert.Equal(t, http.StatusCreated, response.Code, "Invalid response code")
	assert.Nil(t, err, "Invalid response body format")
	assert.Equal(t, mArticles[0], out, "Invalid response body content")
}

func TestCreateErrorBody(t *testing.T) {
	var out Error

	for i := 0; i < len(badArticles); i++ {
		article, _ := json.Marshal(badArticles[i])
		req, _ := http.NewRequest("POST", "/article", bytes.NewBuffer(article))
		response := executeRequest(req)

		err := json.Unmarshal(response.Body.Bytes(), &out)

		assert.Equal(t, http.StatusBadRequest, response.Code, "Invalid response code")
		assert.Nil(t, err, "Invalid response body format")
		assert.Equal(t, mErrors[0], out, "Invalid response body content")
	}
}

// List
func TestList(t *testing.T) {
	Articles = append(Articles, mArticles[1])

	req, _ := http.NewRequest("GET", "/articles", nil)
	response := executeRequest(req)

	var out []Article
	err := json.Unmarshal(response.Body.Bytes(), &out)

	assert.Equal(t, http.StatusOK, response.Code, "Invalid response code")
	assert.Nil(t, err, "Invalid response body format")
	assert.Equal(t, 2, len(out), "Invalid response body size")
	assert.Equal(t, mArticles, out, "Invalid response body content")
}

// Get
func TestGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/article/"+mArticles[0].Id, nil)
	response := executeRequest(req)

	var out Article
	err := json.Unmarshal(response.Body.Bytes(), &out)

	assert.Equal(t, http.StatusOK, response.Code, "Invalid response code")
	assert.Nil(t, err, "Invalid response body format")
	assert.Equal(t, mArticles[0], out, "Invalid response body content")
}

func TestGetErrorId(t *testing.T) {
	var out Error
	var str string

	for i := 0; i < len(badIds); i++ {
		str = fmt.Sprintf("%v", badIds[i])
		req, _ := http.NewRequest("GET", "/article/"+str, nil)
		response := executeRequest(req)

		err := json.Unmarshal(response.Body.Bytes(), &out)

		assert.Equal(t, http.StatusBadRequest, response.Code, "Invalid response code")
		assert.Nil(t, err, "Invalid response body format")
		assert.Equal(t, mErrors[1], out, "Invalid response body content")
	}
}

// Update
func TestUpdate(t *testing.T) {
	var out Article

	for i := 0; i < len(updArticles); i++ {
		article, _ := json.Marshal(updArticles[i])
		req, _ := http.NewRequest("PUT", "/article/"+mArticles[0].Id, bytes.NewBuffer(article))
		response := executeRequest(req)

		err := json.Unmarshal(response.Body.Bytes(), &out)
		for k, v := range updArticles[i] {
			reflect.ValueOf(&mArticles[0]).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}

		assert.Equal(t, http.StatusOK, response.Code, "Invalid response code")
		assert.Nil(t, err, "Invalid response body format")
		assert.Equal(t, mArticles[0], out, "Invalid response body content")
	}
}

// Delete
func TestDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/article/"+mArticles[0].Id, nil)
	response := executeRequest(req)

	assert.Equal(t, http.StatusNoContent, response.Code, "Invalid response code")
	assert.Equal(t, 0, len(response.Body.Bytes()), "Invalid response body content")

	req2, _ := http.NewRequest("GET", "/article/"+mArticles[0].Id, nil)
	response2 := executeRequest(req2)
	var out Error
	err := json.Unmarshal(response2.Body.Bytes(), &out)
	assert.Equal(t, http.StatusBadRequest, response2.Code, "Invalid response code")
	assert.Nil(t, err, "Invalid response body format")
	assert.Equal(t, mErrors[1], out, "Invalid response body content")
}
