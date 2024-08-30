## The Smooth Operator

> Mindful AI mode

Unlike the song, smooth operators in JavaScript help you perform various calculations and manipulations with ease.

### AI-Powered Learning Techniques

**Step-by-Step Instruction Technique:**

This type of prompt encourages the AI to provide detailed, step-by-step instructions for learning new concepts.

> Find the examples across the subject ;)

## Concepts:

### Math Operators

In JavaScript, operators are symbols that perform operations on variables and values. Let's delve into the most common types of operators you'll encounter.

There are other operators other than assignment, for now let's focus on the one you
probably already know:

- `+` Addition
- `-` Subtraction
- `/` Division
- `*` Multiplication

Those operators are used the same way we would write them in math:

```js
console.log(5 + 7); // -> 12
console.log(5 * 5); // -> 25
console.log(7 - 5); // -> 2
console.log(9 / 3); // -> 3
```

Operators are evaluated using classic priority:

```js
console.log(1 + 5 * 10); // -> 51
```

you can use parentheses `()` to enforce priority:

```js
console.log((1 + 5) * 10); // -> 60
```

And they result in a value, so they can be assigned to variables:

```js
let halfMyAge = 33 / 2;
let twiceMyAge = 33 * 2;
```

#### **`Prompt example`**:

"Can you provide step-by-step examples of basic math operations in JavaScript?"

### Placeholders

JavaScript allows you to include expressions within strings using template literals. This is done using backticks ``(`)`` and the ``${}`` syntax to include expressions.

#### Example

```js
console.log(`5 + 10 = ${5 + 10} = 15`); // -> 5 + 10 = 15 = 15
```

**Note that it only works using:** the `` ` `` backtick, not the `"` or `'`
quotes.

#### **`Prompt example`**:

"Can you provide a step-by-step guide on how to use template literals to create a string that includes variable values in JavaScript?"

### Instructions

#### Task 1:

Your code must use the given variable `smooth` as our initial value

> When in doubt, always test your code with console.log() and the Run button.  
> But, when the platform gives you an already existing variable to manipulate, like the `smooth` variable here, if you want to use/display it, you have to do so with the submit button.  
> You'll then see the result in the code editor console output, as this variable is not available in `Run` button mode, but only in `Submit` button mode.

```js
console.log("smooth = ", smooth);
let lessSmooth = smooth - 5;
console.log("lessSmooth = ", lessSmooth);
```

You will declare a few variables:

- `lessSmooth` that is just `1` less than `smooth`
- `semiSmooth` that is half the amount of `smooth` _(it's still pretty
  smooth)_
- `plus11` that is `smooth` plus `11`
- `ultraSmooth` that is the square of smooth _(now that's smooth !)_

#### Task 2:

We will provide a variable `name` and `age`. They will be pre-declared by us.

Declare your robot's `presentation` variable of the string:

> `Hello, my name is` **name** `and I'm` **age** `years old`

But use placeholders to build the string you will put inside the `presentation`.  
Put the values of the provided variables `age` and `name` inside those placeholders.

---

> BGM:
> [Sade - Smooth Operator - Official - 1984](https://www.youtube.com/watch?v=4TYv2PhG89A)
