#### Functional

> In order to run and hot reload the app either on emulator or device, follow the [instructions](https://docs.flutter.dev/get-started/test-drive?tab=androidstudio#run-the-app).

###### Does the app run without crashing?

##### Confirm if the app contains all the UI screens requested.

- A Main menu
- A Game board
- A waiting room

###### Does the app have all the UI screens requested?

###### Does the app have a multiplayer mode for real-time matches which enables players to challenge and play against each other??

###### Can players invite each other to a game or join a public game?

###### Does the app have a waiting room for players to wait for opponents and initiate game sessions?

###### When two players join a game session, does the player with the white pieces play first?

###### Does the app have a notification system to alert players when it's their turn to make a move?

##### Try performing illegal moves (e.g., moving a pawn 3 squares up, castling after moving the king, moving a rook diagonally).

###### Does the app prevent illegal moves?

###### After a move is played, does the other player receive the move?

##### Play the game until one player is mated.

###### Does the game end and the player who mated the other win?

###### When the game reaches a terminal state (checkmate, stalemate, or draw), does the app send a message to both players indicating that the game is over and the reason of being over?

###### Does the backend generate a unique identifier (UUID) for each game session when a player joins the waiting room, ensuring proper identification and management of individual game sessions?

#### Bonus

###### +Does the app allow players to restart the game?

###### +Is there a single-player story mode where players can engage in chess matches against AI opponents, progressing through a series of challenges or levels?

###### +Are there AI-powered bots that players can play against in offline mode or when waiting for online opponents?
