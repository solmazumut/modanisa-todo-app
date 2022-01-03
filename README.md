# modanisa-todo-app
-------

This is an Modanisa Assignment - Web Based ToDo List Application. This project is an example of simple REST API implementation with clean architecture written in Go with complete Dependency Injection along with Mocking example, following SOLID principles and a simple to-do list app built with Vue.js

Get Started:

 - [Install](https://github.com/solmazumut/modanisa-todo-app/#install)
 - [Introduction](https://github.com/solmazumut/modanisa-todo-app/#introduction)
 - [Backend Folder Structure](https://github.com/solmazumut/modanisa-todo-app/#folder-structure)
 - [Testing](https://github.com/solmazumut/modanisa-todo-app/#testing)


----------

[Install](https://github.com/solmazumut/modanisa-todo-app/#install)
-------

Start a couchbase instance locally

    docker run -d --name casedb -p 8091-8094:8091-8094 -p 11210:11210 couchbase
    
- Visit the admin client from http://localhost:8091/
- Follow setup a new cluster. You can give it any name.
- While configuring disk, memory and services uncheck 
search, analytics and eventing services and finish with other 
default configurations. You can check what these services are good 
for from Couchbase Documentation.
- Go to Buckets tab from the left navigation and click "ADD BUCKET" 
on top right of the page.
- Give it "Todo" as name
- Create index


Run server

    go run .

Run the app

    npm start
    
    
----------

[Introduction](https://github.com/solmazumut/modanisa-todo-app/#introduction)
-------

While developing the backend code, Dependency injection, one of the SOLID principles, was implemented. The goal here is to minimize coupling and maximize cohesion. With Abstraction, code fragments can be quickly replaced without any concerns.

- Independent of frameworks
- All code can be easily mocked and tested thanks to dependency injections
- Independent of the database. Couchabase is used in the project, but thanks to the decoupling in the project, databases such as MongoDB and PosgreSQL can be easily switched. Because database codes are not dependent on business logic.
- Every implementation is made using interfaces, so by creating mock objects, unit tests can be made and dependency can be injected.

----------

[Backend Folder Structure](https://github.com/solmazumut/modanisa-todo-app/#folder-structure)
-------
### controllers

They handle requests from the router

### infrasctructures

There is the necessary code for the server to connect to the database. Like Couchbase, PostgreSQL configuration files

### interfaces

All systems are implemented using interfaces to bridge different domains.

### models

They are equivalents of data structures in code

### repositories

This is where data access layers are implemented. It ensures that the queries work properly regardless of the database.

### services

where business logic is applied

----------

[Testing](https://github.com/solmazumut/modanisa-todo-app/#testing)
-------

Backend testing (unit and contract testing)

    go test ./...
    
Frontend testing (unit - UI - Consumer Acceptance Test)

    npm test
