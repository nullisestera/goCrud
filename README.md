# goCrud

1) Create a DB in MySQL named `gocrud`
2) Change to your DB credentials in `utils/connection.go` and `seeder/seeder.go`
3) Run `go run main.go`
4) Close connection and `cd seeder` and run `go run seeder.go`
5) `cd ..` and run `go run main.go`
6) Check in Browser some endpoint `localhost:4000/api/contacts`
7) For all endpoints, check `routes/routes.go`