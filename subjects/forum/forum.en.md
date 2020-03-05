## forum

### Objectives

This project consists in creating a web forum that allows :

- communication between users and the community through the creation of posts/comments.
- associate posts to categories.
- non-registered users to only see posts/comments.
- likes and dislikes in posts and comments, you must take in consideration that the number of likes and dislikes must be accounted.
- searching, inside your forum, for a specific category or post
  - The search bar must have typing suggestions as you write.

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
sqlite> CREATE TABLE car (id INTEGER PRIMARY KEY, brand TEXT, year INTEGER);
sqlite> INSERT INTO car (brand, year) VALUES ("Mercedes", 2010);
sqlite> INSERT INTO car (brand, year) VALUES ("Volvo", 2018);
sqlite> INSERT INTO car (brand, year) VALUES ("Nissan", 1999);
sqlite> SELECT * FROM car;
1|Mercedes|2010
2|Volvo|2018
3|Nissan|1999
sqlite> DELETE FROM car WHERE year>2000;
sqlite> SELECT * FROM car;
3|Nissan|1999
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

This project will help you learn about:

- Client utilities.
- The basics of web :
  - HTML
  - HTTP
  - Sessions and cookies
- Using and [setting up Docker](https://docs.docker.com/get-started/) :
  - Services and dependencies.
  - Containerizing an application.
  - Compatibility/Dependency.
  - Creating images.
- SQLite language.
  - Manipulation of databases.
- The basics of encryption.

### Instructions

- You must use **SQLite**.
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
