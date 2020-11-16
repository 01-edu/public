const SIZE = 100
const MAP = new Int8Array(SIZE * SIZE) // State of the Map
const isFree = ({ x, y }) => MAP[y * SIZE + x] === 0 // 0 = block free
const isOccupied = ({ x, y }) => MAP[y * SIZE + x] === 1 // 1 = block occupied
const isMe = ai => ai.me // is true if this AI is you

// `inBounds` check if our coord (n) is an existing index in our MAP
const inBounds = n => n < SIZE && n >= 0

// `isInBounds` check that properties x and y of our argument are both in bounds
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)

// `pickRandom` get a random element from an array
const pickRandom = arr => arr[Math.floor(Math.random() * arr.length)]

// `addToMap` save the new positions into the map
const addToMap = ({ x, y }) => MAP[y * SIZE + x] = 1

// `update` this function is called at each turn
const update = state => {
  // update is called with a state argument that has 3 properties:
  //   ais: an array of all the AIs
  //   ai: the current AI

  // Each AIs contains:
  //   color: A number that represent the color of a AI
  //   name: A string of the AI name
  //   coords: An array of 4 coordinates of the nearest blocks

  // Each coordinate (and AI) contains:
  //   x: The horizontal position
  //   y: The vertical position
  //   index: The computed index of the coord (x * 100 + y)
  //   cardinal: A number between 0 and 3 that represent the cardinal
  //     [ 0: NORTH, 1: EAST, 2: SOUTH, 3: WEST ]
  //        N
  //     W  +  E
  //        S
  //
  //   direction: A number between 0 and 3 that represent the direction
  //     [ 0: FORWARD, 1: RIGHT, 2: BACKWARD, 3: LEFT ]

  // note that AIs initial cardinal is random (depending on the seed)
  // your AI cardinal is the way you are facing

  // Saving state between each updates:
  // I update the MAP with the new position of each AIs
  console.log(state)
  state.ais.forEach(addToMap)

  // Actual AI logic:
  // I filter my array of coords to keep only those that are in bounds
  const coordsInBound = state.ai.coords.filter(isInBounds)

  // I filter again to keep coords that are free
  const available = coordsInBound.filter(isFree)

  // And I return a random available coord
  return pickRandom(available)
}
