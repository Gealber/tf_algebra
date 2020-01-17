package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"regexp"
	"time"
)

// User contains all the data from a user
type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Nickname    string        `json:"nickname" bson:"nickname"`
	Password    string        `json:"password" bson:"password"`
	Score       int           `json:"score" bson:"score"`
	Status      bool          `json:"status" bson:"status"`
	Rank        int           `json:"rank" bson:"rank"`
	Performance [3]int        `json:"performance" bson:"performance"`
	CreatedOn   time.Time     `json:"createdOn" bson:"createdOn"`
	Role        string        `json:"role" bson:"role"`
}

// Question contains the data format of a question
// the key fields in this struct are the Statement,
// Result and the Subject to which belong
type Question struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Statement string        `json:"statement" bson:"statement"`
	Result    string        `json:"result" bson:"result"`
	Selected  bool          `json:"selected" bson:"selected"`
	Subject   string        `json:"subject" bson:"subject"`
}

// Topic represent a particular topic
// e.g. Algebra, Statistics, etc...
type Topic struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Questions []interface{} `json:"questions" bson:"questions"`
}
// Questions is a slice of Question
type Questions []Question


type newAnswerResponse struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	StartedAt int64         `json:"startedAt" bson:"startedAt"`
}

type newAnswerReq struct {
	UserName string `json:"userName" bson:"userName"`
	Answer   bool   `json:"answer" bson:"answer"`
}

func (n newAnswerReq) isValid() (valid bool) {
	valid = true
	if n.UserName == "" {
		valid = false
	}
	if fmt.Sprintf("%T", n.Answer) != "bool" {
		valid = false
	}
	return valid
}

type newQuestionRepository interface {
	addQuestion(q Question) (err error)
	getQuestions() (questions Questions, err error)
	getQuestion(id string) (q Questions, err error)
	updateQuestion(id string) (err error)
}

// Response it was created for testing purpose only
type Response struct {
	Text string `json:"text" bson:"text"`
}

// Token contains the jwt-token
type Token struct {
	Token string `json:"token" bson:"token"`
}

// CustomClaims contains the neccessary info
// to create the jwt-token, mmm... so so
type CustomClaims struct {
	CustomUserInfo struct {
		Name string
		Role string
	}
	TokenType string
	*jwt.StandardClaims
}

//This next types correspond to the MongoDB database

// Repo is gonna be used to deal with the DB
// contains only two fields one is the logger
// to display all the posible logs when handling 
// with the DB, the other is session which contains
// a pointer to MongoDB session
type Repo struct {
	logger  *log.Logger
	session *mgo.Session
}

// FindAllUsers returns all the users in the DB
func (r Repo) FindAllUsers(selector map[string]interface{}) ([]User, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("appDb").C("Users")

	var users []User
	err := coll.Find(selector).All(&users)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return users, nil
}

// FindAllQuestions returns all the questions according
// to the selector
func (r Repo) FindAllQuestions(selector map[string]interface{}) ([]Question, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("appDb").C("Questions")

	var questions []Question
	err := coll.Find(selector).All(&questions)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return nil, err
	}
	return questions, nil
}

// GetTopic return the topic by the specified name
func (r Repo) GetTopic(name string) (Topic, error) {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("appDb").C("Topics")

	var topic Topic 
    selector := map[string]interface{}{
		"name": name,
	}
	err := coll.Find(selector).One(&topic)
	if err != nil {
		r.logger.Printf("error: %v\n", err)
		return Topic{}, err
	}

	return topic, nil
}
// UpdateUser update the user's info
func (r Repo) UpdateUser(selector interface{}, update interface{}) error {
	session := r.session.Copy()
	defer session.Close()
	coll := session.DB("appDb").C("Users")

	err := coll.Update(selector, update)
	if err != nil {
		r.logger.Printf("error updating user: %v\n", err)
		return err
	}
	return nil
}

// Middleware is used to handle the CORS Preflight
type Middleware struct {
	OriginRule string
}

func (m *Middleware) allowedOrigin(origin string) bool {
	if m.OriginRule == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(m.OriginRule, origin); matched {
		return true
	}
	return false
}
