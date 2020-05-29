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

// filter values
t(({ ctx, eq }) =>
  eq(
    filterValues(ctx.groceriesCart, (v) => v < 80),
    { oil: 50, garlic: 22 },
  ),
)
t(({ ctx, eq }) =>
  eq(
    filterValues(
      nutritionDB,
      (v) => Object.entries(filterValues(v, (ele) => ele === 0)).length !== 0,
    ),
    ctx.have0,
  ),
)

// map value
t(({ ctx, eq }) =>
  eq(
    mapValues(
      filterValues(ctx.groceriesCart, (v) => v >= 200),
      (ele) => ele - 100,
    ),
    { tomato: 100, onion: 120 },
  ),
)

// reduce value
t(({ ctx }) => reduceValues(ctx.groceriesCart, (acc, cr) => acc + cr) === 572)
t(() => reduceValues({ a: 1, b: 2, c: 3 }, (acc, cr) => acc + cr) === 6)
t(() => reduceValues({ a: 1, b: 2, c: 3 }, (acc, cr) => acc + cr, 3) === 9)

Object.freeze(tests)

export const setup = () => ({
  groceriesCart: {
    tomato: 200,
    vinegar: 80,
    oil: 50,
    onion: 220,
    garlic: 22,
  },
  have0: {
    vinegar: {
      calories: 20,
      protein: 0.04,
      carbs: 0.6,
      sugar: 0.4,
      fiber: 0,
      fat: 0,
    },
    oil: { calories: 48, protein: 0, carbs: 0, sugar: 123, fiber: 0, fat: 151 },
    onion: { calories: 0, protein: 1, carbs: 9, sugar: 0, fiber: 0, fat: 0 },
    paprika: {
      calories: 282,
      protein: 14.14,
      carbs: 53.99,
      sugar: 1,
      fiber: 0,
      fat: 12.89,
    },
    sugar: {
      calories: 387,
      protein: 0,
      carbs: 100,
      sugar: 100,
      fiber: 0,
      fat: 0,
    },
  },
})
