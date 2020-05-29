export const tests = []
const t = (f) => tests.push(f)

// prettier-ignore
t(({ eq }) => eq(chunk(['a', 'b', 'c', 'd'], 2), [['a', 'b'], ['c', 'd']]))
t(({ eq }) => eq(chunk(['a', 'b', 'c', 'd'], 3), [['a', 'b', 'c'], ['d']]))

Object.freeze(tests)
