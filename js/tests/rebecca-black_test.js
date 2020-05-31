export const tests = []
const t = (f) => tests.push(f)

// isFriday
t(() => isFriday(new Date('2014-08-29')))
t(() => isFriday(new Date('2020-07-17')))
t(() => !isFriday(new Date('1992-08-26')))
t(() => !isFriday(new Date('2000-08-26')))

// isWeekend
t(() => isWeekend(new Date('2014-09-06')))
t(() => isWeekend(new Date('2020-05-30')))
t(() => !isWeekend(new Date('1929-02-13')))
t(() => !isWeekend(new Date('1990-01-30')))

// isLeapYear
t(() => isLeapYear(new Date('1804-02-01')))
t(() => isLeapYear(new Date('2008-02-01')))
t(() => isLeapYear(new Date('2216-02-01')))
t(() => !isLeapYear(new Date('1993-02-01')))
t(() => !isLeapYear(new Date('1999-02-01')))

// isLastDayOfMonth
t(() => isLastDayOfMonth(new Date('2020-02-29')))
t(() => isLastDayOfMonth(new Date('2020-12-31')))
t(() => isLastDayOfMonth(new Date('2019-02-28')))
t(() => isLastDayOfMonth(new Date('1998-02-28')))
t(() => isLastDayOfMonth(new Date('1980-02-29')))
t(() => !isLastDayOfMonth(new Date('2020-12-30')))
t(() => !isLastDayOfMonth(new Date('2020-02-28')))
t(() => !isLastDayOfMonth(new Date('2019-02-29')))

Object.freeze(tests)
