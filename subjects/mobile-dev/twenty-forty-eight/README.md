## twenty-forty-eight

### Introduction:

Let us imagine that today is the april of 2014 and we know that the game "2048" is not implemented yet, but will be within a month. As we are aware the game will be a success and it is not such a difficult project to implement. Your task is to implement the game before the author of the game does so.

The game shoud start with a 4x4 table. Every new tile must be valued as either 2 or 4. Start game with 3-4 tiles, which are located randomly on the board. By swiping, tiles are moved towards any of 4 directions as far as possible and if there are two tiles that collide with the same value, they merge into one tile with twice the value, and as a result, the score is updated.
After each move, new tile appears randomly in an empty slot.
You must retain the best score. Score is attained when tiles are merged, i.e. if tiles with a value of 4 are merged, add 4 to current score.

You can use colors of your choice, but colors of the original game are preferred.

Use any [animation](https://docs.flutter.dev/development/ui/widgets/animation) tile movement.

Game should stop when no legal movement is possible, and let user restart game. When game is over, update the best score if it is required.

<center>

<img src="./resources/2048.01.png?raw=true" style = "width: 840px !important; height: 420px !important;"/>

</center>
