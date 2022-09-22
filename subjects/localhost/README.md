## Localhost

Finally you are going to understand how internet works from the server side. The Hypertext Transfer Protocol was created in order to ensure a reliable way to communicate on a request/response base.
It is used by servers and clients (usually browsers) to serve content and it is the backbone of the World Wide Web, still it is also used in many other cases that are far beyond the scope of this exercise.
Here you will learn the basics of the protocol and a good place to start could be the [RFC](https://www.rfc-editor.org/rfc/rfc9112.html).


### Instructions

- The project can be written in one of these languages [`Rust`, `C++`, `C`].

#### The Server

- Your server should **never** crash.
- All requests should timeout if they are taking too long.
- Your server should be able to listen on multiple ports and instantiate multiple servers at the same time.
- Your server must receive a request from the browser/client and send a response using the `HTTP` header and body.
- You must use only one process and one thread.
- Your server should be compatible with `HTTP/1.1` protocol.
- You can compare your results with `NGINX` which will be used as the reference.
- Your server should be compatible with the last version of your chosen browser.
- Your server should manage at least [`GET`, `POST`, `DELETE`] methods.
- Your server should handle cookies and sessions.
- You should create default error pages for at least the following error codes [400,403,404,405,413,500].
- Your server should call `select` function (or equivalent function) one time only.
- You should manage chunked and unchunked requests.
- You should set the right status for each response.

#### The CGI
- Based on the file extension the server will execute the corresponding `CGI` (for example `.php` or `.py`).
- You need to implement only one `CGI` of your choice.
- You are allowed to fork a new process to run the `CGI`.
- The `CGI` should run in the correct directory for relative path file access.
- You must use the environment parameter of `CGI`.
	
#### Configuration File

This an example of a simple server configuration :

```
server {
    listen                          80;
    
    server_name                     localhost;

    root                            /usr/share/nginx/html;
    
    error_page          Mozilla 404        /usr/share/nginx/errors/404.html;
    
    client_body_size                10m; 
    location /
    {
        http_methods  GET  POST DELETE;
        
        index index.html
        
        upload_status on; 

        upload_path         /usr/share/nginx/upload;
    }
}
```
In the example above if the "upload_status" is "off", there is no need for "upload_path." 

You have to base your server on this example but you must test also more complex configurations with multiple locations and ports and multiple servers such as:

```
{
server {
    listen                              80;
    server_address                      127.0.0.1;
    server_name                         localhost;
    ...
    ...
    ...
    location ...{

    }
    location ...{

    }
}
server {
    listen                              5000;
    listen                              5001;
    server_address                      192.182.2.1;
    server_name                         01edu.com;
    ...
    ...
    ...
    location ...{

    }
}
server {
    listen                          5001;
    listen                          5002;
    listen                          5003;
    server_address                  192.186.2.54;
    server_name                     01talent.com;
    ...
    ...
    ...
    location ...{

    }
    location ...{

    }
	location ...{

    }
}
}
```

The Main Rules :

- Choose the host(server_address) and a port or multiple ports for each server.
- The first server for a host:port will be the default if the "server_name" didn't match any other servers "server_name".
- Setup default error pages.
- Limit client body size for uploads.
- Setup routes with one or multiple of the following rules/configuration (routes won't be using regexp):
  - Define a list of accepted HTTP methods for the route.
  - Define an HTTP redirection.
  - Define a directory or a file from where the file should be searched (for example, if URL /test is rooted to /usr/Desktop, url /usr/Desktop/tmp/folder1/folder2).
  - Turn on or off the directory listing.
  - Set a default file to answer if the request is a directory.
  - No need to manage comments "(#)".

to check the syntax of the config file you must run this : 

```console
$ ./localhost -t GoodConfigFile.conf
Localhost : the configuration file GoodConfigFile.conf syntax is ok
$ ./localhost -t WrongConfigFile.conf
you miss server_addr in server 3
```

#### Testing your server
- Do stress tests (for example with `siege -b [IP]:[PORT]`), it must stay available at all costs (availability should be up to 95.99).
- Create and provide during the audit tests for as many cases as you can (redirections, bad configuration files, static and dynamic pages, default error pages and so on).
- You can use the language you prefer to write tests, as long as they are exhaustive and the auditor can check their behavior.
- Test possible memory leaks before to submit the project.
- Once again, the server should never crash and never leak memory.

### Bonus
- Handle at least one more `CGI`.
- Create a PHP page and connect it with a MySQL database. 
