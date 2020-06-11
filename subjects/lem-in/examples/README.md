## lem-in-examples

- Create your file and copy the examples into it.

### example00

```
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
```

### example01

```
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
```

### example02

```
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
```

### example03

```
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
```

### example04

```
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
```

### example05

```
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
```

### example06

```
100
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
```

### example07

```
1000
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
```

### badexample00

```
0
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
```

### badexample01

```
20
##start
0 2 0
1 7 0
2 13 0
3 18 0
##end
4 23 0
5 7 6
6 10 6
7 13 6
8 16 6
9 7 2
10 7 4
11 13 2
12 15 2
13 17 2
14 19 2
15 21 2
16 9 4
0-1
0-5
0-9
0-10
1-2
2-3
3-3
5-6
6-7
7-8
8-7
9-2
10-16
1-11
11-12
12-13
13-14
14-15
15-1
16-7
```
