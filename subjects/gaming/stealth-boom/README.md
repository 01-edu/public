## stealth-boom

In this project, you will have to create a stealth game using Unreal Engine.

### Objectives

The goal of this project is to create a playable gameplay loop around stealth mechanics.

The basics assets you will need for this project can be found in the [StealthBoomAssets.zip](https://assets.01-edu.org/gamedev/stealth-boom-assets.zip) file. It contains the basic animations, character, enemies and props you will need for this project.

> NOTE: The assets in the file are intended to help streamline the process of locating assets, not to eliminate it.
> TIP: Use [itch.io](https://www.itch.io) to get sound effects (not included in the assets file) or to find extra assets

### Resources

Here are some resources to help you tackle potential project challenges:

- [behavior-tree-in-unreal-engine](https://dev.epicgames.com/documentation/en-us/unreal-engine/behavior-tree-in-unreal-engine---quick-start-guide)
- [decal materials and decal actors](https://dev.epicgames.com/documentation/en-us/unreal-engine/decal-materials-in-unreal-engine?application_version=5.4)
- for inspiration look at games like metal gear solid 1/2/3.

### Instructions

The following aspects should be fulfilled:

- Main Menu:

  - Start the game
  - Adjust the general sound of the game
  - Change the graphics settings
    - When changing the resolution, a pop-up should appear in the middle of the screen asking if the player wants to keep the graphics setting he/she just applied. If `Yes` is clicked within 10 seconds, the settings are set, otherwise, if the 10 seconds delay is over or if the player clicks `No`, the settings go back to the old ones.
  - Change the mouse sensitivity
  - Invert the mouse vertical axis

- Map/Level.

  - This map should contain places for the player to hide from enemies by `Ducking`, `Hiding behind walls and props`
  - Pickable ammunition around the map.
  - Pickable Health around the map.

- Player:

  - Walk and run animations
  - Shoot animation
  - Duck animation
  - Melee attack animation
  - Blood particles when hit
  - A health bar that should go down whenever the player gets damaged. When the player dies, he has the choice to either quit the game, go back to the main menu or start over.
  - The player mission is up to you, it can be some task to fix, kill all enemies without getting caught or collect documents... Whatever you choose, the player should have enemies on his way to divert him away from his objective.
  - When the player successfully completes his mission, a pop up should appear saying that the mission is completed.

- Gun mechanics:

  - a widget displaying the number of bullets
  - when the number of bullets is 0 the player is unable to shoot
  - shooting should have a sound effect and visual effect
  - Bullets visual impact on walls

  - Enemy AI:
    - Your game should contain atleast 2 enemy types `Melee` and `Ranged`
    - enemy AI should be made with behavior trees
    - enemy should be able to patrol around the map in a predefined path;
    - enemy detects the player, if the player crosses its field of view;
    - When the player enters the field of view of an enemy, the enemy enters into chasing mode and must start running after the player.
      - `Ranged` enemies should take cover and shoot at the player.
      - `melee` enemies should run close to the player and perform melee attacks.
    - Enemies in chase mode alert nearby enemies making them enter chase mode as well.
    - If the player leaves the field of view of all enemies for a certain time, the enemies go back to patrol mode.

- Enemies should have sounds effect whenever they change from chase to patrol mode, as well as the other way around.

- Enemies should have a visual indicator showing whether they are in patrol or chase mode

- When pressing `Esc` the game is set on paused and a widget similar to the main menu widget pops up.

  - The player should be able to change the game graphics setting exactly like in the main menu.

- A game should not last longer than 6 minutes. After that the same choices should appear as when the player dies.

### Bonus

- use your own assets to create a game in your own style by either searching online or creating them yourself

- have more enemy types e.g a turret that is stationary but deals great damage, etc...

- have areas in the game that are locked behind doors that require keys which you can obtain from specific enemies

- have multiple different weapon types that you can pickup around the map and use to finish the mission

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)
