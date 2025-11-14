## go-test-practices

**go-test-practices** is a Go-based RESTful API for task management, built with GORM, structured, and used as a personal space for practicing **Golang unit testing**.

The project simulates a small-scale API system that connects the basic parts of a backend service: from database models and data access layers to services, controllers, and routes. 
It’s both a functional app and a learning environment, created to explore how Go testing actually works in real project structures.

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** net/http  
- **ORM:** GORM  
- **Database:** PostgreSQL
- **Architecture:** Layered (Route -> Controller -> Service -> DAO -> Model)
- **Testing:** Go’s built-in `testing` package
