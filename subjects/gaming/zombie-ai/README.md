## ZombieAI

The goal of this exercise is to create a zombie game with a wave system similar to the Call of Duty zombie mode. Do you know what it takes to become a zombie? Dead-ication.

### Objectives

You will use the content in the Resources folder to create a zombie wave game. The player will be a character with one of the guns and zombies will be coming for him from various points in the map.

Also you will use the solution from one of the previous exercises.

### Instructions

These are some the features that the game should have:

- a Menu map where you will display 3 options:

  - Start Game (Load the map and launch the game).
  - Settings (Migrate your WidgetFactory project to integrate it to this one).
  - Quit Game (Exits the game).

After entering the game by clicking the "Start Game" button, these are the game requirements:

- a widget on the screen should be visible for the player life and should also contain the wave number that the player is in.

- an enemy spawning system from a blueprint that can be placed around the world.

  - zombies should never all spawn at the same time. Zombies should spawn one by one every `N` second(s)(you decide what N is).

- create the animation for the Zombie. There should be at least two animations (running and attacking). These actions have to be performed by the Behavior Tree.

  - on spawn, the Zombie should detect the player location and run towards it, following the Behavior Tree that you created earlier.

- when shooting, a fire projectile has to be casted and shot from the muzzle of the rifle.

- when the Zombie is near the player, the zombie should hit the player.
  - when the player is hit, his life bar should decrease.
  - after 5 hits the player should die.
  - the life of the player should slowly get back up, 2 seconds after the hit. If the player gets hit again, wait 2 seconds again to give him back his health.
- on zombie component hit with the projectile, the zombie should lose health when it gets hit and should not die in one shot.

  - after being hit multiple times and dying, the death of the zombie has to be noticeable.

- the zombie spawning and the recovery of health should be handled by the node "Set Timer by Function Name".

- the wave system should be as simple as it sounds. When you kill a certain amount of zombies in the first round, you calculate the number of zombie you want for the second round and spawn them (You can maybe get more creative than zombiesNumber \* 2 or \* wave).

  - more zombies should spawn as the waves increases.

- you have the freedom to create the map and design the level. Just be sure that the Navigation Mesh Bound covers the whole map or else you will have trouble with the zombie not coming for the player.

- every projectile and zombie should be destroyed when they are no more visible/alive.

- the sound from the shot and from the zombie have to be handled.

  - zombie sound should be following the zombie and the sound should increase or decrease depending wether the player gets closer or futher from the zombie.
  - when you shoot, two sounds must be heard: the shot sound and the shells of the round hitting the ground.

- when the player is dead at least two information have to be displayed on the screen:
  - the number of zombies killed.
  - the wave during which the player died.
    - these two stats also have to be displayed when the player presses the TAB button and be hidden when the TAB button is released.

After downloading and unzipping this [file](https://assets.01-edu.org/Unreal-Engine-Piscine/ZombieAI.zip), you can copy its content to your project Content folder.

When finished, your project should look like the [“Expected Result” video](https://youtu.be/d8MqIVuC88k).

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

### Bonus

To make the game more interesting you could:

- Implement a way of receiving power-ups/benefits when a zombie dies (like get more health, make the zombies slower, or other things).
- Make a 3D main menu.
- Implement a point system, in which every time the player kills a zombie, pass wave, etc…, the points gets higher.
- Add creepy music during the game, add sound at each new wave, etc…
- Make the player have to reload the weapon, in other words, make the gun go out of ammo.

[Here](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf9sxZqgfPVM06PrLk8_CWA-) you can find an instructional playlist on Unreal Engine.
