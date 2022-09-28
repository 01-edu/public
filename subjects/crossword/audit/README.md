##### Open the repository of the project and check the submitted files

###### Can you confirm that the `crosswordSolver.js` file is present and you can run the following command `node crosswordSolver.js` adding the following lines to the original `crosswordSolver.js`?
```js
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'alan', 'ciao', 'anta']

crosswordSolver(puzzle, words)
```

##### Try running the function with the arguments:
```js
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'alan', 'ciao', 'anta']
```

```
casa
i..l
anta
o..n
```

###### Does the function display the value above?

##### Try running the function with the arguments:
```js
const puzzle = 
`...1...........
..1000001000...
...0....0......
.1......0...1..
.0....100000000
100000..0...0..
.0.....1001000.
.0.1....0.0....
.10000000.0....
.0.0......0....
.0.0.....100...
...0......0....
..........0....`
const words = [
	'sun',
	'sunglasses',
	'suncream',
	'swimming',
	'bikini',
	'beach',
	'icecream',
	'tan',
	'deckchair',
	'sand',
	'seaside',
	'sandals'
]
```

```
...s...........
..sunglasses...
...n....u......
.s......n...s..
.w....deckchair
bikini..r...n..
.m.....seaside.
.m.b....a.a....
.icecream.n....
.n.a......d....
.g.c.....tan...
...h......l....
..........s....
```

###### Does the function display the value above?

##### Try running the function with the arguments:
```js
const puzzle = 
`..1.1..1...
10000..1000
..0.0..0...
..1000000..
..0.0..0...
1000..10000
..0.1..0...
....0..0...
..100000...
....0..0...
....0......`
const words = [
  'popcorn',
  'fruit',
  'flour',
  'chicken',
  'eggs',
  'vegetables',
  'pasta',
  'pork',
  'steak',
  'cheese',
]
```

```
..p.f..v...
flour..eggs
..p.u..g...
..chicken..
..o.t..t...
pork..pasta
..n.s..b...
....t..l...
..cheese...
....a..s...
....k......
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test mismatch between number of input words and puzzle starting cells
```js
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'casa', 'ciao', 'anta']
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test starting words higher than 2
```js
const puzzle = '0001\n0..0\n3000\n0..0'
const words = ['casa', 'alan', 'ciao', 'anta']
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test words repetition
```js
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['casa', 'casa', 'ciao', 'anta']
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test empty puzzle
```js
const puzzle = ''
const words = ['casa', 'alan', 'ciao', 'anta']
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test wrong format checks
```js
const puzzle = 123
const words = ['casa', 'alan', 'ciao', 'anta']
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test wrong format checks
```js
const puzzle = ''
const words = 123
```

```
Error
```

###### Does the function display the value above?

##### Try running the function with the arguments:
[comment]: <> Test multiple solutions 
```js
const puzzle = '2001\n0..0\n1000\n0..0'
const words = ['aaab', 'aaac', 'aaad', 'aaae']
```

```
Error
```

###### Does the function display the value above?

#### Bonus

###### +Is the project using a backtracking algorithm to solve the problem?