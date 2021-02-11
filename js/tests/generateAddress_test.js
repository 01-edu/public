// /*/ // ⚡

const t = (f) => tests.push(f)

// Check if keys can be imported
t(() => {
    let { address, publicKey, privateKey } = generateAddress()
    return true
})

// Import for following tests
let {address, publicKey, privateKey} = generateAddress()

// Check if keys can be imported and their properties
t(() => crypto.createPublicKey(publicKey).type === 'public')
t(() => crypto.createPublicKey(publicKey).asymmetricKeyType === 'ec')
t(() => crypto.createPrivateKey(privateKey).type === 'private')

// Compare if the provided public key result from the private key 
t(() => crypto.createPublicKey(privateKey).export({type:'spki', format:'pem'})=== crypto.createPublicKey(publicKey).export({type:'spki', format:'pem'}))

// Address starts with 01
t(() => address.slice(0,2) === '01')

// Address is 256 bits long (+'01')
t(() => address.length === 66)

// Address corresponds to key precise or intermediate
t(() => {
    let binPublicKey = crypto.createPublicKey(publicKey).export({type:'spki', format:'der'})
    let xy = binPublicKey.slice(24, 89)
    let approxHash = crypto.createHash("sha256").update(publicKey).digest('hex')
    let IntermHash = crypto.createHash("sha256").update(binPublicKey).digest('hex')
    let preciseHash = crypto.createHash("sha256").update(xy).digest('hex')
    let expectedHash = address.slice(2,256)
    return (preciseHash === expectedHash) || (IntermHash === expectedHash)
})

Object.freeze(tests)
