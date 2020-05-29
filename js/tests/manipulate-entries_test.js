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
  orange:  { calories: 49,  protein: 0.9,   carbs: 13,    sugar: 9,   fiber: 0.2, fat: 0.1   },
}
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

// filter entries
t(({ eq, ctx }) =>
  eq(
    filterEntries(ctx.groceriesCart1, ([, v]) => v < 300),
    { onion: 230, garlic: 220 },
  ),
)

// map entries
t(({ eq, ctx }) =>
  eq(
    mapEntries(ctx.groceriesCart1, ([k, v]) => [
      v > 250 ? `✔️${k}` : `❌${k}`,
      v - 250,
    ]),
    {
      '✔️oil': 250,
      '❌onion': -20,
      '❌garlic': -30,
      '✔️paprika': 230,
    },
  ),
)

// filter and map
t(({ eq, ctx }) =>
  eq(
    mapEntries(
      filterEntries(ctx.groceriesCart1, ([k, v]) => k === 'onion'),
      ([k, v]) => [`✔️${k}`, v - 100],
    ),
    { '✔️onion': 130 },
  ),
)

// reduce entries
t(({ eq, ctx }) =>
  eq(
    reduceEntries(ctx.groceriesCart1, (acc, [k, v]) => acc + k + v, ''),
    'oil500onion230garlic220paprika480',
  ),
)

t(({ eq, ctx }) => eq(lowCarbs(ctx.groceriesCart1), { oil: 500, onion: 230 }))
t(({ eq, ctx }) =>
  eq(lowCarbs(ctx.groceriesCart2), { vinegar: 120, tomato: 700 }),
)

t(({ eq, ctx }) => eq(totalCalories(ctx.groceriesCart1), 1921.4))
t(({ eq, ctx }) => eq(totalCalories(ctx.groceriesCart2), 370.5))

t(({ eq, ctx }) => eq(cartTotal(ctx.groceriesCart1), ctx.total1))
t(({ eq, ctx }) => eq(cartTotal(ctx.groceriesCart2), ctx.total2))

Object.freeze(tests)

export const setup = () => ({
  groceriesCart1: { oil: 500, onion: 230, garlic: 220, paprika: 480 },
  groceriesCart2: { tomato: 700, vinegar: 120, orange: 450 },
  total1: {
    oil: {
      calories: 240,
      protein: 0,
      carbs: 0,
      sugar: 615,
      fiber: 0,
      fat: 755,
    },
    onion: {
      calories: 0,
      protein: 2.3,
      carbs: 20.7,
      sugar: 0,
      fiber: 0,
      fat: 0,
    },
    garlic: {
      calories: 327.8,
      protein: 14.08,
      carbs: 72.6,
      sugar: 2.2,
      fiber: 4.62,
      fat: 1.1,
    },
    paprika: {
      calories: 1353.6,
      protein: 67.872,
      carbs: 259.152,
      sugar: 4.8,
      fiber: 0,
      fat: 61.872,
    },
  },
  total2: {
    orange: {
      calories: 220.5,
      carbs: 58.5,
      fat: 0.45,
      fiber: 0.9,
      protein: 4.05,
      sugar: 40.5,
    },
    tomato: {
      calories: 126,
      carbs: 27.3,
      fat: 1.4,
      fiber: 8.4,
      protein: 6.3,
      sugar: 18.2,
    },
    vinegar: {
      calories: 24,
      carbs: 0.72,
      fat: 0,
      fiber: 0,
      protein: 0.048,
      sugar: 0.48,
    },
  },
})
