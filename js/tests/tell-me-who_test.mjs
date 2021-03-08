import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const mkdir = fs.mkdir
const rmdir = fs.rmdir
const writeFile = fs.writeFile

const exec = promisify(cp.exec)

export const tests = []

// maybe get the sames from an api? like https://parser.name/
const guests = [
  'Shyam Langley',
  'Zishan Randolph',
  'Connor Connolly',
  'Edna Floyd',
  'Robyn Jimenez',
  'Carly Alexander',
  'Jed Newman',
  'Marianna Sullivan',
  'Glyn Townsend',
  'Montana Mooney',
  'Amelia-Rose Trejo',
  'Micah Whittle',
  'Nola Sherman',
  'Gregory Vu',
  'Lili Griffiths',
  'Tasnia Hughes',
  'Trixie Pennington',
  'Ava Meyer',
  'Gabriela Tucker',
  'Kiri Wilcox',
]
const shuffle = (arr) => {
  let i = arr.length
  let j, tmp
  while (--i > 0) {
    j = Math.floor(Math.random() * (i + 1))
    tmp = arr[j]
    arr[j] = arr[i]
    arr[i] = tmp
  }
  return arr
}
const ranStr = () =>
  Math.random()
    .toString(36)
    .substring(7)
const getRandomList = (names) =>
  shuffle(names).slice(0, Math.floor(Math.random() * (names.length - 10) + 10))
const getExpected = (list) =>
  list
    .map((g) =>
      g
        .split(' ')
        .reverse()
        .join(' '),
    )
    .sort()
    .map((g, i) => `${i + 1}. ${g}`)
    .join('\n')

export const setup = async ({ path }) => {
  const dir = `${tmpdir()}/tell-me-who`

  await mkdir(dir)
  const randomList = getRandomList(guests)
  const expected = getExpected(randomList)

  const createFilesIn = async ({ fileNames, folderPath }) =>
    await Promise.all(
      fileNames.map(
        async (fileName) => await writeFile(`${folderPath}/${fileName}`, ''),
      ),
    )
  await createFilesIn({
    fileNames: randomList.map((n) => `${n.replace(' ', '_')}.json`),
    folderPath: dir,
  })

  const run = async (cmd) => {
    const cmdPath = isAbsolute(cmd || '') ? cmd : join(dir, cmd || '')
    const { stdout } = await exec(`node ${path} ${cmdPath}`)

    return { stdout: stdout.trim() }
  }

  return { tmpPath: dir, expected, run, createFilesIn }
}

tests.push(async ({ path, eq, ctx }) => {
  const fileNames = [
    'Ubaid_Ballard.json',
    'Victoria_Chan.json',
    'Dominika_Mullen.json',
    'Heath_Denton.json',
    'Lilith_Hamilton.json',
  ]
  const folderName = `tell-me-who-${ranStr()}`
  const folderPath = join(ctx.tmpPath, `../${folderName}`)
  await mkdir(folderPath)
  await ctx.createFilesIn({ folderPath, fileNames })

  const { stdout } = await ctx.run(`../${folderName}`)
  return eq(
    [
      `1. Ballard Ubaid`,
      `2. Chan Victoria`,
      `3. Denton Heath`,
      `4. Hamilton Lilith`,
      `5. Mullen Dominika`,
    ],
    stdout.split('\n'),
  )
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script in a folder named `tell-me-who`
  // '../tell-me-who' in the argument passed
  // `tell-me-who` folder has a random file number
  const { stdout } = await ctx.run('../tell-me-who')
  return eq(stdout, ctx.expected)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  const { stdout } = await ctx.run()
  return eq(stdout, ctx.expected)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `tell-me-who` folder's absolute path as argument
  const { stdout } = await ctx.run(ctx.tmpPath)
  return eq(stdout, ctx.expected)
})

Object.freeze(tests)
