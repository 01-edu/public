import { places } from './data.js'

const body = document.querySelector('body')

export const explore = () => {
  createSections()

  const location = document.createElement('div')
  location.className = 'location'
  setLocation(location)

  const direction = document.createElement('div')
  direction.className = 'direction'

  body.append(location, direction)

  document.addEventListener('wheel', (event) =>
    setLocation(location, direction, event),
  )
  setTimeout(() =>
    document.removeEventListener('wheel', (event) =>
      setLocation(location, direction, event),
    ),
  )
}

const createSections = () => {
  const sorted = places.sort(
    (a, b) => getDegree(b.coordinates) - getDegree(a.coordinates),
  )

  sorted.map(({ name, color }) => {
    const nameDashCase = name
      .toLowerCase()
      .split(',')[0]
      .split(' ')
      .join('-')

    const url = `https://raw.githubusercontent.com/MarieMalarme/dom-js/master/assets/images/${nameDashCase}.jpg`

    const section = document.createElement('section')
    section.style.background = `center / cover url(${url})`
    body.append(section)
  })
}

const getDegree = (coordinates) => {
  const north = coordinates.includes('N')
  const degree = coordinates.split("'")[0].replace('°', '.')
  return north ? degree : -degree
}

const setLocation = (location, direction, event) => {
  const { name, coordinates, color } = getLocation()
  location.textContent = `${name}\n${coordinates}`
  location.style.color = color
  location.onclick = () => {
    window.open(`https://www.google.com/maps/place/${coordinates}`, '_blank')
  }

  if (!event) return
  const scrollUp = event.deltaY < 0
  direction.innerHTML = `<div style="transform: rotate(${
    scrollUp ? -90 : 90
  }deg)">➢</div><div>${scrollUp ? 'N' : 'S'}</div>`
}

const getLocation = () => {
  const { innerHeight, scrollY } = window
  const index = Math.ceil((scrollY - innerHeight / 2) / innerHeight)
  return places[index]
}
