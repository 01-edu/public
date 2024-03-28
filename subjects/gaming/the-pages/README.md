## the-pages

In this exercise, one will create a mini-horror game inspired on ["Slender: The Eight Pages"](https://en.wikipedia.org/wiki/Slender:_The_Eight_Pages) game.

### Objectives

Just like in the Slender game, the goal of your game should be, while being chased by a monster/creature (a "Creep"), to get all the 8 pages spread around the map in different locations.

To find models of scary monsters/creatures:

- you can easily choose checking the [Mixamo website](https://www.mixamo.com/#/).

- alternatively, you can use the models on the ["Creeps.zip"](https://assets.01-edu.org/Unreal-Engine-Projects/ThePages/Creeps.zip) file.

- you can also check in the [assets](https://assets.01-edu.org/Unreal-Engine-Projects/ThePages/ThePages.zip) (basics animations, character, enemies and props) you will need for this project.

### Instructions

On startin, the game should ask the player to press a button to enter the game. Once the player does that, the actual gameplay level should be loaded.

The map for your game should be a large squared landscape terrain, scaled at 10 by 10 units. The action should be happening at night time in a forest-like environment. Your map should contain:

- trees, some little facilities (like shelters or little buildings), monuments etc... You should use the [Foliage Tool](https://docs.unrealengine.com/en-US/BuildingWorlds/Foliage/index.html) to spread trees, rocks and grass around the landscape and to make them have different sizes, scales and rotations.
  - To create buildings, you will be using the built in [BSP](https://dev.epicgames.com/documentation/en-us/unreal-engine/geometry-brush-actors-in-unreal-engine). This method will allow you to create complex meshes and apply textures to it (Useful when you have no 3D modeling skills).
- the map should also contain 8 pages spread around it in the facilities or monuments you created, but should be placed in different spots of those locations in each game.
- the map should have at least three different kind of floor surface: dirt, grass and concrete.

For the creature or monster that you choose to design:

- It can only move (very slowly so it does not catch up to the player that quickly) when it is not within the player field of view. This means if the player sees it, it should stop.
  - Randomly, the "Creep" should teleport to another location.
- The "Creep" never appears before the player collects at least 1 page on the map.
- While completing the task of collecting the 8 pages, the monster should randomly disturb our experience by appearing on the player screen, making that task more difficult.
- When the creep is within the player's sight, a disturbing music/sound must be played and the camera should start stuttering like in the original game.

The player should control a FPS (First Person Shooter) Character that:

- has to collect 8 pages placed around the map.
  - The player can collect the pages by left-clicking on them.
  - A sound should be played when the player collects a page and a text should appear temporarily on screen showing how many pages are missing (example: 5 / 8 Pages).
- can toggle a flashlight on and off by pressing the "F" key.
  - A sound should be played when toggling the flashlight on or off.
- when walking should produce the sound of footsteps on the floor.
  - The sound of player footsteps should be different on at least 3 different surfaces: concrete, dirt and grass.
  - Once the player collects the 8 pages, the forest turns from night to day for a while and then the game ends and the credits show up on the screen.
    - The credits should mention the names of the developers.

If you want you could try to play the real game [Slender: The Eight Pages](https://pt.wikipedia.org/wiki/Slender:_The_Eight_Pages) to get a feel of what your game should look like.

> Do not forget to zip up the project compile and save everything for peer correction.
