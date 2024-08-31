# Entain 

Entain is one of the biggest companies, which hosts casinos like Olybet, Optibet etc. They are looking for a software engineer to join their team. 

## Task

### Tools and technologies used:

1. Go
2. PostgreSQL
3. Docker
4. Makefile
5. Postman

All database tables are in `migrations` folder. To run them, use `make` command.

1. `make migrate-up` - to run up migrations
2. `make migrate-down` - to run down migrations
3. `make migrate-force` - to force run migrations if you have some errors like `error: Dirty database version -1. Fix and force version.`

### Tables:

1. `users` - contains users data
2. `transaction` - contains transactions data

### Endpoints to test:

1. `GET /users` - to get all users
2. `GET /users/{user_id}` - to get user by id, check his balance
3. `GET /transactions/{user_id}` - to get all transactions by user id (check if user has any transactions)
4. `POST /process-record/{user_id}` - to process record by user id

Process record request body example:

```
{
    "amount": 10,
    "transaction_id": "64338a05-81e5-426b-b01e-927e447c9e33",
    "state": "win"
}
```

Transaction id is unique, so you can't process the same transaction twice, provide UUID v4 format.
State can be `win` or `lose`.
Amount is a number should be positive but to have a negative balance you should provide a `lose` state.

### Required header for all endpoints:

1. `Source-Type: game` - available values: `game`, `server`, `payment`

Postman collection is in `postman` folder to test endpoints.

## To run the app locally:

1. Create `.env` file in root folder and add all required variables from `.env.example` file
2. To run migrations you should have migrate tool installed. You can install it with `brew install golang-migrate` (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
3. To run any `make` command you should have `make` tool installed. You can install it with `sudo apt install make` command (https://linuxhint.com/install-make-ubuntu/)
4. Run `make migrate-up` command to run migrations and create all tables with test user (user_id: `63e83104-b9a7-4fec-929e-9d08cae3f9b9`)
5. Run `make run` command to run application
6. Take a look at `postman` folder to take collection for testing all endpoints

Test user with id `63e83104-b9a7-4fec-929e-9d08cae3f9b9` will be created automatically when you run migrations.
This user has 50 amount of his balance for testing.

## To run application in docker container:

1. Create `.env` file in root folder and add all required variables from `.env.example` file
2. To run docker container you should have `docker` and `docker-compose` tools installed (Tested on `Docker version 26.1.3, build b72abbb` and `Docker Compose version v2.27.1`)
3. `docker-compose up` - to run application in docker container
4. `docker-compose down` - to stop application in docker container