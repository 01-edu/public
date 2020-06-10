const SIZE = 100
const h = SIZE / 2
const m = h * 0.8
const max = m => n => n > m ? max1(n - m) : n
const max1 = max(1)
const max2PI = max(Math.PI * 2)
const toInt = (r, g, b) => (r << 16) | (g << 8) | b
const toRange = n => Math.round(n * 0xFF)
const hue2rgb = (p, q, t) => {
  if (t < 0) t += 1
  if (t > 1) t -= 1
  if (t < 1/6) return p + (q - p) * 6 * t
  if (t < 1/2) return q
  if (t < 2/3) return p + (q - p) * (2/3 - t) * 6
  return p
}

const hslToRgb = (h, s, l) => {
  if (!s) return toInt(toRange(l), toRange(l), toRange(l))

  const q = l < 0.5 ? l * (1 + s) : l + s - l * s
  const p = 2 * l - q
  const r = hue2rgb(p, q, h + 1/3)
  const g = hue2rgb(p, q, h)
  const b = hue2rgb(p, q, h - 1/3)

  return toInt(toRange(r), toRange(g), toRange(b))
}

export const init = ({ players, seed }) => {
  const rand = () => {
    let t = seed += 0x6D2B79F5
    t = Math.imul(t ^ t >>> 15, t | 1)
    t ^= t + Math.imul(t ^ t >>> 7, t | 61)
    return ((t ^ t >>> 14) >>> 0) / 4294967296
  }

  const angle = (Math.PI * 2) / players.length
  const rate = (SIZE / players.length / SIZE)
  const shift = angle * rand()

  // shuffle using seeded random
  players.sort((a, b) => a.name - b.name)
  let i = players.length, j, tmp
  while (--i > 0) {
    j = Math.floor(rand() * (i + 1))
    tmp = players[j]
    players[j] = players[i]
    players[i] = tmp
  }

  return players.map((name, i) => {
    const jsonName = `"name":${JSON.stringify(name)}`
    const hue = max1(i * rate + 0.25)
    const p = {
      hue,
      name,
      x: Math.round(max2PI(Math.cos(angle * i + shift)) * m + h),
      y: Math.round(max2PI(Math.sin(angle * i + shift)) * m + h),
      cardinal: 0,
      direction: 0,
      color: hslToRgb(hue, 1, 0.5),
      toString: () => `{${jsonName},"dead":${!!p.dead},"cardinal":${p.cardinal},"direction":${p.direction},"color":${p.color},"x":${p.x},"y":${p.y},"coords":[{"x":${p.x},"y":${p.y - 1},"cardinal":0,"direction":${(4 - p.cardinal) % 4}},{"x":${p.x + 1},"y":${p.y},"cardinal":1,"direction":${(5 - p.cardinal) % 4}},{"x":${p.x},"y":${p.y + 1},"cardinal":2,"direction":${(6 - p.cardinal) % 4}},{"x":${p.x - 1},"y":${p.y},"cardinal":3,"direction":${(7 - p.cardinal) % 4}}]}`,
    }
    return p
  })
}

export const injectedCode = `
if (typeof update !== 'function') throw Error('Update function not defined')
addEventListener('message', self.init = initEvent => {
  const { seed, id } = JSON.parse(initEvent.data)
  const isOwnPlayer = p => p.name === id

  Math.random = () => {
    let t = seed += 0x6D2B79F5
    t = Math.imul(t ^ t >>> 15, t | 1)
    t ^= t + Math.imul(t ^ t >>> 7, t | 61)
    return ((t ^ t >>> 14) >>> 0) / 4294967296
  }

  removeEventListener('message', self.init)
  addEventListener('message', ({ data }) => {
    const players = JSON.parse(data)
    const player = players.find(isOwnPlayer)
    player.isOwnPlayer = true

    try { postMessage(JSON.stringify(update({ players, player }))) }
    catch (err) {
      console.error(err)
      throw err
    }
  })
  postMessage('loaded') // Signal that the loading is over
})
`
