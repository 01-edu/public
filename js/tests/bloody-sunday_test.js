export const tests = []
const t = (f) => tests.push(f)

t(() => bloodySunday(new Date('0001-01-01')) === 'Monday')
t(() => bloodySunday(new Date('0001-01-02')) === 'Tuesday')
t(() => bloodySunday(new Date('0001-01-03')) === 'Wednesday')
t(() => bloodySunday(new Date('0001-01-04')) === 'Thursday')
t(() => bloodySunday(new Date('0001-01-05')) === 'Friday')
t(() => bloodySunday(new Date('0001-01-06')) === 'Saturday')
t(() => bloodySunday(new Date('0001-01-07')) === 'Monday')

t(() => bloodySunday(new Date('0001-12-01')) === 'Friday')
t(() => bloodySunday(new Date('1664-08-09')) === 'Saturday')

t(() => bloodySunday(new Date('2020-01-01')) === 'Monday')
t(() => bloodySunday(new Date('2048-12-08')) === 'Thursday')

Object.freeze(tests)
