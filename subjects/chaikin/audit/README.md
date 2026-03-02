#### Functional

##### Run the program using `cargo run`.

###### Does it compile and run without warnings?

##### Left click on the canvas to set one or more control points.

###### Does the program allow you to mark these control points?

###### Can you confirm that the program draws a small circle around the control points in order to identify them?

##### Left click on the canvas to set three or more control points and press `Enter`.

###### Does the animation of the Chaikin's algorithm start?

##### Press `Escape` to exit the program.

###### Does the program exit without errors?

##### Start the program and left click on the canvas to set just one control point and press `Enter`.

###### Can you confirm that only the control point is shown and nothing changes?

##### Restart the program and left click on the canvas to set just two control points and press `Enter`.

###### Can you confirm that only a straight line was drawn?

##### Restart the program and left click on the canvas to set three or more control points and press `Enter`.

###### Does the animation complete 7 steps before restarting?

##### Restart the program and left click on the canvas to set three or more control points and press `Enter`. Then press `Escape` to exit the program.

###### Does the program exit without errors?

##### Start the program and press `Enter` without selecting any points.

###### Does the program continue without errors?

###### After you pressed `Enter` before selecting points, is it possible to place points without needing to kill the program?

##### Unit Tests

##### Run the command `cargo test`.

###### Do all tests pass without errors?

###### Is there a test that verifies the math of a single Chaikin iteration (calculating the 1/4 and 3/4 points)?

###### Is there a test checking that the point count grows correctly according to the algorithm?

###### Are there tests ensuring the algorithm handles empty or single-point inputs safely?

##### Bonus

###### +When you pressed `Enter` without drawing any points, was a message displayed to inform you that you forgot to draw any points?

###### +Is it possible to clear the screen and add new control points without killing and relaunching the program?

###### +Is it possible to drag the control points in real time and get a new curve?
