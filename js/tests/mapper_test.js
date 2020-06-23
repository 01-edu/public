Array.prototype.map = undefined
Array.prototype.flatMap = undefined
Array.prototype.flat = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

const add1 = (el) => el + 1
const sub3 = (el) => el - 3
const mult2 = (el) => el * 2

const doAll = (el) => sub3(mult2(add1(el)))
const posValsIndex = (el, i) => (el >= 0 ? `${i}: ${el}` : undefined)
const indexValsArray = (el, i, arr) =>
  `${el} is at index: ${i} out of ${arr.length - 1}`

const arrayFormatSentence = (item, index, arr) => {
  if (index === arr.length - 2) return `and ${arr[arr.length - 2]} `
  if (index === arr.length - 1) {
    return `are ${String(arr.length - 1)} ${item}.`
  }
  return `${item}, `
}

// map
t(({ eq, ctx }) =>
  eq(map(ctx.numbers, add1), [11, -9, 21, -94, 87, 103, 36, 90, 111])
)
t(({ eq, ctx }) =>
  eq(map(ctx.numbers, mult2), [20, -20, 40, -190, 172, 204, 70, 178, 220])
)

t(({ eq, ctx }) =>
  eq(map(ctx.numbers, sub3), [7, -13, 17, -98, 83, 99, 32, 86, 107])
)

t(({ eq, ctx }) =>
  eq(map(ctx.numbers, doAll), [19, -21, 39, -191, 171, 203, 69, 177, 219])
)

t(({ eq, ctx }) =>
  eq(map(ctx.numbers, posValsIndex), [
    '0: 10',
    undefined,
    '2: 20',
    undefined,
    '4: 86',
    '5: 102',
    '6: 35',
    '7: 89',
    '8: 110',
  ])
)

t(({ eq, ctx }) =>
  eq(map(ctx.numbers, indexValsArray), [
    '10 is at index: 0 out of 8',
    '-10 is at index: 1 out of 8',
    '20 is at index: 2 out of 8',
    '-95 is at index: 3 out of 8',
    '86 is at index: 4 out of 8',
    '102 is at index: 5 out of 8',
    '35 is at index: 6 out of 8',
    '89 is at index: 7 out of 8',
    '110 is at index: 8 out of 8',
  ])
)

t(({ eq, ctx }) =>
  eq(
    map(ctx.sentences[0], arrayFormatSentence).join(''),
    'Colombia, Mexico, and El Salvador are 3 Spanish speaking countries.'
  )
)

t(({ eq, ctx }) =>
  eq(
    map(ctx.sentences[1], arrayFormatSentence).join(''),
    'Perou, Brazil, Argentina, and Venezuela are 4 countries in South America.'
  )
)

t(({ eq, ctx }) =>
  eq(
    map(ctx.sentences[2], arrayFormatSentence).join(''),
    'France, Portugal, and Italy are 3 members of the EU.'
  )
)

t(({ eq }) =>
  // map should not flat
  eq(
    map([1, 2, 3], (n) => [n, n]),
    [
      [1, 1],
      [2, 2],
      [3, 3],
    ]
  )
)

// flatMap
t(({ eq }) =>
  // flatMap should flatten the result of map
  eq(
    flatMap([1, 2, 3], (n) => [n, n]),
    [1, 1, 2, 2, 3, 3]
  )
)

t(({ eq, ctx }) =>
  eq(flatMap(ctx.mixed, add1), ['101', -9, 21, -94, 87, '1021', '35,891', 111])
)

t(({ eq, ctx }) =>
  eq(flatMap(ctx.mixed, posValsIndex), [
    '0: 10',
    undefined,
    '2: 20',
    undefined,
    '4: 86',
    '5: 102',
    undefined,
    '7: 110',
  ])
)

t(({ eq, ctx }) =>
  eq(flatMap(ctx.nested, indexValsArray), [
    '5 is at index: 0 out of 7',
    '4 is at index: 1 out of 7',
    '-3 is at index: 2 out of 7',
    '20 is at index: 3 out of 7',
    '17 is at index: 4 out of 7',
    '-33 is at index: 5 out of 7',
    '-4 is at index: 6 out of 7',
    '18 is at index: 7 out of 7',
  ])
)

Object.freeze(tests)

export const setup = () => {
  const numbers = [10, -10, 20, -95, 86, 102, 35, 89, 110]
  const mixed = [[10], -10, 20, -95, 86, [102], [35, 89], 110]
  const nested = [[5], [4], [-3], [20], [17], [-33], [-4], [18]]
  const sentences = [
    ['Colombia', 'Mexico', 'El Salvador', 'Spanish speaking countries'],
    ['Perou', 'Brazil', 'Argentina', 'Venezuela', 'countries in South America'],
    ['France', 'Portugal', 'Italy', 'members of the EU'],
  ]

  Object.getPrototypeOf([]).proto = ' [avoid for..in] '
  Object.freeze(numbers)
  Object.freeze(mixed)
  Object.freeze(nested)
  Object.freeze(sentences[0])
  Object.freeze(sentences[1])
  Object.freeze(sentences[2])

  return { numbers, mixed, nested, sentences }
}
