export const tests = []
const t = (f) => tests.push(f)

// more is declared and is a function and take 1 argument
t(() => typeof more === 'function')
t(() => more.length === 1)

// more works with positive numbers
t(() => more(5) === 6)
t(() => more(more(more(5))) === 8) // more more more !!

// ok that's enough, let's do less, is it even declared ?
t(() => typeof less === 'function')
t(() => less.length === 1)

// less decrease argument by 1
t(() => less(5) === 4)
t(() => less(1) === more(-1))

// let's add now
t(() => typeof add === 'function')
t(() => add.length === 2)
t(() => add(3, 10) === 13)
t(() => add(-1, -1) === -2)

// do the sub thingy
t(() => typeof sub === 'function')
t(() => sub.length === 2)
t(() => sub(-1, -1) === 0)
t(() => sub(3, 10) === -7)

// I can combine all of them
t(() => add(more(10), sub(less(5), 10)) === 5)

const rand = Math.random()
t(() => less(rand) === rand - 1)
t(() => more(rand) === rand + 1)

Object.freeze(tests)
