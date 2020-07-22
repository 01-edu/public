## Social Network

### Objectives

You will have to create a social network that presents both backend and frontend.


### Instructions

### Frontend

Frontend development is the art of creating sites and web applications that render on the "client-side". It includes everything that users experience directly: text colors and styles, images, graphs and tables, buttons, colors, and navigation menu. It focus on making request to the backend in order to get specific chunks of data to be used or send data to be stored on the backend.

HTML, CSS, and Javascript are the languages used for frontend development. Responsiveness and performance are two main objectives of the frontend. It can be used frontend frameworks to simplify a developers work.

#### Framework

You will have to use a JS framework. It is up to you to choose which one you are going to use.

Frameworks will help you to organize and implement the features you want on your project, so that you can get more work done in a easier and faster way.

Some of the most known JS frameworks around are:

- [React](https://reactjs.org/)
- [Angular](https://angular.io/)
- [Vue.js](https://vuejs.org/)
- [Ember.js](https://emberjs.com)
- [Meteor](https://www.meteor.com/)

Caution: Note that JS frameworks are different from JS libraries. JS libraries contain code snippets that are used to perform common JavaScript functions, while frameworks will help you by laying out the groundwork/building the bases for your JS project.

----

### Backend

The back-end is all of the technology required to process the incoming request and generate and send the response to the client.
This is typically divided into three major parts:

- **Server**, this is the computer that receives request. Though there are machines made and optimized for this particular purpose, you will use your own computer.

- **App**, this is the application running on the server that listens for requests, retrieves information from the database and sends a response. The server runs an app that contains all the logic about how to respond to various requests based on the HTTP or other types of protocols. Some of these handlers functions will be middleware. Middlewares is any code that executes between the server receiving a request and sending a response. These middleware functions might modify the request object, query the database, or otherwise process the incoming request. Middleware functions typically end by passing control to the next middleware function, rather than by sending a response.

- **Database**, are used as you already now, to organize and persist data. Many requests sent to the server might require a database query. A client might request information that is stored in the database, or a client might submit data with their request to be added to the database.

#### Sqlite

In order to store the data in your social network, you will use the database library SQLite.
To structure your database and to achieve better performance we highly advise you to take a look at the [entity relationship diagram](https://www.smartdraw.com/entity-relationship-diagram/) and build one based on your own database.

To know more about SQLite you can check the [SQLite page](https://www.sqlite.org/index.html).

#### Migrate

You will have to create migrations for this project so every time the application runs it creates the specific tables to make the project work properly

For this you must focus on a folder structure similar to this one:

```console
student$ tree .
backend
├── pkg
│   ├── db
│   │   ├── migrations
│   │   │   └── sqlite
│   │   │       ├── 000001_create_users_table.down.sql
│   │   │       ├── 000001_create_users_table.up.sql
│   │   │       ├── 000002_create_posts_table.down.sql
│   │   │       └── 000002_create_posts_table.up.sql
│   │   └── sqlite
│   |       └── sqlite.go
|   |
|   └── ...other_pkgs.go
|
└── server.go
```

The folder structure is organized in a way that helps you to **understand** and **use** migrations, where you can apply it using a simple path, for example: `file://backend/pkg/db/migrations/sqlite`. It can be organized as you wish but **do not forget that the application of migrations and the file organization will be tested**.

For migrations you can use [golang-migrate](https://github.com/golang-migrate/migrate) package or other package that better suites you project

All migrations should be stored on a specific folder, like as above. The `sqlite.go` should present the connection to the database, the applying of the migrations and other useful functionalities that you may need to implement.

----

### docker

You must create two images where one will serve the backend and the other will serve the frontend.

Both back and front must communicate, for that is the purpose of having them. For this you will have to create a network to make sure that your containers are isolated, you can see more about `docker network` [here](https://docs.docker.com/network/).
This network must have both back and front end where they will communicate between them.

Only one will be published to a port on the host machine, being the client, so that you can be able to access the port that is exposed.

----

### Authentication

To be able for the users to use the social network they will have to make an account. So you will have to make a registration and login form. To register, every user should provide at least:

- Email
- Password
- First Name
- Last Name
- Date of Birth
- Avatar/Image
- Nickname (Optional)
- Relationship Status (Optional)
- About Me (Optional)

Note that the **Nickname**, **Relationship Status** and **About Me** should be present in the form but the user can skip the filling of those fields.

When you login you should stay logged in until you choose a log out option that should be available at all times. For this you will have to implement [sessions]() and [cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies).

----

### Followers

When navigating the social network you should be able to follow and unfollow other users. Needless to say that to unfollow a user you have to be following him/her.

In order to follow someone the user first needs to send a request to the user he/she wants to follow. Then the other user should be able to accept the request or decline it. If the second user has a public profile (explained in the next topic) this step is skipped and the sending of the request is skipped.

----

### Profile

Every profile should contain :

- User information (every information requested in the register form apart from the **Password**, obviously)
- User activity
  - Every post made by the user
  - Every comment made by the user (and the corresponding post, if the post is from a public profile or followed user)
- Followers and following users (display the users that are following the owner of the profile and who he/she is following)

There are two types of profiles: a public profile and a private profile. A public profile will display the information specified above to every user on the social network, while the private profile will only display that same information to their followers only.

When the user is in their own profile it should be available an option that allows the user to turn its profile public or private.

----

### Posts

After a user is logged in he/she can create posts and comments, on already created posts. While creating a post or a comment, the user can include an image or GIF.

The user must be able to specify the privacy of the post:

- public (all users in the social network will be able to see the post)
- private (only followers of the creator of the post will be able to see the post)
- almost private (only selected users by the creator of the post will be able to see the post)

----

### Groups

A user must be able to create a group. The group should have a title and a description given by the creator and he/she can invite other users to join the group.

The invited users need to accept the invitation to be part of the group. They can also invite other people once they are already part of the group. Another way to enter the group is to request to be in it and only the creator of the group would be allowed to accept or refuse the request.

When in a group, a user can create posts and comment the posts already created. These posts and comments will only be displayed to members of the group.

A user belonging to the group can also create an event, making it available for the other group users. An event should have:

- Title
- Description
- Day/Time
- 2 Options (at least):
  - Going
  - Not going

After creating the event every user can choose one of the options for the event.

----

### Chat

The user should be able to send private messages to users that he/she is following. It should be able for the users to send images, GIFs and emojis to each other.

The user that the message was sent to, will receive the message instantly if he/she is following the user that sent the message or if the user has a public profile.

Groups should have their own chat too, so if a user is a member of the group he/she should be able to send and receive messages to this group chat.

----

### Notifications

A user must be able to see the notifications in every page of the project and a change on the site should be displayed whenever the user receives a new notification. New notifications are different from new private messages and should be displayed in different ways!

A user should be notified if:

- he/she has a private profile and:
  - a different user (that does not follow him/her) tries to send him/her a message
  - some other user sends him/her a following request
- he/she receives a group invitation, so he can refuse or accept the request
- he/she is the creator of a group and another user requests to join the group, so he can refuse or accept the request
- he/she is member of a group and an event is created

Every other notification created by you that isn't on the list is welcomed too.
