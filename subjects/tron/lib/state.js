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

export const init = ({ ais, seed }) => {
  let w = (123456789 + seed) & 0xffffffff
  let z = (987654321 - seed) & 0xffffffff

  const rand = () => {
    z = (36969 * (z & 65535) + (z >>> 16)) & 0xffffffff
    w = (18000 * (w & 65535) + (w >>> 16)) & 0xffffffff
    return (((z << 16) + (w & 65535)) >>> 0) / 4294967296
  }

  const angle = (Math.PI * 2) / ais.length
  const rate = (SIZE / ais.length / SIZE)
  const shift = angle * rand()

  // shuffle using seeded random
  ais.sort((a, b) => a.name - b.name)
  let i = ais.length, j, tmp
  while (--i > 0) {
    j = Math.floor(rand() * (i + 1))
    tmp = ais[j]
    ais[j] = ais[i]
    ais[i] = tmp
  }

  return ais.map((name, i) => {
    const jsonName = `"name":${JSON.stringify(name)}`
    const hue = max1(i * rate + 0.25)
    const ai = {
      hue,
      name,
      x: Math.round(max2PI(Math.cos(angle * i + shift)) * m + h),
      y: Math.round(max2PI(Math.sin(angle * i + shift)) * m + h),
      color: hslToRgb(hue, 1, 0.5),
      toString: () => `{${jsonName},"dead":${!!ai.dead},"color":${ai.color},"x":${ai.x},"y":${ai.y}}`,
    }
    return ai
  })
}

export const injectedCode = `
if (typeof update !== 'function') throw Error('Update function not defined')
addEventListener('message', self.init = initEvent => {
  let { seed, id } = JSON.parse(initEvent.data)

  const r4 = () => Math.floor(Math.random() * 4)
  let w = (123456789 + seed) & 0xffffffff
  let z = (987654321 - seed) & 0xffffffff
  Math.random = () => {
    z = (36969 * (z & 65535) + (z >>> 16)) & 0xffffffff
    w = (18000 * (w & 65535) + (w >>> 16)) & 0xffffffff
    return (((z << 16) + (w & 65535)) >>> 0) / 4294967296
  }

  const prev = {}
  let me
  removeEventListener('message', self.init)
  addEventListener('message', ({ data }) => {
    const ais = JSON.parse(data)
    me || (ais
      .sort((a, b) => a.name - b.name)
      .forEach(a => prev[a.name] = [{...a, cardinal: r4(), direction: 0 }]))

    for (const ai of ais) {
      const { x, y, name, dead } = ai
      if (dead) {
        ai.coords = []
        continue
      }

      name === id && (me = ai)
      const { cardinal, direction } = prev[name].find(c => c.x === ai.x && c.y === ai.y)

      ai.index = ai.x * 100 + ai.y
      ai.direction = direction
      ai.cardinal = cardinal
      ai.coords = prev[name] = [
        { index: x*100+(y-1), x, y: y - 1, cardinal: 0, direction: (4 - cardinal) % 4 },
        { index: (x+1)*100+y, x: x + 1, y, cardinal: 1, direction: (5 - cardinal) % 4 },
        { index: x*100+(y+1), x, y: y + 1, cardinal: 2, direction: (6 - cardinal) % 4 },
        { index: (x-1)*100+y, x: x - 1, y, cardinal: 3, direction: (7 - cardinal) % 4 },
      ]
    }
    me.me = true

    try { postMessage(JSON.stringify(update({ ais, ai: me }))) }
    catch (err) {
      console.error(id, err)
      throw err
    }
  })
  postMessage('loaded') // Signal that the loading is over
})
`
