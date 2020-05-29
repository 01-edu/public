// small database with nutrition facts, per 100 grams
// prettier-ignore
const nutritionDB = {
  tomato:  { calories: 18,  protein: 0.9,   carbs: 3.9,   sugar: 2.6, fiber: 1.2, fat: 0.2   },
  vinegar: { calories: 20,  protein: 0.04,  carbs: 0.6,   sugar: 0.4, fiber: 0,   fat: 0     },
  oil:     { calories: 48,  protein: 0,     carbs: 0,     sugar: 123, fiber: 0,   fat: 151   },
  onion:   { calories: 0,   protein: 1,     carbs: 9,     sugar: 0,   fiber: 0,   fat: 0     },
  garlic:  { calories: 149, protein: 6.4,   carbs: 33,    sugar: 1,   fiber: 2.1, fat: 0.5   },
  paprika: { calories: 282, protein: 14.14, carbs: 53.99, sugar: 1,   fiber: 0,   fat: 12.89 },
  sugar:   { calories: 387, protein: 0,     carbs: 100,   sugar: 100, fiber: 0,   fat: 0     },
  orange:  { calories: 49,  protein: 0.9,   carbs: 13,    sugar: 12,  fiber: 0.2, fat: 0.1   },
}
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

// filter keys
t(({ eq, ctx }) =>
  eq(
    filterKeys(ctx.cart, (k) => k.length <= 6),
    ctx.filtered,
  ),
)
t(({ eq, ctx }) =>
  eq(
    filterKeys(ctx.cart, (k) => /onion/.test(k)),
    { onion: 200 },
  ),
)

// map keys
t(({ eq, ctx }) =>
  eq(
    mapKeys(ctx.cart, (k) => `✔️${k}`),
    ctx.mapped,
  ),
)
t(({ eq, ctx }) =>
  eq(
    mapKeys(
      filterKeys(ctx.cart, (k) => k === 'onion'),
      (k) => (k = 'orange'),
    ),
    { orange: 200 },
  ),
)
t(({ eq, ctx }) =>
  eq(
    mapKeys(
      filterKeys(nutritionDB, (k) => k === 'tomato'),
      (k) => `${k}DB`,
    ),
    ctx.combo,
  ),
)

// reduce keys
t(({ eq, ctx }) =>
  eq(
    reduceKeys(ctx.cart, (acc, cr) => acc.concat(', ', cr)),
    'vinegar, sugar, oil, onion, garlic, paprika',
  ),
)

t(({ eq, ctx }) =>
  eq(
    reduceKeys(ctx.cart, (acc, cr) => `${acc}${cr}:`, ':'),
    ':vinegar:sugar:oil:onion:garlic:paprika:',
  ),
)

const join = (acc, cr) => (acc == null ? cr : `${acc}:${cr}`)
t(({ eq, ctx }) =>
  eq(
    reduceKeys(nutritionDB, join, null),
    'tomato:vinegar:oil:onion:garlic:paprika:sugar:orange',
  ),
)
t(({ eq, ctx }) =>
  eq(
    reduceKeys(ctx.cart, join, undefined),
    'vinegar:sugar:oil:onion:garlic:paprika',
  ),
)

t(({ eq, ctx }) =>
  eq(
    reduceKeys(ctx.cart, (acc, cr) => (acc += (cr.length <= 4) & 1), 0),
    1,
  ),
)

Object.freeze(tests)

export const setup = () => ({
  cart: {
    vinegar: 80,
    sugar: 100,
    oil: 50,
    onion: 200,
    garlic: 22,
    paprika: 4,
  },
  filtered: { sugar: 100, oil: 50, onion: 200, garlic: 22 },
  mapped: {
    '✔️vinegar': 80,
    '✔️sugar': 100,
    '✔️oil': 50,
    '✔️onion': 200,
    '✔️garlic': 22,
    '✔️paprika': 4,
  },
  combo: {
    tomatoDB: {
      calories: 18,
      protein: 0.9,
      carbs: 3.9,
      sugar: 2.6,
      fiber: 1.2,
      fat: 0.2,
    },
  },
})
