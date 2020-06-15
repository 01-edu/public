import { gossips as archived } from './data.js'

const body = document.querySelector('body')

const ranges = document.createElement('div')
ranges.className = 'ranges'
body.append(ranges)

const inputs = [
  { props: ['width'], min: 200, max: 800, value: 250 },
  { props: ['fontSize', 'lineHeight'], min: 20, max: 40, value: 25 },
  { props: ['background'], min: 20, max: 75, value: 60 },
]

let gossips = new Proxy(archived, {
  set: (target, prop, value) => {
    target[prop] = value
    createGossip(value, true)
    return true
  },
})

export const grid = () => {
  inputs.forEach((input) => createInput(input))
  createAddGossip()
  gossips.forEach((g) => createGossip(g))
}

const createGossip = (g, isNew = false) => {
  const gossip = document.createElement('div')
  const addGossip = document.getElementById('add-gossip')
  const { fontSize, lineHeight, width, background } = addGossip.style
  gossip.className = 'gossip'
  gossip.textContent = g
  if (isNew) {
    gossip.style.fontSize = fontSize
    gossip.style.lineHeight = lineHeight
    gossip.style.width = width
    gossip.style.background = background
    gossip.classList.add('fade-in')
    body.insertBefore(gossip, addGossip.nextElementSibling)
  } else {
    body.append(gossip)
  }
}

const createAddGossip = () => {
  const addGossip = document.createElement('div')
  addGossip.className = 'gossip'
  addGossip.id = 'add-gossip'

  const newInput = document.createElement('textarea')
  newInput.autofocus = true
  newInput.placeholder = 'Got a gossip to share ?'
  newInput.addEventListener('keyup', (e) => addNewGossip(newInput, e))

  const button = document.createElement('div')
  button.className = 'button'
  button.textContent = 'Share gossip!'
  button.addEventListener('click', (e) => addNewGossip(newInput))

  addGossip.append(newInput, button)
  body.append(addGossip)
}

const addNewGossip = (input, event) => {
  const noValue = !input.value
  const notEnterKey = event && event.keyCode !== 13
  if (notEnterKey || noValue) {
    input.focus()
    return
  }
  gossips[gossips.length] = input.value
  input.value = ''
  input.focus()
}

const createInput = ({ props, min, max, value }) => {
  const range = document.createElement('div')
  range.className = 'range'

  const input = document.createElement('input')
  input.type = 'range'
  input.min = min
  input.max = max
  input.value = value
  input.addEventListener('input', (e) => customize(e, ...props))

  const propLabel = document.createElement('label')
  propLabel.textContent = props[0]

  const valueLabel = document.createElement('span')
  valueLabel.textContent = value

  range.append(propLabel, input, valueLabel)
  ranges.append(range)
}

const customize = ({ target }, ...props) => {
  const gossips = [...document.querySelectorAll('.gossip')]
  gossips.forEach((gossip) => {
    props.forEach((prop) => {
      const updatedValue =
        (prop === 'lineHeight' && `${Number(target.value) * 1.5}px`) ||
        (prop === 'background' && `hsl(280, 50%, ${target.value}%)`) ||
        `${target.value}px`
      gossip.style[prop] = updatedValue
    })
  })
  const valueLabel = target.nextElementSibling
  valueLabel.textContent = target.value
}
