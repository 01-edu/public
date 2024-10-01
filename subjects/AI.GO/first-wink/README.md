## First wink

> JSPowered Mode

### Context

You're making fantastic progress! You've reached a crucial stage in your journey where you’ll learn how to make your robot alive using the power you gained, JavaScript basics.

Don't worry if things feel a bit challenging—that's part of the process! Just remember to take it step by step and never forget what you learned on using Generative AI along with your peers to seek clarifications, and you'll be amazed at how quickly you'll see results.

> Let's use all your power of learning and searching to make your robot alive!

#### Reminder:

- Before you start coding, it's okay to take a moment to think about what you want to do. You can even write out a rough plan or "PseudoCode" to help organize your thoughts. It makes the coding part much easier!

> We can mention thing you do not know; but by this time, you know what to do! Search for it, ask your peers and use clever prompts ;)

- **You need to continue on the HTML, CSS, JS code you submitted for the exercise `first-move`, but with an empty JavaScript file. Do not forget to change the name of the linked files to the name of this exercise as well as following the new instructions!**

### Resources

We provide you with some content to get started smoothly, check it out!

- [Video](https://www.youtube.com/watch?v=m34qd7aGMBo&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=13) on `querySelector`
- [Video](https://www.youtube.com/watch?v=ydRv338Fl8Y) DOM JS - Add an `event listener` to an element
- [Video](https://www.youtube.com/watch?v=4O6zSVR0ufw&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=15) DOM JS - Set an `element's properties`
- [Video](https://www.youtube.com/watch?v=amEBcoTYw0s&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=21) DOM JS - `classList`: toggle, replace & contains
- [Video](https://www.youtube.com/watch?v=pxlYKvju1z8&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=16) DOM JS - Set an element's `inline style`
- [Memo DOM JS](https://github.com/nan-academy/js-training/blob/gh-pages/examples/dom.js)

### Instructions

#### Task 1:

Let's put a button on the top right corner of the page with the `id` set to "eye-btn", that will toggle (close or open) the left eye when clicked.

Add it in the HTML structure:

```js
<button id="eye-btn">Click to close the left eye</button>
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

Select the button in your JavaScript file by its `id`. That will allow the user to control the robot’s left eye.

```js
// Select the button element using its ID so we can interact with it in our JavaScript

//Example of selecting a button called btn-example
const myButton = document.getElementById("btn-example");
```

**`Prompt Example:`**

- "How do I use `getElementById` to select an HTML element by its ID?"

#### Task 3:

Write a function that will be triggered when the button is clicked.

This function will make the robot "wink" by toggling the `eye-closed` class on the left eye and change the `button` text based on the current state of the eye.

1- It changes the text content of the button: if the eye is open, write "Click to close the left eye", if the eye is closed, write "Click to open the left eye".

2- It toggles the class `eye-closed` in the `classList` of the `eye-left` HTML element.

3- It changes the background color of the `eye-left`: if the eye is open, to "red", if the eye is closed, to "black"

**Code Example:**

```js
const button = document.getElementById('eye-btn')

const handleClick = (event) => {

   // Select left eye by its ID and assign it to the variable eyeLeft
    const eyeLeft = ...

  // Check if the eye is currently closed by looking at the eyeLeft background color, if it's 'black', that means it's closed.
  if (...) {
    /*
     - Add to the 'button' element the text: "Click to close the left eye"
     - Change the 'eyeLeft' element background color to red
    */

  } else {
    /*
    If the eye is open:
    - Add to the 'button' element the text: "Click to open the left eye"
    - Change the 'eyeLeft' element background color to black
    */
  }
    // Toggle the 'eye-closed' class on the 'eye-left' div
    eyeLeft.classList.toggle("eye-closed")
};

/* Register the event:
 here we ask the button to call our `handleClick` function
on the 'click' event, so every time it's clicked
*/
button.addEventListener('click', handleClick)
```

### Expected result

You can see an example of the expected result [here](https://youtu.be/IQ6-3X3JBss)

**`Prompt Examples:`**

- "As a beginner, explain to me how can I change the text content and the background color of an `HTML element` using JavaScript?"

- "As a beginner, explain to me how do I use `addEventListener` to make a button respond to a click event in JavaScript?"
