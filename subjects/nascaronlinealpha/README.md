## NascarOnlineAlpha

In this exercise, you will create a racing online game. The player will be able to race with another player, to see who is faster. When you feel stuck remember that once, a man couldn’t work out how to fasten his seatbelt. But then it suddenly clicked!

### Objectives

You will have to create a car game using an already existing project that you will have to download.

You will learn to analyze the code and understand what the developer did before you and adapt to extend the gameplay to continue the development of the game.

### Instructions

In order to start the project, you will have to open the Epic Games Launcher, switch to the Learn tab and search for the Vehicle Game project. Once you find it and download it, you can create a new Vehicle Game project and that will be the base of your project.

As the game is done, you only will need to make the multiplayer part of it. Basically, you will have to:

- As soon as the game starts it should ask for a name to assign to the player.
- Create a main menu map where the player can choose either to Host/Join a session.
  - If the player chooses to Host, the game should ask you for a server name, number of players allowed on the server, if the player wants to host via LAN or ONLINE and a host button that launches the player into a multiplayer Lobby.
    - Max. number of players should be 4 and the minimum should be 2.
    - If the server name is not provided, the player can not yet host the server and a warning message should appear, requesting to insert a name for the server.
  - If you press the Join button, the game should display to the player, a list of available sessions he/she can join. Also a refresh button should be available to search for sessions again.
    - A button with the server name in it should appear for each session found. By clicking on it, the player is directly transferred to the host lobby.
    - It must be presented to the player the option to do the search for LAN sessions or ONLINE sessions.
  - In the start menu, there should be also a button that enables the player to quit the game.
    - In every other menu, there should be a button that allows the player to go back.
- On the host lobby, before the creator of the session starts the game, the player can drive around the map.
  - A widget should pop on the screen saying “Waiting for the host to start” if you are not the host and “Click START to start” otherwise. Therefore the host of the server should have available a Start button.
- Once the game started, all players spawn on the map and the race starts after a countdown (A little delay may occur for the client but it quickly goes back to normal so don’t worry for that.)
- The first to finish the race has a text on the screen written “YOU WIN” and “YOU LOSE” for the other players.
- Once the player crosses the finish line, the player input is disabled (player can no longer move).
- Use the provided loading screen to hide all screen/player switching between maps and screens.

As for the multiplayer aspect, here are some hints that you will have to follow:

- Use the Advanced Session plugin to have more host options when creating and searching for a session.
- Use replicated Custom Events to share information.
- Use Game Instance to handle all widgets navigation.
- Use Game State and replicated variables to know who is the winner of the game. If you have more info to share for all players on the game, do it through Game State using server replication.
- Use "Has Authority" checks to know if it is the server or the client that is trying to execute an action.

When finished, your project should look like the executable example on the folder or the [“Expected Result” video](https://youtu.be/s56rio0bw0U).

> Don’t forget to zip up the project compile and save everything for peer correction.

### Bonus

As bonus for this exercise you can:

- create a restart race button when all players crossed the finish line.
- display the winner on every player screen when crossing the finish line.
- make players join a server by IP address, from another LAN network (two players connected to different networks, for example).
