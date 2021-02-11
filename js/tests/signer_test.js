
// /*/ // ⚡

setup =  function (params) {

} 

const t = (f) => tests.push(f)

// init are unique
t(() => init() !== init())


// Signatures are unique
t(() =>{
  init()
  return signer("O ALquimista") !== signer("Война и миръ") 
})

// Verifier identifies correct signature
t(() =>{
  const message = "O ALquimista"
  const pubKey = init()
  const signature = signer(message)
  return verifier(message,pubKey,signature)
})

// Verifier rejects wrong signatures
t(() =>{
  const pubKey = init()
  const signature = signer("O ALquimista")
  return !verifier("Война и миръ",pubKey,signature)
})

t(() =>{
  const message = "O ALquimista"
  const pubKey = init()
  const signature = "3040021e131c347199b207b4e7de85141c6d606f874a866e86f4cc6d85a680d58726021e1f7bbe71bd76f623170dcf56a51bf0fd1557c6b61a6b19c48d232e63e50f"
  return !verifier(message,pubKey,signature)
})


Object.freeze(tests)



