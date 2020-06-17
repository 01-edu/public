const body = document.querySelector('body')

const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

export const generateLetters = () => {
  body.append(...[...Array(120).keys()].map((c) => {
    const shape = document.createElement('div')

    shape.textContent = alphabet[Math.floor(Math.random() * alphabet.length)]
    shape.style.fontSize = `${c + 11}px`
    shape.style.fontWeight = [300,400,600][Math.floor(c / 40)]
    return shape
  }))
}
