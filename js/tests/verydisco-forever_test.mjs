import fs from 'fs/promises'
import { join, resolve } from 'path'
import { tmpdir } from 'os'
const readFile = fs.readFile
const mkdir = fs.mkdir
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

export const tests = []
const randomLetters = (number) =>
  Math.random()
    .toString(36)
    .substring(0, number)

export const setup = async ({ path }) => {
  const run = async (word) => {
    await exec(`node ${path} ${word}`)
    const fileContent = await readFile('verydisco-forever.txt', 'utf8').catch((err) =>
      err.code === 'ENOENT' ? 'output file not found' : err,
    )

    return { data: fileContent }
  }

  return { run }
}

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`discovery`)
  return eq(data, 'verydisco')
})

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`"kiss cool"`)
  return eq(data, "sski olco")
})

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`kiss cool`)
  return eq(data, "sski")
})

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`"Node is awesome"`)
  return eq(data, "deNo si omeawes")
})

tests.push(async ({ ctx, eq }) => {
  const tic = randomLetters(7)
  const tac = randomLetters(7)
  const { data } = await ctx.run(`"${tic}${tac}"`)
  return eq(data, `${tac}${tic}`)
})

tests.push(async ({ ctx, eq }) => {
  const ying = randomLetters(8)
  const yang = randomLetters(7)
  const { data } = await ctx.run(`"${ying}${yang}"`)
  return eq(data, `${yang}${ying}`)
})

tests.push(async ({ ctx, eq }) => {
  const tic = randomLetters(5)
  const tac = randomLetters(5)
  const ying = randomLetters(3)
  const yang = randomLetters(2)
  const { data } = await ctx.run(`"${tic}${tac} ${ying}${yang}"`)
  return eq(data, `${tac}${tic} ${yang}${ying}`)
})

Object.freeze(tests)
