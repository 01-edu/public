import * as cp from 'child_process'
import { readdir, rm, mkdir, writeFile, readFile } from 'fs/promises'
import { tmpdir } from 'os'
import { promisify } from 'util'

const exec = promisify(cp.exec)
export const tests = []

export const setup = async ({ path }) => {
  const tmpPath = `${tmpdir()}/happiness-manager`

  await mkdir(tmpPath)
  await mkdir(`${tmpPath}/guests`)
  const run = async (dir, file) => {
    const output = await exec(
      `node ${path} ${tmpPath}/${dir} ${tmpPath}/${file}`,
    )
    const fileContent = await readFile(
      `${tmpPath}/${file}`,
      'utf8',
    ).catch(err => (err.code === 'ENOENT' ? 'output file not found' : err))

    return {
      data:
        fileContent === 'output file not found'
          ? fileContent
          : JSON.parse(fileContent),
      stdout: output.stdout.trim(),
    }
  }
  const resetAnswersIn = async ({ dir }) => {
    const files = await readdir(`${tmpPath}/${dir}`)
    await Promise.all(files.map(file => rm(`${tmpPath}/${dir}/${file}`)))
  }
  const createAnswers = (nb, elem) => [...Array(nb).keys()].map(() => elem)
  const setAnswersIn = async ({ answers, dir }) => {
    await resetAnswersIn({ dir })
    await Promise.all(
      answers.map(
        async (content, idx) =>
          await writeFile(
            `${tmpPath}/${dir}/${idx}.json`,
            JSON.stringify(content, null, '\t'),
            'utf8',
          ),
      ),
    )
  }

  return { run, tmpPath, createAnswers, resetAnswersIn, setAnswersIn }
}

tests.push(async ({ eq, ctx }) => {
  // test with no vips (no {answer: yes})
  // no file should be created, a special message should appear in console
  const answers = ctx.createAnswers(2, { answer: 'no' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { stdout, data } = await ctx.run('guests', 'happy-list.json')
  return eq(
    { stdout, data },
    { stdout: 'No one is coming.', data: 'output file not found' },
  )
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'carnivores' }
  // should create a list with burgers and potatoes
  const answers = ctx.createAnswers(2, { answer: 'yes', food: 'carnivore' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-carn-list.json')
  return eq(data, { burgers: 2, potatoes: 2 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'fish' }
  // should create a list with sardines and potatoes
  const answers = [
    { answer: 'no', food: 'fish' },
    ...ctx.createAnswers(3, { answer: 'yes', food: 'fish' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-fish-list.json')
  return eq(data, { potatoes: 3, sardines: 3 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'everything' }
  // should create a list with kebabs and potatoes
  const answers = [
    { answer: 'no', food: 'everything' },
    ...ctx.createAnswers(3, { answer: 'yes', food: 'everything' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-omni-list.json')
  return eq(data, { potatoes: 3, kebabs: 3 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'iced-tea' }
  // should create a list with iced-tea-bottles and potatoes
  const answers = [
    { answer: 'no', drink: 'beer' },
    ...ctx.createAnswers(1, { answer: 'yes', drink: 'iced-tea' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-iced-tea-list.json')
  return eq(data, { potatoes: 1, 'iced-tea-bottles': 1 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'iced-tea' }
  // should create a list with iced-tea-bottles and potatoes
  const answers = [
    { answer: 'no', drink: 'iced-tea' },
    ...ctx.createAnswers(6, { answer: 'yes', drink: 'iced-tea' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-iced-tea-bottles-list.json')
  return eq(data, { potatoes: 6, 'iced-tea-bottles': 1 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'sparkling-water' }
  // should create a list with sparkling-water-bottles and potatoes
  const answers = [
    ...ctx.createAnswers(3, { answer: 'no', drink: 'sparkling-water' }),
    ...ctx.createAnswers(5, { answer: 'yes', drink: 'sparkling-water' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-sparkling-water-list.json')
  return eq(data, { potatoes: 5, 'sparkling-water-bottles': 2 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'sparkling-water' }
  // should create a list with sparkling-water-bottles and potatoes
  const answers = ctx.createAnswers(8, { answer: 'yes', drink: 'sparkling-water' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-sparkling-water-bottles-list.json')
  return eq(data, { potatoes: 8, 'sparkling-water-bottles': 2 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'water' }
  // should create a list with water-bottles and potatoes
  const answers = [
    ...ctx.createAnswers(2, { answer: 'no', drink: 'water' }),
    ...ctx.createAnswers(2, { answer: 'yes', drink: 'water' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-water-list.json')
  return eq(data, { potatoes: 2, 'water-bottles': 1 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'water' }
  // should create a list with water-bottles and potatoes
  const answers = ctx.createAnswers(7, { answer: 'yes', drink: 'water' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-water-bottle-list.json')
  return eq(data, { potatoes: 7, 'water-bottles': 2 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'soft' }
  // should create a list with soft-bottles and potatoes
  const answers = [
    ...ctx.createAnswers(8, { answer: 'no', drink: 'soft' }),
    ...ctx.createAnswers(12, { answer: 'yes', drink: 'soft' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-soft-list.json')
  return eq(data, { potatoes: 12, 'soft-bottles': 3 })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { drink: 'soft' }
  // should create a list with soft-bottles and potatoes
  const answers = ctx.createAnswers(13, { answer: 'yes', drink: 'soft' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-soft-bottle-list.json')
  return eq(data, { potatoes: 13, 'soft-bottles': 4 })
})

// tests with veggstuff
// 1) vegan but no veggie
tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'vegan' }
  // should create a list with eggplants, mushrooms, courgettes and potatoes
  const answers = [
    ...ctx.createAnswers(2, { answer: 'no', food: 'vegan' }),
    ...ctx.createAnswers(4, { answer: 'yes', food: 'vegan' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-vegan-list.json')
  return eq(data, {
    potatoes: 4,
    mushrooms: 4,
    eggplants: 2,
    courgettes: 2,
    hummus: 2,
  })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'vegan' }
  // should create a list with eggplants, mushrooms, hummus, courgettes and potatoes
  const answers = ctx.createAnswers(6, { answer: 'yes', food: 'vegan' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-vegan-list.json')
  return eq(data, {
    potatoes: 6,
    mushrooms: 6,
    eggplants: 2,
    courgettes: 2,
    hummus: 2,
  })
})

// 2) veggie but no vegan
tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'veggie' }
  // should create a list with eggplants, mushrooms, courgettes and potatoes
  const answers = [
    ...ctx.createAnswers(2, { answer: 'no', food: 'veggie' }),
    ...ctx.createAnswers(4, { answer: 'yes', food: 'veggie' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-veggie-list.json')
  return eq(data, {
    potatoes: 4,
    mushrooms: 4,
    eggplants: 2,
    courgettes: 2,
    hummus: 2,
  })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'veggie' }
  // should create a list with eggplants, mushrooms, hummus, courgettes and potatoes
  const answers = ctx.createAnswers(6, { answer: 'yes', food: 'veggie' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-veggie-list.json')
  return eq(data, {
    potatoes: 6,
    mushrooms: 6,
    eggplants: 2,
    courgettes: 2,
    hummus: 2,
  })
})

// 3) both
tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'vegan' } and { food: 'veggie' }
  // should create a list with eggplants, mushrooms, courgettes and potatoes
  const answers = [
    ...ctx.createAnswers(4, { answer: 'yes', food: 'vegan' }),
    ...ctx.createAnswers(2, { answer: 'yes', food: 'veggie' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-vegg-list.json')
  return eq(data, {
    potatoes: 6,
    mushrooms: 6,
    eggplants: 2,
    courgettes: 2,
    hummus: 2,
  })
})

tests.push(async ({ eq, ctx }) => {
  // test when vips answer { food: 'vegan' } and { food: 'veggie' }
  // should create a list with eggplants, mushrooms, hummus, courgettes and potatoes
  const answers = [
    ...ctx.createAnswers(6, { answer: 'yes', food: 'vegan' }),
    ...ctx.createAnswers(1, { answer: 'yes', food: 'veggie' }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'happy-vegg-list.json')
  return eq(data, {
    potatoes: 7,
    mushrooms: 7,
    eggplants: 3,
    courgettes: 3,
    hummus: 3,
  })
})

// test with existing file
tests.push(async ({ eq, ctx }) => {
  // test with an existing file
  // should add elems to the existing list
  await writeFile(
    `${ctx.tmpPath}/old-happy-list.json`,
    JSON.stringify({ candies: 2000 }),
  )
  const answers = ctx.createAnswers(1, { answer: 'yes', food: 'vegan' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'old-happy-list.json')
  return eq(data, {
    candies: 2000,
    potatoes: 1,
    mushrooms: 1,
    eggplants: 1,
    courgettes: 1,
    hummus: 1,
  })
})

tests.push(async ({ eq, ctx }) => {
  // test with an existing file
  // should replace elems in the existing list (if already there)
  await writeFile(
    `${ctx.tmpPath}/old-happy-list.json`,
    JSON.stringify({ candies: 2000, potatoes: 32 }),
  )
  const answers = ctx.createAnswers(1, { answer: 'yes', food: 'vegan' })
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'old-happy-list.json')
  return eq(data, {
    candies: 2000,
    potatoes: 1,
    mushrooms: 1,
    eggplants: 1,
    courgettes: 1,
    hummus: 1,
  })
})

// test with a little bit of everything
tests.push(async ({ eq, ctx }) => {
  // test with mix of everything
  await writeFile(
    `${ctx.tmpPath}/party.json`,
    JSON.stringify({ 'super-gift': 1, balloons: 100, 'flower-bouquet': 7 }),
  )
  const answers = [
    ...ctx.createAnswers(7, { answer: 'no', food: 'vegan', drink: 'water' }),
    ...ctx.createAnswers(2, {
      answer: 'yes',
      food: 'carnivore',
      drink: 'soft',
    }),
    ...ctx.createAnswers(6, { answer: 'yes', food: 'vegan', drink: 'water' }),
    ...ctx.createAnswers(2, { answer: 'yes', food: 'veggie', drink: 'water' }),
    ...ctx.createAnswers(3, { answer: 'yes', food: 'veggie', drink: 'iced-tea' }),
    ...ctx.createAnswers(11, { answer: 'yes', food: 'fish', drink: 'sparkling-water' }),
    ...ctx.createAnswers(4, {
      answer: 'yes',
      food: 'everything',
      drink: 'iced-tea',
    }),
  ]
  await ctx.setAnswersIn({ dir: 'guests', answers })

  const { data } = await ctx.run('guests', 'party.json')
  return eq(data, {
    'super-gift': 1,
    balloons: 100,
    'flower-bouquet': 7,
    potatoes: 28,
    mushrooms: 11,
    eggplants: 4,
    courgettes: 4,
    hummus: 4,
    sardines: 11,
    burgers: 2,
    kebabs: 4,
    'iced-tea-bottles': 2,
    'sparkling-water-bottles': 3,
    'water-bottles': 2,
    'soft-bottles': 1,
  })
})

Object.freeze(tests)
