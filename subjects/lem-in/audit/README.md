#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

##### Try to run `go run . example00.txt`.

###### Is the program able to read the ant farm in this standard input?

###### Does the program accept only the commands `##start` and `##end`?

##### Try running the program with a valid example of your choice.

```console
number_of_ants
the_rooms
the_links

Lx-y
Lx-y Lx-y
Lx-y Lx-y Lx-y
Lx-y Lx-y
Lx-y
```

###### Are the commands and the ants movements printed with the right format? One line per turn, `N` movements per turn, movements defined by `Lx-y` `x` being the ant and `y` being the room, as it shows above?

##### Try running the program with [example00](../examples/README.md).

```
$ go run . example00.txt
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
$
```

###### Does it present the quickest path possible with at most 6 turns?

##### Try running the program with [example01](../examples/README.md).

```
$ go run . example01.txt
10
##start
start 1 6
0 4 8
o 6 8
n 6 6
e 8 4
t 1 9
E 5 9
a 8 9
m 8 6
h 4 6
A 5 2
c 8 1
k 11 2
##end
end 11 6
start-t
n-e
a-m
A-c
0-o
E-a
k-end
start-h
o-n
m-end
t-E
start-0
h-A
e-end
c-k
n-m
h-n

L1-t L2-h L3-0
L1-E L2-A L3-o L4-t L5-h L6-0
L1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0
L1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t
L1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E
L4-end L5-end L6-end L7-m L8-k L9-e L10-a
L7-end L8-end L9-end L10-m
L10-end
$
```

###### Does it present the quickest path possible with at most 8 turns?

##### Try running the program with [example02](../examples/README.md).

```
$ go run . example02.txt
20
##start
0 2 0
1 4 1
2 6 0
##end
3 5 3
0-1
0-3
1-2
3-2

L1-3 L2-1
L2-2 L3-3 L4-1
L2-3 L4-2 L5-3 L6-1
L4-3 L6-2 L7-3 L8-1
L6-3 L8-2 L9-3 L10-1
L8-3 L10-2 L11-3 L12-1
L10-3 L12-2 L13-3 L14-1
L12-3 L14-2 L15-3 L16 -1
L14-3 L16-2 L17-3 L18-1
L16-3 L18-2 L19-3
L18-3 L20-3
$
```

###### Does it present the quickest path possible with at most 11 turns?

##### Try running the program with [example03](../examples/README.md).

```
$ go run . example03.txt
4
4 5 4
##start
0 1 4
1 3 6
##end
5 6 4
2 3 4
3 3 1
0-1
2-4
1-4
0-2
4-5
3-0
4-3

L1-1
L1-4 L2-1
L1-5 L2-4 L3-1
L2-5 L3-4 L4-1
L3-5 L4-4
L4-5
$
```

###### Does it present the quickest path possible with at most 6 turns?

##### Try running the program with [example04](../examples/README.md).

```
$ go run . example04.txt
9
##start
richard 0 6
gilfoyle 6 3
erlich 9 6
dinish 6 9
jimYoung 11 7
##end
peter 14 6
richard-dinish
dinish-jimYoung
richard-gilfoyle
gilfoyle-peter
gilfoyle-erlich
richard-erlich
erlich-jimYoung
jimYoung-peter

L1-gilfoyle L3-dinish
L1-peter L2-gilfoyle L3-jimYoung L5-dinish
L2-peter L3-peter L4-gilfoyle L5-jimYoung L7-dinish
L4-peter L5-peter L6-gilfoyle L7-jimYoung L9-dinish
L6-peter L7-peter L8-gilfoyle L9-jimYoung
L8-peter L9-peter
$
```

###### Does it present the quickest path possible with at most 6 turns?

##### Try running the program with [example05](../examples/README.md).

```
$ go run . example05.txt
9
#rooms
##start
start 0 3
##end
end 10 1
C0 1 0
C1 2 0
C2 3 0
C3 4 0
I4 5 0
I5 6 0
A0 1 2
A1 2 1
A2 4 1
B0 1 4
B1 2 4
E2 6 4
D1 6 3
D2 7 3
D3 8 3
H4 4 2
H3 5 2
F2 6 2
F3 7 2
F4 8 2
G0 1 5
G1 2 5
G2 3 5
G3 4 5
G4 6 5
H3-F2
H3-H4
H4-A2
start-G0
G0-G1
G1-G2
G2-G3
G3-G4
G4-D3
start-A0
A0-A1
A0-D1
A1-A2
A1-B1
A2-end
A2-C3
start-B0
B0-B1
B1-E2
start-C0
C0-C1
C1-C2
C2-C3
C3-I4
D1-D2
D1-F2
D2-E2
D2-D3
D2-F3
D3-end
F2-F3
F3-F4
F4-end
I4-I5
I5-end

L1-A0 L4-B0 L6-C0
L1-A1 L2-A0 L4-B1 L5-B0 L6-C1
L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0
L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1
L2-end L3-A2 L4-D3 L5-D2 L6-I4 L7-A1 L8-A0 L9-E2
L3-end L4-end L5-D3 L6-I5 L7-A2 L8-A1 L9-D2
L5-end L6-end L7-end L8-A2 L9-D3
L8-end L9-end
$
```

###### Does it present the quickest path possible with at most 8 turns?

##### Try running the program with [badexample00](../examples/README.md).

```
$ go run . badexample00.txt
ERROR: invalid data format
$
```

###### Does it present the right result as above?

##### Try running the program with [badexample01](../examples/README.md).

```
$ go run . badexample01.txt
ERROR: invalid data format
$
```

###### Does it present at least the result above?

##### Try running the program with [example06](../examples/README.md) and with 100 ants.

###### Is the real time less than 1.5 minutes?

##### Try running the program with [example07](../examples/README.md) and with 1000 ants.

###### Is the real time less than 2.5 minutes?

##### Try running the program with a valid example of your choice.

###### Are the ants alone in each room?

###### Is each tunnel only used once per turn?

##### Try running the program with a valid example of your choice.

###### At the end can you confirm that all the ants are in the `##end` room?

##### Try to run the program with an example of your choice multiple times.

###### Are the results always correct?

###### Does the program handles all the error cases with a message?

###### Is the solution making the ants cross the farm from `##start` to `##end` properly?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Does the program present an ant farm visualizer?

###### +Is it possible to see the ants moving?

###### +Is the error output more specific? (example: `"ERROR: invalid data format, invalid number of Ants"` or `"ERROR: invalid data format, no start room found"`)

###### +Is the visualizer in 3D?

#### Basic

###### +Does the project run quickly and effectively (favoring of recursive, no unnecessary data requests, etc.)?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Does the code obey the [good practices](../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
