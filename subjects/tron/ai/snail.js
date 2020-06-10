const SIZE = 100
const FREE = 0
const UNSAFE = -1
const FILLED = 1
const MAP = new Int8Array(SIZE * SIZE)
const isFree = ({ x, y }) => MAP[y * SIZE + x] < FILLED
const isUnsafe = ({ x, y }) => MAP[y * SIZE + x] === UNSAFE
const setUnsafe = ({ x, y }) => MAP[y * SIZE + x] = UNSAFE
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
const isOtherPlayer = player => !player.isOwnPlayer

// `snailIt` goes with the form of a snail
const snailIt = arr => goRight(arr) || goForward(arr) || goLeft(arr)

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

  const available = snailIt(coordsInBound.filter(isFree))
  const lastResort = snailIt(coordsInBound.filter(isUnsafe))
  return available || lastResort
}
