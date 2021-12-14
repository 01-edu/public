## lem-in

### Objectives

This project is meant to make you code a digital version of an ant farm.

Create a program `lem-in` that will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, `lem-in` will display the content of the file passed as argument and each move the ants make from room to room.

How does it work?

- You make an ant farm with tunnels and rooms.
- You place the ants on one side and look at how they find the exit.

You need to find the quickest way to get `n` ants across a colony (composed of rooms and tunnels).

- At the beginning of the game, all the ants are in the room `##start`. The goal is to bring them to the room `##end` with as few moves as possible.
- The shortest path is not necessarily the simplest.
- Some colonies will have many rooms and many links, but no path between `##start` and `##end`.
- Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no `##start` or `##end`, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message `ERROR: invalid data format`. If you wish, you can elaborate a more specific error message (example: `ERROR: invalid data format, invalid number of Ants` or `ERROR: invalid data format, no start room found`).

You must display your results on the standard output in the following format :

```console
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by `"name coord_x coord_y"`, and will usually look like `"Room 1 2", "nameoftheroom 1 6", "4 6 7"`.

- The links are defined by `"name1-name2"` and will usually look like `"1-2", "2-5"`.

Here is an example of this in practice :

```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```

Which corresponds to the following representation :

```console
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

### Instructions

- You need to create tunnels and rooms.
- A room will never start with the letter `L` or with `#` and must have no spaces.
- You join the rooms together with as many tunnels as you need.
- A tunnel joins only two rooms together never more than that.
- A room can be linked to an infinite number of rooms and by as many tunnels as deemed necessary.
- Each room can only contain one ant at a time (except at `##start` and `##end` which can contain as many ants as necessary).
- To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
- You will only display the ants that moved at each turn, and you can move each ant only once and through a tunnel (the room at the receiving end must be empty).
- The rooms names will not necessarily be numbers, and in order.
- Any unknown command will be ignored.
- The program must handle errors carefully. In no way can it quit in an unexpected manner.
- The coordinates of the rooms will always be `int`.
- Your project must be written in **Go**.
- The code must respect the [**good practices**](../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed.

### Usage

Example 1 :

```
$ go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
```

Example 2 :

```
$ go run . test1.txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
$
```

Example 3 :

```
$ go run . test1.txt
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3

L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
$
```

### Bonus

- As a bonus you can create an ant farm visualizer that shows the ants moving trough the colony.

  - Here is an usage example : `./lem-in ant-farm.txt | ./visualizer`

  - The coordinates of the room will be useful only here.

This project will help you learn about :

- Algorithmic
- Ways to receive data
- Ways to output data
- Manipulation of strings
- Manipulation of structures
