import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const mkdir = fs.mkdir
const rmdir = fs.rmdir
const readFile = fs.readFile

const exec = promisify(cp.exec)

export const tests = []
const name = 'personal-shopper'
const getRandomElem = () =>
  Math.random()
    .toString(36)
    .substring(7)
export const setup = async () => {
  const dir = tmpdir()
  const elems = {}
  const randomElem = getRandomElem()
  const anotherElem = getRandomElem()
  const commands = ['add', 'rm', 'delete', 'create', 'ls', 'help']
  // check if already exists and rm
  const exists = await fs
    .stat(`${dir}/${name}`)
    .catch((err) => (err.code === 'ENOENT' ? 'file not found: good.' : err))
  if (Object.keys(exists).length) {
    await rmdir(`${dir}/${name}`, { recursive: true })
  }

  await mkdir(`${dir}/${name}`)

  return { tmpPath: `${dir}/${name}`, elems, commands, randomElem, anotherElem }
}
const assistInShopping = async ({
  file,
  keyword,
  elem,
  number,
  ctx,
  path,
  eq,
}) => {
  // check if file exists before testing create - if exists, rm it
  if (keyword === 'create') {
    const out = await fs
      .stat(`${ctx.tmpPath}/${file}`)
      .catch((err) =>
        err.code === 'ENOENT' ? 'output file not found: normal' : err,
      )
    if (out.birthtime) await fs.rm(`${ctx.tmpPath}/${file}`)
  }

  const scriptPath = join(resolve(), path)
  const { stdout, stderr } = await exec(
    `node ${scriptPath} ${file} ${keyword} ${elem || ''} ${number || ''}`,
    {
      cwd: ctx.tmpPath,
    },
  )
  const out = await readFile(`${ctx.tmpPath}/${file}`, 'utf8').catch((err) =>
    err.code === 'ENOENT' ? 'output file not found' : err,
  )

  if (keyword === 'create') return eq(out, '{}')

  if (keyword === 'add' || keyword === 'rm') {
    if (!elem) return eq(stderr.trim(), 'No elem specified.')
    if (number && isNaN(Number(number)) && keyword === 'rm') {
      return eq(stderr.trim(), 'Unexpected request: nothing has been removed.')
    }
    const content = JSON.parse(out)

    const addValue = !number || isNaN(Number(number)) ? 1 : number
    const rmValue = !number ? ctx.elems[elem] : number
    const newNumber =
      keyword === 'add'
        ? (ctx.elems[elem] || 0) + addValue
        : (ctx.elems[elem] || 0) - rmValue
    const state = newNumber > 0 ? content[elem] === newNumber : !content[elem]
    ctx.elems[elem] = newNumber > 0 ? newNumber : 0
    return state
  }

  if (keyword === 'ls') {
    const content = JSON.parse(out)
    const entries = Object.entries(content)
    return entries.length
      ? eq(entries.map(([k, v]) => `- ${k} (${v})`).join('\n'), stdout.trim())
      : eq(stdout.trim(), 'Empty list.')
  }

  if (keyword === 'help' || !keyword) {
    return ctx.commands.every((c) => stdout.includes(c))
  }

  if (keyword === 'delete') return eq(out, 'output file not found')
}

tests.push(async ({ path, eq, ctx }) => {
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'create',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // no elem specifiel, should log an error
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add NaN value, should add 1
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: 'someNaN',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add without number, should add 1
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add random value, between 5 and 15
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: Math.floor(Math.random() * (15 - 5) + 5),
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add negative value, bigger than current value - should delete the entry
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: -Math.floor(Math.random() * (25 - 16) + 16),
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add value to prepare next test
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: Math.floor(Math.random() * (25 - 16) + 16),
  })
})
tests.push(async ({ path, eq, ctx }) => {
  // add negative exact nm of elem -> 0, should delete the entry
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: -ctx.elems[ctx.randomElem],
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // no elem specified, should log an error
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    // no elem, no number
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm elem which is not in the list, should do nothing
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
    // no number
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm negative value, should add it
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
    number: -Math.floor(Math.random() * (25 - 16) + 16),
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm positive value, smaller than current value -> should substract it from current value
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
    number: Math.floor(Math.random() * (10 - 5) + 5),
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the list with the elem as expected inside
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'ls',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm without number, should delete the entry
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // should print an empty list
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'ls',
  })
})

// add 2 elems - tests: ls several values, rm more or same value as current
tests.push(async ({ path, eq, ctx }) => {
  // add value to prepare next test
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.randomElem,
    number: Math.floor(Math.random() * (25 - 16) + 16),
  })
})
tests.push(async ({ path, eq, ctx }) => {
  // add value to prepare next test
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'add',
    elem: ctx.anotherElem,
    number: ctx.elems[ctx.randomElem], // add same value than randomelem
  })
})
tests.push(async ({ path, eq, ctx }) => {
  // test ls with several elems
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'ls',
  })
})
tests.push(async ({ path, eq, ctx }) => {
  // test where rm exat nm of elem -> 0, delete the entry
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
    number: ctx.elems[ctx.randomElem],
  })
})
tests.push(async ({ path, eq, ctx }) => {
  // test where rm more than nm of elem -> <0, delete the entry
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'rm',
    elem: ctx.randomElem,
    number: 2000,
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the helper
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'help',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the helper
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: '',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  return assistInShopping({
    path,
    eq,
    ctx,
    file: `${name}.json`,
    keyword: 'delete',
  })
})

Object.freeze(tests)
