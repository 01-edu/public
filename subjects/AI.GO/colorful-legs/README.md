## Colorful legs

> JSPowered Mode

### Context

Your robot is cool! But you can make it even more striking! Or at least, make it show off its colorful legs to everyone!

> Follow the instructions, ask your peers if you are stuck and stay motivated because you are close to the goal!
> Follow every hint you have in the subject!
> Continue on the code from last exercise, and change the file names!

### Instructions

Just like you turned your robot's arms special in the last exercise, you will do the same now with the legs!

Follow the same

#### Task 1:

Let's put another button in the top right corner of the page with an id `leg-color`, that will change the legs color. Add it in the HTML structure:

```html
<button id="leg-color">robot's leg colors</button>
```

Add the button style in the CSS file:

```css
#leg-color {
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

- The button with the ID `leg-color`.
- The left and right leg elements with the IDs `leg-left` and `leg-right`.
- Apply the random color to both legs by changing their `backgroundColor`.

### Expected result

You can see an example of the expected result [here](https://youtu.be/QQ0GKGuXgBw)

**`Prompt Example:`**

- "What mistakes should I prevent while changing the background color of multiple elements in JavaScript?"
