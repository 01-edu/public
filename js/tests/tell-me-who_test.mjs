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
const name = 'tell-me-who'
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
const shuffle = arr => {
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
  list
    .map((g) =>
      g
        .split(' ')
        .reverse()
        .join(' '),
    )
    .sort()
    .map((g, i) => `${i +1}. ${g}`)
    .join('\n')

export const setup = async () => {
  const dir = tmpdir()

  // check if already exists and rm?
  await mkdir(`${dir}/${name}`)
  const randomList = getRandomList(guests)
  const expected = getExpected(randomList)
  console.log({randomList, expected})
  await Promise.all(
    randomList.map(
      async (n) =>
        await writeFile(
          `${dir}/${name}/${n.replace(' ', '_')}.json`,
          '',
          'utf8',
        ),
    ),
  )

  return { tmpPath: dir, expected }
}
const printGuestList = async ({ arg, ctx, path, eq }) => {
  const scriptPath = join(resolve(), path)
  const { stdout } = await exec(`node ${scriptPath} ${arg}`, {
    cwd: arg ? `${ctx.tmpPath}` : `${ctx.tmpPath}/${name}`,
  })
  console.log({stdout})

  // await rmdir(`${ctx.tmpPath}/${name}`, { recursive: true })
  return eq(stdout.trim(), ctx.expected)
}

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with "tell-me-who" as an argument
  // `tell-me-who` folder has random files names
  return printGuestList({
    path,
    eq,
    ctx,
    arg: 'tell-me-who',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  // in the `tell-me-who` folder
  return printGuestList({
    path,
    eq,
    ctx,
    arg: '',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `tell-me-who` folder's absolute path as argument
  return printGuestList({
    path,
    eq,
    ctx,
    arg: `${ctx.tmpPath}/tell-me-who`,
  })
})

Object.freeze(tests)
