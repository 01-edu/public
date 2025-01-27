## Social Network

### Objectives

You will have to create a Facebook-like social network that will contain the following features:

- Followers
- Profile
- Posts
- Groups
- Notification
- Chats

### Instructions

### Frontend

Frontend development is the art of creating sites and web applications that render on the "client-side". It includes everything that users experience directly: text colors and styles, images, graphs and tables, buttons, colors, and navigation menus. It focuses on making requests to the backend in order to get specific chunks of data to be used or send data to be stored on the backend.

HTML, CSS, and Javascript are the languages used for frontend development. Responsiveness and performance are two main objectives of the frontend. Frontend frameworks may be used to simplify a developers work.

#### Framework

You will have to use a JS framework. It is up to you to choose which one you are going to use.

Frameworks will help you to organize and implement the features you want on your project, so that you can get more work done in a easier and faster way.

Some of the most known JS frameworks around are:

- [Next.js](https://nextjs.org/)
- [Vue.js](https://vuejs.org/)
- [Svelte](https://svelte.dev/)
- [Mithril](https://mithril.js.org/)

Caution: Note that JS frameworks are different from JS libraries. JS libraries contain code snippets that are used to perform common JavaScript functions, while frameworks will help you by laying out the groundwork/building the bases for your JS project.

---

### Backend

The backend is all of the technology required to process the incoming requests and generate and send responses to the client.
This is typically divided into three major parts:

- **Server**: this is the computer that receives requests. It acts as the entry point for all incoming requests. While there are specialized machines designed for this purpose, you can use your own computer as a server.

- **App**: is the application running on the server that listens for requests, retrieves information from the database and sends responses. This is where the core logic of your social network resides. It contains the logic for handling various requests based on HTTP or other protocols. Some of these functions are known as middleware, which execute between receiving a request and sending a response.

- **Database**: As you may already know, the database is used to organize and store data. Many requests sent to the server involve database queries. Clients may request information stored in the database or submit data to be added to it.

#### App

The backend may consist, like said above, of an **app** containing all the backend logic. This logic will therefore have several middleware, for example:

- Authentication, since HTTP is a stateless protocol, we can use several ways to overcome and authenticate a client/user. You must use [sessions](https://allaboutcookies.org/cookies/session-cookies-used-for.html) and cookies.
- Images handling, supporting various types of extensions. In this project you have to handle at least JPEG, PNG and GIF types. You will have to store the images, it can be done by storing the file/path in the database and saving the image in a specific file system.
- Websocket, handling the connections in real time, between clients. This will help with the private chats.

For the web server you can take a look at [caddy](https://caddyserver.com/docs/), it can serve your site, services and apps and its written in **Go**. Or you are free to create your own web server.

#### Sqlite

In order to store the data in your social network, you will use the database library SQLite.
To structure your database and to achieve better performance we highly advise you to take a look at the [entity relationship diagram](https://www.smartdraw.com/entity-relationship-diagram/) and build one based on your own database.

To know more about SQLite you can check the [SQLite page](https://www.sqlite.org/index.html).

#### Migrate

You will have to create migrations for this project so every time the application runs, it creates the specific tables to make the project work properly.

For this, you must focus on a folder structure similar to this one:

```console
student$ tree .
backend
├── pkg
│   ├── db
│   │   ├── migrations
│   │   │   └── sqlite
│   │   │       ├── 000001_create_users_table.down.sql
│   │   │       ├── 000001_create_users_table.up.sql
│   │   │       ├── 000002_create_posts_table.down.sql
│   │   │       └── 000002_create_posts_table.up.sql
│   │   └── sqlite
│   |       └── sqlite.go
|   |
|   └── ...other_pkgs.go
|
└── server.go
```

The folder structure is organized in a way that helps you to **understand** and **use** migrations, where you can apply it using a simple path, for example: `file://backend/pkg/db/migrations/sqlite`. It can be organized as you wish but **do not forget that the application of migrations and the file organization will be tested**.

For migrations you can use [golang-migrate](https://github.com/golang-migrate/migrate) package or other package that better suits your project.

All migrations should be stored on a specific folder, as above. The `sqlite.go` should present the connection to the database, the applying of the migrations and other useful functionalities that you may need to implement.

This migration system can help you manage your time and testing, by filling your database.

---

### docker

**Containerizing the Backend and Frontend**

Your project should consist of two Docker images, one for the backend and another for the frontend. Each of these containers serves a specific purpose and communicates with each other to provide a functional social network application.

**Backend Container:**

Create a Docker image for the backend of your social network application. This container will run the server-side logic of your application, handle requests from clients, and interact with the database.

**Frontend Container:**

Create a Docker image for the frontend of your social network application. This container will serve the client-side code, like HTML, CSS, and JavaScript files, to users browsers. It will also communicate with the backend via `HTTP requests`.

**Tips:**

- Name your frontend Docker image appropriately.
- Make sure that the backend container exposes the necessary ports for communication with the frontend and external clients and the frontend container exposes the appropriate port to serve the frontend content to users' browsers.

---

### Authentication

In order for the users to use the social network they will have to make an account. So you will have to make a registration and login form. To register, every user should provide at least:

- Email
- Password
- First Name
- Last Name
- Date of Birth
- Avatar/Image (Optional)
- Nickname (Optional)
- About Me (Optional)

Note that the **Avatar/Image**, **Nickname** and **About Me** should be present in the form but the user can skip the filling of those fields.

When the user logins, he/she should stay logged in until he/she chooses a logout option that should be available at all times. For this you will have to implement [sessions](https://allaboutcookies.org/cookies/session-cookies-used-for.html) and [cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies).

You can implement your own package for sessions and cookies or you can take a look at some packages to help you.

---

### Followers

When navigating the social network the user should be able to follow and unfollow other users. Needless to say that to unfollow a user you have to be following him/her.

Regarding following someone, the user must initiate this by sending a follow request to the desired user. The recipient user can then choose to "accept" or "decline" the request. However, if the recipient user has a public profile (as explained in the next section), this request-and-accept process is bypassed and the user who sent the request automatically starts following the user with the pubic profile.

---

### Profile

Every profile should contain :

- User information (every information requested in the register form apart from the **Password**, obviously)
- User activity
  - Every post made by the user
- Followers and following users (display the users that are following the owner of the profile and who he/she is following)

There are two types of profiles: a public profile and a private profile. A public profile will display the information specified above to every user on the social network, while the private profile will only display that same information to their followers only.

When the user is in their own profile it should be available an option that allows the user to turn its profile public or private.

---

### Posts

After a user is logged in he/she can create posts and comments on already created posts. While creating a post or a comment, the user can include an image or GIF.

The user must be able to specify the privacy of the post:

- public (all users in the social network will be able to see the post)
- almost private (only followers of the creator of the post will be able to see the post)
- private (only the followers chosen by the creator of the post will be able to see it)

---

### Groups

A user must be able to create a group. The group should have a title and a description given by the creator and he/she can invite other users to join the group.

The invited users need to accept the invitation to be part of the group. They can also invite other people once they are already part of the group. Another way to enter the group is to request to be in it and only the creator of the group would be allowed to accept or refuse the request.

To make a request to enter a group the user must find it first. This will be possible by having a section where you can browse through all groups.

When in a group, a user can create posts and comment the posts already created. These posts and comments will only be displayed to members of the group.

A user belonging to the group can also create an event, making it available for the other group users. An event should have:

- Title
- Description
- Day/Time
- 2 Options (at least):
  - Going
  - Not going

After creating the event every user can choose one of the options for the event.

---

### Chat

Users should be able to send private messages to other users that they are following or being followed, in other words, at least one of the users must be following the other.

When a user sends a message, the recipient will instantly receive it through Websockets if they are following the sender or if the recipient has a public profile.

It should be able for the users to send emojis to each other.

Groups should have a common chat room, so if a user is a member of the group he/she should be able to send and receive messages to this group chat.

---

### Notifications

A user must be able to see the notifications in every page of the project. New notifications are different from new private messages and should be displayed in a different way!

A user should be notified if he/she:

- has a private profile and some other user sends him/her a following request
- receives a group invitation, so he can refuse or accept the request
- is the creator of a group and another user requests to join the group, so he can refuse or accept the request
- is member of a group and an event is created

Every other notification created by you that isn't on the list is welcomed too.

#### Allowed Packages

- The [standard Go](https://golang.org/pkg/) packages are allowed
- [Gorilla](https://pkg.go.dev/github.com/gorilla/websocket) websocket
- [golang-migrate](https://github.com/golang-migrate/migrate/)
- [sql-migration](https://pkg.go.dev/github.com/rubenv/sql-migrate)
- [migration](https://pkg.go.dev/github.com/Boostport/migration)
- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [UUID](https://github.com/gofrs/uuid)

This project will help you learn about:

- Authentication :
  - Sessions and cookies
- Using and [setting up Docker](https://docs.docker.com/get-started/)
  - Containerizing an application
  - Compatibility/Dependency
  - Creating images
- SQL language
  - Manipulation of databases
  - Migrations
- The basics of encryption
- Websocket
