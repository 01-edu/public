## Widget Factory

The focus of this exercise is to learn to display information, buttons, images and whatever you will need on the screen using Widgets. Kinda like those displays in Iron Man, but not really.

### Objectives

You will create a first-person style game with only directional and mouse input. The main goal is to create a pause menu where you can change the graphics settings such as resolution, resolution scale, render distance, shadow, texture quality and so on…

### Instructions

For this project you will start the game with an empty project and be sure to create it including the starter content. It is requested that you:

- Create a PlayerCharacter blueprint and give him the ability to move forward, backward, left and right. Also, use the mouse to look around.
- Create a start menu by pressing Escape.
- This Start menu should pause the game and the mouse should show up.
  - The Start menu should have a minimum of 4 buttons:
    - Resume, that resumes the game when clicked
    - A button for each of the students in the group with their username
      - Each button should link directly to each student Intra profile.
    - Graphics Settings, explained afterwards
    - Quit to desktop, that on release should exit the game

When the player clicks on the Graphics Settings button, another widget should appear where it contains:

- a slider that lets the player choose the percentage of the resolution scale.
- a dropdown menu where the player can choose the resolution.
  - At least three different resolutions should be displayed in the dropdown
- a button at the bottom of the Widget to apply the two above resolution settings.
- a checkbox that when switched on shows the game in full screen and in windowed mode otherwise
- four choices for PostProcessQuality, ShadowQuality, TextureQuality, EffectQuality and FoliageQuality. These four choices are LOW, MEDIUM, HIGH and EPIC.
  - Theses settings got to be executed with the “Execute Console Command” node.
- PostProcessAAQuality setting should have the following options as buttons: Off, 2x, 4x, 6x.
  - You can use the “append (string)” node to switch between all four choices.
- A “Back” button that bring us back to the main Start menu.
- No more than 25 string variables are allowed in the Graphics Setting widget file.

While controlling the player in the game, a widget should be visible at all times. This widget should display on a corner of the screen the rounded X, Y, Z (as integers) coordinates of the player. Just make sure it doesn’t disturb the visibility of the other widgets. If you want, after opening the pause menu you can make the player position widget disappear, making it appear again when you return to the game.

On the level you should display some grass, textured objects and post process on the terrain to show the settings changes. You can find some stuff on the StarterContent folder.

When finished, your project should look like the executable example on the folder or the [“Expected Result” video](https://youtu.be/0v-spQJwDwM).

> Don’t forget to zip up the project compile and save everything for peer correction.
