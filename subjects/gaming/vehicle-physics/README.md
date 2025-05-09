## Vehicle Physics

In this exercise you will learn to use the heritage system to create multiple blueprints from a single blueprint file. Also, you will learn to create the vehicle gameplay from scratch using blueprints. Remember to always wear a seat belt and do not text and drive.

#### Objectives

The point of this exercise is to create three cars and be able to enter and exit them as a third person character and drive them around.

#### Instructions

The player starts the game by controlling a Third Person character and as soon he gets close to one of the doors of one the three cars present in the level, a message should appear saying to press 'E' to enter the car. As soon as the player presses 'E', the corresponding door of the corresponding car should open and the player should then control the vehicle. Once inside the vehicle the player can press 'E' again to exit the car. Once pressed, the left door of the vehicle should open and the player should appear next to the car.

As for the vehicle, these are the expectations:

- create a blueprint named BaseCar, that will be the parent of all vehicles classes.
  - BaseCar should be a fully functional and realistic vehicle. In other words, it should accelerate, brake, have a handbrake, the back wheels should not turn, neither should the front wheels exceed a certain angle, etc...
  - BaseCar should also have dynamic instance for the lights.
    - The vehicle lights needs to simulate the real behavior of a real vehicle. When pressing the letter 'L' the front lights and the rear presence lights should turn on or off, when braking three braking lights (left, right and center) should turn on and the rear lights should turn on when reversing.
    - The BaseCar blueprint should contain "Set Material" nodes to change the headlamp material in order to have the effect where the lights are turned on.
- use PhysicMaterials to simulate rain driving terrain.

- only one AnimationBlueprint can be used for this project.

- when inside the vehicle, if the 'C' key is pressed, the camera should change from the third person view of the vehicle to the driver view inside the car.
  - when in the camera of the driver, the player must be able to see the gears and the speedometer in the dashboard.

There must be three models of cars on the map the character can choose from. All of them should inherit the BaseCar class. These three models should be:

- a Peugeot 3008 year 2018 car.
- an Audi A8 year 2011 car.
- a Tesla Model 3 car.
- The Mass, Acceleration, Number of gears, Max speed, Torque Curve, etc. needs to be really close of the selected vehicle.

To create the map you should use the Landscape tool. Also add a track using spline roads.

After downloading and unzipping this [file](https://assets.01-edu.org/Unreal-Engine-Piscine/VehiclePhysics.zip), you can copy its content to your project Content folder.

When finished, your project should look like the [“Expected Result” video](https://youtu.be/4dXjFKh_jjY).

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

#### Bonus

As bonus objectives you can:

- Create steam/flames particles going out from the car exhaust.
- Animate the character entering and sitting in the car.
- Add a button to open the doors while inside the car without exiting it.

[Here](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf9sxZqgfPVM06PrLk8_CWA-) you can find an instructional playlist on Unreal Engine.
