

export const tests = []

tests.push(async ({ eq }) => {
  // check the CSS stylesheet is linked in the head tag

  return eq.$('head link', {
    rel: 'stylesheet',
    href: 'http://localhost:9898/select-and-style/select-and-style.css',
  })
})

tests.push(async ({ eq }) => {
  // check the universal selector has been declared properly

  return eq.css('*', {
    margin: '0px',
    opacity: '0.85',
    boxSizing: 'border-box',
  })
})
