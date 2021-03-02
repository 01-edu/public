import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const mkdir = fs.mkdir
const rmdir = fs.rmdir
const writeFile = fs.writeFile
const readFile = fs.readFile

const exec = promisify(cp.exec)

export const tests = []
const name = 'tell-me-vip'
// maybe get the sames from an api? like https://parser.name/
const guests = [
  'Shyam Langley',
  'Austin Harwood',
  'Reem Morgan',
  'Neal Chamberlain',
  'Ryan Walters',
  'Ocean Battle',
  'Ubaid Ballard',
  'Victoria Chan',
  'Dominika Mullen',
  'Heath Denton',
  'Lilith Hamilton',
  'Aisling Bailey',
  'Maizie Love',
  'Nathanial Franco',
  'Charmaine Bernard',
  'Sohail Downes',
  'Rabia Gomez',
  'Brendan Brennan',
  'Shannen Atherton',
  'Esa Villarreal',
  'Kayla Wynn',
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
  'Olli Obrien',
  'Jagoda Avalos',
  'Bethanie Emery',
  'Kenya Medina',
  'Ava-Mai Estes',
  'Robyn Jimenez',
  'Carly Alexander',
  'Jed Newman',
  'Marianna Sullivan',
  'Alicja Scott',
  'Isaac Guerrero',
  'Dion Huff',
  'Milly Quintero',
  'Kwabena Cairns',
  'Rukhsar Conley',
  'Glyn Townsend',
  'Colby Holmes',
  'Zeynep East',
  'Miriam Higgins',
  'Kaelan Clegg',
  'Sharna English',
  'Uma Ortega',
  'Crystal Bird',
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
const getExpected = (list) =>
  Object.entries(list)
    .filter(([n, { answer }]) => answer === 'yes')
    .map(([n, _]) =>
      n
        .split(' ')
        .reverse()
        .join(' '),
    )
    .sort()
    .map((g, i) => `${i + 1}. ${g}`)
    .join('\n')
const generateObj = () => ({
  answer: ['yes', 'no'][Math.floor(Math.random() * 2)],
})

export const setup = async () => {
  const dir = tmpdir()

  // check if already exists and rm?
  await mkdir(`${dir}/${name}`)
  const randomList = getRandomList(guests)
  const randomAnswers = Object.fromEntries(
    randomList.map((g) => [g, generateObj()]),
  )
  const expected = getExpected(randomAnswers)
  await Promise.all(
    Object.entries(randomAnswers).map(
      async ([n, answer]) =>
        await writeFile(
          `${dir}/${name}/${n.replace(' ', '_')}.json`,
          JSON.stringify(answer, null, '\t'),
          'utf8',
        ),
    ),
  )

  return { tmpPath: dir, expected }
}
const printVIPGuestList = async ({ arg, ctx, path, eq }) => {
  const scriptPath = join(resolve(), path)
  const cwd = arg ? `${ctx.tmpPath}` : `${ctx.tmpPath}/${name}`
  const { stdout } = await exec(`node ${scriptPath} ${arg}`, {
    cwd,
  })
  const out = await readFile(`${cwd}/vip.txt`, 'utf8').catch((err) =>
    err.code === 'ENOENT' ? 'output file not found' : err,
  )
  await exec(`rm vip.txt`, { cwd })
  return eq(out, ctx.expected)
}

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with "tell-me-vip" as an argument
  // `tell-me-vip` folder has random files names
  return printVIPGuestList({
    path,
    eq,
    ctx,
    arg: 'tell-me-vip',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  // in the `tell-me-vip` folder
  return printVIPGuestList({
    path,
    eq,
    ctx,
    arg: '',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `tell-me-vip` folder's absolute path as argument
  return printVIPGuestList({
    path,
    eq,
    ctx,
    arg: `${ctx.tmpPath}/tell-me-vip`,
  })
})

Object.freeze(tests)
