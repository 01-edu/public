export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) =>
  // simple object
  copyAndCompare(eq, { user: 'mika', age: 37 })
)

t(({ eq }) =>
  // simple array
  copyAndCompare(eq, [1, 'a'])
)

t(({ eq }) =>
  // works with any value type
  copyAndCompare(eq, [console.log, /hello/])
)

t(({ eq }) =>
  // nesting object
  copyAndCompare(eq, { a: { b: { c: 1 } } })
)

t(({ eq }) =>
  // nesting array
  copyAndCompare(eq, [1, [2, [true]]])
)

t(({ eq }) =>
  // mixed nesting
  copyAndCompare(eq, [{ a: () => {} }, ['b', { b: [3] }]])
)

t(({ eq }) =>
  // undefined value
  copyAndCompare(eq, { undef: undefined })
)

t(({ eq }) => {
  // check deep freeze (caution: might stuns a target for 4 seconds)
  const r = Math.random()
  const obj = [r, Object.freeze([r, Object.freeze([r])])]
  const copy = deepCopy(obj)
  eq(copy, obj)
  return obj[1][1] !== copy[1][1]
})

Object.freeze(tests)

const copyAndCompare = (eq, obj) => {
  const copy = deepCopy(Object.freeze(obj))
  eq(copy, obj)
  return copy !== obj
}
