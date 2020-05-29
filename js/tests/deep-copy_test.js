export const tests = []
const t = (f) => tests.push(f)

// simple object
t(({ eq }) => copyAndCompare(eq, { user: 'mika', age: 37 }))

// simple array
t(({ eq }) => copyAndCompare(eq, [1, 'a']))

// works with any value type
t(({ eq }) => copyAndCompare(eq, [console.log, /hello/]))

// nesting object
t(({ eq }) => copyAndCompare(eq, { a: { b: { c: 1 } } }))

// nesting array
t(({ eq }) => copyAndCompare(eq, [1, [2, [true]]]))

// mixed nesting
t(({ eq }) => copyAndCompare(eq, [{ a: () => {} }, ['b', { b: [3] }]]))

// undefined value
t(({ eq }) => copyAndCompare(eq, { undef: undefined }))

// check deep freeze (caution: might stuns a target for 4 seconds)
t(({ eq }) => {
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
