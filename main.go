package main

import (
	"context"
	"fmt"

	"github.com/colinnewell/jenkins-queue-health/analysis"
	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
)

func main() {
	client, err := kivik.New("couch", "http://jenkins-info:FIXMEPLEASE@192.168.128.2:5984/")
	if err != nil {
		panic(err)
	}

	db := client.DB(context.TODO(), "jenkins-builds")

	var b analysis.AnalysedBuild
	b.ID = "3"

	rev, err := db.Put(context.TODO(), b.ID, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Build inserted with revision %s\n", rev)
}
