const { createGzip } = require("zlib")

// /*/ // ⚡
setup = function (params) {
  let ctx = {}

  ctx.crypto = require("crypto")
  ctx.fs = require("fs")

  ctx.h = e => crypto.createHash("sha256").update(e).digest("hex")

  ctx.path1 = './test1.txt'
  let intro = "Desocupado lector: sin juramento me podrás creer que quisiera que este libro, como hijo del entendimiento, fuera el más hermoso, el más gallardo y más discreto que pudiera imaginarse."
  fs.writeFileSync(ctx.path1, intro)  
  ctx.hash1 = ctx.h(intro)

  ctx.path2 = './test2.txt'
  let forest = ""
  for (let index = 0; index < Math.ceil(Math.random()*100); index++) {
    forest += "🌴"
  }
  ctx.hash2 = ctx.h(forest)
  fs.writeFileSync(ctx.path2, forest)

  return ctx
}

t(({ctx}) => hashFile(ctx.path1).length)
t(({ctx}) => hashFile(ctx.path1)===ctx.hash1)
t(({ctx}) => hashFile(ctx.path2)===ctx.hash2)
t(({ctx}) => hashFile(ctx.path2)!==ctx.hash1)

Object.freeze(tests)
