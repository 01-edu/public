## Using Filter

### Instructions

- Create a function `filterShortStateName` that takes an array of
  strings and that returns the ones with less than 7 characters.

> Example: `'Iowa'` only contains 4 characters

- Create a function `filterStartVowel` that takes an array of strings
  and that returns only the ones that start with a vowel (a,e,i,o,u).

> Example: `'Alabama'` starts with a vowel

- Create a function `filter5Vowels` that takes an array of strings
  and that returns only the ones which contain at least 5
  vowels (a,e,i,o,u).

> Example: `'California'` contains at least 5 vowels

- Create a function `filter1DistinctVowel` that takes an array of
  strings and that returns only the ones which vowels are of only
  one distinct one (a,e,i,o,u).

> Example: `'Alabama'` only contains 1 distinct vowels `'a'`.

- Create a function `multiFilter` that takes an array of
  objects and that returns only the ones which:

- the key `capital` contains at least 8 characters.
- the key `name` does not start with a vowel
- the key `tag` has at least one vowel.
- the key `region` is not `'South'`

> Example of an array of objects matching the criterias:

```js
[
  { tag: 'CA', name: 'California', capital: 'Sacramento', region: 'West' },
  { tag: 'PA', name: 'Pennsylvania', capital: 'Harrisburg', region: 'Northeast' }
]
```

#### Special instruction

The goal of this exercise is to learn to use `filter`, as such all your
solution **MUST** use `filter`

### Notions

- [devdocs.io/javascript/global_objects/array/filter](https://devdocs.io/javascript/global_objects/array/filter)
