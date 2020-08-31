# jenkins-couch-link

Program to upload Jenkins build info that has been analysed up to CouchDB

## To build

    go build

## Setup

Ensure there is a DB and user for the program to use.  Creating the DB can be done easily from the /_util page.
Creating a user can be done by using something like the `setup.sh` script doing a curl.

## To use

   ./jenkins-couch-link -user jenkins -password somepassword -host 192.168.128.2 builds.json 

Note that documents can only be inserted, it won't update existing builds in the database.
