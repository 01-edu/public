## ascii-art-web

### Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web **GUI** (graphical user interface) version of your last project, [ascii-art](../ascii-art).

Your webpage must allow the use of the different banners:

- [shadow](../ascii-art/shadow.txt)
- [standard](../ascii-art/standard.txt)
- [thinkertoy](../ascii-art/thinkertoy.txt)

Implement the following HTTP endpoints:

1. GET `/`: Sends HTML response, the main page.\
   1.1. GET Tip: [go templates](https://golang.org/pkg/html/template/) to receive and display data from the server.

2. POST `/ascii-art`: that sends data (text and specified banner) to the Go server.\
   2.1. POST Tip: use form and other types of [tags](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) to make the post request.

The way you display the result from the POST is up to you. What we recommend are one of the following :

- Display the result in the route `/ascii-art` after the POST is completed. So going from the home page to another page.
- Or display the result of the POST in the home page. This way appending the results in the home page.

The main page must have:

- a text input field
- radio buttons, select object or another method to switch between banners
- a button that sends a POST request to '/ascii-art' and displays the result on the page.

### HTTP status code

Your endpoints must return appropriate HTTP status codes.

- 200 OK, if everything went without errors.
- 404 Not Found, if nothing is found, for example templates or banners.
- 400 Bad Request, for incorrect requests.
- 500 Internal Server Error, for unhandled errors.

## Markdown

In the root project directory create a `README.MD` file with the following sections and contents:

- Description
- Authors
- Usage: how to run
- Implementation details: algorithm

### Instructions

- HTTP server must be written in **Go**.
- HTML templates must be in the project root directory _templates_.
- The code must respect the [good practices](../good-practices/README.md).

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Usage

- [Here's an example](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20).
