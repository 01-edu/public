## tron

### Objectives

In this project you will have to create your own Tron Ai snake

### Getting started

First, clone this repo : https://github.com/01-edu/tronBlank

then execute the command below.

```sh
http-server tronBlank/
```

if http-server is not installed do this command:

```sh
npm install -g http-server
```

### Controls

- `space` toggles auto-play
- `right arrow` plays one move
- `up arrow` increases the auto-play speed
- `down arrow` lowers the auto-play speed
- `R` reloads the same play
- `S` loads a new play (new seed)

### Rules

- Your Ai has to move every turn *(it can not stay still)*
- Every time the Ai moves somewhere the Ai leaves a color trail.
- the Ai can only move to a blank tile.
- the Ai can not move out of the map *(100 x 100)*
- the Ai can only move to its `left`, `forward` or its `right`.
  *(Moving `backward` is suicide as it would hit its own trail !)*
- If too much CPU power is required to decide where to go, the Ai dies.
- If two Ais moved to the same spot, both of them die.
- **The Ai has to survive as long as it can.**

### The game ends

- Once no players can make a move
- The player with the longest trail wins

### How to write your AI

- Copy the file [/ai/random.js](https://github.com/01-edu/tronBlank/blob/master/ai/random.js) to /ai/**GITHUB_LOGIN**.js
- You may now edit the `update` function which is called each turn

### How to test your AI

- Suppose the link provided by http-server is : `http://127.0.0.1:8080/?seed=somevalue&users=random`
- You need to add your ai as a user in that link,
- Example: if your ai file's name is **/ai/Frenchris.js**
the link becomes:
`http://127.0.0.1:8080/?seed=somevalue&users=Frenchris,random`

- Open the inspector of the browser used and **disable the cache**

- let's change the update function so that your ai only goes forward.

Replace this line just before the `return` of the update function

```js
 const available = coordsInBound.filter(isFree)
 ```

With this line

```js
const available = coordsInBound.filter(isFree).filter(el => el.direction === 0)
```

- save the file and rerun the game in the browser. If the cache was correctly disabled,
you have changed your ai behaviour from a random pick of available moves to only going
forward.

- To understand better the way of controlling your AI, read the comments inside the Ai file and do a lot of testing.

- When peer-corrected, you Ai will be competing against other Ais.
Be aware that there will be the possibility for the peer-correcter to use his or her own Ai.

May the best tron win :)

Have fun and good luck.
