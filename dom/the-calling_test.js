export const tests = []

tests.push(async ({ page, eq }) => {
  // check the 3 sections have the correct id and no text
  const elements = await page.$$eval('body', (nodes) =>
    [...nodes[0].children].map((node) => ({
      tag: node.tagName.toLowerCase(),
      text: node.textContent,
      id: node.id,
    })),
  )
  eq(expectedSections, elements)
})

const expectedSections = [
  { tag: 'section', text: '', id: 'face' },
  { tag: 'section', text: '', id: 'upper-body' },
  { tag: 'section', text: '', id: 'lower-body' },
]
