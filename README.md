# Introduction
Example of basic structure of api rest on golang with dummy DataBase.

# Requirements
1) Install golang >1.13.1
2) Clone repo on your $GOPATH
3) go run main.go

Then, server is expose on localhost:8080/

# Methods Allowed
Methods 
## GET
-   Get all DataBase -> /responses
-   Get one row of DataBase -> /response/{id}

## POST
- Insert new row -> /response

## PATCH
-   Update one row -> /response/{id}

## DELETE
-   Delete one row -> /response/{id}
-   Delete all -> -> /responses