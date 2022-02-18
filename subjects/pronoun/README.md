## Pronoun

### Instructions

Create a function called `pronoun` that has a string as parameter. This function returns an object
that will have all the personal pronouns, present in the string, as keys. Each key will have a sub object with the
first word after each of the personal pronouns found in the string.
Also, a property `count` must be added, to the sub object, with the amount of occurrences of the pronoun.

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
