#### Functional

##### Try running `./TCPChat`.

###### Is the server listening for connections on the default port?

##### Try running `./TCPChat 2525 localhost`.

```
[USAGE]: ./TCPChat $port
```

###### Did the server respond with usage, as above?

##### Try running `./TCPChat 2525`.

###### Is the server listening for connections on the port 2525?

##### Try opening 3 terminals, run on the first terminal the command `./TCPChat <port>` and on the second and third terminal run the command `nc <host ip> <port>`.

###### Do both clients connect to the server with success?

##### Try creating a server and 2 Clients.

###### Did the server respond with a linux logo and ask for the name?

##### Try creating a server and 2 Clients.

###### Do all Clients receive a message informing them that the Client joined the chat?

##### Try creating a server and 2 Clients and send a message using the first Client.

###### Does the second Client receive the message?

##### Try creating a server and 1 Client and send some messages using this Client. Then create a new Client.

###### Can the new Client see all the previous messages?

##### Try creating a server and 3 Clients and send a message using the second Client.

###### Did all the Clients (first, second and third) received the same message?

##### Try creating a server and use 2 or 3 different computers and create one Client for each computer.

###### Did the server/Clients connect with success?

##### Try creating a server and 4 Clients and disconnect one of the Clients.

###### Do the rest of the Clients stay connected?

##### Try creating a server and 3 Clients and disconnect one of the Clients.

###### Do the rest of the Clients receive a message notifying that the Client left?

##### Try creating a server and 3 Clients. Then send messages between the Clients.

```
[2020-01-20 15:48:41][client.name]:[client.message]
```

###### Are the messages identified by the name of each Client and the time that the messages were sent, as above?

###### Are the connections between server and Clients well established?

###### Does the project present go routines?

###### Does the project use channels or mutexes?

##### Are the students using only the allowed packages?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Can the Clients change their names?

###### +Is the chat group informed if a Client changes his name?

###### +Does the server produce logs about Clients activities?

###### +Are the server logs saved into a file?

###### +Is there more NetCat flags implemented?

###### +Does the project present a Terminal UI using JUST this package : https://github.com/jroimartin/gocui?

#### Basic

###### +Does the project run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc...)

###### +Does the code obey the [good practices](../../good-practices/README.md)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
