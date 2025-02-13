#### Authentication

###### Are an email and a password asked for in the registration?

###### Does the project detect if the email or password are wrong?

###### Does the project detect if the email or user name is already taken in the registration?

##### Try to register as a new user in the forum.

###### Is it possible to register?

##### Try to login with the user you created.

###### Can you login and have all the rights of a registered user?

##### Try to login without any credentials.

###### Does it show a warning message?

###### Are sessions present in the project?

##### Try opening two different browsers and login into one of them. Refresh the other browser.

###### Can you confirm that the browser non logged remains unregistered?

##### Try opening two different browsers and login into both of them. Refresh both browsers.

###### Can you confirm that only one of those browsers has an active session?

##### Try opening two different browsers and login into one of them. Then create a new post or just add a comment. Refresh both browsers.

###### Does it present the comment/post on both browsers?

#### SQLite

###### Does the code contain at least one CREATE query?

###### Does the code contain at least one INSERT query?

###### Does the code contain at least one SELECT query?

##### Try registering in the forum, open the database with `sqlite3 <database_name.db>` and perform a query to select all the users (Example: SELECT \* FROM users;).

###### Does it present the user you created?

##### Try creating a post in the forum, open the database with `sqlite3 <database_name.db>` and perform a query to select all the posts (Example: SELECT \* FROM posts;).

###### Does it present the post you created?

##### Try creating a comment in the forum, open the database with `sqlite3 <database_name.db>` and perform a query to select all the comments (Example: SELECT \* FROM comments;).

###### Does it present the comment you created?

#### Docker

###### Does the project have Dockerfiles?

##### Try to run the command `"docker image build [OPTINS] PATH | URL | -"` to build the image using the project Dockerfiles and run the command `"docker images"` to see images.

```
student$ docker images
REPOSITORY              TAG                             IMAGE ID            CREATED             SIZE
<name of the image>     latest                          85a65d66ca39        7 seconds ago       795MB
```

###### Did all images build as above?

##### Try running the command `"docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]"` to start the containers using the images just created and run the command `"docker ps -a"` to see containers.

```
student$ docker ps -a
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                    NAMES
cc8f5dcf760f        <name of the image>    "./server"               6 seconds ago       Up 6 seconds        0.0.0.0:8080->8080/tcp   forum
```

###### Are the Docker containers running as above?

###### Does the project have no [unused objects](https://docs.docker.com/config/pruning/)?

#### Functional

##### Enter the forum as a non-registered user and try to create a post.

###### Are you forbidden from creating a post?

##### Enter the forum as a non-registered user and try to create a comment.

###### Are you forbidden from creating a comment?

##### Enter the forum as a non-registered user and try to like a comment.

###### Are you forbidden from liking a post?

##### Enter the forum as a non-registered user and try to dislike a comment.

###### Are you forbidden from disliking a comment?

##### Enter the forum as a registered user, go to a post and try to create a comment for it.

###### Were you able to create the comment?

##### Enter the forum as a registered user, go to a post and try to create an empty comment for it.

###### Were you forbidden from creating the empty comment?

##### Enter the forum as a registered user and try to create a post.

###### Were you able to create a post?

##### Enter the forum as a registered user and try to create an empty post.

###### Were you forbidden from creating the empty post?

##### Try creating a post as a registered user and try to choose several categories for that post.

###### Were you able to choose several categories for that post?

##### Try creating a post as a registered user and try to choose a category for that post.

###### Were you able to choose a category for that post?

##### Enter the forum as a registered user and try to like or dislike a post.

###### Can you like or dislike the post?

##### Enter the forum as a registered user and try to like or dislike a comment.

###### Can you like or dislike the comment?

##### Enter the forum as a registered user, try liking or disliking a post and then refresh the page.

###### Does the number of likes/dislikes change?

##### Enter the forum as a registered user and try to like and then dislike the same post.

###### Can you confirm that it is not possible that the post is liked and disliked at the same time?

##### Enter the forum as a registered user and try seeing all of your created posts.

###### Does it present the expected posts?

##### Enter the forum as a registered user and try seeing all of your liked posts.

###### Does it present the expected posts?

##### Navigate to a post of your choice and see its comments.

###### Are all users (registered or not) able to see the number of likes and dislikes that comment has?

##### Try seeing all posts from one category using the filter.

###### Are all posts displayed from that category?

###### Did the server behaved as expected?(did not crashed)

###### Does the server use the right [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)?

###### Are all the pages working? (Absence of 404 page?)

###### Does the project handle [HTTP status 400 - Bad Requests](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)?

###### Does the project handle [HTTP status 500 - Internal Server Errors](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)?

###### Are only the allowed packages being used?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Does the project present a script to build the images and containers? (using a script to simplify the build)

###### +Is the password encrypted in the database?

#### Basic

###### +Does the project run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

###### +Does the code obey the [good practices](../../good-practices/README.md)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
