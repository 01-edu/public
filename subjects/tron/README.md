## Tron

### Objectives

In this project you will have to create your own Tron AI snake

### Getting started

You will need to create a public repository with the name `tron`. Next you need to create a file named `ai.js`. It must respect the instructions given

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

### How to write your AI

- Copy the code on the file [random.js](https://raw.githubusercontent.com/01-edu/public/master/subjects/tron/ai/random.js) to your file, `ai.js`
- You may now edit the `update` function which is called each turn

> ⚠️ Do not rename the `update` function ⚠️ \
> as it's the function that the worker will try to run to test your AI.

### How to test your AI

- You may use this link [tron](/public/subjects/tron/?ai=&seed=1653547275), to test your AI
- You need to add your AI as a user in that link
  > Example:
  - if your git login is **Frenchris** and you want to test against **LEEDASILVA** the link becomes: [/public/subjects/tron/?ai=Frenchris@master+LEEDASILVA@master](/public/subjects/tron/?ai=Frenchris@master+LEEDASILVA@master)
  - if you want to test against the default `/random.js` AI the link becomes: [/public/subjects/tron/?ai=Frenchris@master+/random.js](/public/subjects/tron?ai=Frenchris@master+/random.js)
- Open the inspector of the browser used and **disable the cache**
- let's change the update function so that your AI only goes forward.

Replace this line just before the `return` of the update function:

```js
const available = coordsInBound.filter(isFree)

// And I return a random available coord
return pickRandom(available)
```

...with this line:

```js
// always return the first free coordinates
return coordsInBound.filter(isFree)[0]
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
