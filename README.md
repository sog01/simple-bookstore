# Simple Bookstore with Elasticsearch and Go

A mini project that builds a simple bookstore application using Elasticsearch and Go. 

## Pre-requisite
- [Go](https://go.dev/doc/install)
- [Docker](https://www.docker.com/)
- [Postman](https://www.postman.com/downloads/)

## How to Run this Project
- Start the Elasticsearch with docker compose
```
make upES
```
- Populate Elasticsearch data
```
make populateData
```
- Run this project and open localhost:8080
```
make run
```

## Stop Elasticsearch
```
make downES
```