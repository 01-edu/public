const { strictEqual } = require("assert")
const { createContext } = require("vm")

// /*/ // ⚡
setup = function () {
    let ctx = {}
    ctx.h = e => crypto.createHash("sha256").update(e).digest("hex")
    return ctx
}

t(({ctx}) => ctx.h(semiBrute("ff")).substring(0,2)=="ff")
t(({ctx}) => ctx.h(semiBrute("e2")).substring(0,2)=="e2")
t(({ctx}) => ctx.h(semiBrute("c6")).substring(0,2)!="c3")

Object.freeze(tests)
