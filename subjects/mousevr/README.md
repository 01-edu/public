## MouseVR

On this one you will learn the basics of interaction using objects, learn the use of blueprint interface, learn stereoscopic sphere texturing. How, you ask? By creating a game in which you'll be in a city and will be looking for interaction points in order to travel through the map. And do not to look straight at the end of the subject.

### Objectives

You will have to create a game that uses the mouse as if it were the player head in a VR game. You will navigate through the map using the mouse, focusing in certain points in order to "teleport" to that location.

### Instructions

After creating a blank project and importing the content from the provided folder, you should focus on the following points to accomplish the previous explained features:

- Use Scene Capture to take a 360 degrees screenshot, that you’ll need later on every targeted location on the map.
  - With this, you can create a new material and apply it to a sphere.
  - Find a way to display the material inside and outside the sphere to create a VR world illusion.
- Create a pawn that can move his head using the mouse inside the spheres.
- In order to play the game, the player only needs to use the mouse and never the keyboard.
- Create your own interaction point to indicate where the player is heading.
  - The design of the interaction point is up to you, but it should display a text underneath it, indicating where this interaction point is gonna lead the player to (burger store, metro, park, etc...)
  - There should always be an interaction point visible to the player, somewhere in the world that lets the player quit, no matter where the player is in the map.
- There should be no brutal transitions in the teleportation phase.
- Use an Interface Blueprint to handle interactions: a Blur, a Focus and an Active interaction.
- Use SphereTraceByChannel node to be able to look at objects.
  - Use casting with your interaction point to know which object you are looking at.
- Use interaction timer to set how many time should the player look to an interaction point, in order to travel to the next space (Like on the video example). In other words, the player should look for a while at the interaction point before getting teleported to the next scene.
- Except for the interaction point and the spheres containing the panoramic images, no 3D objects are allowed on the map.

When finished, your project should look like the executable example on the folder or the “Expected Result” video.

> Don’t forget to zip up the project compile and save everything for peer correction.

### Bonus

As bonus for this project there are the following:

- Create a circle around the interaction points, that indicates when the interaction point will be triggered.
- Animate the interaction points.
- Add looping background music during the experience.

Ah, made you look.
