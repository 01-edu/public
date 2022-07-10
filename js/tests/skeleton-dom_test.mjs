export const tests = []

tests.push(async ({ page, eq }) => {
  // check that the title tag is present & is set with some text
  const title = await page.$eval('title', (node) => node.textContent)
  if (!title.length) throw Error('missing title')
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
  // check the lower-body, my favorite part

  return eq.$('section:nth-child(3)', { textContent: 'lower-body' })
})
