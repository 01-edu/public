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
// TODO: could make the ai harder if left an empty space so that it could get out

// this functions will find the best path, so the path that has more empty spaces
const findBestPath = ({ ai }) => {
  let arr = []
  let car = ai.cardinal
  // if it as a block on the symmetric position it must
  // simulate the symmetric position and see witch path is the best
  // from the symmetric position
  //  o→
  // ←↓oooooooooo
  //   ↓
  if (
    (car === 3 || car === 0) &&
    !isFree({ x: ai.x - 1, y: ai.y - 1 }) &&
    isFree({ x: ai.x, y: ai.y - 1 }) &&
    isFree({ x: ai.x - 1, y: ai.y })
  ) {
    let xad = ai.x - 1
    let yad = ai.y - 1
    // [right, line0, line3, down]
    let choose = [
      calDistance(xad + 1, yad, 1, 0),
      calDistance(ai.x, ai.y - 1, 0, 0),
      calDistance(ai.x - 1, ai.y, 3, 0),
      calDistance(xad, yad + 1, 2, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1 ? ai.coords[0] : ai.coords[3]
  }
  if (
    (car === 1 || car === 0) &&
    !isFree({ x: ai.x + 1, y: ai.y - 1 }) &&
    isFree({ x: ai.x, y: ai.y - 1 }) &&
    isFree({ x: ai.x + 1, y: ai.y })
  ) {
    let xad = ai.x + 1
    let yad = ai.y - 1
    // choose will save the biggest path to be chosen
    // [ down, line1, line0, left ]
    let choose = [
      calDistance(xad, yad + 1, 2, 0),
      calDistance(ai.x + 1, ai.y, 1, 0),
      calDistance(ai.x, ai.y - 1, 0, 0),
      calDistance(xad - 1, yad, 3, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1 ? ai.coords[1] : ai.coords[0]
  }
  if (
    (car === 2 || car === 1) &&
    !isFree({ x: ai.x + 1, y: ai.y + 1 }) &&
    isFree({ x: ai.x, y: ai.y + 1 }) &&
    isFree({ x: ai.x + 1, y: ai.y })
  ) {
    let xad = ai.x + 1
    let yad = ai.y + 1
    // choose will save the biggest path to be chosen
    // [ left, line2, line1, up ]
    let choose = [
      calDistance(xad - 1, yad, 3, 0),
      calDistance(ai.x, ai.y + 1, 2, 0),
      calDistance(ai.x + 1, ai.y, 1, 0),
      calDistance(xad, yad - 1, 0, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1 ? ai.coords[2] : ai.coords[1]
  }
  if (
    (car === 2 || car === 3) &&
    !isFree({ x: ai.x - 1, y: ai.y + 1 }) &&
    isFree({ x: ai.x, y: ai.y + 1 }) &&
    isFree({ x: ai.x - 1, y: ai.y })
  ) {
    let xad = ai.x - 1
    let yad = ai.y + 1
    // choose will save the biggest path to be chosen
    // [ right, line2, line3, up ]
    let choose = [
      calDistance(xad + 1, yad, 1, 0),
      calDistance(ai.x, ai.y + 1, 2, 0),
      calDistance(ai.x - 1, ai.y, 3, 0),
      calDistance(xad, yad - 1, 0, 0),
    ]
    let index = choose.indexOf(Math.max(...choose))
    return index === 0 || index === 1 ? ai.coords[2] : ai.coords[3]
  }

  for (const { x, y, cardinal } of ai.coords) {
    // if everything is ok it must continue with the best path
    arr.push(calDistance(x, y, cardinal, 0))
  }
  return ai.coords[arr.indexOf(Math.max(...arr))]
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

const addToMap = ({ x, y }) => (MAP[y * SIZE + x] = 1)
const update = (state) => {
  state.ais.forEach(addToMap)
  return findBestPath(state)
}
