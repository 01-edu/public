# Object Attributes System

> This document cover two notions, the edition of attributes and the impact it has relatively to the others Objects.

## Edition of Attributes

> There is no limit to how many attributes can be defined to an Object.

In the "Object Attributes" section of the "Object Edit" Page, the first row is a form to create and append a new attribute. It requires two elements, the name of the attribute and its type (`String`, `Number`, `Boolean`, `Array`, `Object`, `Function`, `Null`). Click 'Add' to create the attribute.

> Within a same Object, each attribute's name must be uniq.

Once created, the new attributes appears right bellow and the ability to associate a value to it is now available. Depending on the type of the attribute, the interface will vary.

- String value input is type String.
- Number value input is type Number.
- Booleans value input appears as a switch, true by default.
- Arrays and Objects content are hideable / showable via the "Show/Hide content" button on the right of the attribute. There is no limit on the depth of Object/Array, however, after a certain level, the interface will start to feel narrow.
- String value input is type String.
- Null will not display any input.
- Function will offer to select from all available functions, save on select.

Any attribute can be delete by clicking on the 'trash' icon on the right hand of it.

Here an example of how the section looks like.
![object-attributes](https://user-images.githubusercontent.com/15313830/56677487-88675600-66b8-11e9-9781-26dc0ee6301d.png)

## Attributes and RelationShips

When an attributes is set to an Object, other Objects, associated to this particular Object, will have access to it. Which means that, if an Object A is added as a child of an Object B, A will embed its attributes within the instance of B.

Object's attributes follow a hierarchy when associated to an other Object.
The Original Attributes of a child, the ones defined in the original Object are the weakest ones. A Children Attribute is applied to all the children and overload the Original Attributes. Finally, child attribute is the strongest one, it overlaod Original Attributes and Children Attributes.

When an object and its relationship are resolved, the three structures ("attrs", "childrenAttrs", "childAttrs") are flattened.

The following json shows how the next object would be resolved.

```json
{
  "childrenAttrs": {
    "xp": 800
  },
  "children": {
    "printalphabet": {
      "duration": 3600,
      "xp": 800,
      "isBonus": true
    }
  }
}
```

Children
![children](https://user-images.githubusercontent.com/15313830/56679319-b189e580-66bc-11e9-8f2a-3d51eb1486d4.png)

Child
![chilld-capture](https://user-images.githubusercontent.com/15313830/56679320-b189e580-66bc-11e9-90ab-c8f69f531876.png)
