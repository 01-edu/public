export const tests = []
const t = (f) => tests.push(f)

// bigger
t(() => typeof biggie !== 'undefined')
t(() => biggie > 1.7976931348623157e308)

// smaller
t(() => typeof smalls !== 'undefined')
t(() => smalls < -1.7976931348623157e308)

Object.freeze(tests)

/*
  “Damn right I like the life I live,
   because I went from negative to positive.”

      ― The Notorius B.I.G
*/
