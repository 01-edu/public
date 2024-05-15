## stealth-boom

In this project, you will have to create an entire stealth game using Unreal Engine.

### Objectives

The idea of this project is to create a little 10 minutes gameplay game with missions, with stealth based gameplay and with an AI patrolling NPC.

The basics assets you will need for this project can be found in the [Stealth-Boom.zip](https://assets.01-edu.org/Unreal-Engine-Projects/StealthBoom.zip) file. It contains the basic animations, character, enemies and props you will need for this project.

### Instructions

The following aspects should be fulfilled:

- Create a map where the player can walk around.

  - This map should contain places for the player to hide from enemies by crouching, hide behind walls, and all other props you may use to help it make a stealth game.
  - Buildings with at least 2 floors.
  - Pickable ammunition and weapons around the map.

- For the player you should add:

  - Walk and run animations
  - Reload animation
  - Aim animation
  - Shoot animation
  - Crouch animation
    - the player should be able to do the above six actions while crouching
  - Melee attack animation
  - Gun sound when firing
  - Bullets visual impact on walls (see decals documentation)
  - Blood particles when hit

- The game should contain a main menu where the player can:

  - Start the game
  - Adjust the general sound of the game
  - Change the graphics settings
    - When changing the resolution, a pop-up should appear in the middle of the screen asking if the player wants to keep the graphics setting he/she just applied. If `Yes` is clicked within 10 seconds, the settings are set, otherwise, if the 10 seconds delay is over or if the player clicks `No`, the settings go back to the old ones.
  - Change the mouse sensitivity
  - Invert the mouse vertical axis

- You should have at least 3 types of enemies: `Guards` (who patrol around and are lookig for intruders), `Drones` (same as Guards but which can fly), and `Cameras` (stationary and looking for intruders). More enemies can be added if you want to.

  - Guards AI:
    - Guards should be able to patrol around the map;
    - A Guard is able to see the player, if the player crosses his field of view;
    - When the player enters the field of view of a Guard, the Guard enters into chasing mode and must start running after the player, takes cover and shoots at the player.
    - If the player leaves the field of view for a certain time, the Guard goes back to patrol mode.
  - Drones AI:
    - Drones should be able to patrol around the map;
    - A light color should determine the state of the drone (Blue for patrolling, Red for chasing the player);
    - Once the player crosses the drone camera, the drone light turns red and the drone enters chasing mode;
    - When a drone is in chasing mode, all the guards on the area are alerted, and should enter chasing mode as well;
    - When the player is out of the drone sight, the drone turns back to patrol mode;
    - The sight radius should be inferior on the drones that on the guards.
  - Camera AI:
    - Cameras should be placed on walls;
    - Cameras should have the same light sign as the drone, so when the player is in the camera sight, the camera light turns red and all Guards enter in chasing mode;
    - Like the Drones, Cameras warn guards, whenever the player passes through the camera field of view;
    - Some Cameras should lock access of certain areas of the map (for example, close doors) when the player is detected by that camera.

- Drones, Guards and Cameras should have sounds effect whenever they change from chase to patrol mode, as well as the other way around.

- All AI should be implemented using Behavior Trees except for the Camera.

- The player mission is up to you, either it can be some task to fix, kill all guards without getting caught or collect documents... Whatever you choose, the player should have enemies on his way to divert him away from his objective.

  - When the player successfully completes his mission, a pop up should appear saying that the mission is completed.

- The player has a health bar that should go down whenever the player gets shot. When the player dies, he has the choice to either quit the game, go back to the main menu or start over.

  - If the player starts over, the level should not be reloaded. The player should spawn back to the starting point.

- When pressing `Esc` the game is set on paused and the main menu widget pops up.

  - The player should be able to change the game graphics setting exactly like in the main menu.

- A game should not last longer than 6 minutes. After that the same choices should appear as when the player dies.

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)
