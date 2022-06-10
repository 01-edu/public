## Tron

### Objectives

Create your own Tron AI snake. You will be provided with a game engine, and your objective will be to create an AI player which will battle against other AI players to win the game.

### Useful files
- [index.html (Game Engine)](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/game_students/index.html)
- [hard.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/hard.js)
- [license-to-kill.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/license-to-kill.js)
- [random.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/random.js)
- [right.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/right.js)
- [snail.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/snail.js)

### Rules
The objective is for the AI player to stay alive as long as possible. The game ends when no player can move, or both players are dead.

The player with the highest score wins.

- AI players cannot stay still during a turn, they must make a move.
- Every time the AI player moves, it leaves a colorful trail.
- The AI player can only move to a blank tile, and cannot move outside of the 100 x 100 map.
- AI players can move `left`, `right` or `forward`. Moving backwards kills the AI player, because of the colorful trail.
- If too much CPU power is required to decide where to move, the AI player dies.
- If two AI players move to the same spot, both of them die.

### Controls
You can use the arrows or scroll to move step by step through the game. `shift` will make it fast. You can also click anywhere on the progress bar to seek through the history.

### How to write your AI
- Create `ai.js` at the root of your repository.
- Copy the contents of `random.js`, and paste it to `ai.js`.
- Edit the `update` function which is called each turn.

> Do not rename the `update` function. It is the function that the worker will try to run to test your AI player.

### Preparing to test your AI player
Let's say that you want to battle `ai.js` against `random.js`. You will need to run the game engine, and specify the local path or web path to both of those files.

For simplicity, we'll use a single directory to demonstrate. So place `index.html` and `random.js` in the same directory as `ai.js`.

You'll need to create a simple web server. From the command line, `cd` into that directory, and execute the following command:

```sh
$ python3 -m http.server
Serving HTTP on :: port 8000 (http://[::]:8000/)
```

Now open your browser at the specified port. For us, that will be where the `http.server` is running:
```
http://localhost:8000
```

You can modify the path to load the players using a relative path:
```
http://localhost:8000/?ai=random.js+ai.js
```

Each game has a `seed` which will spawn the players in different positions. When you modify the path, it will add a `seed` to the query path. You can change the `seed` to spawn the players elsewhere, or remove it for a random one to be generated.
```
http://localhost:8000/?ai=random.js+ai.js&seed=207734936
```

You can also specify a web path to some `.js` file like so:
```
http://localhost:8000/?ai=https://example.com/example.js+ai.js
```

Open the inspector of the browser and **disable the cache**.

### Getting started with improving your AI player
Let's change the `update` function so that your AI only goes forward.

Replace this line just before the `return` of the update function:

```js
const available = coordsInBound.filter(isFree);

// And I return a random available coord
return pickRandom(available);
```

...with this line:

```js
// always return the first free coordinates
return coordsInBound.filter(isFree)[0];
```

Save the file, and re-run the game. If the cache was correctly disabled, you will have changed the behavior of your AI player from random movements, to only making forward moves.

To better understand controlling your AI player, read the comments inside `ai.js`, and do a lot of testing.

During the audit, you will be competing against other AI players. Be aware, that there is a possibility for the auditor to use their own AI.

_May the best tron win :)_

Have fun and good luck.
