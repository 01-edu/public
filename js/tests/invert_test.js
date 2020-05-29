export const tests = []
const t = (f) => tests.push(f)

// It works with a single property
t(({ eq }) => eq(invert({ language: 'english' }), { english: 'language' }))

// It works with multiple properties
t(({ eq }) =>
  eq(invert({ firstName: 'John', lastName: 'Doe', age: 32 }), {
    John: 'firstName',
    Doe: 'lastName',
    32: 'age',
  }),
)

// Last similar value should override the others
t(({ eq }) =>
  eq(
    invert({ brand: 'ford', motor: 'v8', year: 2000, fast: true, eco: true }),
    {
      ford: 'brand',
      v8: 'motor',
      2000: 'year',
      true: 'eco',
    },
  ),
)

// It should ignore properties from the prototype chain
t(({ eq }) => eq(invert({ f: 5, __proto__: { d: 6 } }), { 5: 'f' }))

Object.freeze(tests)
