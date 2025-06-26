#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

###### Does the project have a DockerFile?

##### Try running the [command](https://docs.docker.com/engine/reference/commandline/image_build/) `"docker image build [OPTIONS] PATH | URL | -"` to build the image using the project Dockerfile.

##### Example :

`"docker image build -f Dockerfile -t <name_of_the_image> ."`

```
student$ docker images
REPOSITORY               TAG                             IMAGE ID            CREATED             SIZE
ascii-art-web-docker     latest                          85a65d66ca39        7 seconds ago       795MB
```

###### Run the command `"docker images"` to see all images. Is the docker image built as above?

##### Try running the [command](https://docs.docker.com/engine/reference/commandline/container_run/) `"docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]"` to start the container using the image just created. (example : `"docker container run -p <port_you_what_to_run> --detach --name <name_of_the_container> <name_of_the_image>"`)

```
student$ docker ps -a
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                    NAMES
51c2efe2d366        ascii-art-web-docker   "./server"               6 seconds ago       Up 6 seconds        0.0.0.0:8080->8080/tcp   dockerize
```

###### Run the command `"docker ps -a"` to see all containers. Is the docker container running as above?

##### Try running the [command](https://docs.docker.com/engine/reference/commandline/exec/) `"docker exec [OPTIONS] CONTAINER COMMAND [ARG...]"`. (example : `"docker exec -it <container_name> /bin/bash"`) and do a `"ls -l"` to see the file system.

```
student$ docker exec -it dockerize /bin/bash
51c2efe2d366:/app$ ls -l
-rwxr-xr-x   1 root root 10833387 Sep  8 10:31 server
drwxr-xr-x   2 root root     4096 Sep  8 10:51 static
drwxr-xr-x   2 root root     4096 Sep  8 10:51 templates
51c2efe2d366:/app$ exit
exit
student$
```

###### Does the container's file system contain the expected files and directories?

###### Does the DockerFile contain some [metadata](https://docs.docker.com/config/labels-custom-metadata/) applied to the docker object?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Does the project present a script to build the images and containers? (using a script to simplify the build)

#### Basic

###### +Does the server run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

###### +Does the code obey the [good practices](../../../good-practices/README.md)?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Are the instructions in the website clear?

###### +Does the project run using an API?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
