Array.prototype.forEach = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

export const setup = () => {
  const arr = [1, 2, 3, 4, 5, Math.random(), 7, 10, -10, 20, -95]
  Object.getPrototypeOf([]).proto = ' [avoid for..in] '
  return { arr }
}

// callback is call with the item value
t(({ eq, ctx }) => {
  const result = []
  forEach(ctx.arr, (value) => result.push(value))
  return eq(result, ctx.arr)
})

// callback second parameter is the index
t(({ eq, ctx }) => {
  const result = []
  forEach(ctx.arr, (_, index) => result.push(index))
  return eq(result, [...ctx.arr.keys()])
})

// callback third parameter is the array
t(({ eq, ctx }) => {
  const result = []
  forEach(ctx.arr, (_, __, arr) => result.push(arr))
  return eq(result, Array(ctx.arr.length).fill(ctx.arr))
})

Object.freeze(tests)
