const vertexArray = new Float32Array(100 * 100 * 12)
const colorArray = new Float32Array(100 * 100 * 6)
const state = new Float32Array(100 * 100 * 2)
const [canvas] = document.getElementsByTagName('canvas')
const gl = canvas.getContext('webgl2', { antialias: false })
const S = 0.02

const applyState = (x, y, turn) => {
  const index = x * 100 + y
  const color = state[index * 2 + 0] > turn ? 0 : state[index * 2 + 1]
  colorArray[index * 6 + 0] = color
  colorArray[index * 6 + 1] = color
  colorArray[index * 6 + 2] = color
  colorArray[index * 6 + 3] = color
  colorArray[index * 6 + 4] = color
  colorArray[index * 6 + 5] = color
}

export const colorize = (x, y, color) => state[(x * 100 + y) * 2 + 1] = color
export const move = (x, y, color, turn) => {
  const index = (x * 100 + y) * 2
  if (state[index]) return state[index]
  state[index] = turn
  state[index + 1] = color
}

const loop = fn => {
  let x = -1, y = -1
  while (++x < 100) {
    y = -1
    while (++y < 100) fn(x, y)
  }
}

const compileShader = (type, script) => {
  const shader = gl.createShader(type)
  gl.shaderSource(shader, script.trim())
  gl.compileShader(shader)
  if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
    throw gl.getShaderInfoLog(shader)
  }
  return shader
}

// program
const program = gl.createProgram()
gl.attachShader(program, compileShader(gl.VERTEX_SHADER, `
#version 300 es

in vec2 a_position;
in float a_color;
out float v_color;

void main() {
  gl_Position = vec4(a_position * vec2(1, -1), 0, 1);
  v_color = a_color;
}`))

gl.attachShader(program, compileShader(gl.FRAGMENT_SHADER, `
#version 300 es

precision mediump float;
in float v_color;
out vec4 outColor;

vec4 unpackColor(float f) {
  vec4 color;
  color.r = floor(f / 65536.0);
  color.g = floor((f - color.r * 65536.0) / 256.0);
  color.b = floor(f - color.r * 65536.0 - color.g * 256.0);
  color.a = 256.0;
  return color / 256.0;
}

void main() {
  outColor = unpackColor(v_color);
}`))

gl.linkProgram(program)
if (!gl.getProgramParameter(program, gl.LINK_STATUS)) {
  throw gl.getProgramInfoLog(program)
}

gl.useProgram(program)

// initialize state
loop((x, y) => {
  const x1 = ((x + 1) - 50) / 50 - S
  const y1 = ((y + 1) - 50) / 50 - S
  const x2 = x1 + S
  const y2 = y1 + S
  const index = (x * 100 + y) * 12
  vertexArray[index + 0x0] = x1
  vertexArray[index + 0x1] = y1
  vertexArray[index + 0x2] = x2
  vertexArray[index + 0x3] = y1
  vertexArray[index + 0x4] = x1
  vertexArray[index + 0x5] = y2
  vertexArray[index + 0x6] = x1
  vertexArray[index + 0x7] = y2
  vertexArray[index + 0x8] = x2
  vertexArray[index + 0x9] = y1
  vertexArray[index + 0xa] = x2
  vertexArray[index + 0xb] = y2
})

const vertexBuffer = gl.createBuffer()
const a_position = gl.getAttribLocation(program, 'a_position')
gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
gl.vertexAttribPointer(a_position, 2, gl.FLOAT, false, 0, 0)
gl.enableVertexAttribArray(a_position)
gl.bufferData(gl.ARRAY_BUFFER, vertexArray, gl.STATIC_DRAW)
gl.drawArrays(gl.TRIANGLES, 0, 60000)

// color buffer
const colorBuffer = gl.createBuffer()
const a_color = gl.getAttribLocation(program, 'a_color')
gl.bindBuffer(gl.ARRAY_BUFFER, colorBuffer)
gl.vertexAttribPointer(a_color, 1, gl.FLOAT, false, 0, 0)
gl.enableVertexAttribArray(a_color)
gl.bindBuffer(gl.ARRAY_BUFFER, colorBuffer)

export const update = (turn) => {
  loop((x, y) => applyState(x, y, turn))
  gl.bufferData(gl.ARRAY_BUFFER, colorArray, gl.STATIC_DRAW)
  gl.drawArrays(gl.TRIANGLES, 0, 60000)
}

export const reset = () => {
  state.fill(0)
  update(0)
}
