## App image

### Objectives

For this optional, `cross platform app image`, you will have to implement a desktop app using [Electron](https://www.electronjs.org/):

This desktop app should have as principle objective the creation of a messenger, just like facebook or discord. It should be able to run in multiple platforms : windows, linux and macOS.

You will have to create:

- A way to see which users are online (able to talk)
- A way to notify the user whenever he/she receives a message
- A real time communication between the users that are chatting
- A section for emojis, where users can send to each other
- An offline possibility here you can see all messages from all users, but can not send messages to them or receive. You must inform the user that he/she is offline/online
- A search engine to search for a messages

We encourage you to add any other additional features that you find relevant.

### Instructions

You must use a method of authentication for the app.

To be able for the users to use the app you must create a login form, the user should provide:

- Email
- Password

If the user does not present a registration, then the app should redirect to the social network website (mandatory project) so that the user can register himself.

When the user logs in or registers he should stay logged in until he/she chooses a logout option that should be available at all times. Even when the user exits the app and opens it again it should continue logged in until the [session](https://www.electronjs.org/docs/api/session) expires (the time limit is up to you to decide) or the user decides to logout.

#### Websocket

The use of websocket must be present in this project:

- You must be able to send messages in real time just like the [mandatory project](https://public.01-edu.org/subjects/social-network/), you should be able to send chat messages using the website as an user, to the desktop app as another user.
- To see the status of a user, whenever a user goes online/offline the status must change automatically, in real time. So when a user goes offline all the followers must be able to see that the user went offline and its not available to talk. The same happens when the user goes online, all followers must see that he/she is online.

---

#### Offline

If the internet connection goes down or the user does not have internet, the app should warn the user that there is no connection by displaying a message to say so. The offline mode allows the user to view all the messages sent, but if the user tries to send a message it should display the same error message as the first one.

Here's a tip [offline/online](https://www.electronjs.org/docs/tutorial/online-offline-events)

The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices/).

---

#### Search

The search should be interactive, in other words, the results should be displaying as you write, not needing a button for you to click.

This project will help you learn about:

- Data manipulation and local storage
- Authentication
- Desktop Applications:
  - [Electron](https://www.electronjs.org/docs)
- Websocket
