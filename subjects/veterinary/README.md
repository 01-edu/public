## veterinary

### Instructions

There is a new veterinary in town!
To help him get introduced to your neighborhood, create an object `veterinary` with the following properties:

- an array of string `animalKnowledge` that will store all the animal species the new veterinary can treat,

- a function `canTreat` that will receive a string as an argument and will return a boolean indicating if the string is present in the veterinary `animalKnowledge` array,

- a function `respondClient` that will receive two string, the pet name and the animal species of the pet, and will return a string that will include `Yes` or `No` plus the pet name depending on whether the pet can be treated or not, as in the example below.

### Example

```js
veterinary.animalKnowledge.push('dog')
console.log(veterinary.canTreat('dog')) // true
console.log(veterinary.respondClient('Matias', 'dog')) // 'Yes, I can treat Matias'
console.log(veterinary.respondClient('Jack', 'parrot')) // No, I cannot treat Jack'
```

### Notions

- [this](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/this)
