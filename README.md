### Advanced Databases Course MIM UW Task 5

Program made during the Advanced Databases Course.
The task was to create a graphql resolver in a "reasonable" language (= not C).
I achieved it using the gqlgen library and pgx library to communicate with postgres db.
It exposes graphql api under `/gql` path and visual frontend at `/`.

### Before launching

- set the `DATABASE_URL` env variable with access credentials to the database
- set up the db (using the create.sql file)
- set the `PORT` env variable (on which service should listen)

### Exposed apis (by suffixes)

- user - add users
- ad - add ads
- views - add information about ad view

Each of the programs may be built using the command `go build server.go` and run with
`go run server.go`. Attached dockerfiles builds appropriate images.
TODO in the future: add docker-compose

Graphql specs may be found in schema.graphqls in each app directory

### Important information

Application is not maintained and may not run as is (however all functionalities are provided).
Feel free to contact is sth is not described properly.