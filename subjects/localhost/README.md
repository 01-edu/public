## Localhost

Finally you are going to understand how internet works from the server side. The Hypertext Transfer Protocol was created in order to ensure a reliable way to communicate on a request/response base.

This protocol is used by servers and clients (usually browsers) to serve content and it is the backbone of the World Wide Web, still it is also used in many other cases that are far beyond the scope of this exercise.

Here you will learn the basics of the protocol and a good place to start could be the [HTTP/1.1 RFC](https://www.rfc-editor.org/rfc/rfc9112.html).


### Instructions

- The project must be written in `Rust`.
- You can use the `libc` crate for functions as `epoll` or equivalents and any other necessary system call that Rust do not implement natively.
- You may need to use `unsafe` keyword for some operations (don't abuse it).

> You can't use crates that already implement server features like `tokio` or `nix`.

#### The Server

You goal is to write your very own server, which will use `HTTP` protocol to serve static web pages to browsers.

For your server you must guarantee the following behavior:
- It **never** crashes.
- All requests timeout if they are taking too long.
- It can listen on multiple ports and instantiate multiple servers at the same time.
- You use only one process and one thread.
- It receives a request from the browser/client and send a response using the `HTTP` header and body.
- It is compatible with `HTTP/1.1` protocol.
- You can compare your results with `NGINX` which will be used as the reference.
- It is compatible with the last version of your chosen browser.
- It manages at least [`GET`, `POST`, `DELETE`] methods.
- It is able to receive file uploads made by the client.
- It handles cookies and sessions.
- You should create default error pages for at least the following error codes [400,403,404,405,413,500].
- It calls `epoll` function (or equivalent) only once for each client/server communication.
- All reads and writes should pass by `epoll` or equivalent API.
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

> There is no need to pass through `epoll` when reading the configuration file.

#### Testing your server
- Do stress tests with `siege -b [IP]:[PORT]`, it must stay available at all costs (availability should be up to 99.5, it will be tested during audits).
- Create tests for as many cases as you can (redirections, bad configuration files, static and dynamic pages, default error pages and so on).
- You will be requested to provide and explain your tests during the audits.
- You can use the language you prefer to write tests, as long as they are exhaustive and the auditor can check their behavior.
- Test possible memory leaks before to submit the project.
- Once again, the server should never crash and never leak memory.

> Attention: `siege` is a stressing tool, use it ONLY to test your own server. Do **NEVER** use it on any server/website without the owner's permission. If you do so you would have illegally DDoSed a server and could face serious troubles.

### Unit Tests

You must implement unit tests within your `Localhost` project to ensure your protocol parsing and configuration logic are robust. Specifically, your tests should:

- Verify **HTTP Request Parsing** by ensuring headers, methods (GET, POST, DELETE), and body content (including chunked encoding) are correctly extracted from raw byte streams.
- Verify **Configuration Validation** by ensuring the server correctly identifies conflicting ports, invalid paths, and out-of-bounds body size limits during startup.
- Test **Route Matching** logic to confirm that requests are mapped to the correct directory roots, CGI scripts, or redirections based on the configuration file.
- Verify **Status Code Generation** by ensuring the internal logic returns the correct `4xx` or `5xx` codes for malformed requests or missing files before the response is sent.

### Bonus
- Handle at least one more `CGI`.
- Rewrite the project in another programming language (can be `C++` or `C`).
