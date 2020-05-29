export const tests = []
const t = (f) => tests.push(f)

// handle simple array
t(() => getLength([2, 42]) === 2)

// handle mixed array
t(() => getLength(['pouet', 4, true]) === 3)

t(() => getLength(Array(100)) === 100) // handle holey array
t(() => getLength('salut') === 5) // handle strings
t(() => getLength([]) === 0) // handle empty arrays
t(() => getLength('') === 0) // handle empty strings

Object.freeze(tests)
