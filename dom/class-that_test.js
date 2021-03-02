export const tests = []

tests.push(async ({ page, eq }) => {
  // check the class 'eye' has been declared properly in the CSS
  eq.css('.eye', {
    width: '60px',
    height: '60px',
    backgroundColor: 'red',
    borderRadius: '50%',
  })
})

tests.push(async ({ page, eq }) => {
  // check the class 'arm' has been declared properly in the CSS
  eq.css('.arm', { backgroundColor: 'aquamarine' })
})

tests.push(async ({ page, eq }) => {
  // check the class 'leg' has been declared properly in the CSS
  eq.css('.leg', { backgroundColor: 'dodgerblue' })
})

tests.push(async ({ page, eq }) => {
  // check the class 'body-member' has been declared properly in the CSS
  eq.css('.body-member', { width: '50px', margin: '30px' })
})

tests.push(async ({ page, eq }) => {
  // check that the targetted elements have the correct class names
  await eq.$('p#eye-left', { className: 'eye' })
  await eq.$('p#eye-right', { className: 'eye' })
  await eq.$('div#arm-left', { className: 'arm body-member' })
  await eq.$('div#arm-right', { className: 'arm body-member' })
  await eq.$('div#leg-left', { className: 'leg body-member' })
  await eq.$('div#leg-right', { className: 'leg body-member' })
})
