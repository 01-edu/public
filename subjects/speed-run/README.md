## Speed run

### Instructions

You are making the score board of a speed run event.

Speed running means doing something (usually a game) in the least amount of time
possible.

Declare a function `scoreboard` that takes an array of `scores`

Each `score` is an object with 2 properties:

- `name` the name of the participant
- `duration` the duration of the performance of the participant (in seconds)

Operations that must be done by your function:

- Use `.filter` to only select `scores` with a `duration` lower than 12 minutes
- Use `.map` format each results in a string

> format example: `{ name: 'Noura', duration: 235 }` -> `Noura, 235 seconds`

**Complete Example**:

```js
let result = scoreboard([
  { name: 'Noura', duration: 235 },
  { name: 'Maitha', duration: 927 },
  { name: 'Muntaser', duration: 122 },
])

console.log(result) // ['Noura, 235 seconds', 'Muntaser, 122 seconds']
```
