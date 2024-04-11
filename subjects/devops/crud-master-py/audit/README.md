#### Functional

##### Before we start, let's check we have the necessary tools to perform the tests. You should make sure VirtualBox (or an equivalent software), Vagrant and Postman are installed, if not you should install them.

#### Documentation

##### Take a look at the `README.md` file provided by the student.

###### Does the file include enough context and details to understand and run the project?

#### VM Configuration

##### In the directory of the project run `vagrant up` and then run `vagrant status`.

###### Can you confirm that the three VMs (`gateway-vm`, `inventory-vm` and `billing-vm`) are up and running?

##### Locate the `.env` file in the root of the project, run `cat .env`:

###### Does the output contain all the necessary credentials for the microservices to run properly?

###### Is the source code free from any credential that could have been added to the `.env` file?

###### Is the student able to explain the commands included in `/scripts` directory and why they are used?

#### Inventory API Endpoints

##### Open Postman and make a `POST` request to `http://[GATEWAY_IP]:[GATEWAY_PORT]/api/movies/` address with the following body as `Content-Type: application/json`:

```json
{
  "title": "A new movie",
  "description": "Very short description"
}
```

###### Can you confirm the response was the success code `201`?

##### In Postman make a `GET` request to `http://[GATEWAY_IP]:[GATEWAY_PORT]/api/movies/` address.

###### Can you confirm the response was success code `200` and the body of the response is in `json` with the information of the last added movie?

##### Ask to locate the Postman configuration file in the files committed by the student and import this file in Postman.

###### Can you confirm the imported endpoints includes all methods supported by both APIs and that all of those methods are returning the expected response? (use the subject as a reference)

#### PostgreSQL database for Inventory

##### Run `vagrant ssh inventory-vm` to enter into the VM, then run `sudo -i -u postgres`, then `psql` and once in the database enter `\l`.

###### Can you confirm the `movies` database is listed?

##### Still in `psql` run `\c movies` and then `TABLE movies;`.

###### Can you confirm the entries are presents and reflect the calls you made when checking the endpoints for this API?

#### Billing API Endpoints

##### Open Postman and make a `POST` request to `http://[GATEWAY_IP]:[GATEWAY_PORT]/api/billing/` address with the following body as `Content-Type: application/json`:

```json
{
  "user_id": "20",
  "number_of_items": "99",
  "total_amount": "250"
}
```

###### Can you confirm the response was success code `200`?

##### Run `vagrant ssh billing-vm` to interact with the proper VM. Run `sudo pm2 stop billing_app` and then `sudo pm2 list`.

###### Can you confirm the `billing_app` API was correctly stopped?

##### Open Postman and make a `POST` request to `http://[GATEWAY_IP]:[GATEWAY_PORT]/api/billing/` address with the following body as `Content-Type: application/json`:

```json
{
  "user_id": "22",
  "number_of_items": "10",
  "total_amount": "50"
}
```

###### Can you confirm the response was success code `200` even if the `billing_app` is not working?

#### PostgreSQL database for Billing

##### Run `vagrant ssh billing-vm` to enter into the VM, then run `sudo -i -u postgres`, then `psql` and once in the database enter `\l`.

###### Can you confirm the `orders` database is listed?

##### Still in `psql` run `\c orders` and then `TABLE orders;`.

###### Can you confirm the order with `user_id = 20` is listed properly?

###### Can you confirm the order with `user_id = 22` is NOT listed?

#### Check resilience of messaging queue

##### Run `sudo pm2 start billing_app` to start again the Billing API. At this point enter again in the database following the same instructions of the previous section.

###### Can you confirm the order with `user_id = 22` is now listed properly?
