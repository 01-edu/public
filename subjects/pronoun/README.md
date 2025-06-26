## Pronoun

### Instructions

Create a function named `pronoun` that accepts a string as its parameter.
This function should return an object where each key is a personal pronoun found in the string. The value of each key should be a sub-object containing:

- a `word` array, which includes the first word that follows each occurrence of the pronoun (excluding other pronouns)
- a `count` property, indicating how many times the pronoun appears in the string.

> Note: If the word that follows a pronoun is also a pronoun, it should be ignored and not included in the word array.

Pronouns:

- i
- you
- he
- she
- it
- they
- we

#### Example

```js
const ex = 'Using Array Destructuring, you you can iterate through objects easily.'

{ you: { word: [ 'can' ], count: 2 } }

const ex = 'If he you want to buy something you have to pay.'

{
  he: { word: [], count: 1}
  you: { word: [ 'want', 'have' ], count: 2 }
}

```
