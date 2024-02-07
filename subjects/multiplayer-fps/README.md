## multiplayer_fps

### Instructions

Write your own version of the game [maze wars](https://www.youtube.com/watch?v=5V5X5SbSjns). You should recreate all the elements of the game, but you have freedom to implement the user interface.

### Objectives

#### User Interface

The game should present a specific User Interface, in which the minimum requirements are:

- A mini map where the player can see his own position and the whole "game world".
- The graphics of the game (walls and other players) should be similar to the original game (see [maze_wars](https://www.youtube.com/watch?v=5V5X5SbSjns) for more details)
- Finally you have to display the frame rate of the game on the screen.

#### Architecture

- Implement a client-server architecture where clients connect to a central server to play the game.
- Your implementation should allow one client and the server to run on the same machine, with other clients connecting from different machines.
- Use the UDP protocol to enable the communication between the clients and the server.
- The game should have at least 3 levels with increasing difficulty (with difficulty we mean, making the maze harder, with more dead ends).

You will have to develop the game server and also a client application:

- The server must accept as much connections as possible (the minimum should be 10).
- When the client is initialized, the game should ask for:
  - The IP address of the server, allowing the client application to connect to any server.
  - A username for identification.

After providing the above information, the game should start and open the graphical interface, allowing the player to join and start playing the game.

Example:
Assuming that you can connect to a server in the same computer.

```console
$ cargo run
Enter IP Address: 198.1.1.34:34254
Enter Name: name
Starting...
$
```

#### Performance

The game should always have a frame rate above 50 fps (frames per second).

### Bonus

As bonus for this project here are some ideas:

- Implement a level editor to allow players to create their own mazes.
- Implement an algorithm that generates automatically new mazes.
- Implement A.I. players to allow playing the game without having to wait for more people to join to the server.
- For the basic implementation you can initialize the game from the command line. As a bonus you can implement the initialization of the game as part of the graphical interface and save a history of the hosts with an alias so it's easier to reconnect to known servers.

This project will help you learn about:

- GUI applications
- Game mechanics
- [UDP protocol](https://searchnetworking.techtarget.com/definition/UDP-User-Datagram-Protocol)
