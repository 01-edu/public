## Smooth Operator

![sade](https://user-images.githubusercontent.com/231748/112029913-a0101900-8b31-11eb-8f59-cd7d68d7269b.jpg)

### Math Operators

They are other operators than assignation, for now let's focus on the one you
probably already know:

- `+` Addition
- `-` Substraction
- `/` Division
- `*` Multiplication

Those operators are used the same way we would write them in math:

```js
console.log(5 + 7) // -> 12
console.log(5 * 5) // -> 25
console.log(7 - 5) // -> 2
console.log(9 / 3) // -> 3
```

Operators are evaluated using classic priority:

```js
console.log(1 + 5 * 10) // -> 51
```

you can use parens `()` to enforce priority:

```js
console.log((1 + 5) * 10) // -> 60
```

And they are resulting in a value, as such they can be assigned to variables:

```js
let halfMyAge = 33 / 2
let twiceMyAge = 33 * 2
```

### Instructions

Your code must use the given variable `smooth` as our initial value

You will declare a few variables:

- `lessSmooth` that is just `1` less than `smooth`
- `semiSmooth` that is the half the amount of `smooth` _(it's still pretty
  smooth)_
- `plus11` that is `smooth` plus `11`
- `ultraSmooth` that is the square of smooth _(now that's smooth !)_

> BGM:
> [Sade - Smooth Operator - Official - 1984](https://www.youtube.com/watch?v=4TYv2PhG89A)
