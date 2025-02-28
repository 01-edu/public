## scoreboard

### Objectives

You must follow the same [principles](../README.md) as the first subject.

For this project you must implement:

- **Scoreboard**
- **Go API service** that allows saving the data from the game in JSON format
  - The API should accept POST and GET requests from the client side for the scoreboard data

### Instructions

Just like with the first subject you must respect performance.

In order to tell scores apart the player's name should be requested before submitting the score after the game has ended.

After every game, whether the player wins or loses, a scoreboard should be shown with the five highest scores.

For each score, the scoreboard must display **position**, **name**, **score**, **time** in minutes. Scores must be paginated. For the submitted score, the position percentile must be displayed.

For example:

```console
Congrats O.J, you are in the top 6%, on the 2nd position.

Rank| Name | Score  | time
---------------------------
1st | Kave | 233254 | 12:01
2nd | A.J. | 222555 | 03:00
3rd | O.J. | 14356  | 05:40
4th | -.-  | 13663  | 02:34
5th | iris | 2354   | 00:40

     <- Page 1/50 ->
```

The scores should be displayed in descending order, so that the player with the most points appears on the first place.

You will have to create a **go API service**, where you can load the data (POST), and request it (GET). This service will store the information of each play (name, score, and time) in JSON format and return all the information when requested.
The data can be organized in JSON as you wish.

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

This project will help you learn about:

- Creating and using an API
  - POST requests
  - GET requests
- JSON
- Sorting algorithms
