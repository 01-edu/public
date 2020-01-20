#### Functional

##### Try running `"./net-cat -p 8080"`.
###### Is the server listening for connections?

##### Try opening 3 terminals, run on the first terminal the command `"./net-cat -p <port>"` and on the second and third terminal run the command `"./net-cat <host ip> <port>"`.
###### Does both clients connect to the server with success?

##### Try opening 4 terminals, run on the first terminal the command `"./net-cat -u -p <port>"` and on the second, third and fourth terminal run the command `"./net-cat -u <host ip> <port>"`.
###### Does all clients connect to the server with [UDP](https://www.privateinternetaccess.com/blog/2018/12/tcp-vs-udp-understanding-the-difference/) connection?

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
###### Are the messages identified by the name of each Client and the time that the messages was sent?

###### Is the connections between server and Clients well established?

#### General

###### +Can the Clients change their names?
###### +Is the chat group informed if a Client changes his name?
###### +Does the server produce logs about Clients activities?
###### +Is there more NetCat flags implemented?

#### Basic

###### +Does the project runs quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)
###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices.en)?
###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?
###### +Can it be open-sourced / be used for other sources?
###### +Would you recommend/nominate this program as an example for the rest of the school?
