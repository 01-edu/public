import { gossips } from './subjects/gossip-grid/gossip-grid.data.js'

export const tests = []

tests.push(async ({ page, eq }) => {
  // test that the grid is properly generated

  const content = await page.$$eval('.gossip', nodes => nodes.map(n => n.textContent))

  eq(content, ['Share gossip!', ...gossips])
})


tests.push(async ({ page, eq }) => {
  // test that you can add a gossip

  const rand = Math.random().toString(36).slice(2)
  await page.type('textarea', `coucou ${rand}`)
  await page.click('.gossip button')
  const content = await page.$eval('div.gossip', n => n.textContent)

  eq(content, `coucou ${rand}`)
})


const getStyle = (nodes, key) => nodes.map(n => n.style[key])
tests.push(async ({ page, eq }) => {
  // test that you can change the width to the min

  const min = await page.$eval('#width', n => {
    n.value = n.min
    n.dispatchEvent(new Event('input'))
    return Number(n.min)
  })

  eq(min, 200)

  const values = await page.$$eval('div.gossip', getStyle, 'width')

  eq(Array(gossips.length + 1).fill(`${min}px`), values)
})

tests.push(async ({ page, eq }) => {
  // test that you can change the width to the max

  const max = await page.$eval('#width', n => {
    n.value = n.max
    n.dispatchEvent(new Event('input'))
    return Number(n.max)
  })

  eq(max, 800)

  const values = await page.$$eval('div.gossip', getStyle, 'width')

  eq(Array(gossips.length + 1).fill(`${max}px`), values)
})

tests.push(async ({ page, eq }) => {
  // test that you can change the font-size to the min

  const min = await page.$eval('#fontSize', n => {
    n.value = n.min
    n.dispatchEvent(new Event('input'))
    return Number(n.min)
  })

  eq(min, 20)

  const values = await page.$$eval('div.gossip', getStyle, 'fontSize')

  eq(Array(gossips.length + 1).fill(`${min}px`), values)
})

tests.push(async ({ page, eq }) => {
  // test that you can change the font-size to the max

  const max = await page.$eval('#fontSize', n => {
    n.value = n.max
    n.dispatchEvent(new Event('input'))
    return Number(n.max)
  })

  eq(max, 40)

  const values = await page.$$eval('div.gossip', getStyle, 'fontSize')

  eq(Array(gossips.length + 1).fill(`${max}px`), values)
})

const getBackground = (nodes, key) => nodes.map(n => n.style.background || n.style.backgroundColor)
tests.push(async ({ page, eq, rgbToHsl }) => {
  // test that you can change the background to the darkest

  const min = await page.$eval('#background', n => {
    n.value = n.min
    n.dispatchEvent(new Event('input'))
    return Number(n.min)
  })

  eq(min, 20)

  const values = await page.$$eval('div.gossip', getBackground)
  const lightness = values.map(rgbToHsl).map(([h,s,l]) => l)
  eq(Array(gossips.length + 1).fill(min), lightness)
})

tests.push(async ({ page, eq, rgbToHsl }) => {
  // test that you can change the background to the darkest

  const max = await page.$eval('#background', n => {
    n.value = n.max
    n.dispatchEvent(new Event('input'))
    return Number(n.max)
  })

  eq(max, 75)

  const values = await page.$$eval('div.gossip', getBackground)
  const lightness = values.map(rgbToHsl).map(([h,s,l]) => Math.round(l))

  eq(Array(gossips.length + 1).fill(max), lightness)
})
