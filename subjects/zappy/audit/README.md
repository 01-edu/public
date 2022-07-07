#### Functional

**_Guildlines_**

##### Zappy isn't a trivial project to correct. it is also long project to grade. For a full Project you need to spend around:

##### - 20 to 30 minutes for the server testing
##### - 10 to 20 minutes for the graphic client testing
##### - 5 to 10 minutes to test the client's AI

##### take the necessary time to check out the work of your peer carefully, Between 35 and 60 minutes in total

**_resources_**


##### Can you confirm that food and stones exist as resources in the game?

##### Can you confirm that six types of stones are present in the game? (linemate, deraumere, sibur, mendiane, phiras, thystame)



**_The server_**

##### You will have to evaluate the functional and technical quality of the server.

##### Try to run `"./server"`.

```console
$ ./server
 Usage: ./server -p <port> -w <width> -y <height> -nt <team> [<team>] [<team>] ... -c <nb> [-t <t>]
 -p port number
 -w world width
 -y world height
 -nt team_name_1 team_name_2 ...
 -c number of clients authorized at the beginning of the game
 -t [100] time unit divider (the greater t is, the faster the game will go)
```

##### Does it display the correct result as above?

##### Try to run `./server -p 8080 -x 10 -y 10 -c 5 -nt "TeamOfVectory" -t 10` then open another terminal and run `telenet 127.0.0.1 8080`

```console 
$ telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
WELCOME
```
##### Does it display the correct result as above?

##### Use Vscode or a simular program and search about any `exec` functions like [`execve`,`execpe`,`execl`,`execlp`,`execle`,...]
##### Did he/she use one of those function ?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory TeamOfPower "`

```console 
[ SQ ] : FOOD R0 R1 R2 R3 R4 R5 BOTS ...
[ 0] : 0 1 0 0 0 0 0 
[ 1] : 0 3 2 0 0 0 0 
...
[ 97] : 1 3 1 0 0 0 0 
[ 98] : 0 1 0 0 0 0 0 
[ 99] : 0 4 0 0 0 0 0 
```

##### Does it display the correct result as above? check only the first column and the first row of output

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory -t 10"`, then open another terminal and run `"siege -b 127.0.0.1:8080"`

##### The program `./server` is still working ?

##### + Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory"` at 2 terminal separate 

##### Does it display in the second terminal `error: Address already in use` or something similar?

**_Client_**

##### Try to run `"./client"`

```console
$ ./client
Usage: ./client -n <team> -p <port> [-h <hostname>]
 -n team_name
 -p port
 -h name of the host , the default is localhost
```

##### Does it display the correct result as above?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory"` then run `"./client -n TeamOfVectory -p 8080"`

##### Does the program launch without any errors?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory"` then run `"./client -n TeamOfVectory -p 8080 -h 127.0.0.1"`

##### Do the two programs interact with each other?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -nt TeamOfVectory"` then run `"./client -n TeamNotWorking -p 8080 -h 127.0.0.1"`

##### Does the program print `" Error: the team TeamNotWorking does not exist"` in the server part, and the client kick out?

**_Graphic Client_**

##### For the square content, Test the possibility of clicking on a square to see details about it like a floating window, tooltip, or something else?

##### For an Advance client, Is it possible to distinguish the number of similar stones on a square?

##### Does the client connects properly to the server and displays the map?

##### For a graphic client vision, are the players, stones, and food visible?

##### Can you (as a graphic client) click on a player to see his characteristics in a floating window, tooltip, or something else?

##### Is it possible to scroll a map?

##### Does each player starts on level 1?

##### Is the player able to pick up food?

##### Is one food unit equal to 126-time units?

##### Try to run `go run . example00.txt`.

##### Is the player able to pick up stones?

##### Can you confirm that if the player does not eat, he starves?

##### Can you confirm that if the player eats, he will survive longer?

##### Can the player level up when it can confirm the requirements to do so?

##### Does the sight of the players increase with the level?

##### Can the player hatch an egg if he needs a spot in his family?

##### Can you confirm that if the player exits on the right side of the board, he will re-enter on the left side or vice versa?

##### Is there sound management for broadcast?

##### Try to confirm that the rules of the elevation ritual are the same given in the subject.

##### Are the elevation ritual rules the same?

##### Can you confirm that in order to elevate to `level 2`, the player needs the stone `linemate`

##### Can you confirm that to elevate to `level 3`, the player needs a combination of the stones `linemate,` `deraumere,` `sibur,` and two players on the same level?

##### Can you confirm If the movements left and right commands are working?

**_The AI Client_**

##### Test with a t betwen 50 and 100 .

##### Test if the time taken by each movement and action is respected?

##### Try to run `"./server "`

#### General

##### +Is the visualizer in 3D?

#### Basic

##### +Does the project run quickly and effectively (favoring recursive, no unnecessary data requests, etc.)?

##### +Does the code obey the [good practices](../../good-practices/README.md)?

#### Social

##### +Did you learn anything from this project?

##### +Would you recommend/nominate this program as an example for the rest of the school?

