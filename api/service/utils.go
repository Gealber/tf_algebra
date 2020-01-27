package service

import (	
	//"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/acme/autocert"
	"crypto/tls"
	"context"	
	"encoding/json"
	"fmt"
	"log"
	//"net"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
)

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
	t.Claims = &CustomClaims{
		CustomUserInfo: struct {
			Name string
			Role string
		}{user.Name, user.Role},
		TokenType:      "level11",
		StandardClaims: &jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 20).Unix(), Issuer: "admin"},
	}
	return t.SignedString(signKey)

}

func newMongoSession() (*mgo.Session, error) {
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		err := os.Setenv("MONGO_URL", "localhost")
		if err != nil {
			log.Fatal("Enable to set MONGO_URL...:(")
		}
	}

	// dialInfo, err := mgo.ParseURL(mongoURL)
	// if err != nil {
	// 	log.Println(err)
	// }

	// dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
	// 	tlsConfig := getTLSConfig()
	// 	conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	return conn, err
	// }
	
	// session, err := mgo.DialWithInfo(dialInfo)
	// if err != nil {
	// 	log.Println(err)
	// }
	
	return mgo.Dial(mongoURL)

	//return session, err
}

func newMongoRepositoryLogger() *log.Logger {
	return log.New(os.Stdout, "[mongoDB]", 0)
}

func getTLSConfig() *tls.Config {

	dataDir := "./api"
	hostPolicy := func(ctx context.Context, host string) error {
		
		allowedHost := "cybersepro.com"
		if host == allowedHost {
			return nil
		}
		return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
	}

	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: hostPolicy,
		Cache: autocert.DirCache(dataDir),
	}

	return &tls.Config{GetCertificate: m.GetCertificate}

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
