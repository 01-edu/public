export const tests = []

tests.push(async ({ page, eq }) => {
  // check the 3 sections have the correct id and no text
  const sections = await page.$$eval('section', (nodes) =>
    nodes.map((node) => ({
      tag: node.tagName.toLowerCase(),
      text: node.textContent,
      id: node.id,
    })),
  )
  eq(expectedSections, sections)
})

const expectedSections = [
  { tag: 'section', text: '', id: 'face' },
  { tag: 'section', text: '', id: 'upper-body' },
  { tag: 'section', text: '', id: 'lower-body' },
]
