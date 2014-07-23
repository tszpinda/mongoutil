//set of simple mongodb helpers
package mgou

import (
	"log"
	"os"
	"strings"
)

//GetMongoConfg looks up mongo url in various environments:
//- heroku
//- wercker
//- system variable (MONGODB_DB)
func GetMongoConfg(defaultDb string) (url, databaseName string) {
	url = os.Getenv("MONGOHQ_URL")
	if url == "" {
		url = os.Getenv("WERCKER_MONGODB_HOST")
	}
	log.Printf("url set by env:%v ", url)
	if url != "" {
		//only if url already has db
		var urlArr = strings.Split(url, "/")
		if len(urlArr) > 2 {
			databaseName = urlArr[3]
		}
	}
	if os.Getenv("MONGODB_DB") != "" {
		databaseName = os.Getenv("MONGODB_DB")
	}
	if url == "" {
		url = "mongodb://localhost"
	}
	if databaseName == "" {
		databaseName = defaultDb
	}
	log.Println("Connecting to:", url, "db: ", databaseName)
	return url, databaseName
}
