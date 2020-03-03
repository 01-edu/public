#### Functional

##### Enter the website as a non-registered user.

###### Are you able to see posts/comments?

##### Enter the website as a non-registered user.

###### Are you prohibited to create a post/comment?

##### Enter the website as a registered user.

###### Are you able to create a post/comment?

##### Try creating a post as a registered user.

###### Are you able to choose a category for that post?

###### Is the forum composed of services?

###### The communication between the services well stablish? Does the communication use [smart EndPoints](https://medium.com/@nathankpeck/microservice-principles-smart-endpoints-and-dumb-pipes-5691d410700f)(ex: Request-Response, Observer)?

#### SQLite

###### Does the code contain at least one CREATE query?

###### Does the code contain at least one INSERT query?

###### Does the code contain at least one SELECT query?

###### Does the code contain at least one DELETE query?

###### Is the use of the sqlite3 command missing from the code?

##### Try registering in the forum, open the database with `sqlite3 <database_name.db>` and perform a query to select all the users (Example: SELECT * FROM users;).

###### Does it present the user you created?

##### Try creating a post in the forum, open the database with `sqlite3 <database_name.db>` and perform a query to select all the users (Example: SELECT * FROM post;).

###### Does it present the post you created?

#### Authentication

##### Try to register as a new user in the forum.

###### Is it possible to register?

##### Try opening two different browsers and login into one of them. Then create a new post or just add a comment. Refresh the other browser.

###### Can you confirm that the browser non logged remains unregistered?

##### Try to login with the user you created.

###### Can you login and have all the rights of a registered user?

##### Are sessions present in the project?

##### Try to login without any credentials.

###### Does it show a warning message?

##### Try opening two different browsers and login into one of them. Then create a new post or just add a comment. Refresh the other browser.

###### Does it present the changes made in the logged in browser?

###### Can you confirm that it is not possible to do it?

###### Is it asked in the register for an email and a password?

###### Does the project detects if the email or password are wrong?

###### Does the project detects if the email or user name is already taken in the register?

#### Docker

###### Does the project have Dockerfiles?

##### Try to run the command `"docker image build [OPTINS] PATH | URL | -"` to build the image using using the project Dockerfiles?
```
student$ docker images
REPOSITORY              TAG                             IMAGE ID            CREATED             SIZE
<name of the image>     latest                          85a65d66ca39        7 seconds ago       795MB
```
###### Run the command `"docker images"` to see images. Does all images build as above?

##### Try running the command `"docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]"` to start the containers using the images just created.
```
student$ docker ps -a
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                    NAMES
cc8f5dcf760f        ascii-art-web-docker   "./server"               6 seconds ago       Up 6 seconds        0.0.0.0:8080->8080/tcp   ascii-art-web
```
###### Run the command `"docker ps -a"` to see containers. Is the docker containers running as above?

##### Try running the comment `"docker network ls"`
```
NETWORK ID          NAME                DRIVER              SCOPE
24560f95a216        <network-name>      bridge              local
24560f95a216        bridge              bridge              local
3f12239adb4f        host                host                local
159cdc8d8083        none                null                local
```
###### Is the network a bridge driver and does it appear in the list of networks as above?

##### Try openning a shell to the running container, using the [command](https://docs.docker.com/engine/reference/commandline/exec/) `"docker exec [OPTIONS] CONTAINER COMMAND [ARG...]"` (ex: `"docker exec -it <container-name> /bin/bash"`) and ping the other container that are connected to the same network, for [example](https://docs.docker.com/engine/tutorials/networkingcontainers/).

###### Did it ping with success? Are the multiple services/containers communicating between each other?

###### Does the project present no [unused object](https://docs.docker.com/config/pruning/)?

###### Did the server behaved as expected?(did not crashed)

###### Does the server use the right [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)?

###### Has the website runned without crashing at anytime?

###### Are all the pages working? (Absence of 404 page?)

###### Does the project avoid [HTTP status 400](https://kinsta.com/knowledgebase/400-bad-request/#causes)?

###### Does the project avoid [HTTP status 500](https://www.restapitutorial.com/httpstatuscodes.html)?

###### Are the libraries used allowed?

#### General

###### +Is the database a service?

###### +Does the project present a script to build the images and containers? (using a script to simplify the build)

#### Basic

###### +Does the project runs quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices.en)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
