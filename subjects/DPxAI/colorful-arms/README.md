## Colorful legs

> JSPowered Mode

### Context

Your robot is cool, but you can make it even more striking! Or at least, make it show off its colorful legs to everyone!

> Follow the instructions, ask your peers if you are stuck and stay motivated because you are close to the goal!
> Follow every hint you have in the subject!
> Continue on the code from last exercise, and change the file names!

### Instructions

In this exercise, you will use your gained knowledge to make your robot's legs change colors!

#### Task 1:

Let's put another button in the top right corner of the page with an id ``leg-color``, that will change the legs color. Add it in the HTML structure:

```html
 <button id="leg-color">Change robot's leg colors</button>
```

Add the button style in the CSS file:

```css
#leg-color
{
    position: fixed;
    top: 240px;
    right: 30px;
    padding: 10px;
    margin-bottom: 20px; 
    z-index: 2;
}
```

#### Task 2:

In the JS file, like in the previous exercise, get the following elements:

- The button with the ID ``leg-color``.
- The left and right leg elements with the IDs ``leg-left`` and ``leg-right``.
-  Apply the random color to both legs by changing their    ``backgroundColor``.

At this point, you know how it goes! make those legs vibrate with beautiful colors!

### Code Example:

```js
// Select the button with id 'leg-color'
//...Here

// Select the left and right legs elements
//...Here

// Your function that gets triggered when clicking the new button

const handleChangeLegColor = (event) => {
  // Generate a random color
  const randomColor = `#${Math.floor(Math.random() * 16777215).toString(16)}`;

  // Apply the random color to both legs
  //...Here
};

legColorButton.addEventListener("click", handleChangeLegColor);
```

**``Prompt Example:``**

- "What mistakes should I prevent while changing the background color of multiple elements in JavaScript?"
