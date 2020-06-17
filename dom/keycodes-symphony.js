const body = document.querySelector('body')

export const compose = () => {
  document.addEventListener('keydown', (e) => handleKey(e))
  setTimeout(
    () => document.removeEventListener('keydown', (e) => handleKey(e)),
    500,
  )
}

const handleKey = (e) => {
  const notes = [...document.querySelectorAll('.note')]

  if (e.key === 'Backspace') {
    const last = notes[notes.length - 1]
    last && last.remove()
    return
  }

  if (e.key === 'Escape') {
    if (notes.length) {
      notes.forEach((note) => note.remove())
    }
    return
  }

  createNote(e)
}

const createNote = ({ key }) => {
  const number = key.charCodeAt(0) * 2 - 150
  const note = document.createElement('div')
  note.className = 'note'
  note.textContent = key
  note.style.background = `hsl(270, ${number}%, ${number}%)`
  body.append(note)
}
