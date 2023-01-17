<h1 align="center">Welcome to Coplay! - A Trivia Game built in Go :video_game:👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0-blue.svg?cacheSeconds=2592000" />
</p>
 Welcome to my first Golang application!
This is a fullstack app built in using Go Fiber, Docker and PostgreSQL!
Coplay was designed to be an icebreaker to help employees feel comfortable during company meetings. In this first version, users will be able to visualize a list of questions/answers and add a new one to the list!    

## Tech Stack

Project was created using:
* Golang 1.19.3
* Go Fiber v2.40.1 
* PostgreSQL 15.1
* Docker version: 
    Client: 20.10.21
    Server: 20.10.21
    
## Usage
In your IDE terminal run:
```sh
docker compose up
```
## ShowCase
![coplay](https://user-images.githubusercontent.com/40796998/207687448-c02886d3-2519-498b-b9aa-737a4107a6bb.gif)


## Golang : Concepts applied

  - GORM:  It's a Object-relational mapping (ORM)library for Golang. It was used to setup up our database:
     * mapping a database struct
     * connecting to the DB
     * initial migration to our db
     
 - Fiber: It's an Express-inspired Go web framework built of Fasthttp, the fastest HTTP engine for Go. It was used in our project to handle operations such as routing/endpoints and middleware. 
 
 - Fiber Template HTML: It's a template package that provides wrappers for multiple template engines. For our project we've used the engine html.

 - Go Air: It's a command-line utility that provides live reloading for Go applications.


## Acknowledgments (Resources used along the way)
    
   - Docker
     * TechWorld with Nana (Docker Tutorial gor Beginners) https://www.youtube.com/watch?v=3c-iBn73dDE&list=PLy7NrYWoggjxtN4YbSMYFFdpaxb-fR4zC
     * https://nickjanetakis.com/blog/differences-between-a-dockerfile-docker-image-and-docker-container
   - Fullstack: 
     * https://divrhino.com/articles/full-stack-go-fiber-with-docker-postgres/
     * https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54
   - Golang GORM (CRUD Handling)
      https://dev.to/yanoandri/database-handling-with-golang-gorm-crud-handling-4c66
      https://jfrog.com/blog/top-go-modules-golang-web-apis-with-gorm/ 
   - How to build a RESTful API with Docker, PostgreSQL, and go-chi
      https://blog.logrocket.com/how-to-build-a-restful-api-with-docker-postgresql-and-go-chi/   
   - Go Fiber
     * Building an Express-style API in Go with Fiber https://blog.logrocket.com/express-style-api-go-fiber/ 
     * Go Fiber by Examples: How can the Fiber Web Framework be useful? https://dev.to/koddr/go-fiber-by-examples-how-can-the-fiber-web-framework-be-useful-487a
     * https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k
   - Go Templates
     * https://gowebexamples.com/templates/
     * https://tutorialedge.net/golang/basic-rest-api-go-fiber/
     * https://bugz.pythonanywhere.com/gofiber/http-responses
     * https://www.practical-go-lessons.com/chap-32-templates
   - Partials Templates
     https://gohugo.io/templates/partials/#:~:text=Partials%20are%20smaller%2C%20context%2Daware,to%20keep%20your%20templating%20DRY.
   - Go Air (Live-reload)
     * https://betterprogramming.pub/a-good-way-to-do-live-reload-for-go-b3707eb47336
     * https://techinscribed.com/5-ways-to-live-reloading-go-applications/#:~:text=Air%20is%20a%20command%2Dline,github.com%2Fcosmtrek%2Fair
     * https://iaziz786.com/blog/golang-auto-reload-apps/
     * https://thedevelopercafe.com/articles/live-reload-in-go-with-air-4eff64b7a642
  - How I build web frontends in Go
    * https://philipptanlak.com/web-frontends-in-go/
     

    
