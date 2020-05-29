export const tests = []
const t = (f) => tests.push(f)

// work with simple key / values
t(() => get({ key: 'value' }, 'key') === 'value')

// work with nested objects
t(() => get({ nested: { key: 'value' } }, 'nested.key') === 'value')

// return undefined without error if the value do not exist
t(() => get({ key: 'value' }, 'nx') === undefined)
t(() => get({ nested: { key: 'value' } }, 'nested.nx') === undefined)
t(() => get({ nested: { key: 'value' } }, 'nx.nx') === undefined)

// work with nested arrays too
t(() => get({ a: [{ b: t }] }, 'a.0.b') === t)
t(() => get({ a: [{ b: t }] }, 'a.0.b.toString') === t.toString)

Object.freeze(tests)
