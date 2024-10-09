## Colorful arms

> JSPowered Mode

### Context

Your robot is already impressive, but you can make it even more eye-catching! Or at least, make it show off its colorful arms to the world!

> Follow the instructions, ask your peers if you are stuck and stay motivated because you are close to the goal!
> Follow every hint you have in the subject!
> Continue on the code from last exercise, and change the file names!

### Instructions

After finishing these tasks, you will be able make your robot's arms change to beautiful colors!

#### Task 1:

Let's put a third button in the top right corner of the page with an id `arm-color`, that will change the arms color. Add it in the HTML structure:

```html
<button id="arm-color">Change robot's arm colors</button>
```

Add the button style in the CSS file:

```css
#arm-color {
  position: fixed;
  top: 170px;
  right: 30px;
  padding: 10px;
  z-index: 2;
}
```

Also replace the ``button` element styles with this block:

```css
button {
  border-radius: 50px;
  padding: 10px;
  z-index: 1;
  position: fixed;
  top: 30px;
  right: 30px;
}
```

Do you see how prettier the buttons turned to be?

#### Task 2:

In the JS file, like in the previous exercise, get the following elements:

- The button with the ID `arm-color`.
- The left and right arm elements with the IDs `arm-left` and `arm-right`.

#### Task 3:

In the code example below, we show you the variable `randomColor` which stores a random color each time we use it.

- Create a function that applies the random color to both arms by changing their `backgroundColor`every time the function is called.
- Then add an `event listener` on `click event` for the button with ID `arm-color`.

### Code Example:

```js
// Select the button with id 'arm-color'
//...Here

// Select the left and right arm elements
//...Here

// Your function that gets triggered when clicking the new button

const handleChangeArmColor = (event) => {
  // Generate a random color
  const randomColor = `#${Math.floor(Math.random() * 16777215).toString(16)}`;

  // Apply the random color to both arms
  //...Here
};

// Attach the handleChangeArmColor function to the 'arm-color' button
armColorButton.addEventListener("click", handleChangeArmColor);
```

### Expected result

You can see an example of the expected result [here](https://youtu.be/viQymmWw6wo)

**`Prompt Example:`**

- "How do I change the background color of multiple elements in JavaScript?"
