## Firing-range

### Overview

This project involves creating a fully functional firing range in Unreal Engine 5 (UE5). You will focus on implementing shooting mechanics, a weapon system, player input handling, physics for projectiles, and basic AI to control moving targets. Your task is to build a simple environment where the player can shoot targets, and all systems should work seamlessly together.

<center>
<img src="./resources/fgm.jpg?raw=true" style="width: 500px !important; height: 350px !important;"/>
</center>

### Role Play

You are developing a firing range training level for a first-person shooter game. You are required to set up a working shooting system, integrate a functional weapon with reload mechanics, implement a physics system for projectiles, and design basic AI for moving targets. Additionally, you will create a player HUD that tracks their score, accuracy, and remaining ammunition.

### Learning Objective

By the end of this project, you will have implemented:

- A complete shooting mechanic (with accurate aiming, projectiles, and hit detection).
- A weapon system with at least one firearm, including shooting and reload mechanics.
- Basic AI to control target movement.
- Physics-based projectiles that interact with the environment.
- A simple UI displaying relevant gameplay information (accuracy, ammo count, etc.).
- A simple firing range level design that incorporates moving and stationary targets.

### Instructions

#### General Requirements

- Use `Blueprints` or `C++` to implement the game mechanics.
- Implement a firing system where the player can aim downsights and shoot.
- Implement a weapon system that supports shooting, reloading, and proper hit detection.
- Create a firing range level that includes both stationary and moving targets.
- Implement a physics system that governs projectile behavior (trajectory and collisions).
- Implement a HUD that shows the player's score, accuracy, and remaining ammunition.

#### Main Menu

The game's main menu must:

- Be a separate `level/map`.
- Include the following options:
  - **Start Game**: Transitions to the firing range gameplay.
  - **Settings**: Allows the player to adjust mouse sensitivity.
  - **Quit**: Closes the game.

#### Game Screen (HUD)

The game screen must include a HUD that displays:

- A precise crosshair that accurately indicates the bullet's point of impact.
- The player's current accuracy (hits devided by total shots fired).
- Remaining ammunition and reload status for the current weapon.

#### Player Character

The player character must have the following functionalities:

- Basic movement around the firing range.
- Proper input handling for aiming, shooting, and reloading.
- Interact with ammo pickups to replenish ammo.

#### Weapons

You must create at least one firearm (e.g., a pistol) that includes:

- Recoil mechanics.
- Reloading the weapon with an appropriate animation.
- Projectile physics that simulate bullet trajectories.
- Bullet hit detection (with an effect on impact, e.g., sound and visual feedback).

#### Targets and AI

You must implement the following for the targets:

- Create a stationary target.
- Create a moving target with basic AI that follows a predefined path or pattern.
- Implement hit detection for the targets.
- Ensure the targets respawn after a set interval once they are hit.
<center>
  <div style="display: flex; justify-content: center;">
    <img src="./resources/example1.jpg?raw=true" style="width: 500px !important; height: 350px !important;" />
    <img src="./resources/example2.jpg?raw=true" style="width: 500px !important; height: 350px !important;" />
  </div>
</center>

#### Game Loop Logic

The game loop must be continuous and consist of:

- The player starts with a loaded weapon and can shoot at targets.
- The player can open a pause menu at any time, with options to:
  - Restart the game (reset score, targets, and ammo).
  - Return to the main menu.

#### Level Design

You must design a firing range that includes:

- A coherent theme to ensure a consistent and immersive experience.
- Proper lighting and environment setup to ensure the range is clear and visually appealing.
- A section with stationary targets.
- A section with moving targets.
- Ammo pickups where the player can restock.

### Bonus

- Add multiple weapons (e.g., a shotgun, sniper rifle) with unique shooting and reload mechanics.
- Implement more advanced target AI, such as random movement patterns or different difficulty levels.
- Add different hit zones for targets, with headshots awarding more accuracy points.
- Design a timed challenge mode where the player must hit all targets within a set time frame.

### Submission

- You must upload a zip file of your game build in your repository.
  - Ensure that the build works on your platform.
  - If file size is an issue, use GitHub with [Git LFS](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github).

### Resources

- [Unreal Engine's AI Documentation](https://docs.unrealengine.com/en-US/InteractiveExperiences/AI/index.html)
- [Projectile Movement Component](https://docs.unrealengine.com/en-US/Gameplay/Movement/ProjectileMovement/index.html)
- [Mixamo for character animations and models](https://www.mixamo.com/)
- [SketchFab for 3D models](https://sketchfab.com/)
- [itch.io for additional assets](https://itch.io/)
