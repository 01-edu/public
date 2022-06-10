## Tron

### Objectives

In this project you will have to create your own Tron AI snake.

### Getting started

You will need to create a public repository with the name `tron`. Next you need to create a file named `ai.js`. It must respect the instructions given.

### Controls

- `arrows` or `scroll` to move step by step
- `shift` will make it fast
- you can click anywhere on the progress bar to seek into the history

### Rules

- Your AI has to move every turn _(it can not stay still)_
- Every time the AI moves somewhere the AI leaves a color trail.
- the AI can only move to a blank tile.
- the AI can not move out of the map _(100 x 100)_
- the AI can only move to its `left`, `forward` or its `right`.
  _(Moving `backward` is suicide as it would hit its own trail !)_
- If too much CPU power is required to decide where to go, the AI dies.
- If two AIs moved to the same spot, both of them die.
- **The AI has to survive as long as it can.**

### The game ends

- Once no players can make a move the player with the biggest score wins

### Useful files
- [index.html (Game Engine)](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/game_students/index.html)

- [hard.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/hard.js)
- [license-to-kill.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/license-to-kill.js)
- [random.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/random.js)
- [right.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/right.js)
- [snail.js](https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/snail.js)


### How to write your AI

- Create `ai.js` at the root of your repository.
- Copy the contents of `random.js`, and paste it to `ai.js`.
- You may now edit the `update` function which is called each turn.

> ⚠️ Do not rename the `update` function ⚠️ \
> as it's the function that the worker will try to run to test your AI.

### How to test your AI
**AI AT ROOT**
- You may test your ai locally. For that, create a folder and give it a name. Inside the folder you created, insert this file `index.html`.
- After that, create a simple web server by running the following command:

```sh
$ &>/dev/null python3 -m http.server &
```

Now open your browser at the specified port. You'll use an appropriate command for your system:

- Linux: `xdg-open`
- macOS: `open`
- Windows: `start`

```sh
xdg-open 'http://localhost:8000'
```

- You can set a seed by adding the variable `seed` to the url params.
- You can add up to two AI's by adding the variable `ai` to the url params. The AI's will be separated by a `+`

  - You can add a local file by specifying the relative path.
  - You can add a online raw file by specifying the url to that file.

A example of a url with local files using the default AI `ai.js` against the AI `hard.js` would be `http://localhost:8000/?ai=hard.js+ai.js&seed=2077349364`.

A example of a url with online files would be `http:?/localhost:8000/?ai=https://((DOMAIN))/git/root/public/raw/branch/master/subjects/tron/ai/hard.js+ai.js`.

Note: You can test a local AI against and online one.

- Open the inspector of the browser used and **disable the cache**
- let's change the update function so that your AI only goes forward.

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

- save the file, push the changes and re-run the game in the browser.
  If the cache was correctly disabled,
  you have changed your AI behaviour from a random pick of available moves
  to only going forward.

- To understand better the way of controlling your AI,
  read the comments inside the AI file and do a lot of testing.

- When peer-corrected, you AI will be competing against other AIs.
  Be aware that there will be the possibility for the peer-correcter
  to use his or her own AI.

_May the best tron win :)_

Have fun and good luck.
