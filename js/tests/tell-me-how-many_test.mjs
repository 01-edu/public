import * as cp from 'child_process'
import { mkdir, writeFile } from 'fs/promises'
import { join, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'

const exec = promisify(cp.exec)

export const tests = []
export const setup = async ({ path }) => {
  const randomFilesNumber = Math.floor(Math.random() * (30 - 1) + 1)
  const dir = `${tmpdir()}/tell-me-how-many`

  await mkdir(dir)
  for (let i = 0; i < randomFilesNumber; i++) {
    await writeFile(`${dir}/${i}.txt`, '', 'utf8')
  }
  const run = async (cmd) => {
    const cmdPath = isAbsolute(cmd || '') ? cmd : join(dir, cmd || '')
    const { stdout } = await exec(`node ${path} ${cmdPath}`)

    return { stdout: stdout.trim() }
  }

  return { randomFilesNumber, tmpPath: dir, run }
}

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script in a folder named `tell-me-how-many`
  // '../tell-me-how-many' in the argument passed
  // `tell-me-how-many` folder has a random file number
  const { stdout } = await ctx.run('../tell-me-how-many')
  return eq(Number(stdout), ctx.randomFilesNumber)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  const { stdout } = await ctx.run()
  return eq(Number(stdout), ctx.randomFilesNumber)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `tell-me-how-many` folder's absolute path as argument
  const { stdout } = await ctx.run(ctx.tmpPath)
  return eq(Number(stdout), ctx.randomFilesNumber)
})

Object.freeze(tests)
