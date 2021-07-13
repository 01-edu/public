## make-your-game

As times passes and technology evolves, the human brain has been requesting
more and more stimuli to keep the brain occupied. Boredom happens whenever
people don't receive enough of these stimuli. And you are not receiving enough
of it, so you decide to do **your own game**. And you also decide to just use
JavaScript as a challenge.

Your game will be for single player only and you will have to create your
own engine and tools so that the game works. And as you try to search all sort
of information you start to remember all those [60fps memes](https://pics.me.me/60-fps-59-fps-35518800.png).

### Objectives

Here are some of the features you want to implement on your game:

- Run the game at least at **60fps** at all time
- You must not have frame drops!
- Proper use of **RequestAnimationFrame**
- It is very hard to predict performances in JS. So measure performances to see if your code is fast. This will be tested!
- Pause menu, that includes:
  - Continue
  - Restart
- A score board that must present the following tasks:
  - **Countdown clock / Timer**: That will indicate the amount of time the player has until the game ends or the time that the game took to finish
  - **Score**: That will keep a score count of the amount of XP/points that the player got while playing
  - **Lives**: That indicates the amount of lives that the player has left

You must not use frameworks or canvas, the game must be implemented using just plain JS/DOM and HTML

### Instructions

Animation must have consistent motion, so in order to have a smooth animation (without interruptions or better named: jank animation) you must achieve a special number, [**60fps**](https://blog.algolia.com/performant-web-animations/). You can see more about performance [here](../good-practices/README.md)

In order to play the game you must use only the keyboard. The usage of keyboard must be smooth, in other words you must not spam the key to move the player. But instead you must, for example maintain the key pressed and the player must continue to do the proper action. If the key is released the player should stop doing the action.

Basically every motion triggered by a key must not jank or stutter.

For the pause menu you must be able to pause, restart and continue the game whenever you want to do so. The frames should not drop if paused.

### Pre-Approved List

Your game will have to respect the genre of one of these games listed below. In other words, the main goal of the game has to be similar to one of these:

- [Bomberman](https://en.wikipedia.org/wiki/Super_Bomberman)
- [Flipper/ Pinball](https://en.wikipedia.org/wiki/Pinball)
- [Space Invaders](https://en.wikipedia.org/wiki/Space_Invaders)
- [Donkey Kong](https://en.wikipedia.org/wiki/Donkey_Kong)
- [Brick Breaker/ Arkanoid](https://en.wikipedia.org/wiki/Arkanoid)
- [Pac-Man](https://en.wikipedia.org/wiki/Pac-Man)
- [Super Mario](https://en.wikipedia.org/wiki/Super_Mario)
- [Tetris](https://en.wikipedia.org/wiki/Tetris)
- [Duck Hunt](https://en.wikipedia.org/wiki/Duck_Hunt)

### Dev Tools

We highly advise you to use Developer Tools present in every browser (usually with the key combination: Ctrl+Shift+I or in the Tools tab).

Below we briefly explain the tools that will help the most for this project:

- **Page Inspector**: You can view and edit page content and layout.
- **Web Console**: You can see your `console.log`s and interact with the page using JavaScript.
- **Performance Tool**: You can analyze your site's general responsiveness, performance, Javascript and layout performance.

The one tool that most will help you is the Performance Tool. There you can record a sample of your actions on the site and analyze the FPS, check for frame drops, how many time is your functions taking to execute, amongst other useful stuff.

In the developer tools you can also find a Paint Flashing option that will highlight every paint that happens in your page as you perform actions in it.

This project will help you learn about:

- [`requestAnimationFrame`](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame)
- [Event loop](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop)
- FPS
- DOM
- [Jank/stutter animation](https://murtada.nl/blog/going-jank-free-achieving-60-fps-smooth-websites)
- [Transform](https://developer.mozilla.org/en-US/docs/Web/CSS/transform)/ [opacity](https://developer.mozilla.org/en-US/docs/Web/CSS/opacity)
- Tasks
  - JavaScript
  - Styles
  - Layout
  - Painting
  - Compositing
- Developer Tools
  - [Firefox](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools)
  - [Chrome](https://developers.google.com/web/tools/chrome-devtools)
