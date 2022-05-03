## Dr Strange

### Instructions

You have been given a mission to create a new sense of time. Normally a week has 7 days right? Well, that is about to change.

Days will instead have **14** days.

Let me explain; this new week will have 14 days, from `Monday` to `Sunday`, then `secondMonday` to `secondSunday`.

Your purpose is to create a new function named `addWeek`, that takes a `Date` as an argument. Your function should return the weekday as a string, according to our new 14-day week format.
The **epoch** of our new 14-day week is `0001-01-01`.

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

Now imagine you have an appointment with your doctor, and you have to wait some hours, and you do not want to wait. So you decide that you need to create a new function named `timeTravel`, that allows you to change the time according to your needs.

This function will give you the power to go backwards or forwards in time.

Your function will take an `object` as an argument. You can see the `timeTravel` example below to see the structure of your function's argument.

Your objective is to 

and a `Date` as arguments.an argument. T as an argument. This  `{date, hour, minute, second}`. This object will be responsible for changing the `hour`, `minute` and `second` of the given `date`.

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
