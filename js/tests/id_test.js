export const tests = []
const t = (f) => tests.push(f)

// id is declared and is a function
t(() => typeof id === 'function')

// id take 1 argument
t(() => id.length === 1)

// id return numbers back
t(() => id(5) === 5)

// id return strings back
t(() => id('pouet') === 'pouet')

// id return itself, why not
t(() => id(id) === id)

// id return anything really
t((_) => id(_) === _)

Object.freeze(tests)
