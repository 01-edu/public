## Chaikin

### Instructions

Implement the [Chaikin's](http://graphics.cs.ucdavis.edu/education/CAGDNotes/Chaikins-Algorithm/Chaikins-Algorithm.html) algorithm step by step in an animation.

The objective is to create a canvas where the user will be able to draw 1 or more points on it, and then the screen will play an animation of each step that is taken to obtain the final result of a drawn curve. The animation must play for 7 steps and then it must restart.

It is mandatory to use the `Chaikin's` algorithm.

You can see how the application should work [here](https://youtu.be/PbB2eKnA2QI).

- Functionality:

  - The canvas should receive input from the mouse:
    - The user of the program should be able to use the left button to select the control point for the `Chaikin's` algorithm.
    
  - The selected points must be visible in the canvas, so you will have to draw a small circle surrounding the selected points.

  - After pressing `Enter` the program should start the animation of the steps until reaching the 7th step of the `Chaikin's` algorithm. After the 7th step the program must restart.

  - If the canvas only has only one control point, the program must only present that point and do no change at all.

  - If the canvas has only two control points, the program must draw a straight line.

  - Pressing `Escape` should quit the window.

### Note

You are free to use any library to create and handle windows, rendering, keyboard and mouse events.

### Bonus

- Add the ability to clear the screen and add new control points.

- Make it possible to drag the control points in real time.
