package service

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"time"
)

// This function suite better in another module perhaps
func evaluateAnswer(answer string, q Question) bool {
	if q.Result == answer {
		return true
	} 
	return false
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createToken(user User) (string, error) {
	//create a signer for RSA256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims =&CustomClaims{
		CustomUserInfo: struct {
			Name string
			Role string
		}{user.Name, user.Role},
		TokenType: "level11",
		StandardClaims: &jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 20).Unix(), Issuer: "admin"},
	}
	return t.SignedString(signKey)

}

func newMongoSession() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		err := os.Setenv("MONGO_URL","localhost")
		if err != nil {
			log.Fatal("Enable to set MONGO_URL...:(")
		}
	}
	return mgo.Dial(mongoURL)
}

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB]", 0)
}

// NewMongoRepository initialize a new Repo
// to work with the DB
func NewMongoRepository() Repo {
	logger := newMongoRepositoryLogger()
	session, err := newMongoSession()
	if err != nil {
		logger.Fatalf("Couldn't connect to the database: %v\n", err)
	}

	//defer session.Close()

	return Repo{
		logger:  logger,
		session: session,
	}
}