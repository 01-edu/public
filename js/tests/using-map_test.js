export const tests = []
const t = (f) => tests.push(f)

// citiesOnly
t(({ eq, ctx }) =>
  eq(citiesOnly(ctx.states), [
    'Los Angeles',
    'San Francisco',
    'Miami',
    'New York City',
    'Juneau',
    'Boston',
    'Jackson',
    'Utqiagvik',
    'Albuquerque',
  ])
)

t(({ eq, ctx }) =>
  // check that the code did use map properly
  eq(ctx.mapCalls[0], ctx.states)
)

// upperCasingStates
t(({ eq, ctx }) =>
  eq(upperCasingStates(ctx.cities), [
    'Alabama',
    'New Jersey',
    'Alaska',
    'New York',
    'California',
    'New Hampshire',
    'Ohio',
    'Texas',
    'West Virginia',
  ])
)

t(({ eq, ctx }) =>
  // check that the code did use map properly
  eq(ctx.mapCalls.includes(ctx.cities), true)
)

// fahrenheitToCelsius
t(({ eq, ctx }) =>
  eq(fahrenheitToCelsius(ctx.temps), [
    '30°C',
    '37°C',
    '5°C',
    '12°C',
    '-13°C',
    '21°C',
    '-19°C',
  ])
)

t(({ eq, ctx }) =>
  // check that the code did use map properly
  eq(ctx.mapCalls.includes(ctx.temps), true)
)

// trimTemp
t(({ eq, ctx }) =>
  eq(trimTemp(ctx.states), [
    {
      city: 'Los Angeles',
      state: 'california',
      region: 'West',
      temperature: '101°F',
    },
    {
      city: 'San Francisco',
      state: 'california',
      region: 'West',
      temperature: '84°F',
    },
    { city: 'Miami', state: 'Florida', region: 'South', temperature: '112°F' },
    {
      city: 'New York City',
      state: 'new york',
      region: 'North East',
      temperature: '0°F',
    },
    { city: 'Juneau', state: 'Alaska', region: 'West', temperature: '21°F' },
    {
      city: 'Boston',
      state: 'massachussetts',
      region: 'North East',
      temperature: '45°F',
    },
    {
      city: 'Jackson',
      state: 'mississippi',
      region: 'South',
      temperature: '70°F',
    },
    { city: 'Utqiagvik', state: 'Alaska', region: 'West', temperature: '-1°F' },
    {
      city: 'Albuquerque',
      state: 'new mexico',
      region: 'West',
      temperature: '95°F',
    },
  ])
)

t(({ eq, ctx }) =>
  // check that the code did use map properly
  eq(ctx.mapCalls.includes(ctx.states), true)
)

// tempForecasts
t(({ eq, ctx }) =>
  eq(tempForecasts(ctx.states), [
    '38°Celsius in Los Angeles, California',
    '28°Celsius in San Francisco, California',
    '44°Celsius in Miami, Florida',
    '-18°Celsius in New York City, New York',
    '-7°Celsius in Juneau, Alaska',
    '7°Celsius in Boston, Massachussetts',
    '21°Celsius in Jackson, Mississippi',
    '-19°Celsius in Utqiagvik, Alaska',
    '35°Celsius in Albuquerque, New Mexico',
  ])
)

export const setup = () => {
  const mapCalls = []
  const _map = Array.prototype.map
  Array.prototype.map = function () {
    mapCalls.push(this)
    return _map.apply(this, arguments)
  }

  const states = [
    {
      city: 'Los Angeles',
      temperature: '101 °F',
      state: 'california',
      region: 'West',
    },
    {
      city: 'San Francisco',
      temperature: '84 °F',
      state: 'california',
      region: 'West',
    },
    {
      city: 'Miami',
      temperature: ' 112 °F',
      state: 'Florida',
      region: 'South',
    },
    {
      city: 'New York City',
      temperature: ' 0 °F',
      state: 'new york',
      region: 'North East',
    },
    { city: 'Juneau', temperature: ' 21° F', state: 'Alaska', region: 'West' },
    {
      city: 'Boston',
      temperature: '45 °F',
      state: 'massachussetts',
      region: 'North East',
    },
    {
      city: 'Jackson',
      temperature: ' 70°F  ',
      state: 'mississippi',
      region: 'South',
    },
    {
      city: 'Utqiagvik',
      temperature: ' -1 °F',
      state: 'Alaska',
      region: 'West',
    },
    {
      city: 'Albuquerque',
      temperature: ' 95 °F',
      state: 'new mexico',
      region: 'West',
    },
  ]

  const cities = [
    'alabama',
    'new jersey',
    'alaska',
    'new york',
    'california',
    'new hampshire',
    'ohio',
    'texas',
    'west virginia',
  ]

  Object.getPrototypeOf([]).proto = ' [avoid for..in] '
  const temps = ['86°F', '100°F', '41°F', '55°F', '10°F', '70°F', '-2°F']

  Object.freeze(states)
  Object.freeze(cities)
  Object.freeze(temps)

  return { mapCalls, states, cities, temps }
}

Object.freeze(tests)
