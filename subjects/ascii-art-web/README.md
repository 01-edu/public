## ascii-art-web

### Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web **GUI** (graphical user interface) version of your last project, ascii-art.

Your web-page should provide usage of different [banners](https://github.com/01-edu/public/tree/master/subjects/ascii-art).

Implement following HTTP endpoints:

1. GET `/`: Sends HTML response - the main page.
2. POST `/ascii-art`: Receives _JSON_ body with the following data and returns _JSON_ response with the result of _ascii-art_:

Request body:

```json
{
  "banner": "shadow",
  "text": "Your text here"
}
```

Response body:

```json
{
  "result": "..."
}
```

Main page must have:
- text input
- radio buttons, select object or anything else to choose between banners
- button, which sends _AJAX_ request to '/ascii-art' and outputs the result on page.

### HTTP status code

Your endpoints must return appropriate HTTP status codes.

- OK (200), if everything went without errors
- Not Found, if anything is not found, e.g: template, banner etc.
- Bad Request, for incorrect requests
- Internal Server Error, for unhandled errors

## Markdown

In root project directory create `README.MD` file with the following sections and contents:
- Description
- Authors
- Usage: how to run
- Implementation details: algorithm

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Instructions
- HTTP server must be written in _Go_.
- HTML templates must be in project root directory _templates_.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/).

### Usage

- [Here's an example](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20).
