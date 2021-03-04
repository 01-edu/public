import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const { mkdir, rmdir, readFile } = fs
const exec = promisify(cp.exec)

export const setup = async ({ path }) => {
  const tmpPath = `${tmpdir()}/personal-shopper`
  const filePath = `${tmpPath}/personal-shopper.json`

  await mkdir(tmpPath)
  const run = async cmd => {
    const output = await exec(`node ${path} ${filePath} ${cmd}`)
    const fileContent = await readFile(`${filePath}`, 'utf8')

    return {
      data: JSON.parse(fileContent),
      stdout: output.stdout.trim(),
      stderr: output.stderr.trim(),
    }
  }

  return { run }
}

export const tests = []

tests.push(async ({ eq, ctx }) => {
  // create make an empty json file

  const { data, stderr } = await ctx.run('create')
  return eq({ data, stderr }, { data: {}, stderr: '' })
})

tests.push(async ({ path, eq, ctx }) => {
  // no elem specifiel, should log an error

  return eq(await ctx.run('add'), {
    stdout: '',
    stderr: 'No elem specified.',
    data: {},
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // add without number, should add 1

  const { data } = await ctx.run('add beetroot')
  return eq(data, { beetroot: 1 })
})

tests.push(async ({ path, eq, ctx }) => {
  // with a valid number it shoud add the correct amount

  const { data } = await ctx.run(`add beetroot 5`)
  return eq(data, { beetroot: 6 })
})

tests.push(async ({ path, eq, ctx }) => {
  // negative numbers should reduce the amount

  const { data } = await ctx.run(`add beetroot -1`)
  return eq(data, { beetroot: 5 })
})

tests.push(async ({ path, eq, ctx }) => {
  // add NaN value, should add 1

  const { data } = await ctx.run('add beetroot someNaN')
  return eq(data, { beetroot: 6 })
})

tests.push(async ({ path, eq, ctx }) => {
  // add negative exact number of elem (= 0)
  // -> should delete the entry

  const { data } = await ctx.run(`add beetroot -6`)
  return eq(data, {})
})

tests.push(async ({ path, eq, ctx }) => {
  // add negative value, bigger than current value
  // -> should delete the entry

  await ctx.run(`add beetroot 4`)
  const { data } = await ctx.run(`add beetroot -10`)
  return eq(data, {})
})

tests.push(async ({ path, eq, ctx }) => {
  // no elem specified, should log an error

  return eq(await ctx.run('rm'), {
    stdout: '',
    stderr: 'No elem specified.',
    data: {},
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm elem which is not in the list, should do nothing

  const { data, stderr } = await ctx.run(`rm mustard`)
  return eq({ stderr, data}, { stderr: '', data: {} })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm negative value, should add it

  const { data } = await ctx.run(`rm mustard -4`)
  return eq(data, { mustard: 4 })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm positive value, smaller than current value
  // -> should substract it from current value

  const { data } = await ctx.run(`rm mustard 2`)
  return eq(data, { mustard: 2 })
})

tests.push(async ({ path, eq, ctx }) => {
  // rm without number, should delete the entry

  const { data } = await ctx.run(`rm mustard`)
  return eq(data, {})
})

tests.push(async ({ path, eq, ctx }) => {
  // should print an empty list

  const { stdout } = await ctx.run('ls')
  return eq(`Empty list.`, stdout.trim())
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the list with the elem as expected inside

  await ctx.run(`add mustard 2`)
  const { stdout } = await ctx.run('ls')
  return eq(`- mustard (2)`, stdout.trim())
})

tests.push(async ({ path, eq, ctx, randStr, between}) => {
  // should print the list with the elem as expected inside

  const randomElem = randStr(7)
  const randomAmount = between(1000, 5000)

  await ctx.run(`add mustard 2`)
  await ctx.run(`add beetroot 1`)
  await ctx.run(`add chips ${randomAmount}`)
  await ctx.run(`add ${randomElem} 4`)
  const { stdout } = await ctx.run('ls')

  return eq(
    [
      `- mustard (4)`,
      `- beetroot (1)`,
      `- chips (${randomAmount})`,
      `- ${randomElem} (4)`,
    ],
    stdout.split('\n'),
  )
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the helper that include every commands

  const { stdout } = await ctx.run(`help`)
  return ['add', 'rm', 'delete', 'create', 'ls', 'help'].every(c =>
    stdout.includes(c),
  )
})

tests.push(async ({ path, eq, ctx }) => {
  // should print the helper

  const { stdout } = await ctx.run(``)
  return ['add', 'rm', 'delete', 'create', 'ls', 'help'].every(c =>
    stdout.includes(c),
  )
})

tests.push(async ({ path, eq, ctx }) => {
  // delete should remove the file

  return ctx.run(`delete`)
    .then(() => Promise.reject(Error('json file still present after delete')))
    .catch(err => err.code === 'ENOENT' || err)
})

Object.freeze(tests)
