## stealth-boom

A stealth game.

<center>
<img src="./resources/mgsmeme.png?raw=true" style = "width: 700px !important; height: 464px !important;"/>
</center>

### Objectives

The goal of this project is to create a playable gameplay loop around stealth mechanics.

### Scenario

You are an Unreal Engine developer tasked with creating a full stealth game from scratch. You need to implement everything, including AI behavior trees, player animations, and gun mechanics. As you progress, you realize balancing these elements—ensuring the AI responds accurately to player movements while maintaining smooth gameplay—is more challenging than expected. Every feature, from the main menu to mission completion, must work seamlessly to deliver a polished and playable game within the given constraints.

Good luck and remember to have fun while making the game !

### Resources

Here are some resources to help you tackle potential project challenges:

- [behavior-tree-in-unreal-engine](https://dev.epicgames.com/documentation/en-us/unreal-engine/behavior-tree-in-unreal-engine---quick-start-guide)
- [decal materials and decal actors](https://dev.epicgames.com/documentation/en-us/unreal-engine/decal-materials-in-unreal-engine?application_version=5.4)
- for inspiration look at games like metal gear solid 1/2/3.

- The basics assets you will need for this project can be found in the [StealthBoomAssets.zip](https://assets.01-edu.org/gamedev/stealth-boom-assets.zip) file. It contains the basic animations, character, enemies and props you will need for this project.

> NOTE: The assets in the file are intended to help streamline the process of locating assets, not to eliminate it.
> TIP: Use [itch.io](https://www.itch.io) to get sound effects (not included in the assets file) or to find extra assets

### Instructions

The following aspects should be fulfilled:

#### Main Menu:

- Option to start the game.
- Adjust the general sound of the game.
- Change graphics settings:
  - When changing the resolution, a pop-up should appear in the middle of the screen, asking if the player wants to keep the newly applied graphics settings. If the player clicks "Yes" within 10 seconds, the settings are confirmed. If the 10 seconds pass or the player clicks "No," the settings revert to the previous ones.
- Change the mouse sensitivity.
- Option to invert the mouse vertical axis.

#### Map/Level.

- The map should include areas where the player can hide from enemies by `ducking` or taking cover `behind walls and props`.
- There should be pickable ammunition scattered throughout the map.
- Health packs should be placed around the map for the player to collect.

#### Player:

- The player should have animations for `walking`, `running`, `shooting`, `ducking`, and performing `melee attacks`.
- Blood particles should appear when the player is hit.
- A health bar should decrease whenever the player takes damage.
- Upon death, the player should have the option to quit the game, return to the main menu, or start over.
- The player's mission is flexible and can be any of the following: completing a task, eliminating all enemies without being detected, or collecting documents. Regardless of the mission, the player will encounter enemies that attempt to hinder their progress.
- When the player successfully completes his mission, a pop up should appear saying that the mission is completed.

#### Gun mechanics:

- The player should be able to shoot.
- A widget should display the current number of bullets available to the player.
- wWhen the bullet count reaches 0, the player should be unable to shoot.
- Shooting should trigger both a sound effect and a visual effect.
- Bullets should have a visual impact on walls.

#### Enemy:

- The game should feature at least two types of enemies: `Melee` and `Ranged`.
- Behavior trees should be used to implement enemy AI.
- Enemies should be able to patrol predefined paths around the map.
- Enemies should detect the player if the player enters their field of view.
- When the player enters the field of view of an enemy, the enemy enters into chasing mode and must start running after the player.
  - `Ranged` enemies should take cover and shoot at the player.
  - `melee` enemies should run close to the player and perform melee attacks.
- Enemies in chase mode alert nearby enemies making them enter chase mode as well.
- If the player leaves the field of view of all enemies for a certain time, the enemies go back to patrol mode.
- Enemies should have sounds effect whenever they change from chase to patrol mode, as well as the other way around.
- Enemies should have a visual indicator showing whether they are in patrol or chase mode

#### Game loop

- Pressing `Esc` pauses the game, bringing up a widget similar to the main menu.
- The player should be able to change the game graphics settings exactly as they can in the main menu.
- The game should last no longer than 6 minutes. After this time, the player is presented with the same choices as when they die: quit the game, return to the main menu, or start over.

### Bonus

- Use your own assets to create a game in your own style by either searching online or creating them yourself

- Have more enemy types e.g a turret that is stationary but deals great damage, etc...

- Have areas in the game that are locked behind doors that require keys which you can obtain from specific enemies

- Have multiple different weapon types that you can pickup around the map and use to finish the mission

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)
