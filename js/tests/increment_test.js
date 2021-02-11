// /*/ // ⚡

t(() => Buffer.compare(increment("01"), Buffer.from("02","hex"))===0)
t(() => Buffer.compare(increment("10"), Buffer.from("11","hex"))===0)
t(() => Buffer.compare(increment("ff"), Buffer.from("0100","hex"))===0)
t(() => Buffer.compare(increment("d537"), Buffer.from("d538","hex"))===0)
t(() => Buffer.compare(increment("17ffff"), Buffer.from("180000","hex"))===0)

Object.freeze(tests)

