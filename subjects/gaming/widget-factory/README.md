## Widget Factory

The focus of this exercise is to learn to display information, buttons, images and whatever you will need on the screen using Widgets. It is kinda like those displays in Iron Man, but not really.

### Objectives

You will create a first-person style game with only directional and mouse input. The main goal is to create a pause menu where you can change the graphics settings such as resolution, resolution scale, render distance, shadow, texture quality and so on…

### Instructions

For this project you will start the game with an empty project and be sure to create it including the starter content. It is requested that you:

- Create a PlayerCharacter blueprint and give him the ability to move forward, backward, left and right. Also, the mouse has to be used to look around.

- Create a start menu by pressing Escape.

- This Start menu should pause the game and the mouse should display.

  - The Start menu should have a minimum of 4 buttons:

    - A Resume button, that resumes the game when clicked.
    - A button for each of the students in the group with their username.
      - Each button should link directly to each student Intra profile.
    - A Graphics Settings button, this is explained afterwards.
    - A Quit to desktop button, that on release should exit the game.

When the player clicks on the Graphics Settings button, another widget should appear where it contains:

- a slider that lets the player choose the percentage of the resolution scale.

- a dropdown menu where the player can choose the resolution.

  - At least three different resolutions should be displayed in the dropdown.

- an apply button at the bottom of the Widget which actualizes the two above resolution settings.

- a checkbox, which when switched on, shows the game in full screen and in windowed mode otherwise.

- four choices for PostProcessQuality, ShadowQuality, TextureQuality, EffectQuality and FoliageQuality. These four choices are LOW, MEDIUM, HIGH and EPIC.

  - These settings have to be executed with the “Execute Console Command” node.

- PostProcessAAQuality setting should have the following options as buttons: Off, 2x, 4x, 6x.

  - You can use the “append (string)” node to switch between all four choices.

- A “Back” button that bring the player back to the main Start menu.

- No more than 25 string variables are allowed in the Graphics Setting widget file.

While controlling the player in the game, a widget should be visible at all times. This widget should display on a corner of the screen the rounded X, Y, Z (as integers) coordinates of the player. Just make sure it does not disturb the visibility of the other widgets. If you want, after opening the pause menu you can make the player position widget disappear, and make it appear again when you return to the game.

On the level you should display some grass, textured objects and post process on the terrain to show the settings changes. You can find some of those in the StarterContent folder.

When finished, your project should look like the [“Expected Result” video](https://youtu.be/0v-spQJwDwM).

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

### Bonus

As bonuses you can:

- Add a resolution confirmation button when changing resolution.
- Use the “On Release” attribute of the buttons, instead of the "On Pressed".
  Automatically set the default screen resolution to get the best computer performance.

[Here](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf9sxZqgfPVM06PrLk8_CWA-) you can find an instructional playlist on Unreal Engine.
