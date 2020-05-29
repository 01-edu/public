String.prototype.repeat = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(() => typeof repeat === 'function')
t(() => repeat.length === 2)
t(() => repeat('a', 3) === 'aaa')
t(() => repeat('ba', 10) === 'babababababababababa')
t(() => repeat('pouet', 2) === 'pouetpouet')
t(() => repeat('haha', 1) === 'haha')
t(() => repeat('hehehe', 0) === '')

Object.freeze(tests)
