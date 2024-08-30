## Robots Harmony

> AI Synergy Mode

## Context

Welcome to your final milestone! After each of you has brought a robot to friend to life, it's time to reunite them together.

Now, all those incredible robots need to be displayed in perfect harmony in our interactive and visually stunning Robots `Gallery`.

As a team, your task is to combine your individual creations into a cohesive display.

This gallery won't just be functional, it will be a fun and visually appealing experience that highlights the creativity and collaboration behind your robots.

> You'll be working on this mission as a team, so effective communication and mutual support are key to bringing your robots together.

> Separate tasks equally, it will make the results better!

> Remember what you learned in asking AI for the precise explanation of concepts, while your team apply those concepts and do the actual fun work!

Go ahead, and let the world see the amazing robots you've created together!

## Setups:

**1- Code Editor:**

First , remember that you will not be using the platform's code editor. Instead, you can use this online [tool](https://jsfiddle.net/) to create and modify your code.

> Feel free to play with the settings and familiarize yourself with the collaborative environment!

> PS: Click on Run to see the result of your code!

**2- How to submit your code:**

Your project must be handed in by the group captain before the time runs out. This is done by publishing the code on the platform.

To do so, go to **gitea** by clicking on the logo at the top right of your platform dashboard. (on the captain's account)  
[![Capture-d-e-cran-2024-08-29-a-19-51-15.png](https://i.postimg.cc/zBRM6BMv/Capture-d-e-cran-2024-08-29-a-19-51-15.png)](https://postimg.cc/N9BprgQv)

Sign in if you are not already logged in.

You must create a repository to submit your work.  
[![Capture-d-e-cran-2024-08-29-a-19-57-32.png](https://i.postimg.cc/59Bt8f3q/Capture-d-e-cran-2024-08-29-a-19-57-32.png)](https://postimg.cc/vDmMsFCc)

Name it **robots-harmony** and check the box that says **Initialize Repository (Adds .gitignore, License and README).**

You can now create all your project files to this repository, using the button New file and finalize this step by clicking Commit changes  
[![Capture-d-e-cran-2024-08-29-a-20-00-08.png](https://i.postimg.cc/G2kqcB15/Capture-d-e-cran-2024-08-29-a-20-00-08.png)](https://postimg.cc/vgZWhZTr)

**3- Expected files:**

You need to submit the following files:

- `robots-harmony.html`
- `robots-harmony.css`
- `robots-harmony.js`

## Instructions

Theses tasks are representing the mandatory part for this raid to be passed successfully. Bonus part is optional but highly encouraged for more fun!

### Task 1: Set Up the Gallery Structure

#### 1- Create a New HTML File:

Inside your HTML file, set up the basic structure of the HTML document with a `<head>` and `<body>` inside an `<html>` tag.

#### 2- Give a gallery title:

- Inside your `<body>` , create an `<h1>` tag inside a `<div>` with an `id` of `title`. Then put the name of your gallery inside it !

```HTML
    <div id="title">
        <h1>your favorite gallery name</h1>
    </div>
```

#### 3- Put your robots inside the gallery:

- Under the `title` div, create a div element with the `id` `gallery`.

- This div will serve as the container for all the robot portraits.

```html
<div id="gallery"></div>
```

- Each team member should copy their robot's HTML structure (the one you provided in the `first-move.html` exercise between the `<section>` `</section>` code) and paste it inside this `gallery` div.

- Ensure each robot is placed inside its own div with the class `name-robot`, do not forget to change the `name` to be your name in the team, for example:

```html
<div class="john-robot">
  <!-- Paste your robot's HTML structure here (face, upper-body, lowe-body)-->
</div>
<div class="sarah-robot">
  <!-- Paste your robot's HTML structure here (face, upper-body, lowe-body) -->
</div>
<div class="bob-robot">
  <!-- Paste your robot's HTML structure here (face, upper-body, lowe-body) -->
</div>
```

#### 4- Add Robot Information:

Under each robot's `</section>` (that you copied from your robot's code in `first-move.html`), add an `<h3>` element for the robot's name and a `<p>` element for a short description of its power.

```html
<h3>Your robot's name</h3>
<p>Your robot's description</p>
```

### Task 2: Style the Gallery

#### 1- Update the CSS File:

- In your HTML file, link your CSS file by pasting the following code in between the `<head></head>` tags:

```html
<link rel="stylesheet" href="robots-harmony.css" />
```

- In your CSS file, add these styles:

```css
* {
  margin: 0;
  box-sizing: border-box;
  opacity: 0.85;
}

body {
  height: 100vh;
}

section {
  padding: 20px;
  width: 100%;
  height: auto;
  display: flex;
  justify-content: center;
}

div,
p {
  border: solid 1px black;
  padding: 10px;
  margin: 0;
  border-radius: 30px;
}

#face {
  align-items: center;
  margin-bottom: 20px;
}

#eyes {
  display: flex;
  background-color: yellow;
  justify-content: space-between;
  align-items: center;
  border-radius: 50px;
  width: 200px;
}

#torso {
  width: 150px;
  height: 90%;
  background-color: violet;
}

.eye {
  width: 60px;
  height: 60px;
  background-color: red;
  border-radius: 50%;
}

.arm {
  background-color: aquamarine;
  width: 50px;
  height: 100px;
  margin: 20px;
}

.leg {
  background-color: dodgerblue;
  width: 50px;
  height: 100px;
  margin: 20px;
}

.body-member {
  width: 50px;
  margin: 30px;
}

h3 {
  margin-bottom: 10px;
}

#title {
  text-align: center;
  margin: 1rem;
  background-color: #383086;
  color: rgb(158, 235, 223);
}
```

> Trust the process!

- Then add this gallery style block, and change the `background-color` of it based on your team's favorite color:

```css
#gallery {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 100px;
  padding: 100px;
  /* Change this color to reflect the color of your gallery */
  background-color: #383086;
}
```

- Add the following block of style to all your class's of `name-robot`. To do so, we follow the rule:

```css
.class-one,
.class-two,
.class-three {
  /* block of style */
}
```

- Name it with your `name-robot` for each member of the team, and put inside the block the following styles:

```css
{
    border: 2px solid #333;
    padding: 20px;
    text-align: center;
    background-color: #fff;
    border-radius: 15px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: transform 0.3s, box-shadow 0.3s;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    min-height: 400px;
}
```

#### 2- Add Animations:

- Add a subtle `hover` animation to the robot portraits to make the gallery more dynamic and engaging.

```css
.class-one:hover,
.class-two:hover,
.class-three:hover {
  transform: scale(1.05);
  box-shadow: 0 6px 10px rgba(0, 0, 0, 0.15);
}
```

> Experiment by changing the colors of the box-shadow and the scale value! Be creative.

- In your `#gallery` CSS rule, add some animated gradient color to the background! You can achieve it by combining CSS properties: `background`, `background-size`, `animation`, `box-shadow`.

> Hint : Do not forget to replace the background property with the new value bellow!

_For Example:_

```css
background: linear-gradient(
  45deg,
  /* Starting angle of the gradient */ red,
  /* First color in the gradient */ blue,
  /* Second color in the gradient */ green
); /* Optional third color in the gradient */
background-size: 300% 300%; /* Increases the background size for smooth animation or effects */
animation: gradientBackground 2s ease infinite; /* Animates the background to create a dynamic effect. */
box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2); /* Adds a shadow around the element. Adjust the offsets, blur radius, and opacity as needed */
```

> You saw no change? well we did not specify yet what to do with all those colors, angles and timing yet!

- Let's make it more exciting, by actually making the colors move! We can do that with `keyframes`! Under the `#gallery` css rule, put the `keyframes`` block and see the magic !

_For Example:_

```css
@keyframes gradientBackground {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}
```

> Play with the values and colors to get the best effect for your gallery!

**`Prompt Example:`**

- "How do I create a smooth gradient background with multiple colors in CSS?"

- "Explain to me as a beginner these CSS properties: `background`, `background-size`, `animation`, `box-shadow`."

### Task 3: Add Interactivity with JavaScript:

#### 1- Add Color Change on Key Press:

Inside your JavaScript file, write functions that change the colors of different parts of your robots when specific keys are pressed.

Follow the following steps:

- **Function to Change Arm Colors**:

  - Create a function named `changeArmColor`.
  - This function should accept a parameter `robotClass` to identify which robot to modify.
  - Inside the function, generate a random color using JavaScript, to both `left` and `right` arms.

  ```js
  const randomColor = `#${Math.floor(Math.random() * 16777215).toString(16)}`;
  ```

  - Use `document.querySelector` to select the `left` and `right` arm elements of the robot.
  - Apply the random color to both arm elements using `style.backgroundColor`.
  - Here is an example, do the same for right arm:

  ```js
  const randomColor = `#${Math.floor(Math.random() * 16777215).toString(16)}`;

  document.querySelector(`.${robotClass} #arm-left`).style.backgroundColor =
    randomColor;

  //do the same to right arm
  ```

- **Function to Change Legs Colors**:

  - Create a function named `changeLegColor`.
  - Follow the same steps as the `changeArmColor` function but target the `legs`.

- **Function to Change Eye Colors**:

  - Create a function named `changeEyeColor`.
  - Follow the same steps as the `changeArmColor` function but target the `eye`.

- **Function to Change Face Colors**:

  - Create a function named `changeFaceColor`.
  - Follow the same steps as the `changeArmColor` function but target the `eyes` `id`. (Yes, pay attention, the `id` that changes the face color is called `eyes`!)

#### 3- Adding Event Listeners for Keyboard Input:

Next, you need to detect when specific keys are pressed and trigger the corresponding function.

- Use `document.addEventListener` to listen for `keydown` events.
- Inside the event listener, use multiple `if` statements to check which `key` was `pressed`.
- Depending on the `key` pressed, call the appropriate `function` and pass the `robot's class` as an `argument`.

_Code example:_

```js
document.addEventListener("keydown", function (event) {
  if (event.key === "1") {
    changeArmColor("john-robot");
  }
  if (event.key === "2") {
    changeLegColor("john-robot");
  }
  if (event.key === "Q" || event.key === "q") {
    changeEyeColor("john-robot");
  }
  if (event.key === "A" || event.key === "a") {
    changeFaceColor("john-robot");
  }

  // Add similar conditions for other robots choosing what keys will change the colors!

  //...Here
});
```

- You should decide with your team mates the keys that will trigger each robot of yours!

**`Prompt Example:`**

- "What is `event.key` in JavaScript and how can it be used to detect keyboard inputs?"
- "How can I use `document.querySelector` to select multiple elements in JavaScript?"

#### 4- Task 5: Final Touches:

- Ensure that all robots are displayed correctly in the gallery.

- Make sure all files (robots-harmony.html, robots-harmony.css, robots-harmony.js) are in the same folder submitted to your `Gitea`.

- Double-check the code to ensure everything is clean and well-organized.

### Expected Output

Your project needs to check all the previous tasks, it will look something close to (and maybe better than) [this](The link of the demo).

### Bonus part

If you would like to make your project even more creative, you can add more features to your `gallery` on top of the mandatory ones!

It could be `music` to the `background`, `pop-out` information window or anything that makes your project `different`!

> Remember to ensure that the mandatory part is working perfectly before adding more effects.

Look at all of you—you've made it to this point, working on uniting all your robot friends after bringing them to life. Now in this final step (or maybe the beginning), use your collaboration skills as much as you can!

**"Alone, we can do so little; together, we can do so much." — Helen Keller**
