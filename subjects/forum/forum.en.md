## forum

### Objectives

This project consists in creating a web forum that allows :

- communication between users and the community through the creation of posts/comments.
- non-registered users to only see posts/comments.
- registered users to like or dislike posts/comments.
- associate posts to categories.

Your forum should work based on services. Using services to create a project means that instead of having a monolith architecture, you actually have various components distributed across a cluster of instances. In other words, dividing the project in smaller "projects", this way it becomes easier to understand and your project will become more scalable.

- For example for a taxi like application you can divide it in : passenger management, billing, notifications, payments, trip management and driver management.

You can learn more about this [here](https://www.nginx.com/blog/introduction-to-microservices/).

#### SQLite

In order to store the data in your forum (like users, posts, comments, etc.) you will use the database library SQLite.

SQLite is a popular choice as embedded database software for local/client storage in application software such as web browsers. It enables you to create a database as well as controlling it by using queries.

To structure your database and to achieve better performance we highly advise you to take a look at the [entity relationship diagram](https://www.smartdraw.com/entity-relationship-diagram/) and build one based on your own database.

- You must use at least one SELECT, one CREATE, one INSERT and one DELETE query.

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

#### Authentication

In this segment the client must be able to `register` as a new user for the forum, by inputting their credentials. You also have to create a `login session` to access the forum and be able to add posts and/or comments.

You should use cookies to allow each user to have only one open session. Each of this sessions must contain an expiration date. It's up to you to decide what time the cookie stays "alive".

Instructions for user registration:

- Must ask for email
- When the email is taken return an error response.
- Must ask for username
- Must ask for password
- The password must be encrypted

The forum must be able to check if the email provided is present in the database and if all credentials are correct. It will check if the password is the same with the one provided and if the password is not the same return an error response.

#### Docker

For the forum project you must use Docker. You can see all about docker basics on the [ascii-art-web-dockerize](https://public.01-edu.org/subjects/ascii-art-web/ascii-art-web-dockerize.en) subject.

You must build a network using containers. In docker you can create your own network of containers through network drivers. Docker provides two network drivers : `bridge` and `overlay`.

- You can see the default network :

```console
student$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
b64130fc5aae        bridge              bridge              local
033f3a191908        host                host                local
159cdc8d8083        none                null                local
```

Your objective is to use the [`bridge` network](https://docs.docker.com/engine/tutorials/networkingcontainers/) to create your own network for the services. Docker connects to this network by default, creating a subnet and a gateway for the `bridge` network. The command `docker run` automatically adds containers to it. To inspect the network for more details you can just run `docker network inspect <network>`.

- Each service will have to run in a container that will be associated to a docker host, which must contain at least two containers/services and these must be connected to a bridge.

- You must be careful, bridges are a private network, so it is restricted to a single Docker host, for instance containers can communicate within networks but not across networks.
  - You may have to use [Port Mapping](https://docker-k8s-lab.readthedocs.io/en/latest/docker/port-mapping.html).
  - Or attaching containers to multiple networks connecting with all containers on all networks creating a ["hub"](https://docs.docker.com/engine/tutorials/bridge3.png).

This project will help you learn about:

- Client utilities.
- The basics of web :
  - Server
  - HTML
  - HTTP
  - Services
  - Sessions and cookies
- Docker network bridge
- Using and [setting up Docker](https://docs.docker.com/get-started/) :
  - Services and dependencies.
  - Containerizing an application.
  - Compatibility/Dependency.
  - Creating images.
- SQLite language.
  - Manipulation of databases.
- How to manage dependencies in Go.
- The basics of encryption.

### Instructions

- You must use **SQLite**.
- You must use **SQLite** queries.
- You must use HTML files.
- You must handle website errors, HTTP status.
- You must handle all sort of technical errors.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommend that the code should present a **test file**.

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.
- github.com/mattn/go-sqlite3
- golang.org/x/crypto/bcrypt
- github.com/satori/go.uuid
