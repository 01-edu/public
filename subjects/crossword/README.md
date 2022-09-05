## crossword

### Instructions

Create the function `crosswordSolver` that is able to solve an empty crossword puzzle 

The function takes two arguments:
1. an empty puzzle, passed as a string and
2. a list of words to fill in the puzzle (no double words allowed)

The function must return the puzzle filled with the input words.

The empty puzzle will be a string with the following rules:
- each character will be either a number, a `.` or a `\n`;
- a number represent the number of words starting from the specific position and a `.` represent a space that does not need to be filled. 

If the puzzle or list of words provided as inputs does not guarantee a unique solution, or any other conditions stated above are not met, the function must return `'Error'`. 

### Examples

```js
console.log(crosswordSolver('2001\n0..0\n1000\n0..0', ['casa', 'alan', 'ciao', 'anta']))
// output: 'casa\ni..l\nanta\no..n'
```
