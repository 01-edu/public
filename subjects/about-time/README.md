## About Time

### Instructions

Create the function `ultimateScoreboard` that takes an array of scores, and a
`count` and produce the **ULTIMATE** score board:

- Only show the TOP `count` peoples
- Format durations in `minutes:seconds`

  > ex: `123` -> `02:03` (123 seconds = 2 minutes and 3 seconds)

- Format each scores with first the position in the scoreboard from `1`, then
  the formated `duration` described previously and finally the name

  > ex: `[{ name: 'Wasan', duration: 93 }]` -> `['#01 - 01:33, Wasan']` \
  > note that numbers lower than 10 must be padded with a leading 0

**Example**:

```js
let scores = [
  { name: 'Wasan', duration: 93 },
  { name: 'Muntaser', duration: 122 },
  { name: 'Noura', duration: 235 },
  { name: 'Maitha', duration: 927 },
]

let scoreboard = ultimateScoreboard(scores, 3)

console.log(scoreboard) /* [
  '#01 - 01:33, Wasan',
  '#02 - 02:02, Muntaser',
  '#03 - 03:55, Noura',
]
```

> This time, no more instructions, find the steps yourselfs, good luck !
