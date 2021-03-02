export const tests = []

tests.push(async ({ page, eq }) => {
  // check that the title tag is present & is set with some text
  const title = await page.$$eval(
    'title',
    (nodes) => nodes[0] && nodes[0].innerHTML,
  )
  const isValidTitle = title !== undefined && title.length !== 0
  eq(isValidTitle, true)
})

tests.push(async ({ page, eq }) => {
  // check the 3 sections have been created with the correct text
  const elements = await page.$$eval('body', (nodes) =>
    [...nodes[0].children].map((node) => ({
      tag: node.tagName.toLowerCase(),
      text: node.textContent,
    })),
  )
  eq(expectedSections, elements)
})

const expectedSections = [
  { tag: 'section', text: 'face' },
  { tag: 'section', text: 'upper-body' },
  { tag: 'section', text: 'lower-body' },
]
