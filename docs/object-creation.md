# Admin object's management - create an object

## Usage

> Elements of the app are managed through objects in _Admin_.

> Objects of the Admin are first created and defined:
>
> - By their **title**,
> - By their **type**.

> Then it can be configured through:
>
> - Attributes,
> - Children.

> This documentation explains how to create an object.

### Create a new object in the admin

> (in _Admin_ > _Add new object_)

<img width="664" alt="Capture d’écran 2019-04-22 à 15 57 37" src="https://user-images.githubusercontent.com/35296671/56507169-6505a500-6518-11e9-89bb-04c7fd9b41ca.png">
<img width="450" alt="Capture d’écran 2019-04-22 à 15 58 21" src="https://user-images.githubusercontent.com/35296671/56507180-6afb8600-6518-11e9-97a5-4dcff8f0a069.png">

- The **title** of your object will be the title displayed to your candidates. Use an intellegible title for your user.

  > NB: you can always edit it in the _Admin_

- The **type** depends on the nature of your object:
  - **Campus** is used to declare a school.
    - Examples: _Alem_, _Madeira_, etc.
    - Campus can contains cursus: _Alem_ contains for example _01-classical_ and _Piscine Go_.
  - **Cursus** is used to declare a course.
    - Examples: _01-classical_, _Piscine Go_, etc.
    - Cursuses can contains cursuses: the main cursus _01-classical_, for example, contains cursuses like _Piscine Go_, but also all the branches that the student have access to, as _Web_, _Security_, _Algorythm_, _Design_, etc.
    - Cursuses can contains quests: _Piscine Go_ of _01-classical_ contains quests like _Quest 1_ or _Quest 2_.
  - **Quest** is used to declare a project.
    - Examples: _Quest 1_, _Quest 2_, etc.
    - Quest contains exercises: _Quest 1_ of _Piscine Go_ contains exercises like _printalphabet_ or _printcomb_.
  - Exercise is used to declare exercises
    - Examples: _printalphabet_, _printcomb_, _atoi_, etc.
    - Exercises doesn't contains any children.
  - Signup is used to declare steps of the registration.
    - Examples: _Using our services_, _Tell us more about you_, etc.
    - One major object _Sign up_ contains all the sign up's modular steps : _Using our services_, _Tell us more about you_, etc.
  - Onbaording is used to declare steps of the onbaording.
    - Examples: _Toad_, _Administration_, _Additional Informations_, _Chart 01_, etc.
    - Three main objects define the major steps of the onboarding : _Toad_, _Administration_, _Piscine_.
    - _Administration_ contains modular steps: _Additional Informations_, _Chart 01_, etc.

> The child object is then available in the _Admin_. It can be found in the section of its type or thanks to the search bar of the cursus object's page.

> More information is available:
>
> - for setting attributes of an object: (soon available)
> - for setting children of an object: [Child object creation](object-child-creation.md)
> - for creation of modular steps in Sign up and onboarding's Administration object: [Modular step management](modular-steps-management.md)
