## Dr Strange

### Instructions

You have been given the mission to create a new sense of time! Normally a week has 7 days right? Well, that is about to change!

Instead of a normal week having only 7 days, it will have 14 days.
Let me explain, this new week will have 14 days, from `Monday` to `Sunday` then `secondMonday` to `secondSunday`.
Your purpose is to create a new function `addWeek`, that takes as parameter a `Date` and that will return what week day the given date is, according to your new week format.
Week number should be count from `0001-01-01`.

```
new week format:

Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
secondMonday
secondTuesday
secondWednesday
secondThursday
secondFriday
secondSaturday
secondSunday
```

Now imagine you have a doctor appointment and you have to wait some hours, but you do not want to wait, so you decided that you need to create a
function `timeTravel` that allows you to change the time according to your needs.
This function will give you the power to go backwards or forwards in time.

So, you will have a function that takes an object with the following `{date, hour, minute, second}`. This object will be responsible for changing the `hour`, `minute` and `second` of the given `date`.

### Example

```js
addWeek(new Date('0001-01-01')) // Output: Monday
addWeek(new Date('0001-01-02')) // Output: Tuesday
addWeek(new Date('0001-01-07')) // Output: Sunday
addWeek(new Date('0001-01-08')) // Output: secondMonday
addWeek(new Date('0001-01-09')) // Output: secondTuesday

// timeTravel({ date, hour, minute, second })

timeTravel({
  date: new Date('2020-05-29 23:25:22'),
  hour: 21,
  minute: 22,
  second: 22,
}).toString()

// Output: Fri May 29 2020 21:22:22 GMT+0100 (Western European Summer Time)
```
