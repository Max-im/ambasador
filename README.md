# Ambassador marketplace

| A Marketplace service example for selling products 

## Stack
#docker #docker-compose #go #fiber #gorm #postgres #adminer #redis


## Using
up and running:

`sudo docker-compose up --build`


seed db
 - `sudo docker-compose exec app sh`
 - `go run src/commands/populateUsers.go`
 - `go run src/commands/populateProducts.go`
 - `go run src/commands/populateOrders.go`