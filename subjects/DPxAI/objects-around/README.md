## Objects Around

> Mindful AI mode

### Context

Think about all the objects around you: like your book or robot friend. Each of these objects has different properties that describe it. For example, a robot can have a type, weight, and operational status. In JavaScript, we use objects to group these related properties together, making it easy to manage and access this information.

### AI-Powered Learning Techniques

`Visualization Technique:`
This type of prompt encourages the AI to explain a concept using diagrams or visual representations to illustrate concepts.

Find the examples across the subject ;)

### Concepts

### JavaScript Objects

Objects in JavaScript are fundamental data structures used to group related values together. They are like a bag of values.

### Example

First, let's look at different types of variables:

```js
let type = "Talkative";
let weight = 120.5;
let isOperational = true;
```

Now, we can group them into an object. Objects are values too, so let's assign one to a robot variable:

```js
let robot = {
  type: "Talkative",
  weight: 120.5,
  isOperational: true,
};
console.log(robot); // This will display the object 'robot'
```

Here, the robot variable is declared, and its value type is an object.

### Object Literal Syntax:

Objects are defined using curly brackets {}.

```js
let emptyRobot = {}; // an empty object
```

### Properties

Objects consist of properties, each having a key and a value:

```js
let robot = {
  type: "Talkative", // 'type' is the key, 'Talkative' is the value
  weight: 120.5,
  isOperational: true,
};
```

Each property is separated by a comma ",". It's good practice to add a trailing comma to every property, though it's not required for the last one.

### Accessing Values

To access values in an object, use the dot notation. For example:

```js
let robot = {
  type: "Talkative",
  weight: 120.5,
  isOperational: true,
};

console.log(robot); // Logs the whole 'robot' object
console.log(robot.weight); // Logs the weight of the robot (120.5)
```

You can use the property values just like any other values:

```js
let efficiency = 1.15; // Efficiency factor
let robot = {
  type: "Talkative",
  weight: 120.5,
  isOperational: true,
};

const adjustedWeight = robot.weight * efficiency;
console.log(adjustedWeight); // Logs the adjusted weight
```

#### **`Prompt Example`**:

- Can you help me visualize a JavaScript object with the following properties: `name`, `age`, `hasEnergy``?

### Instructions

#### Task 1:

Let's declare a variable `myRobot` which has an object as its value with 3 properties.

1. A `name` property of `myRobot`'s name as a String
2. An `age` property of `myRobot`'s age as a Number
3. A `hasEnergy` property as a Boolean indicating if `myRobot` has dangerous features

#### Task 2:

We will provide a `robot` variable of type object just like the one you did in the previous task.

Your job will be to decompose each property into its own variable:

- Define a `name` variable with the value of the `name` property of the `robot` variable.
- Same for `age`.
- And same for `hasEnergy`.
