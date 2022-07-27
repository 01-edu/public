#### Functional

#### Localhost is about program your own HTTP server and test it with an actual browser.
###### Take The necessary time to understand the project and to test it,take a look around the code to ask reasonable questions such as:
###### Explain the basics of an HTTP server.
###### Ask which function they used for I/O Multiplexing.
###### Ask to get an explanation of how select (or equivalent) is working.
###### Ask if they use only one select (or equivalent) and how they've managed the server accept and the client read/write.
###### There should be only one read or one write per client per select (or equivalent). Ask to show you the code that goes from the select (or equivalent) to the read and write of a client.
###### Search for all [read,recv,write,send] and check if the returned value is well checked. (checking only -1 or 0 is not good you should check both)
###### Check return value for O/I functions [read,recv,write,send] . is the audite checking for `-1` and `0` at the same time. 
###### Writing or reading ANY file descriptor without going through the select (or equivalent) is strictly 

### Configuration

#### In the configuration file, check if you can do or check the following:
##### Look at the subject configuration file part to compare it to the project configuration file. is it the same? 
##### Look for the HTTP response status codes list on internet. 
##### Set up multiple servers with different port
##### Set up multiple servers with different hostnames (use something like: curl --resolve example.com:80:127.0.0.1 http://example.com/)
##### Setup default error page
##### Limit the client body (use curl -X POST -H "Content-Type: plain/text" --data "BODY IS HERE write something shorter or longer than body limit")
##### Setup routes in a server to different directories
##### Setup a default file to search for if you ask for a directory
##### Set up a list of methods accepted for a certain route (ex: try to delete something with and without permission)

### Basic checks

###### Using telnet, curl, and prepared files demonstrate that the following features work correctly:
##### Test the GET requests if it works well.
##### Test the POST requests if it works well.
##### Test the DELETE requests if it works well.
##### Test the UNKNOWN requests, it should not work, but the server is not supposed to crash.
##### For every test, there is a status code (RFC) that must be respected.
##### upload some files to the server and get them back.

### Check with a browser

##### Use the reference browser of the team, open the network part of it, and try to connect to the server with it.
##### Look at the request header and response header, It should be compatible with serving a fully static website.
##### Try a wrong URL on the server.
##### Try to list a directory.
##### Try a redirected URL.
##### Try logical things

### Cookies and session

##### It should be a working session and cookies system on the webserver.

### Port issues

##### In the configuration file setup multiple ports and use different websites, use the browser to check that the configuration is working as expected, and show the right website.
##### In the configuration try to setup the same port multiple times. It should not work.
##### Launch multiple servers at the same time with different configurations but with common ports. Is it working? If it is working, ask why the server should work if one of the configurations isn't working?

### Siege & stress test

##### Use Siege to run some stress tests.
##### Availability should be above 99.5% for a simple get on an empty page with a `siege -b [IP]:[PORT]` on that page.
##### Check if there is no memory leak (monitor the process memory usage it should not go up indefinitely).
##### Check if there is no hanging connection.
##### You should be able to use siege indefinitely without restarting the server (look at siege -b).

### Bonus Part

##### + There's more than one CGI system such as [Python, C++,Perle].
##### + Browse http://[IP]:[PORT]/mysql.php and check if the student create a php page that can connect to a database
