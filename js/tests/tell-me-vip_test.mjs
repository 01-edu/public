import * as cp from 'child_process'
import { mkdir, writeFile, readFile } from 'fs/promises'
import { join, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'

const exec = promisify(cp.exec)

export const tests = []
const name = 'tell-me-vip'
// maybe get the sames from an api? like https://parser.name/
const guests = [
  'Gladys Hardy',
  'Laaibah Rogers',
  'Zishan Randolph',
  'Connor Connolly',
  'Arabella Wooten',
  'Edna Floyd',
  'Roksana Montoya',
  'Macauley Ireland',
  'Kennedy Cummings',
  'Emelia Calhoun',
  'Jimmy Hickman',
  'Leela Solomon',
  'Frederick David',
  'Eryk Winters',
  'Christopher Haas',
  'Olivier Galvan',
  'Esha Herring',
  'Montana Mooney',
  'Amelia-Rose Trejo',
  'Micah Whittle',
  'Nola Sherman',
  'Gregory Vu',
  'Lili Griffiths',
  'Tasnia Hughes',
  'Trixie Pennington',
  'Ava Meyer',
  'Konrad Weaver',
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
const getRandomList = (names) =>
  shuffle(names).slice(0, Math.floor(Math.random() * (names.length - 10) + 10))
const getExpected = (list) => list
    .filter(([n, {answer}]) => answer === 'yes')
    .map(([n, _]) => {
      const reversed = n.split('_').reverse()
      return `${reversed[0].slice(0, -5)} ${reversed[1]}`
    })
    .sort()
    .map((g, i) => `${i + 1}. ${g}`)
    .join('\n')
const generateObj = () => ({
  answer: ['yes', 'no'][Math.floor(Math.random() * 2)],
})

const ranStr = () =>
  Math.random()
    .toString(36)
    .substring(7)
export const setup = async ({ path }) => {
  const dir = `${tmpdir()}/tell-me-vip`

  await mkdir(`${dir}/guests`, { recursive: true })
  const randomList = getRandomList(guests)
  const randomAnswers = randomList.map((g) => [
    g.replace(' ', '_').concat('.json'),
    generateObj(),
  ])

  const expected = getExpected(randomAnswers)

  const createFilesIn = async ({ files, folderPath }) =>
    await Promise.all(
      files.map(
        async ([fileName, content]) =>
          await writeFile(`${folderPath}/${fileName}`, JSON.stringify(content)),
      ),
    )
  await createFilesIn({
    files: randomAnswers,
    folderPath: `${dir}/guests`,
  })

  const run = async (cmd) => {
    const cmdPath = isAbsolute(cmd) ? cmd : join(dir, cmd)
    const { stdout } = await exec(`node ${path} ${cmdPath}`)
    const fileContent = await readFile(`vip.txt`, 'utf8').catch((err) =>
      err.code === 'ENOENT' ? 'output file not found' : err,
    )
    return { data: fileContent }
  }

  return { tmpPath: dir, expected, run, createFilesIn }
}

tests.push(async ({ path, eq, ctx }) => {
  // test when no answers in the folder
  const folderName = `guests-${ranStr()}`
  const folderPath = join(ctx.tmpPath, folderName)
  await mkdir(folderPath)

  const { data } = await ctx.run(folderName)
  return eq('', data)
})

tests.push(async ({ path, eq, ctx }) => {
  // test when no one said yes
  const files = [
    ['Ubaid_Ballard.json', { answer: 'no' }],
    ['Victoria_Chan.json', { answer: 'no' }],
    ['Dominika_Mullen.json', { answer: 'no' }],
    ['Heath_Denton.json', { answer: 'no' }],
    ['Lilith_Hamilton.json', { answer: 'no' }],
  ]
  const folderName = `guests-${ranStr()}`
  const folderPath = join(ctx.tmpPath, folderName)
  await mkdir(folderPath)
  await ctx.createFilesIn({ folderPath, files })

  const { data } = await ctx.run(folderName)
  return eq('', data)
})

tests.push(async ({ path, eq, ctx }) => {
  const files = [
    ['Ubaid_Ballard.json', { answer: 'yes' }],
    ['Victoria_Chan.json', { answer: 'yes' }],
    ['Dominika_Mullen.json', { answer: 'no' }],
    ['Heath_Denton.json', { answer: 'no' }],
    ['Lilith_Hamilton.json', { answer: 'yes' }],
  ]
  const folderName = `guests-${ranStr()}`
  const folderPath = join(ctx.tmpPath, folderName)
  await mkdir(folderPath)
  await ctx.createFilesIn({ folderPath, files })

  const { data } = await ctx.run(folderName)
  return eq(
    [`1. Ballard Ubaid`, `2. Chan Victoria`, `3. Hamilton Lilith`],
    data.split('\n'),
  )
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script in a folder named `guests`
  // 'guests' in the argument passed
  // `guests` folder has a random file number with random answers 'yes' or 'no'
  const { data } = await ctx.run('guests')
  return eq(data, ctx.expected)
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `guests` folder's absolute path as argument
  const { data } = await ctx.run(`${ctx.tmpPath}/guests`)
  return eq(data, ctx.expected)
})

// test error when no arg?...

Object.freeze(tests)
