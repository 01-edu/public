export const tests = []
const t = (f) => tests.push(f)

t(() => !(inequality === inequality))

Object.freeze(tests)
