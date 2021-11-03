## Using Reduce

### Instructions

Create three functions :

- `adder` that receives an array and adds its elements.

- `sumOrMul` that receives an array and adds or multiplies its elements
  depending on whether the element is an odd or an even number, where:
    - even = multiply
    - odd = add

- `funcExec` that receives an array of functions and executes them.

All functions may or may not receive an extra argument that should be the
initial value for the functions execution.

Example:

```js
sumOrMul([1, 2, 3, 5, 8], 5)
  // -> ((((5 + 1) * 2) + 3) + 5) * 8
  // -> 160
````

#### Special instruction

The goal of this exercise is to learn to use `reduce`, as such all your
solution **MUST** use `reduce`

### Notions

- [devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)
