export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(getAcceleration({}), 'impossible'))
t(({ eq }) => eq(getAcceleration({ d: 10, f: 2, Δv: 100 }), 'impossible'))
t(({ eq }) => eq(getAcceleration({ f: 10, Δv: 100 }), 'impossible'))
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5 }), 2))
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5, Δv: 100, Δt: 50 }), 2))
t(({ eq }) => eq(getAcceleration({ Δv: 100, Δt: 50 }), 2))
t(({ eq }) => eq(getAcceleration({ f: 10, Δv: 100, Δt: 50 }), 2))
t(({ eq }) => eq(getAcceleration({ f: 10, m: 5, Δt: 100 }), 2))
t(({ eq }) => eq(getAcceleration({ d: 10, t: 2, Δv: 100 }), 5))
t(({ eq }) => eq(getAcceleration({ d: 100, t: 2, f: 100 }), 50))

Object.freeze(tests)
