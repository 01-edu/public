## Play with variables

> Mindful AI mode

### Context

If things get a little hectic at times, take the time to get closer to your peers so that you can think, share and move forward together.

> Keep Going!

### AI-Powered Learning Techniques

**Clarification Technique:**

This type of prompt encourages the AI to explain a concept in detail, helping you gain a deeper understanding.

> Find the examples across the subject ;)

## Concepts

### Escape characters

**Quote delimiters** can be one of the tricky things to deal with.

Since they are used for delimiting text, they need a trick to include them in
our text.

For example, we want a `'` _(single quote)_ in our text, but use them as
delimiters:

```js
console.log('I keep trying , I can't give up! ')
// too bad a single quote, ruined the quote, get it ?
```

The `\` _(backslash)_ is used for that:

Every time there is an _extra special_ character into your string, putting a `\`
in front of it will **escape** it and doing so will let JS understand you meant
the **literal** following character and not the delimiter, _or whatever else
the character normally means for a string_

```js
console.log("I keep trying , I can't give up! ");

// Output: I keep trying, I can't give up!
```

#### **`Prompt Example`**:

"As a beginner, how do I include special characters in a string in JavaScript? Give me simple examples too."

### Assign re-assign

Remember the `let` keyword used to declare new variables.

> Note that we can't have multiple variables with the same identifier otherwise JS wouldn't know which one is which.

Redeclaring a variable will crash !

But it is still possible to use the = (assignment operator) to change its value !

> Note that sometimes you may find variable declared with `const`. This means that the assignation is constant and can never be re-assigned !

> It is used to protect your code against errors, but you can always use `let` in its place..

> Also you may find online old code using var. We are trying to get rid of `var`s since 2015. It's ancient syntax and it was pretty problematic. Never use it! If you see code using it, try to find a more recent example. That one is outdated.

#### **`Prompt Example`**:

- "As a beginner, what is the difference between let and `const` in JavaScript?"
- "As a beginner, how do I reassign a value to an already declared variable in JavaScript?"

### Instructions

#### Task 1:

- Create a `escapeFromDelimiters` variable that includes all 3 quotes _(`` ` ``, `"` and
  `'`)_.

- Create a `escapeTheEscape` variable that includes a backslash _(`\`)_.

#### Task 2:

- The variable `power` has been declared and will be used during the tests.

- You must try to re-assign the power variable to the string value `levelMax`. But without re-declaring it !

---

> “How did I escape? With difficulty. How did I plan this moment? With
> pleasure.” \
> ― Alexandre Dumas, The Count of Monte Cristo
