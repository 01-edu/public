String.prototype.split = undefined
String.prototype.match = undefined
RegExp.prototype.exec = undefined
Array.prototype.join = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(split('a b c', ' '), ['a', 'b', 'c']))
t(({ eq }) => eq(split('ggg - ddd - b', ' - '), ['ggg', 'ddd', 'b']))
t(({ eq }) => eq(split('ee,ff,g,', ','), ['ee', 'ff', 'g', '']))
t(({ eq }) => eq(split('Riad', ' '), ['Riad']))
t(({ eq }) => eq(split('rrrr', 'rr'), ['', '', '']))
t(({ eq }) => eq(split('rrirr', 'rr'), ['', 'i', '']))
t(({ eq }) => eq(split('Riad', ''), ['R', 'i', 'a', 'd']))
t(({ eq }) => eq(split('', 'Riad'), ['']))

t(() => join(['ee', 'ff', 'g', ''], ',') === 'ee,ff,g,')
t(() => join(['ggg', 'ddd', 'b'], ' - ') === 'ggg - ddd - b')
t(() => join(['a', 'b', 'c'], ' ') === 'a b c')
