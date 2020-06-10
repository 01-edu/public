const SIZE = 100
const FREE = 0
const FILLED = 1
const MAP = new Int8Array(SIZE * SIZE)
const isFree = ({ x, y }) => MAP[y * SIZE + x] === FREE
const setFilled = ({ x, y }) => MAP[y * SIZE + x] = FILLED
const inBounds = n => n < SIZE && n >= 0
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)
const isNotBackward = el => el.direction !== 2
const isForward = el => el.direction === 0
const isRight = el => el.direction === 1
const isLeft = el => el.direction === 3
const goForward = arr => arr.find(isForward)
const goRight = arr => arr.find(isRight)
const goLeft = arr => arr.find(isLeft)
const isAlive = ai => !ai.dead
const isOtherAI = ai => !ai.me

// `snailIt` goes with the form of a snail
const snailIt = arr => goRight(arr) || goForward(arr) || goLeft(arr)

const update = ({ ai, ais }) => {
  ais.forEach(setFilled)

  const enemyCoordsIndices = ais
    .filter(isOtherAI)
    .filter(isAlive)
    .flatMap(ai => ai.coords)
    .map(coord => coord.index)

  const possibleCoords = ai.coords
    .filter(isInBounds)
    .filter(isFree)

  const safeCoords = possibleCoords.filter(c => !enemyCoordsIndices.includes(c.index))

  const coordsInBound = ai.coords
    .filter(isInBounds)
    .filter(isNotBackward)

  return snailIt(safeCoords) || snailIt(possibleCoords)
}
