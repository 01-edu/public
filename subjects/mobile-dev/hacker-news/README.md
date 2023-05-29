## Hacker news

Hacker News is a social news website that mainly focuses on computer science and entrepreneurship. The objective of this task is to create a client with all the main features of the [HackerNewsAPI](https://github.com/HackerNews/API).

### Instructions

Your app should have the following functionalities:

- Login button. Once you tap it a log in page must be displayed.
- Show a list of all fetched `posts` on the main screen.
- Each `post` should have an `upvote` option for voting.
- Each `post` should display the total number of votes and comments they have.
- Display the `username` of the `post` creator.
- Display the submission time of the `post`.
- Open the link in a [WebView](https://codelabs.developers.google.com/codelabs/flutter-webview#0) when a `post` is tapped.
  - The user must be able to return to the main page once in the `WebView`.

On the website, users can `Post`, `Comment`, or `Reply` **only when they have an account**.

- You have to register using the [website](https://news.ycombinator.com/). In other words, you must be able to login with your actual login and password from the Hacker News website.
- Login using your app. Observe how the login is done on the [website](https://news.ycombinator.com/) and emulate it.

When users are logged in, they should be able to:

- `Create` new posts which will have:
  - `title`, `URL`, and a `description`.
- `Delete` their own posts.
- `Upvote` or hide/remove their votes.
- `Log out`.

> Note: If the user is not logged in but tries to access any of the above-mentioned actions, the user should be sent to the login page of the app.
