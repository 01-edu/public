const body = document.querySelector('body')

const box = document.createElement('div')
box.className = 'box'
body.append(box)

const { top, bottom, left, right } = box.getBoundingClientRect()

const diameter = 50
const radius = diameter / 2

const insideX = (clientX) => clientX > left + radius && clientX < right - radius
const insideY = (clientY) => clientY > top + radius && clientY < bottom - radius

let isInside = false

export const createCircle = () => {
  document.addEventListener('click', (e) => create(e))
  setTimeout(() => document.removeEventListener('click', create), 500)
}

const create = ({ clientX, clientY }) => {
  const circle = document.createElement('div')
  circle.className = 'circle'
  body.append(circle)
  circle.style.top = `${clientY - radius}px`
  circle.style.left = `${clientX - radius}px`
  circle.style.background = 'white'
  const hasEntered = insideX(clientX) && insideY(clientY)
  if (hasEntered) {
    circle.style.background = 'var(--purple)'
  }
  isInside = false
}

export const moveCircle = () => {
  document.addEventListener('mousemove', (e) => move(e))
  setTimeout(() => document.removeEventListener('mousemove', move), 500)
}

const move = (e) => {
  const circles = [...document.getElementsByClassName('circle')]
  const circle = circles[circles.length - 1]
  if (!circle) return
  position(e, circle)
}

const position = ({ clientX, clientY }, circle) => {
  const hasEntered = insideX(clientX) && insideY(clientY)

  if (hasEntered) {
    isInside = true
    circle.style.background = 'var(--purple)'
  }

  if (isInside) {
    if (insideY(clientY)) {
      circle.style.top = `${clientY - radius}px`
    }
    if (insideX(clientX)) {
      circle.style.left = `${clientX - radius}px`
    }
  } else {
    circle.style.top = `${clientY - radius}px`
    circle.style.left = `${clientX - radius}px`
  }
}
