import { colors } from '../subjects/fifty-shades-of-cold/fifty-shades-of-cold.data.js'

colors.sort()

const cold = ['aqua', 'blue', 'turquoise', 'green', 'purple', 'cyan', 'navy']

const isCold = color => cold.some(coldColor => color.includes(coldColor))
const toClass = (a, b) => `.${a} { background: ${b} }`
const toDiv = (a, b) => `<div class="${a}">${b}</div>`

export const tests = []

tests.push(async ({ page, eq }) => {
  // check that the css is properly generated

  const style = await page.$$eval('style', nodes => nodes[1].innerHTML)
  const classes = style
    .split('}')
    .map(s => s.replace(/(\.|{|:|;|\s+)/g, ''))
    .filter(Boolean)
    .sort()

  for (const [i, c] of colors.entries()) {
    if (!classes[i]) throw Error(`Not enough class (expected: ${c})`)
    const [a, b] = classes[i].split('background')
    eq(toClass(a, b), toClass(c, c))
  }
})

tests.push(async ({ page, eq }) => {
  // check that the divs are properly generated

  const values = await page.$$eval('div', nodes =>
    nodes.map(n => [n.className, n.textContent]),
  )
  const divs = values.map(v => toDiv(...v)).sort()

  let skipped = 0
  for (const [i, c] of colors.entries()) {
    if (isCold(c)) {
      if (!values[i - skipped]) throw Error(`Not enough div (expected: ${c})`)
      eq(divs[i - skipped], toDiv(c, c))
      continue
    }
    skipped++
    if (divs.includes(toDiv(c, c))) {
      throw Error(`div ${toDiv(c, c)} is not very cold`)
    }
  }
})

tests.push(async ({ page, eq }) => {
  // test that clicking update the color accordingly

  const coldColors = colors.filter(isCold)
  for (const c of coldColors) {
    await page.$$eval(
      'div',
      (nodes, c) => nodes.find(n => n.textContent === c).click(),
      c,
    )

    const count = await page.$$eval(
      'div',
      (nodes, c) => nodes.filter(n => n.className === c).length,
      c,
    )

    eq(count, coldColors.length)
  }
})
