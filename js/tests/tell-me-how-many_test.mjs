import * as cp from 'child_process'
import { mkdir, writeFile } from 'fs/promises'
import { join, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'

const exec = promisify(cp.exec)

export const tests = []
export const setup = async ({ path, between }) => {
  const randomFilesNumber = between(30, 6)
  const dir = `${tmpdir()}/tell-me-how-many`

  await mkdir(dir)

  const run = async (cmd) => {
    const cmdPath = isAbsolute(cmd || '') ? cmd : join(dir, cmd || '')
    const { stdout } = await exec(`node ${path} ${cmdPath}`)

    return { stdout: stdout.trim() }
  }
  const createXFilesIn = async ({ numberOfFiles, dirPath }) => {
    for (let i = 0; i < numberOfFiles; i++) {
      await writeFile(`${dirPath}/${i}.txt`, '', 'utf8')
    }
  }
  await createXFilesIn({ numberOfFiles: randomFilesNumber, dirPath: dir })

  return { randomFilesNumber, tmpPath: dir, run, createXFilesIn }
}

tests.push(async ({ path, eq, ctx, between }) => {
  const numberOfFiles = between(5, 13)
  const dirName = `tell-me-how-many-${numberOfFiles}`
  const dirPath = join(ctx.tmpPath, `../${dirName}`)
  await mkdir(dirPath)
  await ctx.createXFilesIn({ dirPath, numberOfFiles })

  const { stdout } = await ctx.run(`../${dirName}`)
  return eq(Number(stdout), numberOfFiles)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script in a directory named `tell-me-how-many`
  // '../tell-me-how-many' in the argument passed
  // `tell-me-how-many` directory has a random file number
  const { stdout } = await ctx.run('../tell-me-how-many')
  return eq(Number(stdout), ctx.randomFilesNumber)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  const { stdout } = await ctx.run()
  return eq(Number(stdout), ctx.randomFilesNumber)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `tell-me-how-many` directory's absolute path as argument
  const { stdout } = await ctx.run(ctx.tmpPath)
  return eq(Number(stdout), ctx.randomFilesNumber)
})

Object.freeze(tests)
