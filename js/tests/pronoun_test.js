export const tests = []
const t = (f) => tests.push(f)

// no pronouns
t(({ eq }) =>
  eq(
    pronoun(`Your reducer function's returned value is assigned to the accumulator,
whose value is remembered across each iteration throughout the array and
ultimately becomes the final, single resulting value.`),
    {},
  ),
)

// simple count
t(({ eq }) =>
  eq(
    pronoun(`The seal method seals an object, preventing new properties from being
 added to it and marking all existing properties as non-configurable. Values of present 
properties can still be changed as long as they are writable.`),
    {
      it: { word: ['and'], count: 1 },
      they: { word: ['are'], count: 1 },
    },
  ),
)

// multiple pronouns
t(({ eq }) =>
  eq(pronoun('I buy,\ni to,\nYOU buy,\nit have,\nIt buys,\nit is,\nyou go'), {
    i: { word: ['buy', 'to'], count: 2 },
    you: { word: ['buy', 'go'], count: 2 },
    it: { word: ['have', 'buys', 'is'], count: 3 },
  }),
)

// repetition of pronouns
t(({ eq }) =>
  eq(pronoun(`it i it she is gone`), {
    it: { word: [], count: 2 },
    i: { word: [], count: 1 },
    she: { word: ['is'], count: 1 },
  }),
)

// repetition of same pronoun
t(({ eq }) =>
  eq(pronoun('she she she she is'), { she: { word: ['is'], count: 4 } }),
)

// pronoun with out verb
t(({ eq }) =>
  eq(pronoun('we will rock you'), {
    we: { word: ['will'], count: 1 },
    you: { word: [], count: 1 },
  }),
)

Object.freeze(tests)
