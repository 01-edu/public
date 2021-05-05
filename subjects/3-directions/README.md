## 3 Directions

### Accessing an array value `[index]`

Since numbers are not valid identifiers we can not use the `.` operator to
access a value in an array, but we can use the `[index]` square brackets to
access a value.

Example:

```js
let numberList = [10, 20, 30, 40]
console.log(numberList[0]) // -> 10
console.log(numberList[3]) // -> 40
console.log(numberList[6]) // -> undefined
```

Since we start at `0`, `[0]` will get the value at the first index.

> note that if we try to access an element that doesn't exist we will get
> `undefined` as a value, just like non existing properties for an object.

#### Using the `.length` property

Another difference of the arrays is that they always keep track of how many
elements are inside them.

You can use the `.length` property to get this value:

```js
console.log([].length) // -> 0
console.log([1].length) // -> 1
console.log([1, 1, 1, 1].length) // -> 4
```

### Instructions

We provide you a variable `list` that contains some elements, you will have to
access them and assign their values to variables:

- a variable `first` of the first element of the `list`
- a variable `last` of the last element of the `list`
- a variable `kiss` of an array of 2 elements, the last and the first element of
  the `list`, in that order.

**Example:** if `list` is `[1,2,3]`

- `first` would be `1`
- `last` would be `3`
- `kiss` would be `[3, 1]`

>     🧑‍🎤 ............ Oh, ........... 🧑‍🎤
>     🎶 .. I wanna be last, yeah ... 🎶
>     🎵 . Baby let me be your last . 🎵
>     ✨ ... Your last first kiss ... ✨
>     ― One Direction
