export const tests = []

tests.push(async ({ page, eq }) => {
  // check that the title tag is present & is set with some text
  const title = await page.$eval('title', (node) => node.textContent)
  if (!title.length) throw Error('missing title')
})

tests.push(async ({ page, eq }) => {
  // check that the title tag is set with text from the given list
  const title = await page.$eval('title', (node) => node.textContent)
  if (title !== 'invisibility' && title !== 'light-speed' && title !== 'super-strength' && title !== 'advanced-healing' && title !== 'mind-link') {
    throw Error('wrong title, pick one of the list')
  }
  // invisibility
  // light-speed
  // super-strength
  // advanced-healing
  // mind-link
})

tests.push(async ({ page, eq }) => {
  // check the face

  return eq.$('section:nth-child(1)', { textContent: 'face' })
})

tests.push(async ({ page, eq }) => {
  // check the upper-body

  return eq.$('section:nth-child(2)', { textContent: 'upper-body' })
})

tests.push(async ({ page, eq }) => {
  // check the lower-body

  return eq.$('section:nth-child(3)', { textContent: 'lower-body' })
})
