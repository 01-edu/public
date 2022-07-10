export const tests = []

const between = (expected, min, max) => expected >= min && expected <= max

export const setup = async ({ page, rgbToHsl }) => ({
  bodyBgRgb: async () =>
    rgbToHsl(await page.$eval('body', (body) => body.style.background)),
})

tests.push(async ({ page, eq }) => {
  // check that the background color is changing on mouse move
  // by simulating 20 moves, so there must be 20 different background colors
  let move = 50
  let bgs = []
  while (move < 250) {
    move += 10
    const x = move
    const y = move * 2
    await page.mouse.move(x, y)
    const bodyBg = await page.$eval('body', (body) => body.style.background)
    const validColor = bodyBg.includes('rgb')
    if (!validColor) continue
    bgs.push(bodyBg)
  }
  const differentBgs = [...new Set(bgs)].length
  eq(differentBgs, 20)
})

const near = (a, b) => a < b + 2.5 && a > b - 2.5
const numVal = (n) => n.textContent.replace(/[^0-9,]/g, '')
const generateCoords = (random) => [
  [random(125, 500), random(125, 400)],
  [random(125, 500), random(125, 400)],
  [random(125, 500), random(125, 400)],
]

tests.push(async ({ page, eq, bodyBgRgb, random }) => {
  // check that the hsl value displayed matches the current background color
  for (const move of generateCoords(random)) {
    await page.mouse.move(...move)
    const a = await bodyBgRgb()
    const b = (await page.$eval('.hsl', numVal)).split(',')
    if (a.every((v, i) => near(v, Number(b[i])))) continue
    throw Error(`hsl(${a.map(Math.round)}) to far from hsl(${b})`)
  }
})

tests.push(async ({ page, eq, bodyBgRgb, random }) => {
  // check that the hue value displayed matches the current background color
  for (const move of generateCoords(random)) {
    await page.mouse.move(...move)
    const [h] = await bodyBgRgb()
    const hue = await page.$eval('.hue', numVal)

    if (!near(h, Number(hue))) {
      console.log({ h, hue, c: near(h, Number(hue)) })
      throw Error(`hue ${Math.round(h)} to far from ${hue}`)
    }
  }
})

tests.push(async ({ page, eq, bodyBgRgb, random }) => {
  // check that the luminosity value displayed matches the current background color
  for (const move of generateCoords(random)) {
    await page.mouse.move(...move)
    const [, , l] = await bodyBgRgb()
    const lum = await page.$eval('.luminosity', numVal)

    if (!near(l, Number(lum))) {
      throw Error(`luminosity ${Math.round(l)} to far from ${lum}`)
    }
  }
})

tests.push(async ({ page, eq, bodyBgRgb, random }) => {
  // check that the hsl value is copied in the clipboard on click
  // Override readText if writeText is used due to a puppeteer bug
  await page.evaluate(() => {
    window.navigator.clipboard.writeText = async (text) => {
      window.navigator.clipboard.readText = async () => text
    }
  })
  for (const move of generateCoords(random)) {
    await page.mouse.click(...move)
    const clipboard = await page.evaluate(() =>
      window.navigator.clipboard.readText()
    )
    const hslValue = await page.$eval('.hsl', (hsl) => hsl.textContent)
    eq(hslValue, clipboard)
  }
})

tests.push(async ({ page, eq, bodyBgRgb, random }) => {
  // check that each svg axis is following the mouse
  const [[mouseX, mouseY]] = generateCoords(random)
  await page.mouse.move(mouseX, mouseY)
  const axisX = await page.$eval('#axisX', (line) => [
    line.getAttribute('x1'),
    line.getAttribute('x2'),
  ])
  const axisY = await page.$eval('#axisY', (line) => [
    line.getAttribute('y1'),
    line.getAttribute('y2'),
  ])

  const checkAxisCoords = (coords) => Number([...new Set(coords)].join())

  eq(checkAxisCoords(axisX), mouseX)
  eq(checkAxisCoords(axisY), mouseY)
})
