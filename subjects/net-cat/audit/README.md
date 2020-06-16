#### Functional

##### Try running `"./TCPChat"`.

###### Is the server listening for connections on the default port?

##### Try running `"./TCPChat" 2525 localhost`.

```
[USAGE]: ./TCPChat $port
```

###### Did server respond with usage, as above?

##### Try running `"./TCPChat 2525"`.

###### Is the server listening for connections on the port 2525?

##### Try opening 3 terminals, run on the first terminal the command `"./TCPChat <port>"` and on the second and third terminal run the command `"nc <host ip> <port>"`.

###### Does both clients connect to the server with success?

##### Try creating a server and 2 Clients.

###### Did the server responded with a linux logo and asked for the name?

##### Try creating a server and 2 Clients.

###### Do all Clients receive a message informing that the Client joined the chat?

##### Try creating a server and 2 Clients and send a message using the first Client.

###### Does the second Client receive the message?

##### Try creating a server and 1 Client and send some messages using this Client. Then create a new Client.

###### Can the new Client see all the previous messages?

##### Try creating a server and 3 Clients and send a message using the second Client.

###### Does all the Clients (first, second and third) received the same message?

##### Try creating a server and use 2 or 3 different computers and create for each computer one Client.

###### Did the server/Clients connected with success?

##### Try creating a server and 4 Clients and disconnect one of the Clients.

###### Does the rest of the Clients stay connected?

##### Try creating a server and 3 Clients and disconnect one of the Clients.

###### Does the rest of the Clients receive a message notifying that the Client left?

##### Try creating a server and 3 Clients. Then send messages between the Clients.

```
[2020-01-20 15:48:41][client.name]:[client.message]
```

###### Are the messages identified by the name of each Client and the time that the messages was sent, as above?

###### Is the connections between server and Clients well established?

###### Does the project present go routines?

###### Does the project use channels or mutexe?

##### Are the students using just the allowed functions?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Can the Clients change their names?

###### +Is the chat group informed if a Client changes his name?

###### +Does the server produce logs about Clients activities?

###### +Does the server logs saved into a file?

###### +Is there more NetCat flags implemented?

###### +Does the project present a Terminal UI using JUST this package : https://github.com/jroimartin/gocui?

#### Basic

###### +Does the project runs quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices.en)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
