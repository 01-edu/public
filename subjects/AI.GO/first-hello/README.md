## First hello

> JSPowered Mode

### Context

Your robot winked, but you can make it talk too! Or at least, make it say his first hello to the world!

> Follow the instructions, ask your peers if you are stuck and stay motivated because you are close to your goal!
> Follow every hint you have in the subject!
> Continue on the code from last exercise, and change the file names!

### Instructions

Now that you know how to make your creation move, what about making it communicate its first words to the world?

After finishing these tasks, you will be able to display and hide `Hello World` in the `torso` of your robot by clicking on a second button.

#### Task 1:

Let's put a second button in the top right corner of the page, that will add some text when clicked. Add it in the HTML structure:

```html
<button id="speak-button">Click me to speak</button>
```

Add the button style in the CSS file:

```css
button#speak-button {
  top: 100px;
}
```

Also add this class to style the text we will add:

```css
.words {
  text-align: center;
  font-family: sans-serif;
}
```

#### Task 2:

In the JS file, like in the previous exercise, get the HTML `button` element with id `speak-button` and add an `event listener` on `click event`, triggering a `function` that will:

- Select the torso element with id="torso".
- Check if a div with the class words already exists inside the torso element.
- If it exists, remove the existing div.
- Otherwise:
  - Create a new HTML element of type div.
  - Set its text content to "Hello World".
  - Set its class name to words, as defined in the CSS.
  - Use the append method to add the new div inside the torso element.

### Code Example:

```js
// Select the button with id 'speak-button'

//...Here

//Your function that gets triggered when clicking the new button
const handleSpeakClick = (event) => {
  // Select the torso element where the text will be added or removed
  const body = document.querySelector("#torso");

  // Check if a div with the class 'words' already exists inside the torso
  const existingDiv = document.querySelector("#torso .words");

  if (existingDiv) {
    // If the "Hello World" text exists, remove it from the torso
  } else {
    // If the "Hello World" text does not exist, create and append it
    // Create a new div element
    // Add the 'words' class to the div
    // Set the text content to "Hello World!"
    // Append the new div to the torso element
  }
};

// Attach the handleSpeakClick function to the 'speak-button' button
//...Here
```

### Expected result

You can see an example of the expected result [here](https://youtu.be/PuyEdAA0wy4)

**`Prompt Example:`**

- "How can I add and remove text in an element with a button click in JavaScript?"
