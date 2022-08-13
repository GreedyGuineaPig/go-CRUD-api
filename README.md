# go-CRUD-api

## Overview
This is a backend API server with Golang. It has no RDBMS. Also, I developed this according to Git-flow. Made Develop branch from Master, branches for each features. Made pull requests to marge changes to Develop. 

## Functions
- localhost:8080/movies GET-> gives all movies
- localhost:8080/movies/{id} GET-> gives the record of the movie at {id}
- localhost:8080/movies/{id} DELETE-> delete the record of the movie at {id}
- localhost:8080/movies/{id} PUT-> changes the record of the movie at {id}
- localhost:8080/movies POST-> create a new record of the movie

## Technologies
Golang, Git-flow, httptest, GorillaMux, json, remote-container(VScode)
