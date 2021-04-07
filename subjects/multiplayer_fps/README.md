## multiplayer_fps

### Instructions

Write your own version of the game [maze wars](https://www.youtube.com/watch?v=5V5X5SbSjns) you should recreate all the elements of the game, but you have freedom to implement the user interface.

#### Minimum Requirements for the User Interface:

- A mini map where the player can see he's own position and the whole "game world".

- The graphics of the game: the walls and the other players (see [maze_wars](https://www.youtube.com/watch?v=5V5X5SbSjns) for more details)

- Finally you have to display the frame rate of the game on the screen.

- The game should have at least 3 levels with increasing difficulty.

#### Architecture of the application

- A peer-to-peer network that will allow other players to join your server and play against each other.

- Your implementation should allow one client and the server to run in the same machine and all the other clients to connect from different computers.

- Use the UDP protocol to enable the communication between the clients and the server.


#### You have to develop the game server and also a client application:

- The server must accept as much connections as possible (the minimum should be 10).

- When the client is initialized it should ask for:

  - The IP address of the server, allowing the client application to connect to any server.

  - After this, it should ask also for a name.

  - Example:
  Assuming that you can to connect to a server in your same computer.

```console
path/to/client $ cargo run
Enter IP Address: 198.1.1.34:34254
Enter Name: name
Starting...
path/to/client $
```

#### Performance Requirements

- The game should always have a frame rate above 50 fps (frames per second).

#### Bonus

- Implement a level editor to allow player to create they're own mazes.

- Implement an algorithm that generates automatically new mazes.

- Implement A.I. players to allow playing the game without having to wait for more people to join to the server.

- For the basic implementation you can initialize the game from the command line. As a bonus you can implement the initialization of the game as part of the graphical interface and save a history of the hosts with an alias so it's easier to reconnect to known servers.
