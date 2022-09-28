## Localhost

Finally you are going to understand how internet works from the server side. The Hypertext Transfer Protocol was created in order to ensure a reliable way to communicate on a request/response base.

This protocol is used by servers and clients (usually browsers) to serve content and it is the backbone of the World Wide Web, still it is also used in many other cases that are far beyond the scope of this exercise.

Here you will learn the basics of the protocol and a good place to start could be the [HTTP/1.1 RFC](https://www.rfc-editor.org/rfc/rfc9112.html).


### Instructions

- The project can be written in one of these languages [`Rust`, `C++`, `C`].

#### The Server

- Your server should **never** crash.
- All requests should timeout if they are taking too long.
- Your server should be able to listen on multiple ports and instantiate multiple servers at the same time.
- You must use only one process and one thread.
- Your server must receive a request from the browser/client and send a response using the `HTTP` header and body.
- Your server should be compatible with `HTTP/1.1` protocol.
- You can compare your results with `NGINX` which will be used as the reference.
- Your server should be compatible with the last version of your chosen browser.
- Your server should manage at least [`GET`, `POST`, `DELETE`] methods.
- Your server should be able to receive file uploads made by the client.
- Your server should handle cookies and sessions.
- You should create default error pages for at least the following error codes [400,403,404,405,413,500].
- Your server should call `select` function (or `poll` or equivalent) only once for each client/server communication.
- All reads and writes should pass by `select` or equivalent API.
- All I/O operations should be non-blocking.
- You should manage chunked and unchunked requests.
- You should set the right status for each response.

#### The CGI
- Based on the file extension the server will execute the corresponding `CGI` (for example `.php` or `.py`).
- You need to implement only one `CGI` of your choice.
- You are allowed to fork a new process to run the `CGI`.
- `CGI` expects the file to process as first argument and `EOF` as end of the body.
- Pay attention to the directory where the `CGI` will run for correct relative paths handling.
- The `CGI` will check `PATH_INFO` environment variable to define the full path.
	
#### Configuration File

In the file you should be able to specify the following:

- The host (server_address) and one or multiple ports for each server.
- The first server for a host:port will be the default if the "server_name" didn't match any other server.
- Path to custom error pages.
- Limit client body size for uploads.
- Setup routes with one or multiple of the following settings:
  - Define a list of accepted HTTP methods for the route.
  - Define HTTP redirections.
  - Define a directory or a file from where the file should be searched (for example, if `/test` is rooted to `/usr/Desktop`, the URL `/test/my_page.html` will route to `/usr/Desktop/my_page.html`).
  - Define a default file for the route if the URL is a directory.
  - Specify a `CGI` to use for a certain file extension.
  - Turn on or off directory listing.
  - Set a default file to answer if the request is a directory.
- No need to manage comments "(#)".

> Routes won't need to support regular expressions.

> There is no need to pass through `poll` when reading the configuration file.

#### Testing your server
- Do stress tests (for example with `siege -b [IP]:[PORT]`), it must stay available at all costs (availability should be up to 99.5).
- Create tests for as many cases as you can (redirections, bad configuration files, static and dynamic pages, default error pages and so on).
- You will be requested to provide and explain your tests during the audits.
- You can use the language you prefer to write tests, as long as they are exhaustive and the auditor can check their behavior.
- Test possible memory leaks before to submit the project.
- Once again, the server should never crash and never leak memory.

### Bonus
- Handle at least one more `CGI`.
- Write the project in two different programming languages.

> If the two languages are C and C++ the provided solution for C++ should heavily rely on C++ specific features.
