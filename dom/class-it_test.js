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
  // Helper function to check if an element has the expected classes
  async function hasClasses(selector, expectedClasses) {
    const element = await page.$(selector);
    const className = await element.evaluate(el => el.className);
    const classList = className.split(' ').sort();
    const expectedClassList = expectedClasses.split(' ').sort();
    return JSON.stringify(classList) === JSON.stringify(expectedClassList);
  }

  // Check that the target elements have the correct classes
  eq(await hasClasses('p#eye-left', 'eye'), true);
  eq(await hasClasses('p#eye-right', 'eye'), true);
  eq(await hasClasses('div#arm-left', 'arm body-member'), true);
  eq(await hasClasses('div#arm-right', 'arm body-member'), true);
  eq(await hasClasses('div#leg-left', 'leg body-member'), true);
  eq(await hasClasses('div#leg-right', 'leg body-member'), true);
});
