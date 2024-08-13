## First Function

> Mindful AI mode

### Context

Your robot will need to start executing custom tasks, and to do that, you'll need to program its `functions`. Functions are like commands that tell robot exactly what to do.

You can create your own functions to give your robot unique abilities and make it more efficient!

### AI-Powered Learning Techniques

`Reflective Practice Technique:`
This type of prompt encourages learners to reflect on the concepts they’ve just learned, reinforcing their understanding by applying the concepts in different contexts or scenarios.

Find the examples across the subject ;)

### Concepts

Remember this example of function call ?

```js
//       ↙ identifier, like variables
console.log("Hello There !"); // <- function call happening here
//          ↖ open paren + argument + close paren
```

There, we saw how to call and use "built-in" functions.

Here, now, we are going to learn how to declare our owns. This will gives us even more freedom to build our own logic.

### Declaring a function

Here, we’ll learn how to declare a function in a `variable`. This gives your robot more freedom to perform custom tasks.

We'll use the `ArrowFunctionExpression` syntax to declare a function:

```js
//    ↙ normal variable     ↙ beginning of the scope of the function
let myFirstFunction = () => {
  //                    ↖ parens () for arguments and the arrow => for syntax
}; // <-end of the scope of the function
```

### Calling a Function

Once declared, you can call the function using the parentheses `()`:

```js
myFirstFunction(); // This will call the function, but nothing happens yet
```

### Adding Instructions

Very much like an if statement a function has a scope. The scope in between the curly braces `{}` is where the action happens. Let's add something to the scope of our function:

```js
let myFirstFunction = () => {
  console.log("Robot is now active!");
};
```

Now, when you call `myFirstFunction()`, Robot will log a message in the console:

```js
myFirstFunction(); // Output: "RoboGuard is now active!"
```

> We actually declared, then called the function and gave it this single instruction, `console.log('Robot is now active!')`.

#### **`Prompt Example`**:

- "How do you call a function in JavaScript, and what happens if the function contains no instructions?"

#### **`Video Ressource`**:

- [Functions in depth.](https://www.youtube.com/watch?v=0lp0d-Uxy4o)

### Instructions

#### Task 1:

You are a robot made by a scientist called Rick and you want to know your purpose.

- Declare a function named `ask` that `log` 'What is my purpose ?' in the console
- Declare a function named `reply` that `log` 'You pass butter.' in the console
  Then first `call the ask` then `the reply` once, in that order.
