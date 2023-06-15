# Go RESTful API with Gin Web Framework & PostgreSQL
This is an example golang backend application using PostgreSQL database with clean architecture.

## Features
* Go Web Framework ([gin-gonic](https://github.com/gin-gonic/gin))
* Containerize ([docker](https://www.docker.com/))
* Swagger ([swaggo](https://github.com/swaggo/swag))
* CRUD operations
* JWT for authentication
* Mock: [golang/mock](https://github.com/golang/mock)
* PostgreSQL Driver: [GORM](gorm.io/gorm)
* Test Assertions: [stretchr/testify](https://github.com/stretchr/testify)

## Getting Started

```sh
# download the project
git clone https://github.com/aabdullahgungor/personel-resume-api.git

cd personel-resume-api
```

### Build and run image of docker

```bash
docker-compose up  --build  -d
```
### Endpoints

- GET localhost:8000/api/v1/abilities
- GET localhost:8000/api/v1/abilities/:id
- POST localhost:8000/api/v1/abilities
- PUT localhost:8000/api/v1/abilities
- DELETE localhost:8000/api/v1/abilities/:id
- GET localhost:8000/api/v1/experiences
- GET localhost:8000/api/v1/experiences/:id
- POST localhost:8000/api/v1/experiences
- PUT localhost:8000/api/v1/experiences
- DELETE localhost:8000/api/v1/experiences/:id
- GET localhost:8000/api/v1/personals
- GET localhost:8000/api/v1/personals/:id
- POST localhost:8000/api/v1/personals
- PUT localhost:8000/api/v1/personals
- DELETE localhost:8000/api/v1/personals/:id
- GET localhost:8000/api/v1/universities
- GET localhost:8000/api/v1/universities/:id
- POST localhost:8000/api/v1/universities
- PUT localhost:8000/api/v1/universities
- DELETE localhost:8000/api/v1/universities/:id
- 
### Sample API Request and Response

## Open API Doc Preview
http://localhost:8000/api/v1/swagger/index.html

![Swagger](.github/images/Swagger.png)
