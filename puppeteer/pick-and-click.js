const body = document.querySelector('body')

const count = document.createElement('div')
count.className = 'count'
count.textContent = 'hsl(0, 50%, 0%)'

const hueText = document.createElement('div')
hueText.className = 'text hue'
hueText.textContent = 'hue'

const luminosityText = document.createElement('div')
luminosityText.className = 'text luminosity'
luminosityText.textContent = 'luminosity'

const origin = document.createElement('div')
origin.className = 'text origin'

const picked = document.createElement('div')
picked.className = 'text picked'
picked.textContent = 'Color successfully picked!'

const svg = document.createElementNS('http://www.w3.org/2000/svg', 'svg')
svg.setAttributeNS(
  'http://www.w3.org/2000/xmlns/',
  'xmlns:xlink',
  'http://www.w3.org/1999/xlink',
)
svg.setAttribute('width', window.innerWidth)
svg.setAttribute('height', window.innerHeight)
svg.setAttribute('viewBox', `0 0 ${window.innerWidth} ${window.innerHeight}`)

const axisX = document.createElementNS('http://www.w3.org/2000/svg', 'line')
axisX.setAttribute('y1', window.innerHeight)
axisX.setAttribute('y2', 0)
svg.append(axisX)

const axisY = document.createElementNS('http://www.w3.org/2000/svg', 'line')
axisY.setAttribute('x1', window.innerWidth)
axisY.setAttribute('x2', 0)
svg.append(axisY)

body.append(count, hueText, luminosityText, origin, picked, svg)

export const pick = () => {
  document.addEventListener('mousemove', (e) => set(e))
  setTimeout(
    () => document.removeEventListener('mousemove', (e) => set(e)),
    500,
  )

  body.addEventListener('click', click)
  setTimeout(() => document.removeEventListener('click', click), 500)

  body.addEventListener('copy', copy)
  setTimeout(() => document.removeEventListener('copy', copy), 500)
}

const click = (e) => {
  document.execCommand('copy')
  const wave = document.createElement('div')
  wave.className = 'wave'
  wave.style.top = `${e.clientY - 10}px`
  wave.style.left = `${e.clientX - 10}px`
  body.append(wave)
  setTimeout(() => wave.remove(), 150)
}

const copy = (event) => {
  event.preventDefault()
  if (event.clipboardData) {
    event.clipboardData.setData('text/plain', count.textContent)
    picked.classList.add('fade-in')
    setTimeout(() => picked.classList.remove('fade-in'), 1000)
  }
}

const calc = (number, max) =>
  Math.round(Math.min(max, Math.max(0, max * number)))

const set = ({ clientX, clientY }) => {
  const { innerWidth, innerHeight } = window
  const padding = 100
  const mouseX = clientX - padding
  const mouseY = clientY - padding

  const hue = calc(mouseX / (innerWidth - padding * 2), 360)
  const luminosity = calc(mouseY / (innerHeight - padding * 2), 100)
  const color = `hsl(${hue}, 50%, ${luminosity}%)`

  axisX.setAttribute('x1', clientX)
  axisX.setAttribute('x2', clientX)
  axisY.setAttribute('y1', clientY)
  axisY.setAttribute('y2', clientY)

  axisX.setAttribute('stroke', color)
  axisY.setAttribute('stroke', color)

  body.style.color = color
  body.style.background = color
  origin.style.background = color
  count.textContent = color

  hueText.textContent = `hue\n${hue}`
  luminosityText.textContent = `${luminosity}\nluminosity`
}
