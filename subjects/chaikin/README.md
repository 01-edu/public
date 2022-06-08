## Chaikin

### Instructions

Implement [Chaikin's](http://graphics.cs.ucdavis.edu/education/CAGDNotes/Chaikins-Algorithm/Chaikins-Algorithm.html) algorithm as a step by step animation.

You will create a canvas to enable the user to draw 1 or more points. The screen will then animate each step that is taken to obtain the final result of a drawn curve. The animation must play for 7 steps and then it must restart.

You can see how the application should work [here](https://youtu.be/PbB2eKnA2QI).

> Usage of Chaikin's algorithm is mandatory.

#### Functionality

- The canvas must receive input from the mouse. The user must be able to use the left button to place control points for Chaikin's algorithm.
  
- The selected points must be visible on the canvas. You will need to draw a **small** circle surrounding each point.

- After pressing `Enter`, the program should start the animation. It should cycle through the steps until reaching the 7th step of `Chaikin's` algorithm. The program should then restart.

- If the canvas only has only one control point, the program must only present that point and not cycle through the steps.

- If the canvas has only two control points, the program must draw a straight line.

- Pressing `Escape` should quit the window.

> You are free to use any library to create and handle windows, rendering, keyboard and mouse events.

### Bonus

- Make it possible to clear the screen, so that the user can select new control points.
- Make it possible to drag the control points in real time.
