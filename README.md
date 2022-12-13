<h1 align="center">Welcome to Golang - Trivia Game Questions and Answers👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
</p>

> Welcome to my first Golang application. This is a fullstack app built in using Go Fiber, Docker and Postgres! In this first version, you will be able to visualize all questions/answers and add a new one to the list!   

## Tech Stack

Project was created using:
* Go 1.19.3
* PostgreSQL 15.1
* Docker version: 
    Client: 20.10.21
    Server: 20.10.21
    
## Usage
In your IDE terminal run:
```sh
docker compose up
```
## Golang : Concepts applied

  - GORM:  It's a Object-relational mapping (ORM)library for Golang. It was used to setup up our database:
     * mapping a database struct
     * connecting to the DB
     * initial migration to our db
     
 - Fiber: It's an Express-inspired Go web framework built of Fasthttp, the fastest HTTP engine for Go. It was used in our project to handle operations such as routing/endpoints and middleware. 
 
 - Fiber Template HTML: It's a template package that provides wrappers for multiple template engines. For our project we've used the engine html.

 - Go Air: It's a command-line utility that provides live reloading for Go applications.

 - Partials templates: 

## Acknowledgments (Resources used along the way)
    
   - TechWorld with Nana (Docker Tutorial gor Beginners)
      https://www.youtube.com/watch?v=3c-iBn73dDE&list=PLy7NrYWoggjxtN4YbSMYFFdpaxb-fR4zC
   - Div Rhino: 
     * https://divrhino.com/articles/full-stack-go-fiber-with-docker-postgres/
     * https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54
   - Database Handling with Golang GORM (CRUD Handling)
      https://dev.to/yanoandri/database-handling-with-golang-gorm-crud-handling-4c66
   - How to build a RESTful API with Docker, PostgreSQL, and go-chi
      https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/ 
   - Top Go Modules: Golang Web APIs with GORM
      https://jfrog.com/blog/top-go-modules-golang-web-apis-with-gorm/ 
   - Building an Express-style API in Go with Fiber
     https://blog.logrocket.com/express-style-api-go-fiber/ 
   - Go Fiber by Examples: How can the Fiber Web Framework be useful?
     * https://dev.to/koddr/go-fiber-by-examples-how-can-the-fiber-web-framework-be-useful-487a
     * https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k
   - Go Web Examples - Templates
     https://gowebexamples.com/templates/
   - Http responses in go fiber in 3 steps [ Golang ]
     https://bugz.pythonanywhere.com/gofiber/http-responses
   - Practical Go Lessons - Chapter 32: Templates
     https://www.practical-go-lessons.com/chap-32-templates
   - Partials Templates
     https://gohugo.io/templates/partials/#:~:text=Partials%20are%20smaller%2C%20context%2Daware,to%20keep%20your%20templating%20DRY.
   - Here’s a Good Way to Do Live-reload in Go
     *  https://betterprogramming.pub/a-good-way-to-do-live-reload-for-go-b3707eb47336
     *  https://techinscribed.com/5-ways-to-live-reloading-go-applications/#:~:text=Air%20is%20a%20command%2Dline,github.com%2Fcosmtrek%2Fair
     *  https://iaziz786.com/blog/golang-auto-reload-apps/
     *  https://thedevelopercafe.com/articles/live-reload-in-go-with-air-4eff64b7a642
     
  - How I build web frontends in Go
     https://philipptanlak.com/web-frontends-in-go/
     

    
