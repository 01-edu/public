export const tests = []

tests.push(async ({ eq, page }) => {
  // check the JS script has been linked
  await eq.$('script', { type: 'module' })

  // check the JS script has a valid src
  const source = await page.$eval(
    'script',
    (node) => node.src.includes('.js') && node.src,
  )
  if (!source.length) throw Error('missing script src')
})

tests.push(async ({ eq, page }) => {
  // check the class 'eye-closed' has been added in the CSS
  eq.css('.eye-closed', {
    height: '4px',
    padding: '0px 5px',
    borderRadius: '10px',
  })
})

tests.push(async ({ eq, page }) => {
  // check the class of left eye before the JS is loaded
  await page.setJavaScriptEnabled(false)
  await page.reload()
  await eq.$('p#eye-left', { className: 'eye' })
})

tests.push(async ({ eq, page }) => {
  // check the class of left eye has been updated after the JS is loaded
  await page.setJavaScriptEnabled(true)
  await page.reload()
  await eq.$('p#eye-left', { className: 'eye eye-closed' })

  // check the background color of left eye has changed after the JS is loaded
  const eyeLeftBg = await page.$eval(
    '#eye-left',
    (node) => node.style.backgroundColor,
  )
  eq(eyeLeftBg, 'black')
})
