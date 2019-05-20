package main

import (
	"bitbucket.org/albertogviana/go-microservice/names/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type NamesTestSuite struct {
	suite.Suite
}

func (s *NamesTestSuite) Test_HelloServer() {
	r := mux.NewRouter()
	r.HandleFunc("/demo/hello", HelloServer).Methods("GET")

	server := httptest.NewServer(r)
	defer server.Close()

	helloServerUrl := fmt.Sprintf("%s/demo/hello", server.URL)

	request, err := http.NewRequest("GET", helloServerUrl, nil)

	if err != nil {
		s.Error(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		s.Error(err)
	}

	assert.Equal(s.T(), 200, response.StatusCode)

	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(s.T(), []byte("hello, world!\n"), responseBody)
}

func (s *NamesTestSuite) Test_Insert() {
	person := models.Person{
		Firstname: "Roberta",
		Lastname:  "Estronioli",
	}

	data, _ := json.Marshal(person)

	b := bytes.NewBuffer(data)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.Insert).ServeHTTP(rec, req)

	assert.Equal(s.T(), 201, rec.Code)

}

func TestNamesTestSuite(t *testing.T) {
	suite.Run(t, new(NamesTestSuite))
}

type mockDB struct{}

func (mdb *mockDB) InsertName(data models.Person) error {
	return nil
}
