#!/bin/sh
COUCHDB="$1"
curl -X PUT http://admin:password@$COUCHDB:5984/_users/org.couchdb.user:jenkins-info \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"name": "jenkins-info", "password": "FIXMEPLEASE", "roles": [], "type": "user"}'

