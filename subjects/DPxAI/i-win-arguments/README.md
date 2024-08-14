## I Win Arguments

> Mindful AI mode

### Context

You made it to the last mission in getting your full power, to make your robot alive and fully functional!

The final step involves mastering the use of `arguments` in functions. By learning how to use and manage these `arguments` effectively, you can unlock the full potential of your robot and make it later truly come alive.

Let's find out more!

### AI-Powered Learning Techniques

`Code Chunking Technique:`

This type of prompt encourages you to break down larger pieces of code into smaller, digestible chunks.

Each chunk is explained individually, allowing you to understand the purpose and functionality of each part.

Find the examples across the subject ;)

### Concepts

### One Argument

We mentioned it before with methods, functions can take arguments. They are always in between parenthesis `()`.

Let's use the same examples that we used for function calls:

Remember this example of function call ?

```js
//       ↙ method
console.log("Hello There !"); //<-
//                 ↖ The String 'Hello There!' is
//                   the argument of console.log()
```

We are now going to adapt `myFirstFuntion` so that it takes one argument : `arg1`.

```js
let myFirstFunction = (arg1) => {
  //<-arg1 is inputed in between the parenthesis
  console.log(arg1); // arg1 can be use inside the scope of the function
  //            ↖   arg1 is "transfered" to be the arg of console.log()
}; //<-end of the scope of the function
```

Now if the function is called, it displays the output of `console.log(arg1)`.

```js
myFirstFunction("using my first arg"); // "using my first arg"
```

But let's say we want to change what the function logs. Now, instead of modifying `myFirstFunction` we just need to modify the `argument` in the `function call`.

```js
myFirstFunction("another arg"); // "another arg"
myFirstFunction("and another one"); // "and another one"
myFirstFunction("and one more"); // "and one more"
```

### More Arguments

We’ve seen how to add `one` argument to a function. Now, let’s learn how to add `two (or more)` arguments.

All we need to do to add a second argument `arg2` is to add a comma `,` after `arg1` and then add `arg2`.

```js
let myFirstFunction = (arg1, arg2) => {
  //<-arg1 and arg2 are inputed in between the parenthesis
  console.log(arg1, arg2);
  //            ↖   arg1 and arg2 are "transfered" to be the args of console.log()
};
// Now we call the function
myFirstFunction("first arg", "second arg");
// "first arg"
// "second arg"
```

For more args, you will need to simply repeat the same process! Comma `,` then the other argument and it goes on.

> Please note that you can name your arguments however you please. Just make sure that you reuse the proper name inside the scope of your function.

### Return value

In addition to accepting arguments, functions can also `return` values.

Return values are the `outputs` that a function provides after completing its task.

We are now going to adapt `myFirstFunction` so that it `returns` a value instead of just `logging it`.

```js
let myFirstFunction = (arg1) => {
  return arg1; // the function now returns the value of arg1
};
```

Now if the function is called, it returns the value of `arg1`:

```js
let result = myFirstFunction("using my first return");
console.log(result); // "using my first return"
```

But let's say we want to change what the function `returns`. Now, instead of modifying `myFirstFunction`, we just need to modify `the argument` in `the function call`.

```js
let anotherResult = myFirstFunction("another return");
console.log(anotherResult); // "another return"
```

#### **`Prompt Example`**:

- "Can you guide me through creating and using a JavaScript function that takes multiple arguments, starting from a basic function without arguments, then adding single and multiple arguments ?"

### Instructions

#### Task 1:

You are the general's aide responsible for transmitting communications to the other `RoboGuards`.

1- Create the `battleCry` Function:

- This function should take `one argument` and display it in the `console`.
- The battlefield is vast, so ensure that `the argument is uppercased` before displaying it.

2- Create the `secretOrders` Function:

- Sometimes, communications need to be given quietly.

- This function will do the `same` as `battleCry`, except `it will lowercase` the argument before sending it.

> hint: you remember methods?

#### Task 2:

As the leader of the RoboGuard forces, you're not just preparing for battle—you're also forming dynamic duos of robots to work together on special missions.

1- Create the `duos` Function:

- This function will take `two arguments`, representing the **names** of **two robots**.
- It will `log them` together with an **and** and an **exclamation mark**.

> Output's example: "robotOne and robotTwo!"

2- Create the `duosWork` Function:

- This function will take `three arguments`: the **names** of two robots and the **task** they will perform together.

- It will `log them` together in a sentence describing their task.

> Output's example: "robotOne and robotTwo are saying hi!

#### Task 3:

Rick's robot, knows his purpose. (Remember ? 'He passes butter.')

- Define the function `passButter` that returns the string 'The butter'.

** "Your hard work is paying off. The only limit to your impact is your imagination and commitment." – Tony Robbins**
