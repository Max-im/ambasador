# Ambassador marketplace

> A Marketplace service example for selling products 

## Stack
#docker #docker-compose #go #fiber #gorm #postgres #adminer #redis #stripe #react #typescript #mui #next


## Pre Requirenments
- make sure you have golang installed on your mashine `go version`
- run mailHog to be able to send emails `~/go/bin/MailHog`

## Using
up and running:

`sudo docker-compose up --build`


seed db
 - `sudo docker-compose exec app sh`
 - `go run src/commands/populateUsers.go`
 - `go run src/commands/populateProducts.go`
 - `go run src/commands/populateOrders.go`
 - `go run src/commands/updateRanking.go`


 