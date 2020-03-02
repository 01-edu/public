## forum

### Objectives

This project consists in creating a web forum that allows communication between users through the creation of posts and comments. It also allows non-registered users to only see posts and comments.

- Your forum should work based on services. Using services to create a project means that instead of having all of your project in one location, you actually divide it in smaller "projects".
- For example if you want to create a gym application you can split it in: login/register, workouts, stats (weight, muscle mass, etc.), and so on.

You can learn more about services [here](https://medium.com/myntra-engineering/my-journey-with-golang-web-services-4d922a8c9897).

#### Docker

For the Forum project you must use Docker, you can see all about docker basics on the [ascii-art-web-dockerize](https://public.01-edu.org/subjects/ascii-art-web/ascii-art-web-dockerize.en) subject.

You must build a network using containers, in docker you can create your own network of containers through network drivers. Docker provides two network drivers, being `bridge` and `overlay`.

Your objective is to use the [`bridge` network](https://docs.docker.com/engine/tutorials/networkingcontainers/) to create your own network for the services created for the forum project. Docker connected to this network by default, so docker creates a subnet and a gateway for the `bridge` network, the command `docker run` automatically adds containers to it. To inspect the network for more details you can just run `docker network inspect <network>`.

- You can see the default network :

```console
student$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
b64130fc5aae        bridge              bridge              local
033f3a191908        host                host                local
159cdc8d8083        none                null                local
```

- Each service will have to run in a container that will be associated to a docker host, a docker host must contain at least two containers/services, the containers must be connected to a bridge.

- You must be carful, bridges are a private network, so it is restricted to a single Docker host, for instance containers can communicate within networks but not across networks.
  - You may have to use [Port Mapping](https://docker-k8s-lab.readthedocs.io/en/latest/docker/port-mapping.html).
  - Or attaching containers to multiple networks, that way connecting with all of the containers on all networks creating a "hub".

#### Authentication

In this segment you will have to create a page that allows you to `register` new users for the forum, by inputting their credentials. You also have to create a `login session` to access te forum and be able to add posts or comments.

Instructions for user register:

- Must ask for email
- Must ask for password
- When the user is taken return an error response.
- The password must be encrypted

After the registrations the new users must be able to login into the forum. For that they will fill their credentials in the login page.

This page must be able to check if the email provided is present in our database and if all credentials are correct. It will check if the password is the same with the one provided and if the password is not the same return an error response.

If there is an error in the login it must be handled properly:

- If there is an error when casting, return an error response.
- If the email is not present or incorrect return an error response.
- If the password is not present or incorrect return an error response.

#### SQLite

In order to store the data in your forum (like users, posts, comments, etc.) you will use the database library SQLite.

SQLite is a popular choice as embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

- To interact with the database you should use http requests between the server and the database.
- You must use at least one SELECT, one CREATE and one INSERT query.

To know more about SQLite you can check the [SQLite page](https://www.sqlite.org/index.html).

##### SQLite Usage

- You can run queries in your database with the `sqlite3` command.
- You cannot use the command as part of your project. `sqlite3` command will be used as a way to check and test your database.
- Below we have an example on how to use the `sqlite3` command, as well as some query examples:

```console
student$ sqlite3 database.db
SQLite version 3.29.0 2019-07-10 17:32:03
Enter ".help" for usage hints.
sqlite> CREATE TABLE people (id INTEGER PRIMARY KEY, first_name TEXT, age INTEGER, gender TEXT);
sqlite> INSERT INTO people (first_name, age, gender) VALUES ("John", 23, "m");
sqlite> INSERT INTO people (first_name, age, gender) VALUES ("Jade", 36, "f");
sqlite> INSERT INTO people (first_name, age, gender) VALUES ("Kim", 49, "f");
sqlite> SELECT * FROM people;
1|John|23|m
2|Jade|36|f
3|Kim|49|f
sqlite> DELETE FROM people WHERE gender="m";
sqlite> SELECT * FROM people;
2|Jade|36|f
3|Kim|49|f
sqlite> ^C^C^Cstudent$

```

This project will help you learn about:

- Client utilities.
- The basics of web :
  - Server
  - HTML
  - HTTP
  - Services
- Learning what is [docker](https://docs.docker.com).
- Docker network bridge
- Using and [setting up Docker](https://docs.docker.com/get-started/) :
  - Services and dependencies.
  - Containerizing an application.
  - Compatibility/Dependency.
  - Creating images.
  - Docker network.
- SQLite language.
  - Manipulation of databases.
- How to manage dependencies in Go.

### Instructions

- You must use **SQLite**.
- You must use **SQLite** queries.
- You must use HTML files.
- You must handle website errors, HTTP status.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommend that the code should present a **test file**.

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.
- github.com/mattn/go-sqlite3
