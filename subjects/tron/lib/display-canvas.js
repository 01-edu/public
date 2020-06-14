const state = new Float32Array(100 * 100 * 2)
const [canvas] = document.getElementsByTagName('canvas')

const ctx = canvas.getContext('2d')
const memo = fn => (c => a => c[a] || (c[a] = fn(a)))(Object.create(null))
const toHex = memo(color => '#'+ `00000${color.toString(16)}`.slice(-6))
export const colorize = (x, y, color) => {
  ctx.fillStyle = toHex(color)
  ctx.fillRect(x * 12, y * 12, 12, 12)
}

export const move = (x, y, color, turn) => {
  const index = (x * 100 + y) * 2
  if (state[index]) return state[index]
  state[index] = turn
  colorize(x, y, color)
}

export const update = () => {}
export const reset = () => {}
