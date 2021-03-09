import { mkdir, writeFile } from 'fs/promises'
import { tmpdir } from 'os'
import { join, isAbsolute } from 'path'
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

export const tests = []

export const setup = async ({ path }) => {
  const dir = `${tmpdir()}/verydisco-reverso`
  await mkdir(dir)
  const run = async (cmd) => {
    const cmdPath = isAbsolute(cmd || '') ? cmd : join(dir, cmd || '')
    const { stdout } = await exec(`node ${path} ${cmdPath}`)

    return { stdout: stdout.trim() }
  }

  return { tmpPath: dir, run }
}

tests.push(async ({ eq, ctx }) => {
  const trueWords = 'kisscool'
  await writeFile(`${ctx.tmpPath}/fever.txt`, trueWords, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/fever.txt`)
  return eq(stdout, 'coolkiss')
})

tests.push(async ({ eq, ctx }) => {
  const trueWords = 'verydisco'
  await writeFile(`${ctx.tmpPath}/fever.txt`, trueWords, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/fever.txt`)
  return eq(stdout, 'discovery')
})

tests.push(async ({ eq, ctx }) => {
  const trueWords = 'deNo si omeawes'
  await writeFile(`${ctx.tmpPath}/myTruth.txt`, trueWords, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/myTruth.txt`)
  return eq(stdout, 'Node is awesome')
})

tests.push(async ({ randStr, eq, ctx }) => {
  const recto = randStr(3)
  const verso = randStr(3)
  const meaningOfLife = `${recto}${verso}`
  await writeFile(`${ctx.tmpPath}/theTruth.txt`, meaningOfLife, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/theTruth.txt`)
  return eq(stdout, `${verso}${recto}`)
})

tests.push(async ({ randStr, eq, ctx }) => {
  const heads = randStr(3)
  const tails = randStr(4)
  const coinFlipping = `${heads}${tails}`
  await writeFile(`${ctx.tmpPath}/coins-flipping.txt`, coinFlipping, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/coins-flipping.txt`)
  return eq(stdout, `${tails}${heads}`)
})

tests.push(async ({ randStr, eq, ctx }) => {
  const ying = randStr(8)
  const yang = randStr(8)
  const tic = randStr(5)
  const tac = randStr(6)
  const absurd = `${tic}${tac} ${ying}${yang}`
  await writeFile(`${ctx.tmpPath}/absurd.txt`, absurd, 'utf8')

  const { stdout } = await ctx.run(`${ctx.tmpPath}/absurd.txt`)
  return eq(stdout, `${tac}${tic} ${yang}${ying}`)
})

Object.freeze(tests)
