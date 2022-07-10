import { places } from './subjects/where-do-we-go/where-do-we-go.data.js'

export const tests = []

const random = (min, max) => {
  if (!max) {
    max = min
    min = 0
  }
  min = Math.ceil(min)
  max = Math.floor(max)
  return Math.floor(Math.random() * (max - min + 1)) + min
}

const getDegree = coordinates => {
  const north = coordinates.includes('N')
  const degree = coordinates.split("'")[0].replace('°', '.')
  return north ? degree : -degree
}

export const setup = async ({ page }) => {
  return {
    getDirection: async () =>
      await page.$eval('.direction', direction => direction.textContent),
  }
}

const sortedPlaces = places.sort(
  (a, b) => getDegree(b.coordinates) - getDegree(a.coordinates),
)

const dataNames = sortedPlaces.map(({ name }) =>
  name
    .split(',')[0]
    .toLowerCase()
    .split(' ')
    .join('-'),
)

tests.push(async ({ page, eq }) => {
  const { width, height } = await page.evaluate(() => ({
    width: window.innerWidth,
    height: window.innerHeight,
  }))

  const sections = await page.$$eval('section', sections =>
    sections.map(section => {
      return [
        section.getBoundingClientRect().width,
        section.getBoundingClientRect().height,
      ]
    }),
  )

  console.log(`Must contain ${places.length} places`)
  // check that the correct amount of sections has been generated
  eq(sections.length, places.length)
  // check that all the sections are fullscreen-size
  eq([...new Set(...sections)], [width, height])
})

tests.push(async ({ page, eq }) => {
  // check that the sections have been generated with the correponding images as background,
  // and sorted in the right order (from the Northest to the Southest)
  const imageNames = await page.$$eval('section', sections =>
    sections.map(section => {
      const test = section.style.background.split('.jpg')[0].split('/')
      return test[test.length - 1]
    }),
  )

  console.log(`Must be sorted from North to South`)
  console.log(`Must have the right images in background`)
  eq(imageNames, dataNames)
})

tests.push(async ({ page, eq }) => {
  // check that the location indicator is updating according to the image displayed
  let step = 1
  while (step < 6) {
    await page.evaluate(() => {
      window.scrollBy(0, window.innerHeight + 200)
    })

    await page.waitForTimeout(150)

    const location = await page.$eval('.location', location => [
      ...location.textContent.split('\n'),
      location.style.color,
    ])

    const currentLocationIndex = await page.evaluate(() =>
      Math.round(window.scrollY / window.innerHeight),
    )
    const currentLocation = sortedPlaces[currentLocationIndex]
    const { name, coordinates, color } = currentLocation
    const expectedLocation = [name, coordinates, color]

    // check that the location indicator and the displayed location contents are matching
    console.log(`Scroll #${step}: displaying ${location[0]}`)
    eq(location, expectedLocation)
    step++
  }
})

tests.push(async ({ page, eq, getDirection }) => {
  // check that the compass is pointing 'S' when scrolling down
  await page.evaluate(() => {
    window.scrollBy(0, window.innerHeight)
  })

  await page.waitForTimeout(100)

  const direction = (await getDirection()).includes('S')
    ? 'S'
    : await getDirection()

  console.log('Scroll down: pointing', direction)
  eq(direction, 'S')
})

tests.push(async ({ page, eq, getDirection }) => {
  // check that the compass is pointing 'N' when scrolling up
  await page.evaluate(() => {
    window.scrollBy(0, -100)
  })

  await page.waitForTimeout(100)

  const direction = (await getDirection()).includes('N')
    ? 'N'
    : await getDirection()

  console.log('Scroll up: pointing', direction)
  eq(direction, 'N')
})

tests.push(async ({ page, eq }) => {
  // check that the location target attribute is set to '_blank' to open a new tab
  const locationTarget = await page.$eval('.location', ({ target }) => target)
  console.log(
    `Location <a> tag target attribute ${
      locationTarget === '_blank' ? '' : 'not '
    }set to open a new tab`,
  )
  eq(locationTarget, '_blank')
})

tests.push(async ({ page, eq }) => {
  // check that the location href is valid
  const location = await page.$eval('.location', ({ href, textContent }) => ({
    href,
    textContent,
  }))
  const isValidUrl = location.href.includes('google.com/maps')
  const coords = location.textContent.split('\n')[1]
  const convertedUrl = location.href
    .split('%C2%B0')
    .join('°')
    .split('%22')
    .join('"')
    .split('%20')
    .join(' ')
  const isValidCoordinates = convertedUrl.includes(coords)

  console.log('URL', location.href, isValidUrl ? 'valid' : 'invalid')
  eq(isValidUrl, true)
  console.log(
    `Matching coordinates ${coords} ${
      isValidCoordinates ? '' : 'not '
    }found in URL`,
  )
  eq(isValidCoordinates, true)
})
