
# Introduction

Example of basic structure of api rest on golang with dummy DataBase.

  

# Requirements

1) Install golang >1.13.1

2) Clone repo on your $GOPATH

3) go run main.go

  

Then, server is expose on localhost:8080/

  

# Methods Allowed

  

## GET

Get all DataBase
#### endpoint
```
localhost:8080/responses
```
#### body
```json
[{"ID": "x", "Response": "xxxxx xxxx"},{ .... }]
```
Get one row of DataBase
#### endpoint
```
localhost:8080/response/{id}
```
#### body
```json
{"ID": "{id}", "Response": "xxxxx xxxx"}
```
  

## POST
Insert new row

#### endpoint
```
 localhost:8080/response
 ```
#### body
```json
{"ID": "x", "Response": "xxxxx xxxx"}
```
  

## PATCH

Update one row
#### endpoint
```
 localhost:8080/response/{id}
 ```
#### body
```json
{"Response": "xxxxx xxxx"}
```

## DELETE

Delete one row
#### endpoint
```
 localhost:8080/response/{id}
 ```
#### body
```json
{"ID": "{id}", "Response": "Removed"}
```

Delete all
#### endpoint
```
 localhost:8080/responses
 ```
#### body
```json
{"ID": "", "Response": "All removed"}
```