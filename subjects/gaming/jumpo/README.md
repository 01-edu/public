## jumpo

In this project, you will have to create a 2D sidescroller game. You will be using for the first time an external SDK to compile your game: the Android SDK Tool for Unreal Engine, which allows you to create a [.apk file](https://fileinfo.com/extension/apk) that can be installed and played on any Android devices.

### Objectives

In this game for Android, the player will control a character in a similar game than the Google Chrome Running Dinosaur but with coins or gems. So the game will have buttons on screen and that will be bound to character actions like jumping, ducking, etc... Another goal of this project is the sprite and [flipbook](https://docs.unrealengine.com/en-US/AnimatingObjects/Paper2D/Flipbooks/index.html) creation that is used to animate 2D art.

The basic assets you will need for this project can be found in the ["Jumpo.zip"](https://assets.01-edu.org/Unreal-Engine-Projects/Jumpo/Jumpo.zip) file.

### Instructions

The following are the requirements for this project:

- This game should have a main menu that must have a:

  - Start button, that loads the game.
  - Stats button, that shows you how many coins you have collected overall, how many times you died, the number of games done and the highest score. You can add more information if you want to.
  - Leaderboard button, that shows the 5 highest scores ever made.
  - Credits button, that shows the names of the developers that worked on the project.
  - Character button, that shows the locked and unlocked characters.
  - Quit button, that closes the app.

#### Map

These are the requirements regarding the game map:

- Create a 2D map with obstacles that the player has to dodge by jumping or ducking.
- No 3D objects are allowed in any scene of the game.
- The game is meant to be played in landscape mode.

#### Character and Gameplay

For the character these are the requisites:

- You are free to choose which character you want to use in your game. You can use the resource ones or you can check [this website](https://www.spriters-resource.com/) for more characters and sprites.
- The character should be "running" from left to right, so the character must be headed right.
  - The character should actually be static in relation to the background and only the background is moving from right to left, giving the illusion that the character is the one moving.
  - Once a part of the background is out of sight, it gets destroyed or moved to the right.
- If the character does not die, the map should be infinite.
- Each action (running, jumping, ducking and dying) should play a different animation (a different flipbook).
  - Use the .png or .tga files to extract and create a flipbook.
- While running, obstacles should appear on the character path.
  - The character can either jump or roll under obstacles to survive.
- When the character hits an obstacle, it dies and the die flipbook animation gets played.
  - When the character dies, a score screen is displayed and the player needs to enter 3 characters to save his score in the Leaderboard (a sort of nickname).
- The more time the character survives, the more points he earns.
- The character can collect coins/gems to load an invincibility bar.
  - The number of coins collected during that run should appear on the screen as well as the progress of the invincibility bar.
  - Once the invincibility bar is loaded up, the player can press it and the character becomes invincible for a short period of time.
  - When the invincibility bar gets emptied, the character comes back to normal.
- The jump and roll button should be accessible using thumbs when holding the phone (in the lower corners).
- During the game, the highest ever score should appear on the top middle of the screen and the current score right below it.
- A pause button should appear somewhere on the screen but should not disturb the player experience. The top right corner is a good place for it.
- When the player presses the pause button or when the character dies, three buttons should appear:
  - Retry -> Restarts the game and all progress is lost if the game is in pause.
  - Menu -> Goes back to the main menu and all progress is lost if the game is in pause.
  - Quit -> Quits the entire app.

#### Other Game Features

In addition to the requirements of the map and the character, these are other aspects that need to be taken into account:

- When closing the app and launching it again, all the player data should remain as he leaved it before.

On the main menu:

- On the "Character" tab, the player sees grayed-out characters. These are the locked characters.
  - This characters can be unlocked by spending coins/gems collected.
  - When first opening the game, there should be only one character unlocked.
- When a player dies, the Leaderboard should be updated accordingly to the score the player just made if it is on the 5 best score ever.
  - Also the Stats should be updated after every game, whether it is or it is not one of the best 5 scores ever.
- On the Leaderboard and on the Stats tabs, a button to reset all the stats should be available.

#### Android

In order to complete this project you are going to need to run your game in Android devices. For that you will need to install an Android SDK. Here is the Unreal Engine [documentation](https://docs.unrealengine.com/en-US/SharingAndReleasing/Mobile/Android/Setup/AndroidStudio/index.html) for that and a [video](https://www.youtube.com/watch?v=q-mAEnqZb0M) that could also help you.

- You could use your own android device or you could use an emulator to install the app on it and run the app to confirm that everything is working properly on Android devices.

### Hints

If you want a good example of how a 2D game is created you can try the [Tappy Chicken](https://assets.01-edu.org/Unreal-Engine-Projects/Jumpo/TappyChicken.zip) game provided. Or you can create an empty project and choose "2D Sidescroller Game" template which contains a lot of info that will be helpful for this project.

> Do not forget to zip up the project compile and save everything for peer correction.
> If it is not possible to upload files to Gitea due to their size, use GitHub instead and have a look at [Git LSF](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github)
