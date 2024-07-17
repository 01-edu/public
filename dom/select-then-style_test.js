
export const tests = []

tests.push(async ({ eq }) => {
  // check the CSS stylesheet is linked in the head tag

  await eq.$('head link', {
    rel: 'stylesheet',
    href: 'http://localhost:9898/select-and-style/select-and-style.css',
  })
})


tests.push(async ({ eq }) => {
  // check the universal selector has been declared properly

  await eq.css('*', {
    margin: '0px',
    opacity: '0.85',
    boxSizing: 'border-box',
  })
})


tests.push(async ({ eq }) => {
  // check that the body was styled

  await eq.css('body', { height: '100vh' })
})


tests.push(async ({ eq }) => {
  // check that sections elements are styled

  await eq.css('section', {
    padding: '20px',
    width: '100%',
    height: 'calc(33.3333%)',
  })
})


tests.push(async ({ eq }) => {
  // check that the individual sections are styled

  await eq.css('#face', { backgroundColor: 'cyan' })
  await eq.css('#upper-body', { backgroundColor: 'blueviolet' })
  await eq.css('#lower-body', { backgroundColor: 'lightsalmon' })
})
