import { gossips } from './subjects/gossip-grid/data.js'

export const tests = []

tests.push(async ({ page, eq }) => {
  // test that the grid is properly generated

  const content = await page.$$eval('.gossip', nodes => nodes.map(n => n.textContent))

  eq(content, ['Share gossip!', ...gossips])
})


tests.push(async ({ page, eq }) => {
  // test that you can add a gossip

  return new Promise(()=>{})
})


tests.push(async ({ page, eq }) => {
  // test that you can change the width

  return new Promise(()=>{})
})


tests.push(async ({ page, eq }) => {
  // test that you can change the font-size

  return new Promise(()=>{})
})


tests.push(async ({ page, eq }) => {
  // test that you can change the background

  return new Promise(()=>{})
})


