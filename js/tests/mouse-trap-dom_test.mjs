export const tests = []

export const setup = async ({ page }) => ({
  getCirclesPos: () =>
    page.$$eval('.circle', nodes => {
      const circleRadius = 25
      const formatPos = pos => Number(pos.replace('px', '')) + circleRadius
      return nodes.map(node => [
        formatPos(node.style.left),
        formatPos(node.style.top),
      ])
    }),
})

tests.push(async ({ page, eq, ctx }) => {
  // check that a circle is created on click at the mouse position
  const { width, height } = await page.evaluate(() => ({
    width: document.documentElement.clientWidth,
    height: document.documentElement.clientHeight,
  }))

  const clicks = [...Array(10).keys()].map(e => [random(width), random(height)])

  for (const [i, click] of clicks.entries()) {
    const [posX, posY] = click
    await page.mouse.click(posX, posY)
    const currentCircle = (await ctx.getCirclesPos())[i]
    eq(currentCircle, click)
  }
})

tests.push(async ({ page, eq, ctx }) => {
  // check that the last created circle moves along the mouse
  let move = 0
  while (move < 100) {
    move++
    const x = move
    const y = move * 2
    await page.mouse.move(x, y)
    const circles = await ctx.getCirclesPos()
    const currentCirclePos = circles[circles.length - 1]
    eq(currentCirclePos, [x, y])
  }
})

tests.push(async ({ page, eq, ctx }) => {
  // check that a circle is trapped and purple when inside the box
  const box = await page.$eval('.box', box => ({
    top: box.getBoundingClientRect().top,
    right: box.getBoundingClientRect().right,
    left: box.getBoundingClientRect().left,
    bottom: box.getBoundingClientRect().bottom,
  }))

  await page.mouse.click(200, 200)

  let move = 200
  let hasEntered = false

  while (move < 500) {
    const x = move + 50
    const y = move
    await page.mouse.move(x, y)

    const circles = await ctx.getCirclesPos()
    const currentCircle = circles[circles.length - 1]

    const circleRadius = 25

    const bg = await page.$$eval(
      '.circle',
      nodes => nodes[nodes.length - 1].style.background,
    )

    const insideX = x > box.left + circleRadius && x < box.right - circleRadius
    const insideY = y > box.top + circleRadius && y < box.bottom - circleRadius
    const isInside = insideX && insideY

    // check that the background is set to the right color
    if (isInside) {
      hasEntered = true
      eq(bg, 'var(--purple)')
    } else {
      eq(bg, hasEntered ? 'var(--purple)' : 'white')
    }

    // check that the mouse is trapped inside the box
    if (hasEntered) {
      if (insideY) {
        eq(currentCircle[1], y)
      } else {
        const maxY =
          currentCircle[1] === box.top + circleRadius + 1 ||
          currentCircle[1] === box.top + circleRadius ||
          currentCircle[1] === box.bottom - circleRadius ||
          currentCircle[1] === box.bottom - circleRadius - 1
        eq(maxY, true)
      }
      if (insideX) {
        eq(currentCircle[0], x)
      } else {
        const maxX =
          currentCircle[0] === box.left + circleRadius ||
          currentCircle[0] === box.left + circleRadius + 1 ||
          currentCircle[0] === box.right - circleRadius ||
          currentCircle[0] === box.right - circleRadius - 1
        eq(maxX, true)
      }
    }
    move++
  }
})

const random = (min, max) => {
  if (!max) {
    max = min
    min = 0
  }
  min = Math.ceil(min)
  max = Math.floor(max)
  return Math.floor(Math.random() * (max - min + 1)) + min
}
