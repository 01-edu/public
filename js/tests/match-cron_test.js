export const tests = []
const t = (f) => tests.push(f)

t(() => matchCron('* * * * 1', new Date('2020-06-01 00:00:00')))
t(() => matchCron('* * * 2 *', new Date('2021-02-01 00:00:00')))
t(() => matchCron('* * 9 * *', new Date('2020-06-09 00:00:00')))
t(() => matchCron('* 3 * * *', new Date('2020-05-31 03:00:00')))
t(() => matchCron('1 * * * *', new Date('2020-05-30 19:01:00')))
t(() => matchCron('3 3 * 3 3', new Date('2021-03-03 03:03:00')))
t(() => matchCron('* * * * *', new Date()))
t(({ ctx }) => matchCron('* 7 * * *', new Date(`201${ctx}-01-01 07:00:00`)))

t(() => !matchCron('* * * * 1', new Date('2020-06-02 00:00:00')))
t(() => !matchCron('* * * 2 *', new Date('2021-03-01 00:00:00')))
t(() => !matchCron('* * 8 * *', new Date('2020-06-09 00:00:00')))
t(() => !matchCron('* 2 * * *', new Date('2020-05-31 03:00:00')))
t(() => !matchCron('1 * * * *', new Date('2020-05-30 19:00:00')))
t(() => !matchCron('3 3 * 3 3', new Date('2021-03-02 03:03:00')))
t(({ ctx }) => !matchCron('* 7 * * *', new Date(`201${ctx}-01-01 06:00:00`)))

Object.freeze(tests)

export const setup = () => Math.ceil(Math.random() * 9)
