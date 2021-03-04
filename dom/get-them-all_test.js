import { people } from './subjects/get-them-all/get-them-all.data.js'

const getIds = predicate =>
  people
    .filter(predicate)
    .map(e => e.id)
    .sort((a, b) => a.localeCompare(b))

const architects = getIds(p => p.tag === 'a')
const notArchitects = getIds(p => p.tag !== 'a')

const classical = getIds(p => p.classe === 'classical')
const notClassical = getIds(p => p.tag === 'a' && p.classe !== 'classical')

const active = getIds(p => p.classe === 'classical' && p.active)
const notActive = getIds(
  p => p.tag === 'a' && p.classe === 'classical' && p.active === false,
)

const bonanno = people.find(p => p.id === 'BonannoPisano').id
const notBonanno = getIds(
  p =>
    p.tag === 'a' &&
    p.classe === 'classical' &&
    p.active &&
    p.id !== 'BonannoPisano',
)

export const tests = []

tests.push(async ({ eq, page }) => {
  // get architects
  const btnArchitect = await page.$(`#btnArchitect`)
  btnArchitect.click()
  await page.waitForTimeout(150)

  const selected = await page.$$eval('a', nodes =>
    nodes
      .filter(node => node.textContent === 'Architect')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  eq(selected, architects)

  const eliminated = await page.$$eval('span', nodes =>
    nodes
      .filter(node => node.style.opacity === '0.2')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  eq(eliminated, notArchitects)
})

tests.push(async ({ page, eq }) => {
  // get classical
  const btnClassical = await page.$(`#btnClassical`)
  btnClassical.click()
  await page.waitForTimeout(150)

  const selected = await page.$$eval('.classical', nodes =>
    nodes
      .filter(node => node.textContent === 'Classical')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  eq(selected, classical)

  const eliminated = await page.$$eval('a:not(.classical)', nodes =>
    nodes
      .filter(node => node.style.opacity === '0.2')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  eq(eliminated, notClassical)
})

tests.push(async ({ page, eq }) => {
  // check active
  const btnActive = await page.$(`#btnActive`)
  btnActive.click()
  await page.waitForTimeout(150)

  const selected = await page.$$eval('.classical.active', nodes =>
    nodes
      .filter(node => node.textContent === 'Active')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )
  eq(selected, active)

  const eliminated = await page.$$eval('.classical:not(.active)', nodes =>
    nodes
      .filter(node => node.style.opacity === '0.2')
      .map(node => node.id)
      .sort((a, b) => a.localeCompare(b)),
  )

  eq(eliminated, notActive)
})

tests.push(async ({ page, eq }) => {
  // get bonanno
  const btnBonanno = await page.$(`#btnBonanno`)
  btnBonanno.click()
  await page.waitForTimeout(150)

  const selected = await page.$eval('#BonannoPisano', node => {
    if (node.textContent === 'Bonanno Pisano') return node.id
  })

  eq(`bonanno: ${selected}`, `bonanno: ${bonanno}`)

  const eliminated = await page.$$eval(
    'a.classical.active:not(#BonannoPisano)',
    nodes =>
      nodes
        .filter(node => node.style.opacity === '0.2')
        .map(node => node.id)
        .sort((a, b) => a.localeCompare(b)),
  )

  eq(eliminated, notBonanno)
})
