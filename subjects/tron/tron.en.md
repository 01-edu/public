## tron

### Objectives

In this project you will have to create your own Tron AI snake

### Getting started

You will need to create a public repository with the name `tron`. Next you need to create a file named `ai.js`. It must respect the instructions given

### Controls

- `right arrow` plays one moves frontwards
- `left arrow` plays one moves backwards
- You can use the scroll to do the same thing as above
- `shift` will make it fast

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

- Once no players can make a move the player with the biggest score wins

### How to write your AI

- Copy the code on the file [random.js](https://raw.githubusercontent.com/01-edu/public/master/subjects/tron/random.js) to your file, `ai.js`
- You may now edit the `update` function which is called each turn

### How to test your AI

- You may use this link [tron](https://01.alem.school/public/subjects/tron/?ai=&seed=1653547275), to test your AI
- You need to add your AI as a user in that link,
- Example: if your git login is **Frenchris** and you want to test against **LEEDASILVA**
the link becomes:
`https://01.alem.school/public/subjects/tron/?ai=Frenchris@master+LEEDASILVA@master&seed=1653547275`

- Open the inspector of the browser used and **disable the cache**

- let's change the update function so that your AI only goes forward.

Replace this line just before the `return` of the update function

```js
 const available = coordsInBound.filter(isFree)
 ```

With this line

```js
const available = coordsInBound.filter(isFree).filter(el => el.direction === 0)
```

- save the file, push the changes and rerun the game in the browser. If the cache was correctly disabled,
you have changed your AI behaviour from a random pick of available moves to only going forward.

- To understand better the way of controlling your AI, read the comments inside the AI file and do a lot of testing.

- When peer-corrected, you AI will be competing against other AIs.
Be aware that there will be the possibility for the peer-correcter to use his or her own AI.

May the best tron win :)

Have fun and good luck.
