/*******************************
 * functions given to the students
 ********************************/
const SIZE = 100
const MAP = new Int8Array(SIZE * SIZE)
const isFree = ({ x, y }) => MAP[y * SIZE + x] === 0
const isOccupied = ({ x, y }) => MAP[y * SIZE + x] === 1
const inBounds = (n) => n < SIZE && n >= 0
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)
const pickRandom = (arr) => arr[Math.floor(Math.random() * arr.length)]

/***********
 * My functions
 ************/
const isAlley = ({ x, y }) => !isFree({ x, y }) || !isInBounds({ x, y })

// this functions will find the best path, so the path that has more empty spaces
// so use `isFree`,
const findBestPath = (state) => {
  let arr = []
  let car = state.player.cardinal
  // if it as a block on the symmetric position it must
  // simulate the symmetric position and see witch path is the best
  if (
    (car === 3 || car === 0) &&
    !isFree({ x: state.player.x - 1, y: state.player.y - 1 }) &&
    isFree({ x: state.player.x, y: state.player.y - 1 }) &&
    isFree({ x: state.player.x - 1, y: state.player.y })
  ) {
    let xad = state.player.x - 1
    let yad = state.player.y - 1

    let choose = [
      calDistance(xad + 1, yad, 1, 0),
      calDistance(state.player.x, state.player.y - 1, car, 0),
      calDistance(state.player.x - 1, state.player.y, car, 0),
      calDistance(xad, yad + 1, 2, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1
      ? state.player.coords[0]
      : state.player.coords[3]
  }
  if (
    (car === 1 || car === 0) &&
    !isFree({ x: state.player.x + 1, y: state.player.y - 1 }) &&
    isFree({ x: state.player.x, y: state.player.y - 1 }) &&
    isFree({ x: state.player.x + 1, y: state.player.y })
  ) {
    let xad = state.player.x + 1
    let yad = state.player.y - 1
    // choose will save the biggest path to be chosen
    // [ down, line1, line0, left ]
    let choose = [
      calDistance(xad, yad + 1, 2, 0),
      calDistance(state.player.x + 1, state.player.y, car, 0),
      calDistance(state.player.x, state.player.y - 1, car, 0),
      calDistance(xad - 1, yad, 3, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1
      ? state.player.coords[1]
      : state.player.coords[0]
  }
  if (
    (car === 2 || car === 1) &&
    !isFree({ x: state.player.x + 1, y: state.player.y + 1 }) &&
    isFree({ x: state.player.x, y: state.player.y + 1 }) &&
    isFree({ x: state.player.x + 1, y: state.player.y })
  ) {
    let xad = state.player.x + 1
    let yad = state.player.y + 1
    // choose will save the biggest path to be chosen
    // [ left, line2, line1, up ]
    let choose = [
      calDistance(xad - 1, yad, 3, 0),
      calDistance(state.player.x, state.player.y + 1, car, 0),
      calDistance(state.player.x + 1, state.player.y, car, 0),
      calDistance(xad, yad - 1, 0, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1
      ? state.player.coords[2]
      : state.player.coords[1]
  }
  if (
    (car === 2 || car === 3) &&
    !isFree({ x: state.player.x - 1, y: state.player.y + 1 }) &&
    isFree({ x: state.player.x, y: state.player.y + 1 }) &&
    isFree({ x: state.player.x - 1, y: state.player.y })
  ) {
    let xad = state.player.x - 1
    let yad = state.player.y + 1
    // choose will save the biggest path to be chosen
    // [ right, line2, line3, up ]
    let choose = [
      calDistance(xad + 1, yad, 1, 0),
      calDistance(state.player.x - 1, state.player.y, car, 0),
      calDistance(state.player.x, state.player.y + 1, car, 0),
      calDistance(xad, yad - 1, 0, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1
      ? state.player.coords[3]
      : state.player.coords[2]
  }

  for ({ x, y, cardinal } of state.player.coords) {
    // if everything is ok it must continue with the best path
    arr.push(calDistance(x, y, cardinal, 0))
  }
  return state.player.coords[arr.indexOf(Math.max(...arr))]
}

// recursion
const calDistance = (x, y, car, count) => {
  if (car <= 0) {
    if (
      isFree({ x, y }) &&
      isAlley({ x: x + 1, y }) &&
      isAlley({ x, y: y - 1 }) &&
      isAlley({ x: x - 1, y })
    )
      return -1
    return !isFree({ x, y }) || !inBounds(y - 1)
      ? count
      : calDistance(x, y - 1, car, count + 1)
  }
  if (car === 1) {
    if (
      isFree({ x, y }) &&
      isAlley({ x, y: y + 1 }) &&
      isAlley({ x, y: y - 1 }) &&
      isAlley({ x: x + 1, y })
    )
      return -1
    return !isFree({ x, y }) || !inBounds(x + 1)
      ? count
      : calDistance(x + 1, y, car, count + 1)
  }
  if (car === 2) {
    if (
      isFree({ x, y }) &&
      isAlley({ x: x - 1, y }) &&
      isAlley({ x, y: y + 1 }) &&
      isAlley({ x: x + 1, y })
    )
      return -1
    return !isFree({ x, y }) || !inBounds(y + 1)
      ? count
      : calDistance(x, y + 1, car, count + 1)
  }
  if (car === 3) {
    if (
      isFree({ x, y }) &&
      isAlley({ x, y: y - 1 }) &&
      isAlley({ x: x - 1, y }) &&
      isAlley({ x, y: y + 1 })
    )
      return -1
    return !isFree({ x, y }) || !inBounds(x - 1)
      ? count
      : calDistance(x - 1, y, car, count + 1)
  }
}

const addToMap = ({ x, y }) => MAP[y * SIZE + x] = 1
const update = (state) => {
  state.players.forEach(addToMap)
  findBestPath(state)
}
