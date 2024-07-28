## Declare Everything

### Context

In this second quest of our adventure, you are going to power up with JavaScript. And in this exercise you'll use variables to control and manage robots in various scenarios. Let's dive in and get started with the basics of variables in JavaScript!

> Don't worry if you are learning many new concepts, just go with the flow!

### AI-Powered Learning Techniques

**Exploratory Questions Technique:**

This type of prompt encourages the AI to provide various explanations and examples, helping you explore different aspects of a concept.

> Find the examples across the subject ;)

### Directions

Values need a way to be identified, that's why we use variables. They add meaning to a value by pointing to it. It's like a **label**, a way to name things.

If we say `20`, it doesn't carry much meaning, _`20` what?_

Imagine we are talking about robots. You have 20 robot parts.

_Now that's a lot of parts!_

> We defined _what_ we have (robot parts) and its _value_ (20).

## Concepts

An `identifier` is used to define what it is, using this syntax:

```js
let robotParts = 20;
```

> 😱 Woa, what's all this?!

Let's explain each part:

### Keyword: `let`

Firstly, a keyword, `let`.

A keyword is a special word that JS knows about. It is used to tell the computer to perform a specific action.

`let` indicates that the script is defining a new variable.

### Identifier

After that, it needs a valid identifier.

In this case, it's `robotParts`. We choose something meaningful here.

A few rules to apply to make sure an identifier is valid:

- No space allowed (`robot parts` would be 3 distinct identifiers)
- Not starting with a number (that's reserved for number values)
- Not being a reserved keyword (for example, using `let`)
- No special characters

As such, we use what's called `camelCase`.

Note that in JS, it is a convention to not uppercase the first letter as this is reserved for special declarations. We won't go into details now.

```js
let robot parts = 20 // invalid because of spaces
let robotParts = 20 // Just right
```

#### **Prompt example**:

> "Explain how to use `let` to declare variables in JavaScript."

### Operator: `=`

The special character `=` is an **operator**, like in math, they are used to define specific operations. In this case, `=` defines the `assignment` operation. It means assigning a value to our variable. This is what **links** the chosen `identifier` with our `value`.

#### `Prompt example`:

"Why do we use camelCase for variable names in JavaScript?"

### Value

Lastly, a value, like the one you already know: `string`, `number`, and `boolean`.Full example with descriptive comments:

```js
// ↙ keyword        ↙ assignation operator
let comicBookTitle = "Tintin in Tibet";
//       ↖ identifier       ↖ the value (here a string)
```

#### **Prompt example**:

"As a beginner, what types of values can I assign to a variable in JavaScript?"

### Instructions

Declare two variables:

- Use the identifier `seven` with the value being a string of the number 7
- Use the identifier `seventySeven` with the value being a string of the number 77

```js
let seven = "7";
let seventySeven = "77";
```

---

> "When we first begin fighting for our dreams, we have no experience and make many mistakes. The secret of life, though, is to fall seven times and get up eight times." ― Paulo Coelho
