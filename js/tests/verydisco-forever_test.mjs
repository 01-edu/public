import { readFile, mkdir } from 'fs/promises'
import { join, resolve } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)
export const tests = []

export const setup = async ({ path }) => {
  const run = async word => {
    await exec(`node ${path} ${word}`)
    const fileContent = await readFile(
      'verydisco-forever.txt',
      'utf8',
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
  return eq(data, 'sski olco')
})

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`kiss cool`)
  return eq(data, 'sski')
})

tests.push(async ({ ctx, eq }) => {
  const { data } = await ctx.run(`"Node is awesome"`)
  return eq(data, 'deNo si omeawes')
})

tests.push(async ({ ctx, eq, randStr }) => {
  const tic = randStr(7)
  const tac = randStr(7)
  const { data } = await ctx.run(`"${tic}${tac}"`)
  return eq(data, `${tac}${tic}`)
})

tests.push(async ({ ctx, eq, randStr }) => {
  const ying = randStr(8)
  const yang = randStr(7)
  const { data } = await ctx.run(`"${ying}${yang}"`)
  return eq(data, `${yang}${ying}`)
})

tests.push(async ({ ctx, eq, randStr }) => {
  const tic = randStr(5)
  const tac = randStr(5)
  const ying = randStr(3)
  const yang = randStr(2)
  const { data } = await ctx.run(`"${tic}${tac} ${ying}${yang}"`)
  return eq(data, `${tac}${tic} ${yang}${ying}`)
})

Object.freeze(tests)
