# SQL Runner 2000

Welcome to the future of running SQL against teeny tiny databases!

## Requirements

Run `$ ./setup.sh` to install Golang dependencies. The SqlLite3 driver takes a minute to install; it might look like it has hung, but--it probably hasn't.

## How to run SQL commands against the teeny tiny database

`$ go run main.go`

## How to change the commands that will run against the teeny tiny database

Add (or edit) an object declaration in `queries.json`. See the content of `queries.json` for a few examples.

## How to update the database

1. Put all the data you want to be in the database in `data.csv` following the format from the interview's original instructions.
2. Then, run `$ go run createDb.go`.

## Future work

 * Update `main.go` to handle query results that have other types, such as strings.
 * Stick this all in Rocky Linux on Apptainer for mega brownie points
 * Do it agian in Docker for shiggles

## Questions?

Contact Luke:

luke.mizuhashi@gmail.com
(512) 966 - 3777

