## Localhost

The Hypertext Transfer Protocol (HTTP) is an application layer protocol in the Internet protocol suite model for distributed, collaborative, hypermedia information systems.
HTTP is the foundation of data communication for the World Wide Web. Hypertext documents include hyperlinks to other resources that the user can easily access, for example, by a mouse click or tap on the screen in a web browser.

### Instructions

#### Server

The Client can be written in one of these languages [Go,C++,C...]

```console
$ ./localhost [CONFIG FILE]
Server is Runing in http://[IP]:[PORT]
```
- Your server must not crash for any reason.
- Any request to your server should never hang forever.
- Your server can listen on multiple servers and ports at the same time without any conflict 
- You must use the single process and no treads.
- You must use `exec` functions only in `CGI` part
- Your server must be compatible with the last version of your chosen browser 
- You can use `NGINX` to compare headers and answer behaviors with your server
- You must manage at least [`GET`, `POST`, `DELETE`] methods.
- You must handle customization errors **(Binding error, default page error...)**
- You must create error page at leat for [400,311,403,404,405,413]
- You must execute CGI based on certain file extensions [`.php`,`.py`,...]
- You must use the enverement paramatre of cgi. 
- Your server must receive a request from the browser and send a response using the HTTP header and body
- Your server must call `select` function (or equivalent function) one time only

here is an example of an HTTP request with `GET` method

```http

GET /Welcome.html HTTP/1.1
User-Agent: Mozilla/4.0 (compatible; MSIE5.01; Windows NT)
Host: www.01talent.com
Accept-Language: en-us
Accept-Encoding: gzip, deflate
Connection: Keep-Alive
```
and here is an example of an HTTP response : 

```http
HTTP/1.1 404 Not Found
Date: Thu, 19 Jul 2022 10:36:20 GMT
Server: Hserver/1.1.0
```
> Stress tests your server. It must stay available at all cost. use this command to test it  `siege -b [IP]:[PORT]`
	the availability should be up to 95.99
	
> Learn about  [CGI](https://en.wikipedia.org/wiki/Common_Gateway_Interface)


#### Configurations File

This an example of a simple server configuration.
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
        
        upload_status on; //if it's "on," you should add the "upload_path"; otherwise, if it's "off," there is no need for "upload_status." 
        
        upload_path         /usr/share/nginx/upload; //depends on "upload_status"
    }
}
```

You have to base your server on this example but You have tested more complex configurations with multiple locations and ports and multiple servers such as :
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
    server_name                     hello.com;
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
The Main Rules:
- Choose the host(server_address) and a port or multiple ports of each server.
- The first server for a host: port will be the default if the server_name didn't mutch any other servers server_name.
- Setup default error pages.
- Limit client body size In uploading.
- Setup routes with one or multiple of the following rules/configuration (routes won't be using regexp):
  - Define a list of accepted HTTP methods for the route.
  - Define an HTTP redirection.
  - Define a directory or a file from where the file should be searched (for example, if URL /test is rooted to /usr/share, url /test/tmp/folder1/folder2 is /usr/share/test/tmp/folder1/folder2).
  - Turn on or off the directory listing.
  - Set a default file to answer if the request is a directory.

to check the syntax of the config file you must run this : 

```console
$ ./localhost -t GoodConfigFile.conf
Localhost : the configuration file GoodConfigFile.conf syntax is ok
$ ./localhost -t WrongConfigFile.conf
Localhost : the configuration file WrongConfigFile.conf syntax is not ok
```

> If you’ve got a question about one behavior, you should compare your
program behavior with NGINX’s.
