export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(RNA(''), ''))
t(({ eq }) => eq(RNA('TAGC'), 'AUCG'))
t(({ eq }) => eq(RNA(DNA('AUCG')), 'AUCG'))

t(({ eq }) => eq(DNA(''), ''))
t(({ eq }) => eq(DNA('AUCG'), 'TAGC'))
t(({ eq }) => eq(DNA(RNA('TAGC')), 'TAGC'))

Object.freeze(tests)
