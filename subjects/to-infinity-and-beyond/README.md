## To Infinity And Beyond

### The `else if` keyword

Last keyword of the serie, this one combine the two previous keywords to allow
to chain conditions:

```js
if (temperature < 8) {
  console.log('very cold !')
} else if (temperature < 16) {
  console.log('still too cold...')
} else if (temperature < 24) {
  console.log('getting warmer')
} else if (temperature < 32) {
  console.log('nice :)')
} else {
  console.log('Too hot !!!')
}
```

You can chain as mainy as you want, the syntax is similar to the if.

### Instructions

The year is 2042 Elon's dream came true He now has spaceship lines to mars and
to the moon and a big pool of candidates who all want to go to space but there
is still a heavy selection process based on:

- number `candidate.physicalAptitudes` from `0` to `100`
- boolean `candidate.noFamily`

You have to create a condition to go to Mars which is always the most demanded
destination for this destination:

- if the candidate `physicalAptitudes` are below `80`, he must stay on `'earth'`
- only candidates with `noFamily` can go on `'mars'`
- otherwhise they can still go to the `'moon'`

You must log which planet the `candidate` can access

> And there seems to be no sign of intelligent life anywhere. \
> ― Buzz Lightyear
