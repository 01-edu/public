import keptPromise 

const getImportantInfo = _ =>
  new Promise(resolve => {
    setTimeout(_ => resolve(Math.round(Math.random() * 10)), 1000)
  })
