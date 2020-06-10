const SIZE = 100
const [FREE, FILLED] = Array(3).keys()
const [FORWARD, RIGHT, BACKWARD, LEFT] = Array(4).keys()
const [NORTH, EAST, SOUTH, WEST] = Array(4).keys()
const MAP = new Int8Array(SIZE * SIZE)
const isFree = ({ x, y }) => MAP[y * SIZE + x] === FREE
const setFilled = ({ x, y }) => MAP[y * SIZE + x] = FILLED
const inBounds = n => n < SIZE && n >= 0
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)
const isNotBackward = el => el.direction !== BACKWARD
const isForward = el => el.direction === FORWARD
const isRight = el => el.direction === RIGHT
const isLeft = el => el.direction === LEFT
const goForward = arr => arr.find(isForward)
const goRight = arr => arr.find(isRight)
const goLeft = arr => arr.find(isLeft)
const isOtherAI = ai => !ai.me
const isAlive = ai => !ai.dead

const pickDirection = arr => goLeft(arr) || goForward(arr) || goRight(arr)

const getNearestWallCardinal = ai => {
  const n = ai.y
  const e = 99 - ai.x
  const s = 99 - ai.y
  const w = ai.x
  const smallest = Math.min(n, e, s, w)
  if (smallest === n) return NORTH
  if (smallest === e) return EAST
  if (smallest === s) return SOUTH
  if (smallest === w) return WEST
}

let reachForWall = true
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

  if (reachForWall) {
    const cardinal = getNearestWallCardinal(ai)
    const coord = ai.coords.find(c => c.cardinal === cardinal)
    if (safeCoords.includes(coord)) return coord
  }

  reachForWall = false
  return pickDirection(safeCoords) || pickDirection(possibleCoords)
}









// hej