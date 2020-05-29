Array.prototype.flat = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(flat([1]), [1]))
t(({ eq }) => eq(flat([1, [2]]), [1, 2]))
t(({ eq }) => eq(flat([1, [2, [3]]]), [1, 2, [3]]))
t(({ eq }) => eq(flat([1, [2, [3], [4, [5]]]], 2), [1, 2, 3, 4, [5]]))
t(({ eq }) => eq(flat([1, [2, [3], [4, [5]]]], 3), [1, 2, 3, 4, 5]))
t(({ eq }) => eq(flat([1, [2, [3], [4, [5]]]], Infinity), [1, 2, 3, 4, 5]))

Object.freeze(tests)
