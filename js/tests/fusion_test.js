export const tests = []
const t = (f) => tests.push(f)

// simple numbers
t(({ eq }) => eq(fusion({ nbr: 12 }, { nbr: 23 }).nbr, 35))

// handle 0
t(({ eq }) => eq(fusion({ nbr: 0 }, { nbr: 23 }).nbr, 23))
t(({ eq }) => eq(fusion({ nbr: 23 }, { nbr: 0 }).nbr, 23))

t(({ eq }) =>
  // multiply numbers
  eq(fusion({ a: 12, b: 2, c: 43 }, { a: 23, b: 2 }), { a: 35, b: 4, c: 43 })
)

t(({ eq }) =>
  // simple string
  eq(fusion({ str: 'hello' }, { str: 'there' }).str, 'hello there')
)

t(({ eq }) =>
  // handle empty strings
  eq(fusion({ str: 'hello' }, { str: '' }).str, 'hello ')
)

t(({ eq }) =>
  // multiple strings
  eq(fusion({ a: 'A', b: 'B', c: 'C' }, { a: 'B', b: 'C' }), {
    a: 'A B',
    b: 'B C',
    c: 'C',
  })
)

t(({ eq }) =>
  // simple arrays
  eq(fusion({ arr: [1, '2'] }, { arr: [2] }).arr, [1, '2', 2])
)

t(({ eq }) =>
  // multiple arrays
  eq(
    fusion(
      { arr: [], arr1: [1] },
      { arr: [12, 3], arr1: [2, 3], arr2: ['2', '1'] }
    ),
    { arr: [12, 3], arr1: [1, 2, 3], arr2: ['2', '1'] }
  )
)

// different matching
t(({ eq }) => eq(fusion({ a: { b: 1 } }, { a: 1 }).a, 1))
t(({ eq }) => eq(fusion({ a: 1 }, { a: { b: 1 } }).a, { b: 1 }))
t(({ eq }) => eq(fusion({ a: [1, 2] }, { a: 1 }).a, 1))
t(({ eq }) => eq(fusion({ a: 'str' }, { a: 1 }).a, 1))

t(({ eq }) =>
  // deep nested objects
  eq(
    fusion(
      { a: { b: [1, 2], c: { d: 2 } } },
      { a: { b: [0, 2, 1], c: { d: 23 } } }
    ),
    { a: { b: [1, 2, 0, 2, 1], c: { d: 25 } } }
  )
)

Object.freeze(tests)
