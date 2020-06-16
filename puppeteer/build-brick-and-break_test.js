export const tests = []

export const setup = async ({ page }) => {
  const body = await page.$('body')
  const hammer = await page.$('#hammer')
  const dynamite = await page.$('#dynamite')

  return {
    body,
    hammer,
    dynamite,
    getDivsLength: async () =>
      await page.$$eval('div', (nodes) => nodes.length),
    getBricksIds: async () =>
      await page.$$eval('div', (nodes) =>
        nodes.filter((node) => node.id.includes('brick')).map((n) => n.id),
      ),
    getMiddleBricksIds: async () =>
      await page.$$eval('div', (nodes) =>
        nodes
          .filter(
            (node) => node.id.includes('brick') && node.dataset.foundation,
          )
          .map((n) => n.id),
      ),
    getExpectedRepairedIds: async () =>
      await page.$eval('body', (body) => {
        const getIdInt = (str) => str.replace('brick-', '')
        return body.dataset.reparations
          .split(',')
          .sort((a, b) => getIdInt(b) - getIdInt(a))
          .map((id) => {
            const isMiddleBrick = getIdInt(id) % 3 === 2
            const status = isMiddleBrick ? 'in progress' : 'repaired'
            return `${id}_${status}`
          })
      }),
    getRepairedIds: async () =>
      await page.$$eval('div', (nodes) => {
        const getIdInt = (str) => str.replace('brick-', '')
        return nodes
          .filter(
            (node) =>
              node.dataset.repaired === 'true' ||
              node.dataset.repaired === 'in progress',
          )
          .sort((a, b) => getIdInt(b.id) - getIdInt(a.id))
          .map(({ id }) => {
            const isMiddleBrick = getIdInt(id) % 3 === 2
            const status = isMiddleBrick ? 'in progress' : 'repaired'
            return `${id}_${status}`
          })
      }),
  }
}

const between = (expected, min, max) => expected >= min && expected <= max

tests.push(async ({ page, eq, getDivsLength }) => {
  // check that the brick divs are built at a regular interval of 100ms
  // the average of the divs built every 100ms must be close to 10
  let repeat = 0
  let buildSteps = []

  while (repeat < 3) {
    const divs = await getDivsLength()
    buildSteps.push(divs)
    await page.waitFor(1000)
    repeat++
  }

  const diff1 = buildSteps[1] - buildSteps[0]
  const diff2 = buildSteps[2] - buildSteps[1]
  const average = Math.round((diff1 + diff2) / 2)

  if (average < 9) {
    console.log('too slow --> average built bricks / sec:', average)
  }
  if (average > 11) {
    console.log('too fast --> average built bricks / sec:', average)
  }

  eq(between(average, 9, 11), between(10, 9, 11))
})

const allBricksIds = [...Array(54).keys()].map((i) => `brick-${i + 1}`)

tests.push(async ({ page, eq, getBricksIds }) => {
  // check that all the bricks are here and have the correct id
  await page.waitFor(3000)
  const bricksIds = await getBricksIds()
  eq(bricksIds, allBricksIds)
})

tests.push(async ({ eq, getMiddleBricksIds }) => {
  // check that the middle column bricks have the `foundation` attribute to `true`
  const expectedIds = allBricksIds.filter(
    (b) => b.replace('brick-', '') % 3 === 2,
  )
  const middleBricksIds = await getMiddleBricksIds()
  eq(middleBricksIds, expectedIds)
})

tests.push(async ({ eq, hammer, getExpectedRepairedIds, getRepairedIds }) => {
  // check that the bricks to repair have the right repaired attribute
  await hammer.click()
  eq(await getRepairedIds(), await getExpectedRepairedIds())
})

tests.push(async ({ eq, hammer, getExpectedRepairedIds, getRepairedIds }) => {
  // check that the bricks to repair have the right repaired attribute
  await hammer.click()
  eq(await getRepairedIds(), await getExpectedRepairedIds())
})

tests.push(async ({ eq, dynamite, getBricksIds }) => {
  // check that the last brick is removed on each dynamite click
  for (const i of allBricksIds.keys()) {
    await dynamite.click()
    const { length } = allBricksIds
    const expectedRemainingBricks = allBricksIds.slice(0, length - (i + 1))
    eq(await getBricksIds(), expectedRemainingBricks)
  }
})
