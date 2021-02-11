// /*/ // ⚡

setup = function (params) {
    let ctx = {}
    ctx.crypto= require("crypto");
    ctx.fs = require("fs");
    ctx.path = './wallet.pem'
    return ctx
}

// Address starts with 01
t(() => generateAddress().slice(0,2) === '01')

// Address is 256 bits long (+'01')
t(() => generateAddress().length === 66)


// Addresses are distinct
t(() => generateAddress() !== generateAddress())

// Private key file exists
t(({ctx}) => {
    generateAddress()
    return ctx.fs.existsSync(ctx.path)
})

// Private key has the correct format
t(({ctx}) => {
    generateAddress()
    let privKeyPEM = ctx.fs.readFileSync(WALLETPATH, {encoding:'utf8', flag:'r'})
    return ctx.crypto.createPrivateKey(privKeyPEM).type === 'private'
})

// Transaction is has the correct format 
t(({ctx}) => {
    generateAddress()
    let t = createTransaction(7, "01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e169")
    return  t.slice(0,2) === "01" &&
            t.slice(66,68) === "01" &&
            t.slice(132,133) === "7" &&
            t.length > 270
})

// Amount is correctly represented
t(({ctx}) => {
    generateAddress()
    let amount = Math.round(Math.random()*200000)
    let t = createTransaction(amount, "01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e169")
    return t.indexOf(amount.toString(16))!== -1
})

// Address match
t(({ctx}) => {
    let addr = generateAddress()
    let t = createTransaction(200, addr)
    return t.slice(0,66) === addr
})


Object.freeze(tests)
