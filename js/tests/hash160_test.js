// /*/ // ⚡
setup = function (params) {
  let ctx = {}
  ctx.crypto = require("crypto");
  ctx.h = function (entry){
    let hash = ctx.crypto.createHash("sha256").update(Buffer.from(entry)).digest()
    return ctx.crypto.createHash("ripemd160").update(hash).digest()
  }
  return ctx
}

t(() => 0===0)
t(({ctx}) => Buffer.compare(hash160("lala"), ctx.h("lala"))===0)
t(({ctx}) => Buffer.compare(hash160("banana"), ctx.h("banana"))===0)
t(({ctx}) => Buffer.compare(hash160(""), ctx.h(""))===0)

Object.freeze(tests)

