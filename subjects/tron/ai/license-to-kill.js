const SIZE = 100
const MAP = new Int8Array(SIZE * SIZE) // State of the Map
const isFree = ({ x, y }) => MAP[y * SIZE + x] === 0 // 0 = block free
const isOccupied = ({ x, y }) => MAP[y * SIZE + x] === 1 // 1 = block occupied
const inBounds = n => n < SIZE && n >= 0
const isInBounds = ({ x, y }) => inBounds(x) && inBounds(y)
const pickRandom = arr => arr[Math.floor(Math.random() * arr.length)]

// ok check if we can move on this block
const ok = (x = -1, y = -1) => {
  const coords = typeof x !== 'number' ? x : { x, y }
  return isFree(coords) && isInBounds(coords)
}

const isAlley = (card, x, y) => {
  switch (card) {
    case 0:
      if (ok(x, y - 1) && !ok(x + 1, y - 1) && !ok(x - 1, y - 1)) {
        while (ok(x, y - 1) && !hasLateralWalls(0, x, y)) {
          y--
          if (hasLateralWalls(0, x, y)) return true
        }
      }
      return false
    case 1:
      if (ok(x + 1, y) && !ok(x + 1, y + 1) && !ok(x + 1, y - 1)) {
        while (ok(x + 1, y) && !hasLateralWalls(1, x, y)) {
          x++
          if (hasLateralWalls(1, x, y)) return true
        }
      }
      return false
    case 2:
      if (ok(x, y + 1) && !ok(x + 1, y + 1) && !ok(x - 1, y + 1)) {
        while (ok(x, y + 1) && !hasLateralWalls(2, x, y)) {
          y++
          if (hasLateralWalls(2, x, y)) return true
        }
      }
      return false
    case 3:
      if (ok(x - 1, y) && !ok(x - 1, y + 1) && !ok(x - 1, y - 1)) {
        while (ok(x - 1, y) && !hasLateralWalls(3, x, y)) {
          x--
          if (hasLateralWalls(3, x, y)) return true
        }
      }
      return false
  }
}

const hasLateralWalls = (card, x, y) => {
  switch (card) {
    case 0: return !(ok(x + 1, y) || ok(x - 1, y) || ok(x, y - 1))
    case 1: return !(ok(x, y - 1) || ok(x, y + 1) || ok(x + 1, y))
    case 2: return !(ok(x + 1, y) || ok(x - 1, y) || ok(x, y + 1))
    case 3: return !(ok(x, y - 1) || ok(x, y + 1) || ok(x - 1, y))
  }
}

const goDirection = (state, card) =>
  ok(state.coords[card]) &&
  !isAlley(card, state.x, state.y) &&
  state.coords[card]

const findEnemy = state =>
  state.players.filter(p => state.player.name !== p.name)[0]

const seekEnemy = state => {
  if (state.players.length === 1) return
  const enemy = findEnemy(state)
  const xPla = state.player.x
  const yPla = state.player.y
  const xOpo = enemy.x
  const yOpo = enemy.y
  const xDif = xPla - xOpo
  const yDif = yPla - yOpo

  return (
    (Math.abs(xDif) > Math.abs(yDif) &&
      goDirection(state.player, xPla < xOpo ? 1 : 3)) ||
    goDirection(yPla < yOpo ? 2 : 0)
  )
}

const walk = state =>
  seekEnemy(state) ||
  goDirection(state.player, 0) ||
  goDirection(state.player, 1) ||
  goDirection(state.player, 2) ||
  goDirection(state.player, 3)

const addToMap = ({ x, y }) => (MAP[y * SIZE + x] = 1)
const update = state => {
  state.players.forEach(addToMap)
  return walk(state)
}
