// prettier-ignore
const personnel = {
  lukeSkywalker: { id: 5,  pilotingScore: 98, shootingScore: 56, isForceUser: true  },
  sabineWren:    { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
  zebOrellios:   { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false },
  ezraBridger:   { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true  },
  calebDume:     { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true  },
}
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

// default values
t(({ eq }) => eq(defaultCurry({ http: 403 })({}), { http: 403 }))
t(({ eq }) =>
  eq(defaultCurry({ http: 403, connection: 'close' })({ http: 200 }), {
    http: 200,
    connection: 'close',
  }),
)

// object mutation
t(({ eq }) =>
  eq(defaultCurry(Object.freeze({ http: 403 }))(Object.freeze({ http: 200 })), {
    http: 200,
  }),
)

// multiple values
t(({ eq }) =>
  eq(
    defaultCurry({ http: 403, age: 0, connection: 'close' })({
      http: 200,
      age: 30,
      connection: 'keep-alive',
      content_type: 'text/css',
    }),
    { http: 200, age: 30, connection: 'keep-alive', content_type: 'text/css' },
  ),
)

// map curry
t(({ eq }) =>
  eq(mapCurry(([k, v]) => [`${k}🤙🏼`, `${v}🤙🏼`])({ emoji: 'cool' }), {
    'emoji🤙🏼': 'cool🤙🏼',
  }),
)

// reduce curry
t(({ eq }) =>
  eq(
    reduceCurry((acc, [k, v]) => acc.concat(' ', `${k}:${v.id}`))(
      personnel,
      'personnel:',
    ),
    'personnel: lukeSkywalker:5 sabineWren:82 zebOrellios:22 ezraBridger:15 calebDume:11',
  ),
)

// filter curry
t(({ eq }) =>
  eq(filterCurry(([, v]) => v.id > 22)(personnel), {
    sabineWren: {
      id: 82,
      pilotingScore: 73,
      shootingScore: 99,
      isForceUser: false,
    },
  }),
)

// reduce score
t(({ eq }) => eq(reduceScore(personnel, 0), 420))
t(({ eq }) => eq(reduceScore(personnel, 420), 840))

//filter score
t(({ eq, ctx }) => eq(filterForce(personnel), ctx.filter))

// map average
t(({ eq, ctx }) => eq(mapAverage(personnel), ctx.total))

Object.freeze(tests)

// prettier-ignore
export const setup = () => ({
  filter: {
    calebDume:     { id: 11, isForceUser: true,  pilotingScore: 71, shootingScore: 85 },
    sabineWren:    { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
  },
  total: {
    sabineWren:    { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false, averageScore: 86 },
    zebOrellios:   { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false, averageScore: 39.5 },
    lukeSkywalker: { id: 5,  pilotingScore: 98, shootingScore: 56, isForceUser: true,  averageScore: 77 },
    ezraBridger:   { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true,  averageScore: 55 },
    calebDume:     { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true,  averageScore: 78 },
  },
})
