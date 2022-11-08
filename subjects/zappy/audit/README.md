#### Functional

**_Guidelines_**

Zappy is a long project to audit, for a full project review we advise you to spend around:

- 20 to 30 minutes for the server testing.
- 10 to 20 minutes for the graphic client testing.
- 5 to 10 minutes to test the client's AI.

Take the necessary time to check out the work of your peer carefully and ask him for guidance to understand his project.

**_The server_**

##### You will have to evaluate the functional and technical quality of the server.

##### Try to run `"./server"`.

```console
$ ./server
 Usage: ./server -p <port> -x <width> -y <height> -n <team> [<team>] [<team>] ... -c <nb> [-t <t>]
 -p port number
 -x world width
 -y world height
 -n team_name_1 team_name_2 ...
 -c number of clients authorized at the beginning of the game
 -t [100] time unit divider (the greater t is, the faster the game will go)
```

###### Does it display a result like the one presented above?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team -t 10"`. Open another terminal and run `"telnet 127.0.0.1 8080"`

```console
$ telnet 127.0.0.1 8080
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
WELCOME
```

###### Does it display result with a welcome message like the one presented above?

##### Use Vscode or a similar program and search for any `exec` functions like (`execve`,`execpe`,`execl`,`execlp`,`execle`,...).

###### Was the student able to do the project without using such functions?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team -t 10"`. Open another terminal and run `"siege -b 127.0.0.1:8080"`.

###### Is the program `./server` still working?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team"` in two different terminals.

###### Does it display, in the second terminal, `ERROR : Address already in use` or something similar?

#### Let the game run for a while and then check the time taken to execute the movements and actions of the player.

##### Is the timing of each movement and action respected?

###### Can you confirm that if the player does not eat, he starves and die?

###### Can you confirm that if the player eats, he will survive longer?

###### Does the sight of the players increase with the level rising?

###### Can you confirm that if the player exits on the right side of the board, he will re-enter on the left side or vice versa?

**_Client_**

##### Try to run `"./client"`.

```console
$ ./client
Usage: ./client -n <team> -p <port> [-h <hostname>]
 -n team_name
 -p port
 -h name of the host , the default is localhost
```

###### Does it display a result like the one presented above?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team"` then run `"./client -n my_team -p 8080"` on a different terminal.

###### Does the program launch without any errors?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team"` then run `"./client -n my_team -p 8080 -h 127.0.0.1"` on a different terminal.

###### Do the two programs interact with each other without errors?

##### Try to run `"./server -p 8080 -x 10 -y 10 -c 5 -n my_team"` then run `"./client -n wrong_team -p 8080 -h 127.0.0.1"`

###### Does the program print an error message like `"Error: the team wrong_team doesn't exist"` in the server part, and the client is kicked out?

**_Graphic Client_**

##### Ask the owners of the project for help to connect the Graphic Client to the server.

###### Does the client properly connect to the server and display the map?

##### For the square content, test the possibility of clicking on a square to see the details about it. It should look like a floating window, tooltip, or something else.

###### Can you confirm that the square is showing its content somehow?

###### Can you confirm that it is possible to distinguish the number of all stones on a square?

##### Search in the map for the player, stones and food to see if they are visible.

###### Are the players, stones, and food visible in the map?

##### Try to click on a player to see his characteristics.

###### Is it possible to see some kind of floating window, tooltip, or something else with the characteristics of the player?

###### Can you confirm that it is possible to visualize the sounds?

**_PLAYER_**

##### Try to run a game and take a close look at what the players are doing.

###### Does each player starts with 10 food (1260 time units) and 0 stones?

###### Does each player starts on level 1?

###### Is the player able to pick up food?

###### Is the player able to pick up stones?

###### Can the player perform the evolution ritual and level up?

###### Can the player call a ship if he needs a spot in his family?

**_resources_**

###### Can you confirm that food and stones exist as resources in the game?

###### Can you confirm that six types of stones are present in the game? (jade, peridot, amber, amethyst, garnet, ammolite)

###### Are resources randomly generated with some clear logical rules? (ask the students to explain those rules)

**_Food_**

###### Is one food unit equal to 126 time units?

**_Evolution ritual_**

##### Try to confirm that the rules of the evolution ritual are the same given in the subject. (ask the project owners for an explanation if necessary).

###### Are the evolution ritual rules exactly the same as in the table in the subject?

**_Broadcast_**

##### Pay attention to the broadcast messages.

```
broadcast <text>
```

###### Can you confirm that in order to send a message, the client must send the command above to the server.

##### After the client send the command to the server, he will send to all its clients the following line:

```
message <K>,<text>
```

###### Can you confirm that the server replied the above message with K indicating the square where the sound comes from?

#### Bonus

###### +Is the visualizer in 3D?

###### +Is there a flag for a log mode when running the server?

###### +Is there a flag for the number of resources and food density?

###### +Is there a flag for the seed in the server and client so that specific scenarios could be reproduced?
