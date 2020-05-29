export const tests = []
const t = (f) => tests.push(f)

t(() => cutFirst('abcdef') === 'cdef')
t(() => cutFirst('a') === '')

t(() => cutLast('abcdef') === 'abcd')
t(() => cutLast('a') === '')

t(() => cutFirstLast('abcdef') === 'cd')
t(() => cutFirstLast('af') === '')
t(() => cutFirstLast('afd') === '')
t(() => cutFirstLast('yoafdyo') === 'afd')

t(() => keepFirst('abcdef') === 'ab')
t(() => keepFirst('a') === 'a')

t(() => keepLast('abcdef') === 'ef')
t(() => keepLast('a') === 'a')

t(() => keepFirstLast('abcdef') === 'abef')
t(() => keepFirstLast('af') === 'af')
t(() => keepFirstLast('afd') === 'afd')
t(() => keepFirstLast('yoafdyo') === 'yoyo')

Object.freeze(tests)
