# Objects

> Allow you to create, manage and organize your pedagical and onboarding content.

## Definition

An Object is an highly customizable element, which can be use in many situations. We use it to compose cursuses and onboarding processes.
Objects can be associated together and then share a vertical or horizontal relationship, which allows to build complex structure of multiple objects.

It structure can be visualized in two parts. The first one is the definition of the object itself and attributes, called `attrs`. The second part is the definition of minor relationships, called `children` and attributes applied to them, called `childrenAttrs`.

This is the minimal structure of an object:

- name
- type (`organisation`, `campus`, `onboarding`, `cursus`, `quest`, `exercise`)
- status (`draft`, `online`, `offline`)
- attrs {}
- childrenAttrs {}
- children {}

## Browse Object:

To access your Objects, go to the admin dashboard and then click on the _manage object_ link within the "Object" card.

![go-to-objects](https://user-images.githubusercontent.com/15313830/56653756-46bdb780-6686-11e9-98ba-18e382987e9c.png)
