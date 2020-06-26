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
- You will use **RequestAnimationFrame**
- It is very hard to predict performances in JS. So measure performances,
  to know if your code is fast. This will be tested

### Instructions

Animation must have consistent motion, so in order to have a smooth animation (without interruptions or better named: jank animation) you must achieve a special number, [**60fps**](https://blog.algolia.com/performant-web-animations/).

Performance is essential, so that's why you have to aim for less than 16.7ms(1000ms/60f), because in 16.7ms the browsers job and your work must be completed in each frame.

Use [`requestAnimationFrame`](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame) to sync your changes

- You can try to reuse memory to avoid jank when the garbage collector pass

Jank animation can be caused by loading too much information, for instance:

- **JavaScript**
- **Styles** that considers which style apply to which element
- **Layout** that calculates the geometry of the pages (example: recalculating the width and height of a page)
- **Painting**, normally when layout is triggered we must repaint. Repainting an element every time it animates
- **Compositing** that consists on placing the pages together at the end

In order to improve performance we must remove the causes above, that are more costly: layout and painting.

The best way to remove layout and painting is to use [`transform`](https://developer.mozilla.org/en-US/docs/Web/CSS/transform) and [`opacity`](https://developer.mozilla.org/en-US/docs/Web/CSS/opacity)

Removing layout can be done using transform:

```js
// bad
// this will trigger the layout to recalculate everything and repaint it again
box.style.left = `${x * 100}px`;

// good
// this way its possible to lose the layout
box.style.transform = `translateX(${x * 100}px)`;
```

It is possible to remove painting by adding a layer:

```css
/* this will take care of the painting by creating a layer and transform it*/
  #box {
    width: 100px;
    height: 100px;
    ....
    will-change: transform;
  }
```

By creating a new layer you can remove painting, but "there is always a tradeoff". If we add to much layers it will increase the **composition** and **update tree**. In conclusion you must promote a new layer only if you now you are going to use it. Performance is the key to a good animation. "Performance is the art of avoiding work".

### Pre-Approved List

Your game will have to respect the genre of one of these games listed below. In other words, the main goal of the game has to be similar to one of these:

- [Flipper/ Pinball](https://en.wikipedia.org/wiki/Pinball)
- [Space Invaders](https://en.wikipedia.org/wiki/Space_Invaders)
- [Donkey Kong](https://en.wikipedia.org/wiki/Donkey_Kong)
- [Brick Breaker/ Arkanoid](https://en.wikipedia.org/wiki/Arkanoid)
- [Pac-Man](https://pt.wikipedia.org/wiki/Pac-Man)
- [Super Mario](https://en.wikipedia.org/wiki/Super_Mario)
- [Tetris](https://pt.wikipedia.org/wiki/Tetris)
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
- [Event loop](https://developer.mozilla.org/pt-BR/docs/Web/JavaScript/EventLoop)
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
