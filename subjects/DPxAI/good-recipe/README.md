## Good Recipe

> Mindful AI mode

### Context

Welcome to this quest! I bet you are hungry after getting more and more Javascript power.

Are you ready to cook? But instead of cooking food, we'll be applying recipes to data through methods! In JavaScript, methods are special functions called from another value, allowing you to perform various operations.

Let's find out more!

### AI-Powered Learning Techniques

`Guided Practice Technique:`
This type of prompt encourages the AI to guide learners through hands-on practice exercises, providing immediate feedback and helping them build confidence in their skills.

Find the examples across the subject ;)

### Concepts

### Functions and Methods

In JavaScript, functions are like recipes that perform a task. Methods are special functions associated with objects. We are focusing on methods now.

### Example

```js
console.log("Robot is mixing ingredients");
```

### Multiple Arguments

Functions can take multiple arguments, like adding ingredients to a recipe:

```js
console.log("flour", 200); // both arguments will appear in your console
```

### Methods

Methods are functions called from another value. For example, toFixed is a method that formats numbers:

```js
let num = 10 / 3;
console.log(num.toFixed(2)); // -> '3.33'
```

### String Manipulation with Methods

Using the .slice method to cut parts of a string:

```js
let alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
let cutFirst = alphabet.slice(10); // 'KLMNOPQRSTUVWXYZ'
let cutLast = alphabet.slice(0, -3); // 'ABCDEFGHIJKLMNOPQRSTU'
let cutFirstLast = alphabet.slice(5, -6); // 'FGHIJKLMNOPQRS
```

### Splitting Strings

Splitting a string into an array of words using split:

```js
let sentence = "Add flour, sugar, and eggs";
let words = sentence.split(" "); // ['Add', 'flour,', 'sugar,', 'and', 'eggs']
```

### Splitting text into lines:

```js
let text = "Add flour\nMix well\nBake at 350°F";
let lines = text.split("\n"); // ['Add flour', 'Mix well', 'Bake at 350°F']
```

#### **`Prompt Example`**:

- "Show me step-by-step how to use the `.split` method to break the sentence 'Learning JavaScript is fun' into an array of words."

- "Show me how to use the `.toUpperCase` method to convert the string 'hello' to uppercase."

- "Show me how to use ``Math.max`` to find the biggest number in an array."

### Instructions

#### Task 1:

You need to find the oldest robot in the kitchen to determine the most experienced helper. 

We provide you with three robot objects assigned to their respective variables `martin`, `kevin`, and `stephanie`. Use the `Math.max` function on their age properties to find the oldest robot.

Declare an `oldestAge` variable that uses `Math.max` on the age properties of `martin`, `kevin`, and `stephanie`.

#### Task 2:

You need to slice some virtual vegetables. Using the `.slice` method and the provided alphabet variable, which is a string of all the characters in the alphabet:

- Declare a `cutFirst` variable that removes the first 10 characters of alphabet, simulating the removal of the first few slices.
- Declare a `cutLast` variable that removes the last 3 characters of alphabet, simulating the trimming of the last few slices.
- Declare a `cutFirstLast` variable that removes the first 5 characters and the last 6 characters of alphabet, simulating a precise cut in the middle.

#### Task 3:

You needs to prepare ingredients for a recipe, and part of the preparation involves transforming text. 

For this task, you will use the `toUpperCase` and `toLowerCase` methods on the provided variable message.

- Create a `noCaps` variable with the value of message but in lower case to simulate finely chopped ingredients.

- Create an `allCaps` variable with the value of message but in upper case to simulate ingredients being loudly announced.
