## First wink

> JSPowered Mode

### Context

You're making fantastic progress! You've reached a crucial stage in your journey where you’ll learn how to make your robot alive using the power you gained, JavaScript basics.

Don't worry if things feel a bit challenging—that's part of the process! Just remember to take it step by step and never forget what you learned on using Generative AI along with your peers to seek clarifications, and you'll be amazed at how quickly you'll see results.

> Let's use all your power of learning and searching to make your robot alive!

#### Reminder:

- Before you start coding, it's okay to take a moment to think about what you want to do. You can even write out a rough plan or "PseudoCode" to help organize your thoughts. It makes the coding part much easier!

> We can mention thing you do not know; but by this time, you know what to do! Search for it, ask your peers and use clever prompts ;)

- **You need to continue on the HTML, CSS, JS code you submitted for the exercise `first-move`, but with an empty JavaScript file and do not forget to change the name of the linked files to the name of this exercise!**

### Resources

We provide you with some content to get started smoothly, check it out!

- [Video](https://www.youtube.com/watch?v=m34qd7aGMBo&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=13) on ``querySelector``
- [Video](https://www.youtube.com/watch?v=ydRv338Fl8Y) DOM JS - Add an ``event listener`` to an element
- [Video](https://www.youtube.com/watch?v=4O6zSVR0ufw&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=15) DOM JS - Set an ``element's properties``
- [Video](https://www.youtube.com/watch?v=amEBcoTYw0s&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=21) DOM JS - ``classList``: toggle, replace & contains
- [Video](https://www.youtube.com/watch?v=pxlYKvju1z8&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=16) DOM JS - Set an element's ``inline style``
- [Memo DOM JS](https://github.com/nan-academy/js-training/blob/gh-pages/examples/dom.js)

### Instructions

#### Task 1:

Let's put a button on the top right corner of the page, that will toggle (close or open) the left eye when clicked.

Add it in the HTML structure:

```js
<button>Click to close the left eye</button>
```

And add the style in the CSS file:

```css
button {
  z-index: 1;
  position: fixed;
  top: 30px;
  right: 30px;
  padding: 20px;
}
```

#### Task 2:

Select the button in your JavaScript file that will allow the user to control the robot’s left eye.

```js
// Select the button element using its ID so we can interact with it in our JavaScript

//Example of selecting a button called myButton
const myButton = document.querySelector("button");
```

**`Prompt Example:`**

- "How do I use ``querySelector`` to select an HTML element by its ID?"

#### Task 3:

Write a function that will be triggered when the button is clicked.

This function will make the robot "wink" by toggling the `eye-closed` class on the left eye and change the `button` text based on the current state of the eye.

- It changes the text content of the button: if the eye is open, write "Click to close the left eye", if the eye is closed, write "Click to open the left eye".

- It toggles the class eye-closed in the ``classList`` of the eye-left HTML element.

It changes the background color of the eye-left: if the eye is open, to "red", if the eye is closed, to "black"

```js
const button = document.querySelector('button')

const handleClick = (event) => {

  // Select the left eye element by its ID
    const myDiv = ...

  // Check if the eye is currently closed by looking at its background color
  if (...) {
    // If the eye is closed, open it and update the button text

  } else {
    // If the eye is open, close it and update the button text
  }
    // Toggle the 'eye-closed' class on the 'eye-left' div
};

// register the event:
button.addEventListener('click', handleClick)
// here we ask the button to call our `handleClick` function
// on the 'click' event, so every time it's clicked
```

**`Prompt Examples:`**

- "As a beginner, explain to me what is ``querySelector`` in JavaScript, and how do I use it to select an HTML element by its ID or class?"

- "As a beginner, explain to me how can I change the text content of an ``HTML element`` using JavaScript?"

- "As a beginner, explain to me how do I use ``addEventListener`` to make a button respond to a click event in JavaScript?"

### Expected output

[This](https://www.youtube.com/watch?v=Wkar5SmswDo) is what you should see in the browser.
