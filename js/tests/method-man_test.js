export const tests = []
const t = (f) => tests.push(f)

// words
t(({ eq }) => eq(words('a b c'), ['a', 'b', 'c']))
t(({ eq }) => eq(words('Hello  world'), ['Hello', '', 'world']))
t(({ eq, ctx: r }) => eq(words(`${r} ${r} ${r}`), [r, r, r]))
// sentence
t(({ eq }) => eq(sentence(['a', 'b', 'c']), 'a b c'))
t(({ eq }) => eq(sentence(['Hello', '', 'world']), 'Hello  world'))
t(({ eq, ctx: r }) => eq(sentence([r, r, r]), `${r} ${r} ${r}`))
// yell
t(({ eq }) => eq(yell('howdy stranger ?'), 'HOWDY STRANGER ?'))
t(({ eq }) => eq(yell('Déjà vu'), 'DÉJÀ VU'))
// whisper
t(({ eq }) => eq(whisper('DÉJÀ VU'), '*déjà vu*'))
t(({ eq }) => eq(whisper('HOWDY STRANGER ?'), '*howdy stranger ?*'))

// capitalize
t(({ eq }) => eq(capitalize('str'), 'Str'))
t(({ eq }) => eq(capitalize('qsdqsdqsd'), 'Qsdqsdqsd'))
t(({ eq }) => eq(capitalize('STR'), 'Str'))
t(({ eq }) => eq(capitalize('zapZAP'), 'Zapzap'))
t(({ eq }) => eq(capitalize('zap ZAP'), 'Zap zap'))

Object.freeze(tests)

export const setup = () => Math.random().toString(36)
