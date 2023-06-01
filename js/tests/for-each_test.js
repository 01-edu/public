Array.prototype.forEach = undefined
Array.prototype.map = undefined
// /*/ // ⚡
export const tests = []
const t = f => tests.push(f)

export const setup = () => {
  const arr = [1, 2, 3, 4, 5, Math.random(), 7, 10, -10, 20, -95]
  Object.getPrototypeOf([]).proto = ' [avoid for..in] '
  return { arr }
}

t(({ eq, ctx }) => {
  // callback is call with the item value
  const result = []
  const returned = forEach(ctx.arr, value => result.push(value))
  return eq([returned, result], [undefined, ctx.arr])
})

t(({ eq, ctx }) => {
  // callback second parameter is the index
  const result = []
  const returned = forEach(ctx.arr, (_, index) => result.push(index))
  return eq([returned, result], [undefined, [...ctx.arr.keys()]])
})

t(({ eq, ctx }) => {
  // callback third parameter is the array
  const result = []
  const returned = forEach(ctx.arr, (_, __, arr) => result.push(arr))
  return eq(
    [returned, result],
    [undefined, Array(ctx.arr.length).fill(ctx.arr)],
  )
})

Object.freeze(tests)
