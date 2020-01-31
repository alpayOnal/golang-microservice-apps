# Sample Golang  Microservices with Kafka and Mongodb

Sample Golang Microservices

-------------------------
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)

-------------------------
## Requirements

- Docker

-------------------------
## Installation

    cd micro_app
    docker-compose up -d 

    cd micro_app_k2m 
    docker-compose up -d

-------------------------
## Usage

### API Resources

  - [POST /items](#post-items)
  - [GET /items](#get-items)
  - [GET /items/[id]](#get-item)


### POST /items

Example: http://localhost:8000/items

Request body:
    
    {
        "title": "item title 111",
        "description":"item description 111",
        "company":"item company 3",
        "price":1240,
        "currency":"EUR"
    }



### GET /items

Example: http://localhost:8000/items

Response body:
    
    [
        {
            "_id": "5e329bd8b01422d7e0936ffa",
            "title": "item title ",
            "description": "item description ",
            "company": "item company",
            "price": 1240,
            "currency": "EUR",
            "createdAt": "2020-01-30T09:03:20.707Z"
        }
    ]
    

### GET /items/[id]

Example: http://localhost:8000/items/5e329bd8b01422d7e0936ffa

Response body:
    
    {
        "_id": "5e329bd8b01422d7e0936ffa",
        "title": "item title ",
        "description": "item description",
        "company": "item company",
        "price": 1240,
        "currency": "EUR",
        "createdAt": "2020-01-30T09:03:20.707Z"
    }
