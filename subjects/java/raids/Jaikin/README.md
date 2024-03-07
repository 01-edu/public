## Jaikin

### Objective

The goal of this project is to implement Chaikin's algorithm and create a step-by-step animation of the process using a canvas.

### Instructions

Follow these instructions to create the Chaikin's algorithm step-by-step animation:

1. Create a canvas where the user can draw one or more points. Each point should be represented by a small circle.

2. Receive input from the mouse, allowing the user to place control points for Chaikin's algorithm using the left button.

3. Display the selected points on the canvas as small circles surrounding each point.

4. If the canvas has control points drawn on it, pressing the `Enter` key should initiate the animation. The animation should cycle through the steps of Chaikin's algorithm, proceeding until it reaches the 7th step. After completing the 7th step, the animation should restart.

5. If the user presses `Enter` before any points have been drawn, nothing should happen. However, the user should still be able to draw points, and an optional message may be displayed to remind the user to draw points.

6. If the canvas has only one control point, the program must display that point without cycling through the steps.

7. If the canvas has only two control points, the program must draw a straight line between those two points.

8. Pressing the `Escape` key should close the window.

### Bonus Features

Optionally, you may implement the following bonus features:

1. Allow the user to clear the screen, enabling them to select new control points.

2. Implement real-time dragging of the control points, so the user can adjust their position interactively.

> Note: You are free to choose any library for handling windows, rendering, keyboard, and mouse events.

For more information on Chaikin's algorithm, refer to [this resource](https://www.cs.unc.edu/~dm/UNC/COMP258/LECTURES/Chaikins-Algorithm.pdf). You can also view an example of how the application should work in this [video](https://youtu.be/PbB2eKnA2QI). Remember, the usage of Chaikin's algorithm is mandatory for this implementation.
