## MouseVR

On this one you will learn the basics of interaction using objects, learn the use of blueprint interface and learn stereoscopic sphere texturing. How, you ask? By creating a game in which you will be in a city and will be looking for interaction points in order to travel through the map. And do not to look straight at the end of the subject.

### Objectives

You will have to create a game that uses the mouse as if it were the player head in a VR game. You will navigate through the map using the mouse, focusing in certain points in order to "teleport" to that location.

### Instructions

After creating a blank project and importing the content from the provided folder, you should focus on the following points to accomplish the previous explained features:

- Use Scene Capture to take a 360 degrees screenshot, that you will need later on every targeted location on the map.

  - With this, you can create a new material and apply it to a sphere.
  - Find a way to display the material inside and outside the sphere to create a VR world illusion.

- Create a pawn that can move his head using the mouse inside the spheres.

- In order to play the game, the player only needs to use the mouse and never the keyboard.

- Create your own interaction point to indicate where the player is heading.

  - The design of the interaction point is up to you, but it should display a text underneath it, indicating where this interaction point is going to lead the player to (burger store, metro, park, etc...).
  - There should always be an interaction point visible to the player, somewhere in the world that lets the player quit, no matter where the player is in the map.

- There should be no brutal transitions in the teleportation phase.

- Use an Interface Blueprint to handle interactions: a Blur, a Focus and an Active interaction.

- Use a SphereTraceByChannel node to be able to look at objects.

  - Use casting with your interaction point to know which object you are looking at.

- Use an interaction timer to set how many times should the player look at an interaction point in order to travel to the next space (like on the video example). In other words, the player should look for a while at the interaction point before getting teleported to the next scene.

- Except for the interaction point and the spheres containing the panoramic images, no 3D objects are allowed on the map.

After downloading and unzipping this [file](https://assets.01-edu.org/Unreal-Engine-Piscine/MouseVR.zip), you can copy its content to your project Content folder.

When finished, your project should look like the [“Expected Result” video](https://youtu.be/Tor1Q10NG_Q).

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)

### Bonus

The following are the bonuses for this project:

- Create a circle, around the interaction points, that indicates when the interaction point will be triggered.
- Animate the interaction points.
- Add looping background music during the experience.

[Here](https://www.youtube.com/playlist?list=PLHyAJ_GrRtf9sxZqgfPVM06PrLk8_CWA-) you can find an instructional playlist on Unreal Engine.

Ah! I made you look.
