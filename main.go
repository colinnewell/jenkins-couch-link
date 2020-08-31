package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/colinnewell/jenkins-queue-health/analysis"
	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
)

var password string
var db string
var host string
var user string
var port int

func main() {
	flag.StringVar(&user, "user", "jenkins-info", "Couchdb username")
	flag.StringVar(&password, "password", "FIXMEPLEASE", "Couchdb password")
	flag.StringVar(&host, "host", "192.168.128.2", "Couchdb password")
	flag.StringVar(&db, "db", "jenkins-builds", "Couchdb db name")
	flag.IntVar(&port, "port", 5984, "Couchdb port")
	flag.Parse()

	client, err := kivik.New(
		"couch",
		fmt.Sprintf("http://%s:%s@%s:%d/", user, password, host, port),
	)
	if err != nil {
		panic(err)
	}

	db := client.DB(context.TODO(), db)

	err = processFiles(context.TODO(), os.Args[1:], db)
	if err != nil {
		log.Fatal(err)
	}
}

func processFiles(ctx context.Context, files []string, db *kivik.DB) error {
	var builds []analysis.AnalysedBuild
	if len(files) == 0 {
		dat, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("Failed to read from stdin - %v", err)
		}
		builds, err = readBuild(dat)
		if err != nil {
			return fmt.Errorf("Failed to process - %v", err)
		}
	} else {
		for _, f := range files {
			// FIXME: support - as a filename for stdin
			dat, err := ioutil.ReadFile(f)
			if err != nil {
				return fmt.Errorf("Failed to read %s - %v", f, err)
			}
			b, err := readBuild(dat)
			if err != nil {
				return fmt.Errorf("Failed to process %s - %v", f, err)
			}
			builds = append(builds, b...)
		}
	}
	for _, b := range builds {
		_, err := db.Put(ctx, b.ID, b)
		if err != nil {
			log.Printf("Error storing build %s: %v\n", b.ID, err)
		}
	}
	return nil
}

func readBuild(fileContents []byte) ([]analysis.AnalysedBuild, error) {
	var analysed []analysis.AnalysedBuild
	err := json.Unmarshal(fileContents, &analysed)
	return analysed, err
}
