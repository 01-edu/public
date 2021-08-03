## scoreboard

### Objectives

You must follow the same [principles](../README.md) as the first subject.

For this project you must take into account:

- The usage of **scoreboards**
- Creation of a **go API service** to save the data from the game in JSON formate
  - The API should accept POST and GET request from the client side, this being the scoreboard data

### Instructions

Just like the first subject you must respect performance.

In order to tell apart every score it should be requested a name when the player ends the game.

After every game, either you win or lose, a scoreboard should be shown with the five highest scores of every game made.

The scoreboard must display the **position**, **name**, **score**, **time** in minutes and paginate the results with the rest of the scores. You also should give to the client the percentage and the position in the scoreboard.

For example:

```console
Congrats O.J, you are in the top 6%, in the 2nd position.

Rank| Name | Score  | time
---------------------------
1st | Kave | 233254 | 12:01
2nd | A.J. | 222555 | 03:00
3rd | O.J. | 14356  | 05:40
4th | -.-  | 13663  | 02:34
5th | iris | 2354   | 00:40

     <- Page 1/50 ->
```

The scoreboard should be ordered by descending order, so the player with the most points should appear on first place.

You will have to create a **go API service**, where you can load the data (POST), and request it (GET). This service will store the information of each play (name, score and time) in a JSON formate and returns all the information when requested.
The JSON can be organized as you wish.

Here is an example:

```json
[
  {
    "name": "O.j.",
    "Rank": 3,
    "Score": 14356,
    "time": "05:40"
  },
  {
    "name": "Kave",
    "Rank": 1,
    "Score": 233254,
    "time": "12:01"
  },
  ....
]
```

### Allowed Packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed.
- [Gorilla](https://pkg.go.dev/github.com/gorilla/websocket) websocket

This project will help you learn about:

- Creating and using an API
  - POST requests
  - GET requests
- JSON
- Sorting algorithms
