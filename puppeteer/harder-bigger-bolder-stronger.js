const body = document.querySelector('body')

const shapes = [...Array(100).keys()]

const random = (min, max) => {
  min = Math.ceil(min)
  max = Math.floor(max)
  return Math.floor(Math.random() * (max - min + 1)) + min
}

const alphabet = 'abcdefghijklmnopqrstuvwxyz'

export const generateLetters = () => {
  shapes.forEach((c) => {
    const shape = document.createElement('div')
    const third = shapes.length / 3
    const firstThird = c < third
    const secondThird = c > third && c < third * 2

    shape.textContent = alphabet[random(0, alphabet.length - 1)]
    shape.style.fontSize = `${c + 10 * 2}px`
    shape.style.fontWeight = (firstThird && 300) || (secondThird && 400) || 600
    body.append(shape)
  })
}
