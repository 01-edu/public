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

// handle simple array
t(() => getLength([2, 42]) === 2)
// handle mixed array
t(() => getLength(['pouet', 4, true]) === 3)
t(() => getLength(Array(100)) === 100) // handle holey array
t(() => getLength('salut') === 5) // handle strings
t(() => getLength([]) === 0) // handle empty arrays
t(() => getLength('') === 0) // handle empty strings

Object.freeze(tests)
