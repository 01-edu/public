import { deepStrictEqual } from 'assert'
import puppeteer from 'puppeteer-core'
import people from '../assets/data/get-them-all.js'

const config = {
  headless: false,
  executablePath:
    '/Applications/Google Chrome.app/Contents/MacOS/Google Chrome',
}

const browser = await puppeteer.launch(config)
const [page] = await browser.pages()
await page.goto('http://localhost:8000/dom-js/get-them-all/')

const architects = people
  .filter((p) => p.tag === 'a')
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const notArchitects = people
  .filter((p) => p.tag !== 'a')
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const checkArchitects = async () => {
  const btnArchitect = await page.$(`#btnArchitect`)
  btnArchitect.click()
  await page.waitFor(500)

  const selected = await page.$$eval('a', (nodes) =>
    nodes
      .filter((node) => node.textContent === 'Architect')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  const eliminated = await page.$$eval('span', (nodes) =>
    nodes
      .filter((node) => node.style.opacity === '0.2')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  deepStrictEqual(`architects: ${selected}`, `architects: ${architects}`)
  deepStrictEqual(
    `not architects: ${eliminated}`,
    `not architects: ${notArchitects}`,
  )
}

checkArchitects()
await page.waitFor(1000)

// get classical
const classical = people
  .filter((p) => p.classe === 'classical')
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const notClassical = people
  .filter((p) => p.tag === 'a' && p.classe !== 'classical')
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const checkClassical = async () => {
  const btnClassical = await page.$(`#btnClassical`)
  btnClassical.click()
  await page.waitFor(500)

  const selected = await page.$$eval('.classical', (nodes) =>
    nodes
      .filter((node) => node.textContent === 'Classical')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  const eliminated = await page.$$eval('a:not(.classical)', (nodes) =>
    nodes
      .filter((node) => node.style.opacity === '0.2')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  deepStrictEqual(`classical: ${selected}`, `classical: ${classical}`)
  deepStrictEqual(
    `not classical: ${eliminated}`,
    `not classical: ${notClassical}`,
  )
}

checkClassical()
await page.waitFor(1000)

// get active
const active = people
  .filter((p) => p.classe === 'classical' && p.active)
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const notActive = people
  .filter(
    (p) => p.tag === 'a' && p.classe === 'classical' && p.active === false,
  )
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const checkActive = async () => {
  const btnActive = await page.$(`#btnActive`)
  btnActive.click()
  await page.waitFor(500)

  const selected = await page.$$eval('.classical.active', (nodes) =>
    nodes
      .filter((node) => node.textContent === 'Active')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  const eliminated = await page.$$eval('.classical:not(.active)', (nodes) =>
    nodes
      .filter((node) => node.style.opacity === '0.2')
      .map((node) => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  deepStrictEqual(`active: ${selected}`, `active: ${active}`)
  deepStrictEqual(`not active: ${eliminated}`, `not active: ${notActive}`)
}

checkActive()
await page.waitFor(1000)

// get bonanno
const bonanno = people.find((p) => p.id === 'BonannoPisano').id

const notBonanno = people
  .filter(
    (p) =>
      p.tag === 'a' &&
      p.classe === 'classical' &&
      p.active &&
      p.id !== 'BonannoPisano',
  )
  .map((e) => e.id)
  .sort((a, b) => a.localeCompare(b))

const checkBonanno = async () => {
  const btnBonanno = await page.$(`#btnBonanno`)
  btnBonanno.click()
  await page.waitFor(500)

  const selected = await page.$eval('#BonannoPisano', (node) => {
    if (node.textContent === 'Bonanno Pisano') return node.id
  })

  const eliminated = await page.$$eval(
    'a.classical.active:not(#BonannoPisano)',
    (nodes) =>
      nodes
        .filter((node) => node.style.opacity === '0.2')
        .map((node) => node.id)
        .sort((a, b) => a.localeCompare(b)),
  )

  deepStrictEqual(`bonanno: ${selected}`, `bonanno: ${bonanno}`)
  deepStrictEqual(`not bonanno: ${eliminated}`, `not bonanno: ${notBonanno}`)
}

checkBonanno()
