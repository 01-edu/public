const SIZE = 100
const [FREE, UNSAFE, FILLED] = Array(3).keys()
const MAP = new Int8Array(SIZE * SIZE)
const isFree = ({ x, y }) => MAP[y * SIZE + x] < FILLED
const isUnsafe = ({ x, y }) => MAP[y * SIZE + x] === UNSAFE
const setUnsafe = ({ x, y }) => MAP[y * SIZE + x] = UNSAFE
const setFilled = ({ x, y }) => MAP[y * SIZE + x] = FILLED
const inBounds = n => n < SIZE && n >= 0
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)
const isForward = el => el.direction === 0
const isRight = el => el.direction === 1
const isLeft = el => el.direction === 3
const goForward = arr => arr.find(isForward)
const goRight = arr => arr.find(isRight)
const goLeft = arr => arr.find(isLeft)
const isOtherPlayer = player => !player.isOwnPlayer

let wallReached = false
const pickForwardOrRightOrLeft = arr => {
  if (arr.length === 3 && !wallReached) {
    return goForward(arr) || goRight(arr) || goLeft(arr)
  }

  if (arr.length <= 2 && !wallReached) {
    wallReached = true
    return goRight(arr) || goLeft(arr) || goForward(arr)
  }

  return goLeft(arr) || goForward(arr) || goRight(arr)
}

const update = ({ player, players }) => {
  players.forEach(setFilled)
  players
    .filter(isOtherPlayer)
    .flatMap(({ coords }) => coords)
    .filter(isInBounds)
    .filter(isFree)
    .forEach(setUnsafe)

  const coordsInBound = player.coords
    .filter(isInBounds)
    .filter(isNotBackward)

  const available = pickForwardOrRightOrLeft(coordsInBound.filter(isFree))
  const lastResort = pickForwardOrRightOrLeft(coordsInBound.filter(isUnsafe))
  return available || lastResort
}
