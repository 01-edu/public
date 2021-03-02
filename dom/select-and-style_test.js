export const tests = []

tests.push(async ({ page, eq }) => {
  // check the CSS stylesheet is linked in the head tag
  const CSSLink = await page.$$eval('head', (nodes) =>
    [...nodes[0].children].some(
      (node) => node.tagName === 'LINK' && node.rel === 'stylesheet',
    ),
  )
  eq(CSSLink, true)

  // check the universal selector has been declared properly
  const universalSelectorStyle = await page.evaluate(() => {
    const target = [...window.document.styleSheets[0].cssRules].find(
      (rule) => rule.selectorText === '*',
    )
    const { margin, opacity, boxSizing } = target.style
    return { margin, opacity, boxSizing }
  })
  eq(
    { margin: '0px', opacity: '0.85', boxSizing: 'border-box' },
    universalSelectorStyle,
  )
})
