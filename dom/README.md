# DOM

Tests that use puppeteer to do browser side exercises

## Run test locally

### Installation

> You need node version 14+

```bash
# Clone the repo
git clone https://github.com/01-edu/public.git

# go into the dom directory
cd public/dom

# install puppeteer
npm i puppeteer
```

### Executing a test

```bash
# run a test
SOLUTION_PATH=/user/you/piscine-repo node test.js exercise-name
```

The `SOLUTION_PATH` is the directory where the test should look
for your solution, usualy your piscine repository.

The `exercise-name` argument should match exactly the name of an
exercise, not including `.js`
