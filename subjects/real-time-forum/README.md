## real-time-forum

Remember the forum you did a while ago? Well it's time to make one even better also using JS, with chat rooms, private messages, real time actions, live sharing video and live screen sharing too. Well, maybe not the last two. To get things straight there is a list below of what you will have to do.

### Objectives

On this project you will have to focus on a few points:

- Registration and Login
- Creation of posts
  - Commenting posts
- Private Messages

As you already did the first forum you can use part of the code, but not all of it. Your new forum will have five different parts:

- **SQLite**, in which you will store data, just like in the [previous forum](https://public.01-edu.org/subjects/forum/#communication)
- **Golang**, in which you will handle data and Websockets (Backend)
- **Javascript**, in which you will handle all the Frontend events
- **HTML**, in which you will organize the elements of the page
- **CSS**, in which you will stylize the elements of the page

You will have only one HTML file, so every change of page you want to do, should be handled in the Javascript. This can be called having a [single page application](https://en.wikipedia.org/wiki/Single-page_application).

#### Registration and Login

To be able to use the new and upgraded forum users will have to register and login. This is premium stuff. The registration and login process should take in consideration the following features:

- Users must be able to fill a register form to register into the forum. They will have to provide at least:
  - Nickname
  - Age
  - Gender
  - First Name
  - Last Name
  - E-mail
  - Password
- The user must be able to connect using either the nickname or the e-mail combined with the password.
- The user must be able to log out from any page on the forum.
- The user must be able to consult other people profiles and see everything except the e-mail and the password (obviously).

#### Posts and Comments

This part is pretty similar to the first forum. Users must be able to:

- Create posts
  - Posts will have categories as in the first forum
- Create comments on the posts
- See posts in a feed display
  - See comments only if they click on a post

But there is a twist. If some user creates a post, that post has to be created to all users without refreshing the page and the same goes for the comments. You can achieve this with [WebSockets](https://en.wikipedia.org/wiki/WebSocket).

#### Private Messages

Users will be able to send private messages to each other, so you will need :

- A section of the site to show who is online and able to talk to.
  - The user must be able to send private messages to the users who are online
- A section of the site that shows the people that the user already texted or that texted the user

Both of these sections must be visible at all times, in every pages.

As it is expected, the messages should work on real time, in other words, if a user sends a message, the other user should receive the notification of the new message without refreshing the page. Again this is possible through the use of WebSockets.

This project will help you learn about:

- The basics of web :
  - HTML
  - HTTP
  - Sessions and cookies
  - CSS
  - Backend and Frontend
  - DOM
- [Go routines](https://golangbot.com/goroutines/)
- [Go channels](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)
- [WebSockets](https://en.wikipedia.org/wiki/WebSocket)
- SQL language
  - Manipulation of databases
