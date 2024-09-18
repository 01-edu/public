## 2.5D Adventure

### Overview

This project involves creating a simple 2.5D side-scroller game designed to introduce you to the fundamentals of Unreal Engine 5 and game development, without the complexity of advanced mechanics. It focuses on the essential skills and concepts needed to start your journey as a game developer.

<center>
<img src="./resources/1.jpg?raw=true" style = "width: 500px !important; height: 350px !important;"/>
</center>

### Role play

You've just embarked on your journey into game development, choosing `Unreal Engine 5` as your primary tool. Excited by the possibilities, you decide to create a simple 2.5D game as your first project. However, as you dive into the engine, you realize there's a vast array of tools and features you've never encountered before. Now, you're tasked with navigating this uncharted territory, learning new skills in level design, scripting, and asset management to bring your game to life.

### Learning Objective

At the end of this project you will have learned:
- Setting up a basic player character (input handling, movement, camera, etc.)
- Basic scripting (`Blueprints` and/or `C++`).
- Collision detection.
- Object creation and `instancing`.
- Basic UI/HUD.
- Basic level design
- Game logic and flow (Game Loops).

### Instructions

#### General

The use of `event dispatchers` is mandatory throughout the project, any code coupling will be disqualified.
> Tip: look for the observer pattern.

Your game should follow a coherent `theme` of your choice, whether it's a dark fantasy, sci-fi, or something entirely different. Let your imagination guide you.
> Note: the resources section will list some places where you can browse assets.

> Tip: For maximum learning, it is recommended to start this project with a blank Unreal project and implement everything from scratch.
#### Main Menu

The main menu `widget` should:
- Be on a separate `level/map` from the `main game` map.
- Contain two buttons:
    - Start game button.
    - Exit game button.

#### Player Character

The player `character` should:
- Have a skeletal mesh.
- Move only along two axes: left-right and up-down.
- Have a basic locomotion system with the following animations:
    - `Idle`.
    - `Walking`.
    - `Running`.
    - `Jumping`.
    - `Falling`.
    - `Landing`.
- Transition smoothly between a walking and running animations based on their speed.
<center>
<img src="./resources/locomotion.gif?raw=true" style = "width: 500px !important; height: 350px !important;"/>
</center>

> Tip: Look into the best practices for creating a locomotion system, Animation blueprints and Blend spaces.

#### Collectible

The collectible `Actor` should:
- Have a static mesh.
- Rotate around an axis of your choice.
- Have a box collider that acts like a trigger.
- Be collected when the player enters its box trigger.
- Play a sound when collected
<center>
<img src="./resources/collectible.gif?raw=true" style = "width: 500px !important; height: 350px !important;"/>
</center>

#### HUD

The HUD `widget` should:
- Display a value related to the `collectible`.
- Be updated each time a `collectible` is picked up.
<center>
<img src="./resources/hud.gif?raw=true" style = "width: 500px !important; height: 350px !important;"/>
</center>

#### Enemy

The enemy `character` should:
- Have a skeletal mesh.
- Have at least a walking cycle.
- Have a simple AI that patrols between two set points.
- Kill the player on collision.
- Be killed if the player lands on top of it.
- Play a sound when killed.
- create an instance of `Collectible` at the place of death.
<center>
  <div style="display: flex; justify-content: center;">
    <img src="./resources/enemywalk.gif?raw=true" style="width: 500px !important; height: 350px !important;" />
    <img src="./resources/enemydie.gif?raw=true" style="width: 500px !important; height: 350px !important;" />
  </div>
</center>

#### Game Loop Logic

The game loop consists of a spawn point where the player `starts` and `respawns` when dead.
And an end point that defines the player goal to finish the level.
When the player reaches the end point:
- a menu with options to `restart` or `quit` the game should be displayed.

#### Level design

Your level design shouldn't consist of just a long, empty run to the finish point. However, since creating a level is a creative process, you have the freedom to design it as you see fit. For example, you can include moving platforms or skill-based platforming elements. The choice is yours to provide the player with a fun and challenging experience.
> Tip: look at other games from a similar genre for inspiration.
<img src="./resources/platformer.jpg?raw=true" style="width: 600px; height: 350px;" />

<!-- Expected results video link to be added later when available -->

### Bonus

- Add a health system to both the player character and the enemy character. (look into `Components`)
- Add a simple crouch to your locomotion system.
- more mechanics as you see fit for your game.
- Create a simple `Game Design Document` where you describe your mechanics and how are you planning to implement them.

### Submission

- In your repository there should be a zip file of a build of your game for your target platform.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

### Resources

- [Animation Blueprints](https://dev.epicgames.com/documentation/en-us/unreal-engine/animation-blueprints-in-unreal-engine)
- [Actor Components](https://dev.epicgames.com/documentation/en-us/unreal-engine/components-in-unreal-engine)
- [Mixamo for animations and character models](https://www.mixamo.com/)
- [SketchFab for 3D models](https://sketchfab.com/)
- [itch.io for more assets](https://itch.io/)
