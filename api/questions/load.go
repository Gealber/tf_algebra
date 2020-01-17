package main

import (	
	"bufio"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("Algebra.txt")
	if err != nil {
		log.Fatalf("This error raised %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	c := session.DB("appDb").C("Questions")
	var contentArray []interface{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "---")
		if len(line) > 2 {
			contentArray = append(contentArray, bson.M{
				"statement": line[1],
				"result":    line[0],
				"selected":  false,
				"subject":   line[2],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = c.Insert(contentArray...)
	if err != nil {
		log.Fatal(err)
	}

	c = session.DB("appDb").C("Topics")
	topic := struct{
		Name string 
		Questions []interface{}
	}{
		Name: "Algebra",
		Questions: contentArray,
	}

	err = c.Insert(topic)
	if err != nil {
		log.Fatal(err)
	}
}
