## Filler

### The Game:

What is filler?

Filler is an algorithmic game which consists in filling a grid of a known size in advance with pieces of a random size and shape, provided by the `game_engine`. In this game, two robots fight against each other, one after the other on the Anfield.

Each player will receive a game piece to place on the Anfield with exactly one cell of overlapping with its previous territory (meaning the region you've already placed your pieces on the Anfield). The shape of robots territory must not exceed the area of the Anfield. You cannot overlap your opponent pieces, neither him can overlap yours.

If one player cannot place any piece on the Anfield, he stops there while the other player can keep on placing pieces in order to achieve the maximum score possible, each piece that the player can put correctly on the Anfield will give him more points.
The player who occupy the biggest surface on the Anfield wins.

### The Anfield:

The Anfield is a two-dimensional grid with an arbitrary number of rows and columns where the robots will fight. To launch the game, we will provide you with some initial Anfield that must be passed as an argument to the `game_engine`. If you want you can create your own Anfield, remember it must have the starting position for each of the robots.

Here is an example of an initial Anfield of 30 by 14 :

```
..............................
..............................
..$...........................
..............................
..............................
..............................
..............................
..............................
..............................
..............................
..............................
...........................@..
..............................
..............................

```

### The pieces

The pieces are managed randomly by the `game_engine`. You can’t predict their size or shape until the `game_engine` transmits them to your program. Here are some examples of possible pieces to give you an idea:

```
Piece 2 2:
.#
#.

Piece 5 4:
.##..
.##..
..#..
...#.

Piece 6 3:
.##...
###...
#..#..
```

### The Robots

The principle is simple, two robots compete on the Anfield, and they must place, in turns, the piece that the `game_engine` randomly gives them, earning points each time they place a piece correctly. The game ends when none of the two players can place a piece or if one of the two players return an unexpected error (timeout, segfault, memory issues and so on). If only one player can't place a piece or place it in an incorrect way his move is ignored but the game continue for the other player.

Your goal is to win against all our robots, except the `terminator`. You'll have to know the following:

- The `game_engine` will run your program and write the Anfield and the piece to place in the standard input.

- Each turn, the `game_engine` rewrites the Anfield and includes a new random piece to be placed.

- The first line the `game_engine` will send to the robot is it's player number in the following format `$$$ exec p<number> : [<player path>]`

- In order to place the game piece on the Anfield, the player will have to write it’s coordinates on the standard output.

- The following format must be used `X Y\n`.

- To be able to place a piece on the Anfield, it is mandatory that one, and only one cell of the shape (in your piece) covers the cell of a shape placed previously (your territory).

- You will collect points each time you place a piece.

- If you are the `Player 1` your program will be represented by `a` and `@`. If you are `Player 2`, your program will be represented by `s` and `$`.

- The lowercases (`s` or `a`) will highlight the piece last placed on the Anfield. At the following turns, that same piece will be represented by the symbols (`$` or `@`), as it won’t be the piece last placed anymore.

- If the solution is wrong the `game_engine` won't make the robot play anymore but the game continues for the other player.

- If there is a timeout or any other kind of unexpected error from the player the game stops and the player loses.

- If your robot can't place anymore peaces he should still return a result (even if invalid), our robots for example return `0 0\n`, when they can't place any more pieces.

### The Game_engine

The `game_engine` will be provided to you and it will run in a docker container.

Here are the flags that can be used:

```
  -f, -file string
       Path to map
  -p1, -player1 string
       Path to AI one
  -p2, -player2 string
       Path to AI two
  -q, -quiet
       Quiet mode
  -r, -refresh
       Throttling mode
  -s, -seed int
       Use a random seed number
  -t, -time int
       Set timeout in seconds (default 10)
```

#### Docker

For the filler project you must use Docker. You can read about docker basics in the [ascii-art-web-dockerize](../ascii-art-web/dockerize/README.md) subject.

Follow this instructions to get the game running:

The `docker_image` folder will be provided to you as a zip file, you can download it [here](https://assets.01-edu.org/filler/filler.zip).

Inside the folder you will find the `Dockerfile`, the `game_engine`, the `maps` and the `robots`.

In order to set everything up you will need to follow these steps:

- Go inside the folder and build the image `docker build -t filler .`
- Run the container `docker run -v "$(pwd)/solution":/filler/solution -it filler`. This instruction will open a terminal in the container, the directory `solution` will be mounted in the container as well.
- Your solution should be inside the `solution` directory so it will be mounted inside the container.
- You can code your solution from the host machine with your favorite editor but don't forget to compile it and run it from the container terminal.

Example of how you can run the game inside the container:

`./game_engine -f maps/map01 -p1 robots/bender -p2 robots/terminator`

### Usage

Here is an example of the first input that the `game_engine` will send to player 1:

```
$$$ exec p1 : [robots/bender]
Anfield 20 15:
    01234567890123456789
000 ....................
001 ....................
002 .........@..........
003 ....................
004 ....................
005 ....................
006 ....................
007 ....................
008 ....................
009 ....................
010 ....................
011 ....................
012 .........$..........
013 ....................
014 ....................
Piece 4 1:
.OO.
```

And his response:

```
7 2\n
```

### Unit Tests

You must implement unit tests within your `filler` project to ensure your robot's parsing logic and placement strategy are reliable. Specifically, your tests should:

- Verify **Input Parsing** by ensuring the Anfield grid and the incoming pieces are correctly converted into internal data structures from the standard input string.
- Verify **Valid Placement Logic** by testing that a potential coordinate correctly identifies if it has exactly one overlapping cell with existing territory and zero overlaps with the opponent.
- Test **Boundary Conditions** to ensure the robot never attempts to place a piece that extends beyond the rows or columns of the Anfield.
- Verify **Coordinate Output** by ensuring the final calculation is formatted exactly as `X Y\n` as required by the `game_engine`.

### Bonus

- Create a graphic visualizer for this project.
- Create a player that can beat our `terminator` robot.
