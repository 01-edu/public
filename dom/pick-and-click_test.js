export const tests = []

const between = (expected, min, max) => expected >= min && expected <= max

const random = (min, max) => {
  if (!max) {
    max = min
    min = 0
  }
  min = Math.ceil(min)
  max = Math.floor(max)
  return Math.floor(Math.random() * (max - min + 1)) + min
}

const rgbToHsl = (r, g, b) => {
  r = r / 255
  g = g / 255
  b = b / 255
  const max = Math.max(r, g, b)
  const min = Math.min(r, g, b)
  let h
  let s
  const l = (max + min) / 2

  if (max === min) {
    h = s = 0 // achromatic
  } else {
    const d = max - min
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min)
    switch (max) {
      case r:
        h = (g - b) / d + (g < b ? 6 : 0)
        break
      case g:
        h = (b - r) / d + 2
        break
      case b:
        h = (r - g) / d + 4
        break
    }
    h /= 6
  }

  return [Math.floor(h * 360), Math.floor(s * 100), Math.floor(l * 100)]
}

export const setup = async ({ page }) => {
  return {
    bodyBgRgb: async () =>
      (await page.$eval('body', (body) => body.style.background))
        .replace(/[^0-9,]+/g, '')
        .split(',')
        .map((e) => Number(e)),
  }
}

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

const coords = () => [
  [random(125, 500), random(125, 400)],
  [random(125, 500), random(125, 400)],
  [random(125, 500), random(125, 400)],
]

tests.push(async ({ page, eq, bodyBgRgb }) => {
  // check that the hsl value displayed matches the current background color
  const moves = coords()
  let step = 0
  while (step < moves.length) {
    await page.mouse.move(...moves[step])
    const bgValue = rgbToHsl(...(await bodyBgRgb()))
    const hslValue = await page.$eval('.hsl', (hsl) =>
      hsl.textContent
        .replace(/[^0-9,]+/g, '')
        .split(',')
        .map((e) => Number(e)),
    )

    const matching = hslValue.map((v, i) =>
      between(v, bgValue[i] - 2, bgValue[i] + 2) ? bgValue[i] : v,
    )
    eq(matching, bgValue)
    step++
  }
})

tests.push(async ({ page, eq, bodyBgRgb }) => {
  // check that the hue value displayed matches the current background color
  const moves = coords()
  let step = 0
  while (step < moves.length) {
    await page.mouse.move(...moves[step])
    const bgValue = rgbToHsl(...(await bodyBgRgb()))
    const hueValue = await page.$eval('.hue', (hue) =>
      hue.textContent.replace(/[^0-9,]+/g, ''),
    )

    const matching = between(hueValue, bgValue[0] - 2, bgValue[0] + 2)
      ? bgValue[0]
      : Number(hueValue)
    eq(matching, bgValue[0])
    step++
  }
})

tests.push(async ({ page, eq, bodyBgRgb }) => {
  // check that the luminosity value displayed matches the current background color
  const moves = coords()
  let step = 0
  while (step < moves.length) {
    await page.mouse.move(...moves[step])
    const bgValue = rgbToHsl(...(await bodyBgRgb()))
    const luminosityValue = await page.$eval('.luminosity', (luminosity) =>
      luminosity.textContent.replace(/[^0-9,]+/g, ''),
    )

    const matching = between(luminosityValue, bgValue[2] - 2, bgValue[2] + 2)
      ? bgValue[2]
      : Number(luminosityValue)
    eq(matching, bgValue[2])
    step++
  }
})

tests.push(async ({ page, eq, bodyBgRgb }) => {
  // check that the hsl value is copied in the clipboard on click
  const moves = coords()
  let step = 0
  while (step < moves.length) {
    await page.mouse.click(...moves[step])
    const clipboard = await page.evaluate(() => navigator.clipboard.readText())
    const hslValue = await page.$eval('.hsl', (hsl) => hsl.textContent)
    eq(hslValue, clipboard)
    step++
  }
})

tests.push(async ({ page, eq, bodyBgRgb }) => {
  // check that each svg axis is following the mouse
  const moves = coords()
  let step = 0
  while (step < 1) {
    const [mouseX, mouseY] = moves[step]
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
    step++
  }
})
