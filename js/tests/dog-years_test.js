export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(dogYears('earth', 1000000000), 221.82))
t(({ eq }) => eq(dogYears('mercury', 2134835688), 1966.16))
t(({ eq }) => eq(dogYears('venus', 189839836), 68.45))
t(({ eq }) => eq(dogYears('mars', 2129871239), 251.19))
t(({ eq }) => eq(dogYears('jupiter', 901876382), 16.86))
t(({ eq }) => eq(dogYears('saturn', 2000000000), 15.07))
t(({ eq }) => eq(dogYears('uranus', 1210123456), 3.19))
t(({ eq }) => eq(dogYears('neptune', 1821023456), 2.45))

Object.freeze(tests)
