#### Functional

##### Try to run the command `./game_engine -f maps/map01 -p1 robots/bender -p2 robots/terminator` inside the container.

###### Can you confirm that the student was able to create the image and container correctly?

##### Try to run the student player against one of our players.

###### Can you confirm that the project runs correctly?

##### Try to run the student player against one of our players.

###### Can you confirm that the student player is placing the pieces correctly with the overlapping of just on cell?

##### Try to run `./game_engine -f maps/map00 -p1 <path to student player> -p2 robots/wall_e` five times changing the position of the players each time so that the student player can be the `p1` and the `p2`.

###### Can you confirm that the student player won at least 4 out of 5 times?

##### Try to run `./game_engine -f maps/map01 -p1 <path to student player> -p2 robots/h2_d2` five times changing the position of the players each time so that the student player can be the `p1` and the `p2`.

###### Can you confirm that the student player won at least 4 out of 5 times?

##### Try to run `./game_engine -f maps/map02 -p1 <path to student player> -p2 robots/bender` five times changing the position of the players each time so that the student player can be the `p1` and the `p2`.

###### Can you confirm that the student player won at least 4 out of 5 times?

#### Unit Tests

###### Do all tests pass without errors?

###### Are there specific tests for **Input Parsing** (e.g., verifying the robot correctly reads the Anfield dimensions and the piece shape from stdin)?

###### Are there tests for **Placement Validation** (e.g., checking that a move is rejected if it overlaps two of your own cells or one of the opponent's)?

###### Are there tests for **Boundary Detection** to ensure pieces are never placed partially outside the grid?

#### Basic

###### +Does the code obey the [good practices](../../good-practices/README.md)?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

#### Bonus

###### +Did the student create a visualizer for the project?

##### Try to run `./game_engine -f maps/<map of your choice> -p1 <path to student player> -p2 robots/terminator` five times changing the position of the players each time so that the student player can be the `p1` and the `p2`.

###### +Can you confirm that the student player won at least 4 out of 5 times?
