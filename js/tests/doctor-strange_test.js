export const tests = []
const t = (f) => tests.push(f)
//First week year 0001
t(() => addWeek(new Date('0001-01-01')) === 'Monday')
t(() => addWeek(new Date('0001-01-02')) === 'Tuesday')
t(() => addWeek(new Date('0001-01-03')) === 'Wednesday')
t(() => addWeek(new Date('0001-01-04')) === 'Thursday')
t(() => addWeek(new Date('0001-01-05')) === 'Friday')
t(() => addWeek(new Date('0001-01-06')) === 'Saturday')
t(() => addWeek(new Date('0001-01-07')) === 'Sunday')
t(() => addWeek(new Date('0001-01-08')) === 'secondMonday')
t(() => addWeek(new Date('0001-01-09')) === 'secondTuesday')
t(() => addWeek(new Date('0001-01-10')) === 'secondWednesday')
t(() => addWeek(new Date('0001-01-11')) === 'secondThursday')
t(() => addWeek(new Date('0001-01-12')) === 'secondFriday')
t(() => addWeek(new Date('0001-01-13')) === 'secondSaturday')
t(() => addWeek(new Date('0001-01-14')) === 'secondSunday')

//random years
t(() => addWeek(new Date('2025-08-11')) === 'secondMonday')
t(() => addWeek(new Date('2001-05-11')) === 'secondFriday')
t(() => addWeek(new Date('2001-11-07')) === 'secondWednesday')
t(() => addWeek(new Date('0001-12-01')) === 'secondSaturday')
t(() => addWeek(new Date('1664-08-09')) === 'Saturday')
t(() => addWeek(new Date('1995-11-07')) === 'Tuesday')
t(() => addWeek(new Date('2020-01-01')) === 'Wednesday')
t(() => addWeek(new Date('2048-12-07')) === 'Monday')

//random time traveling
const test = (time, rightTime, { eq }) =>
  eq(time.getTime(), new Date(rightTime).getTime())

t((_) =>
  test(timeTravel('2020-05-29 23:25:22', 21, 22, 22), '2020-05-29 21:22:22', _)
)
t((_) =>
  test(timeTravel('2000-05-09 01:28:02', 21, 22, 22), '2000-05-09 21:22:22', _)
)
t((_) =>
  test(timeTravel('2018-06-04 13:01:00', 10, 16, 11), '2018-06-04 10:16:11', _)
)
t((_) =>
  test(timeTravel('1995-11-07 00:21:12', 23, 12, 18), '1995-11-07 23:12:18', _)
)
t((_) =>
  test(timeTravel('1000-09-19 06:00:00', 22, 10, 21), '1000-09-19 22:10:21', _)
)
t((_) =>
  test(timeTravel('1975-05-10 10:07:56', 17, 15, 14), '1975-05-10 17:15:14', _)
)

Object.freeze(tests)
