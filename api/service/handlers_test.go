package service

import (
	"bytes"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersRepo_FindAll(t *testing.T) {
	repo := NewMongoRepository()
	t.Log("Testing nil argument in FindAll...")
	_, err := repo.FindAllUsers(nil)
	if err != nil {
		t.Errorf("Error during quering nil argument %v", err)
	}
	t.Log("Testing other arguments...")
	user := map[string]interface{}{"name": "Gealber"}
	_, err = repo.FindAllUsers(user)
	if err != nil {
		t.Errorf("Error during quering name argument %v", err)
	}
	user = map[string]interface{}{"name": "GealberEl"}
	t.Log("Testing not valid argument...")
	_, err = repo.FindAllUsers(user)
	if err != nil {
		t.Errorf("Error during quering name argument %v", err)
	}
	t.Log("Now testing if bson.M is a valid argument type...")
	_, err = repo.FindAllUsers(bson.M{"name": "Gealber"})
	if err != nil {
		t.Errorf("Error during quering name argument %v", err)
	}
}

func TestQuestionsRepo_FindAllQuestions(t *testing.T) {
	repo := NewMongoRepository()
	t.Log("Testing nil argument in FindAll...")
	_, err := repo.FindAllQuestions(nil)
	if err != nil {
		t.Errorf("Error during quering nil argument %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	client := &http.Client{}
	server := httptest.NewServer(
		http.HandlerFunc(updateUserHandler))
	defer server.Close()

	body := []byte("{\n		\"nickname\":\"Testín\",\n		\"score\":\"10\"}")
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Error creating POST request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error in POST to updateUserHAndler: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 status code, received %s",err)
	}
}