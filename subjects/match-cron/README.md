## Match Cron

### Instructions

Create a function called `matchCron` it takes a valid cron schedule string
and a valid date. \
It returns true if the date match the pattern

> You only have to implement numbers and `*`. \
> other complex patterns are not required.

Only valid pattern will be tested.

### Example

```js
matchCron('9 * * * *', new Date('2020-05-30 18:09:00')) // -> true
matchCron('9 * * * *', new Date('2020-05-30 19:09:00')) // -> true
matchCron('9 * * * *', new Date('2020-05-30 19:21:00')) // -> false
//         | | | | |
//         | | | | +- Day of the Week   (range: 1-7, 1 standing for Monday)
//         | | | +--- Month of the Year (range: 1-12)
//         | | +----- Day of the Month  (range: 1-31)
//         | +------- Hour              (range: 0-23)
//         +--------- Minute            (range: 0-59)
```

### Notions

- [crontab.guru](https://crontab.guru/)
- [devdocs.io/javascript/global_objects/date](https://devdocs.io/javascript/global_objects/date)
