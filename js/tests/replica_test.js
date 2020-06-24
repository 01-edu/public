export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) =>
  // objects mutability
  eq(
    replica(
      {},
      Object.freeze({ line: 'Replicants are like any other machine' }),
      Object.freeze({ author: 'Rich' })
    ),
    { line: 'Replicants are like any other machine', author: 'Rich' }
  )
)

t(({ eq }) =>
  // different value types
  eq(replica({ con: console.log }, { reg: /hello/ }), {
    con: console.log,
    reg: /hello/,
  })
)

t(({ eq }) =>
  // primitive into an object
  eq(replica({ a: 4 }, { a: { b: 1 } }).a.b, 1)
)

t(({ eq }) =>
  // nested array into a primitive
  eq(
    replica({ a: { b: { c: [123, 1] } } }, { a: { b: { c: '1' } } }).a.b.c,
    '1'
  )
)

t(({ eq }) =>
  // primitive into an array
  eq(replica({ a: 2 }, { a: [4] }).a, [4])
)

t(({ eq }) =>
  // object into an array
  eq(replica({ a: { b: [2] } }, { a: [4] }).a, [4])
)

t(({ eq }) =>
  // array into an object
  eq(replica({ a: [1, 2, 4] }, { a: { b: [4] } }).a, { b: [4] })
)

t(({ eq }) =>
  // nested objects
  eq(replica({ a: { b: 1, c: 2 } }, { a: { c: 23 } }), { a: { b: 1, c: 23 } })
)

t(({ eq }) =>
  // super nested objects
  // letter+number indicates the depth
  eq(
    replica(
      {},
      { a: { b1: 1, c1: 2 } },
      { a: { b1: { d2: 1, e2: 2 } } },
      { a: { b1: { d2: { f3: 1, h3: 1 }, e2: { g3: 2 } } } },
      { a: { b1: { d2: { f3: { i4: 1 }, h3: 1 }, e2: { g3: 2 } } } }
    ),
    { a: { b1: { d2: { f3: { i4: 1 }, h3: 1 }, e2: { g3: 2 } }, c1: 2 } }
  )
)

Object.freeze(tests)
