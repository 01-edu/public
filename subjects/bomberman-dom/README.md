## bomberman-dom

You certainly know [bomberman](https://en.wikipedia.org/wiki/Bomberman) right? Good good. You will make it. Relax, is not that hard, it is only a multiplayer version of it. Ah and forgot to mention that you will need to do it using the framework you created a while ago. Let me explain.

### Objectives

For this project you have to create a bomberman alike game, where multiple players can join in and battle until one of them is the last man standing.

### Instructions

In the beginning there was 4 players, and only one came out alive. Each player will have to start in the different corners of the map and only one will be victorious.

You will have to follow more a less the same principles has the [make-your-game](https://public.01-edu.org/subjects/make-your-game/) project. But we will refresh one of the concepts you will have to respect and deal with:

- [**Performance**](https://public.01-edu.org/subjects/good-practices/), is one of the most important aspects while developing a game, so lets respect it.\
   Just like make-your-game you will have to respect the policy of:

  - Running the game at least at **60fps** at all time
  - No frame drops
  - Proper use of [**`requestAnimationFrame`**](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame)
  - Measuring performance to know if your code is fast

You must not use canvas neither [Web-GL](https://get.webgl.org/) or other framework. For this project you will use the framework you did on the [mini-framework](https://public.01-edu.org/subjects/mini-framework/) project.

You will have to make also a chat that enables the different players to talk to each other. As you will have to use [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API), this **chat can be considered as a "Hello World" of the multiplayer feature** for the **bomberman-dom**.

#### Game Mechanics

1. Players

   1. Nº of players: 2 - 4
   2. Each player must have 3 lives. Then you are out!!

2. The map should be fixed so that every player sees the whole map.

   1. The maps should be generated with a logical random structure (without blocking the player). There will be two types of blocks, the ones that can be destroyed and the ones that can not. Tip: the optional project [different maps](https://public.01-edu.org/subjects/make-your-game/different-maps) can be useful for this part.
   2. The players should be placed in the corners as their starting positions.

3. Power ups (each time a player destroys a block, a random power up may or may not appear):

   1. Bombs: Increases the amount of bombs dropped at a time by 1;
   2. Flames: Increases explosion range from the bomb in four directions by 1 block;
   3. Speed: Increases movement speed;
   4. Bomb Push: Ability to throw a bomb after it has been placed.

When the user opens the game, he/she should be presented to a page where he/she should enter a **nickname** to differentiate users. After selecting a nickname the user should be presented to a waiting page with a **player counter** that ends at 4. Once a user joins, the player counter will increment by 1.

If there are more than 2 players in the counter and it does not reach 4 players before 20 seconds, a 10 second timer starts, to players get ready to start the game.\
If there are 4 players in the counter before 20 seconds, the 10 seconds timer starts and the game starts.

This project will help you learn about:

- [`requestAnimationFrame`](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame)
- [Event loop](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop)
- FPS
- [Jank/stutter animation](https://murtada.nl/blog/going-jank-free-achieving-60-fps-smooth-websites)
- webSockets
- Synchronization
- Developer Tools
  - [Firefox](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools)
  - [Chrome](https://developers.google.com/web/tools/chrome-devtools)
