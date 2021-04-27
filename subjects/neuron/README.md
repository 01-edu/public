## Neuron

### Instructions

Create a function called `neuron` that allows your AI/bot to learn how to data shape a given
dataset into an object so that it can better navigate the data.

### Example

```js
neuron([
  'Questions: what is ounces? - Response: Ounce, unit of weight in the avoirdupois system',
  'Questions: what is ounces? - Response: equal to 1/16 pound (437 1/2 grains)',
  'Questions: what is Mud dauber - Response: Mud dauber is a name commonly applied to a number of wasps',
  'Orders: shutdown! - Response: Yes Sr!',
  'Orders: Quote something! - Response: Pursue what catches your heart, not what catches your eyes.'
])

// output
{
  questions: {
    what_is_ounces: { question: 'what is ounces?', responses: [
        'Ounce, unit of weight in the avoirdupois system',
        'equal to 1/16 pound (437 1/2 grains)'
    ] },
    what_is_mud_dauber: { question: 'what is Mud dauber', responses: [
        'Mud dauber is a name commonly applied to a number of wasps'
    ] }
  },
  orders: {
    shutdown: { order: 'shutdown!', responses: ['Yes Sr!'] },
    quote_something: { order: 'Quote something!', responses: [
        'Pursue what catches your heart, not what catches your eyes.'
    ] }
  }
}
```
