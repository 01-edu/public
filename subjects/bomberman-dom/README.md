## bomberman-dom

You certainly know [bomberman](https://en.wikipedia.org/wiki/Bomberman) right? Good good. You will make it. Relax, is not that hard, it is only a multiplayer version of it. Ah and forgot to mention that you will need to do it using the framework you created a while ago. Let me explain.

### Objectives

For this project you have to create a bomberman alike game, where multiple players can join in and battle until one of them is the last man standing.

### Instructions

In the beginning there are 4 players, and only one came out alive. Each player will have to start in the different corners of the map and only one will be victorious.

You will have to follow more or less the same principles of [make-your-game](../make-your-game/README.md) project. But we will refresh one of the concepts you will have to respect and deal with:

- [**Performance**](../good-practices/README.md), is one of the most important aspects while developing a game, so let's respect it.\
   Just like make-your-game you will have to respect the policy of:

  - Running the game at least at **60fps** at all time
  - No frame drops
  - Proper use of [**`requestAnimationFrame`**](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame)
  - Measuring performance to know if your code is fast

You must not use canvas, neither [Web-GL](https://get.webgl.org/) nor another framework. For this project you will use the framework you did on the [mini-framework](../mini-framework/) project.

You will also have to make a chat that enables the different players to talk to each other. You will have to use [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API). This chat **can be considered as a "Hello World" of the multiplayer feature** for the **bomberman-dom**.

#### Game Mechanics

1. Players

   1. Nº of players: 2 - 4
   2. Each player must have 3 lives. Then you are out!!

2. Map

   1. The map should be fixed so that every player sees the whole map.
   2. There will be two types of blocks, the ones that can be destroyed (blocks) and the ones that can not (walls).
      1. The walls will always be placed in the same place, while the blocks are meant to be generated randomly on the map. Tip: the optional project [different maps](../make-your-game/different-maps/README.md) can be useful for this part.
      2. In the starting positions the players need to be able to survive. For example: if the players place a bomb, they will need to have space to avoid the bomb explosion.
   3. The players should be placed in the corners as their starting positions.

3. Power ups (each time a player destroys a block, a random power up may or may not appear):

   1. Bombs: Increases the amount of bombs dropped at a time by 1;
   2. Flames: Increases explosion range from the bomb in four directions by 1 block;
   3. Speed: Increases movement speed;

When the user opens the game, he/she should be presented to a page where he/she should enter a **nickname** to differentiate users. After selecting a nickname the user should be presented to a waiting page with a **player counter** that ends at 4. Once a user joins, the player counter will increment by 1.

If there are more than 2 players in the counter and it does not reach 4 players before 20 seconds, a 10 second timer starts, to players get ready to start the game.\
If there are 4 players in the counter before 20 seconds, the 10 seconds timer starts and the game starts.

#### Bonus

Although this bomberman already is super cool, it can be always better. Here are some ideas you can implement into the game to make it super awesomely cool:

- **Solo + Co-Op mode**: You are supposed to develop an AI to play against the players. So once the AI is defeated all players involved win.
- **More power ups**:
  - _Bomb Push_: Ability to throw a bomb after it has been placed;
  - _Bomb Pass_: Ability to pass through bombs;
  - _Block Pass_: Ability to pass through blocks (not walls);
  - _Detonator_: Ability to choose when a bomb will explode on a key press;
  - _1 Up_: Gives the player an extra life;
- **Release power ups after death**: When a player dies it drops one of it's power ups. If the player had no power ups, it drops a random power up.
- **Team mode**: Make games with 2v2 (two players versus two players).
- **After death interaction**: When a player dies, they can reappear as a ghost. If a ghost touches another player they come back to life. If a ghost is caught in a bomb explosion, the player controlling the ghost dies permanently.

This project will help you learn about:

- [`requestAnimationFrame`](https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame)
- [Event loop](https://developer.mozilla.org/en-US/docs/Web/JavaScript/EventLoop)
- FPS
- [Animation performance and frame rate](https://developer.mozilla.org/en-US/docs/Web/Performance/Animation_performance_and_frame_rate)
- webSockets
- Synchronization
- Developer Tools
  - [Firefox](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools)
  - [Chrome](https://developers.google.com/web/tools/chrome-devtools)
