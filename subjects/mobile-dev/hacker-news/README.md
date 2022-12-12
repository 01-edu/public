# Hacker news

### Introduction

Create your own hackernews [app](https://news.ycombinator.com/)!

Hacker News is a social news website, which mainly focusses on Computer Science and Entrepreneurship. You will make a client consisting of all the main features of the website.

### Objective

Your objective for this raid is to create UI for [HackerNewsAPI](https://github.com/HackerNews/API).

Your app should have following functionality:

- Show list of all fetched posts on main screen:
  - Login button
  - Each Post has an Upvote option to vote the post
  - Each Post displays Total Votes and Total Comments on them
  - Displays the Username of the Creator
  - Displays the Submission Time
  - Open link in a [Webview](https://codelabs.developers.google.com/codelabs/flutter-webview#0) when post is tapped.

On the website, we can Post, Comment, or reply **only when we have an account**.

- You have to register using the [website](https://news.ycombinator.com/).
- Login using your app. Observe how log in is done on the [website](https://news.ycombinator.com/), and emulate it.

When users logged in, they should be able to:

- Create new posts. They should have:
  - Title, URL and the Description.
- Delete own post.
- Make up-vote or hide their vote.
- Log out.

- Note: if user is not loggen in, but tries to access above mentioned actions, send user to login page of the app.
