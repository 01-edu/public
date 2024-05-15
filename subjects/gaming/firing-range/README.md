## FiringRange

In this exercise, you will learn to use the Unreal Engine and Blueprints to script functionalities of a basic FPS game. The theme is to reproduce a firing range. And remember, weapons are allowed only inside the shooting area!!

### Instructions

The map of this project should be composed of a cube with dimensions of X = 35, Y = 40, Z = 1. It should simulate the floor and other cubes as walls around the floor. In the map there should be a zone where the character is able to walk around and shoot the targets and another zone where the player can not go and where the targets are present.

For this project you will have to create a Blueprint Class target, that will have some of these characteristics. The target should :

- have associated to it the previous created material.
- either be moving from side to side or be stationary.
- be dynamic, using the timeline node.
- use a public variable to set or unset the movement animation of the target.
- rise again after x seconds after being hit it with a projectile, and behave like before it was hit.

Only one class of target is allowed for the whole project.

The previously mentioned projectile should:

- have a size of X = Y = Z = 0,5.
- have a speed of 10000.

After downloading and unzipping this [file](https://assets.01-edu.org/Unreal-Engine-Piscine/FiringRange.zip), you can copy its content to your project Content folder.

When finished, your project should look like the executable example on the folder or the ["Expected Result" video](https://youtu.be/EBibaN-dh_0).

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

#### Bonus

Here are some ideas for improving the game:

- Targets with different speeds.
- Textures on the walls and ground.
- Adding obstacles in front of the targets.

[Here](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf9sxZqgfPVM06PrLk8_CWA-) you can find an instructional playlist on Unreal Engine.
