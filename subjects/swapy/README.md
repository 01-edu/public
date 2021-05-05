## Swapy

Like we did with `Objects`, we can modify our arrays.

### replacing an `Array` value

Let's look at an example of code:

```js
let weekDays = [
  'Monday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
  'Sunday',
]

// Let's say I don't want Monday but a Second Sunday
weekDays[0] = 'Second Sunday'
```

In this example, we select the element at index `0` (with `weekDays[0]`) and
then assign it using the `=` (assign operator) the value `'Second Sunday'`

Now my array look like this:

```js
;[
  'Second Sunday',
  'Tuesday',
  'Wednesday',
  'Thursday',
  'Friday',
  'Saturday',
  'Sunday',
]
```

### Instructions

- You must replace the third element of the provided `replaceMe` array by the
  string `'great'`

Example:

```js
let replaceMe = ['pif', 'paf', 'pom']
//    expect -> ['pif','paf','great']
```

- You must swap the first and second element of the provided `swapMe` array.

Example:

```js
let swapMe = ['pif', 'paf', 'pom']
// expect -> ['paf','pif','pom'] (last element is untouched)
```

> You must modify the `swapMe` array, not create a new one !
