export const tests = []
const t = (f) => tests.push(f)

t(() => sunnySunday(new Date('0001-01-01')) === 'Monday')
t(() => sunnySunday(new Date('0001-01-02')) === 'Tuesday')
t(() => sunnySunday(new Date('0001-01-03')) === 'Wednesday')
t(() => sunnySunday(new Date('0001-01-04')) === 'Thursday')
t(() => sunnySunday(new Date('0001-01-05')) === 'Friday')
t(() => sunnySunday(new Date('0001-01-06')) === 'Saturday')
t(() => sunnySunday(new Date('0001-01-07')) === 'Monday')

t(() => sunnySunday(new Date('0001-12-01')) === 'Friday')
t(() => sunnySunday(new Date('1664-08-09')) === 'Saturday')

t(() => sunnySunday(new Date('2020-01-01')) === 'Monday')
t(() => sunnySunday(new Date('2048-12-08')) === 'Thursday')

Object.freeze(tests)
