export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(flags({}), { alias: { h: 'help' }, description: '' }))

t(({ eq }) =>
  eq(
    flags({
      invert: 'inverts and object',
      'convert-map': 'converts the object to an array',
      assign: 'uses the function assign - assign to target object',
    }),
    $a,
  ),
)

t(({ eq }) =>
  eq(
    flags({
      invert: 'inverts and object',
      'convert-map': 'converts the object to an array',
      assign: 'uses the function assign - assign to target object',
      help: ['assign', 'invert'],
    }),
    $b,
  ),
)

t(({ eq }) =>
  eq(
    flags({
      invert: 'inverts and object',
      'convert-map': 'converts the object to an array',
      assign: 'uses the function assign - assign to target object',
      help: ['invert'],
    }),
    $c,
  ),
)

Object.freeze(tests)

const $a = {
  alias: { h: 'help', i: 'invert', c: 'convert-map', a: 'assign' },
  description: [
    '-i, --invert: inverts and object',
    '-c, --convert-map: converts the object to an array',
    '-a, --assign: uses the function assign - assign to target object',
  ].join('\n'),
}

const $b = {
  alias: { h: 'help', i: 'invert', c: 'convert-map', a: 'assign' },
  description: [
    '-a, --assign: uses the function assign - assign to target object',
    '-i, --invert: inverts and object',
  ].join('\n'),
}

const $c = {
  alias: { h: 'help', i: 'invert', c: 'convert-map', a: 'assign' },
  description: '-i, --invert: inverts and object',
}
