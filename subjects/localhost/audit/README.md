Localhost is about creating your own HTTP server and test it with an actual browser.

##### Take the necessary time to understand the project and to test it, looking into the source code will help a lot.

### Functional

**_Is the student able to justify his choices and explain the following:_**
**_Note:_** Ask the student to show you the implementation in the source code when necessary.

###### How does an HTTP server works?

###### Which function was used for I/O Multiplexing and how does it works?

###### Is the server using only one select (or equivalent) to read the client requests and write answers?

###### Why is it important to use only one select and how was it achieved?

###### Read the code that goes from the select (or equivalent) to the read and write of a client, is there only one read or write per client per select (or equivalent)?

###### Are the return values for I/O functions checked properly?

###### If an error is returned by the previous functions on a socket, is the client removed?

###### Is writing and reading ALWAYS done through a select (or equivalent)?

### Configuration file

Check the configuration file and modify it if necessary.
**_Are the following configurations working properly:_**

###### Setup a single server with a single port.

###### Setup multiple servers with different port.

###### Setup multiple servers with different hostnames (for example: curl --resolve test.com:80:127.0.0.1 http://test.com/). This aims to confirm if your server correctly distinguishes between requests for different hostnames even though they resolve to the same IP and port.

###### Setup custom error pages.

###### Limit the client body (for example: curl -X POST -H "Content-Type: plain/text" --data "BODY with something shorter or longer than body limit").

###### Setup routes and ensure they are taken into account.

###### Setup a default file in case the path is a directory.

###### Setup a list of accepted methods for a route (for example: try to DELETE something with and without permission).

### Methods and cookies

**_For each method be sure to check the status code (200, 404 etc):_**

###### Are the GET requests working properly?

###### Are the POST requests working properly?

###### Are the DELETE requests working properly?

###### Test a WRONG request, is the server still working properly?

###### Upload some files to the server and get them back to test they were not corrupted.

###### A working session and cookies system is present on the server?

### Interaction with the browser

Open the browser used by the team during tests and its developer tools panel to help you with tests.

###### Is the browser connecting with the server with no issues?

###### Are the request and response headers correct? (It should serve a full static website without any problem).

###### Try a wrong URL on the server, is it handled properly?

###### Try to list a directory, is it handled properly?

###### Try a redirected URL, is it handled properly?

###### Check the implemented CGI, does it works properly with chunked and unchunked data?

### Port issues

###### Configure multiple ports and websites and ensure it is working as expected.

###### Configure the same port multiple times. The server should find the error.

What we mean to check is how your server handles configuration errors or conflicts. If you configure the same port multiple times within your server's configuration file for the same or different servers, the expectation is that your server should catch this error and handle it appropriately. It's about ensuring that your server can identify the misconfigurations.

###### Configure multiple servers at the same time with different configurations but with common ports. Ask why the server should work if one of the configurations isn't working.

This aims to validate your server's configuration handling.
You'll configure multiple servers simultaneously, each with different configurations but with shared ports. If one of these configurations isn't valid or encounters an issue, your server should continue to function for the other configurations without being entirely disrupted. The purpose is to ensure that an error in one server's configuration doesn't bring down the entire server if other configurations are correctly set up.

### Siege & stress test

###### Use siege with a GET method on an empty page, availability should be at least 99.5% with the command `siege -b [IP]:[PORT]`.

###### Check if there is no memory leak (you could use some tools like top).

###### Check if there is no hanging connection.

#### General

###### +There's more than one CGI system such as [Python,C++,Perl].

###### +There is a second implementation of the server in a different language (repeat practical tests on it before to validate).
