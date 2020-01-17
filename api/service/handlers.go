package service

import (
	"github.com/gorilla/mux"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"	
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)


// verify key and sign key
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// This is the handler that retrieve all the questions
func questionHandler(w http.ResponseWriter, _ *http.Request) {
	var questions Questions
	// Add all notes from the database or the temporary storage
	// to the questions variable
	repo := NewMongoRepository()
	questions, err := repo.FindAllQuestions(nil)
	if err != nil {
		repo.logger.Fatal("Error selecting all current questions")
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(questions)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		log.Fatalf("Error %v writing response on ResponseWriter w", err)
	}
}

// Function responsible for the creation of users in the fake storage
//Fixed and work properly
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.CreatedOn = time.Now()
	//here comes the work with the database
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//init a session in appDb
	c := session.DB("appDb").C("Users")
	//Inserting data
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	j, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(j)
	if err != nil {
		log.Fatalf("Error %v writing response on ResponseWriter w", err)
	}
}

// Function that retrieve all the available users from the
// fakeUsersStorage. When the database is working we need to make
// a some changes in this function
//Work properly
func usersHandler(w http.ResponseWriter, r *http.Request) {
	var users []User

	repo := NewMongoRepository()
	users, err := repo.FindAllUsers(nil)
	if err != nil {
		repo.logger.Fatal("Error selecting all current users")
	}
	
	j, err := json.Marshal(&users)
	if err != nil {
		panic(err)
	}
	
	w.Header().Set("Content-Type", "application/json")		
	// w.Header().Set("Access-Control-Allow-Headers","*")
	// w.Header().Set("Lola es","Puta")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		log.Fatalf("Error %v writing response on ResponseWriter w", err)
	} 
}

// updateUserHandler update only until now just the score of the user
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	repo := NewMongoRepository()

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	selector := map[string]interface{}{
		"nickname": data["nickname"],
	}

	score, err := strconv.Atoi(data["score"])
	users, err := repo.FindAllUsers(selector)
	users[0].Score = score
	err = repo.UpdateUser(selector, users[0])
	if err != nil {
		repo.logger.Fatal("Error updating database")
	}

	w.WriteHeader(http.StatusOK)
}
// read the key files before starting http handlers
//work properly
func init() {
	var err error

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

}

// reads the login credentials, passed through a POST request(obvious)
// checks them and creates JWT the token to be passed to authHandler
//work properly
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	// Verify if the method is appropriate
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No POST", r.Method)
		return
	}

	// decode the data from the post request into the User
	// struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in request body")
		return
	}

	//validate user credentials
	// this is done against the database
	repo := NewMongoRepository()
	query, err := repo.FindAllUsers(bson.M{"nickname":user.Nickname, "password":user.Password})
	if err != nil {
		repo.logger.Fatalf("Error during quering name argument %v", err)
	}

	//if the user is already in the database we send a 403 error
	//maybe change this to send more info
	if len(query) == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Wrong info")
		return
	}

	//create a signer for rsa 256
	//t := jwt.New(jwt.GetSigningMethod("RS256"))
	// set our claims
	// setting the role
	//iss. Is the issuer of the claim. Connect uses it to identify the application making the call.
	//Create a new token with claims
	// All this is done by the createToken(nice name) method
	tokenString, err := createToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := Token{tokenString}
	//send the token as a HTTP response
	jsonResponse(response, w)
}

// only accessible with a valid token
//remember that after the user is logged in every time
//the user request for a resource he musts passed the token.
// Ok we need a middleware to pass every request through it and verify if the user has
// the JWT token, in this way the authentication is more trusty
// A middleware receive a http.Handler and return a http.HandlerFunc which has a
// ServeHTTP method so this type also satisfies with the interface (http.Handler) contract.
//work properly
func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//validate the token
		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		//if the token is not a valid JWT token we send an Unauthorized HTTP status code
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError: //Something was wrong during the validation
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintln(w, "Token expired, get a ne one.")
					return
				case jwt.ValidationErrorMalformed:
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintln(w,"Something is malformed")
					return

				default:
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w,"Error while parsing the token")
					log.Printf("Validation error: %+v\n", vErr.Errors)
					return
				}
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w,"Error while parsing the token")
				log.Printf("Validation error: %+v\n", err)
				return
			}
		}

		//if token is valid then we send the text Authorized to the system
		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}

// corsHandler manage CORS preflight
func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			m := Middleware{OriginRule:"*"}
			if m.allowedOrigin(r.Header.Get("Origin")) {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			}
			return
		}
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

// topicHandler retieve the information related 
// to the specific topic request in the name parameter in the url
func topicHandler(w http.ResponseWriter, r *http.Request) {
	var topic Topic
	
	repo := NewMongoRepository()
	name := mux.Vars(r)["name"]
	topic, err := repo.GetTopic(name)
	if err != nil {
		repo.logger.Fatal("Error selecting the requested topic")
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(topic)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		log.Fatalf("Error %v writing response on ResponseWriter w", err)
	}
}