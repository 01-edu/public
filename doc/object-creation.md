# Admin object's management - create an object

## Usage
> Elements of the app are managed threw objects in *Admin*.

> Objects of the Admin are first created and defined:
> * By their **title**,
> * By their **type**.

> Then it can be configured threw:
> * Attributes,
> * Children.

> This documentation explains how to create an object. 

### Create a new object in the admin
> (in *Admin* > *Add new object*)

<img width="664" alt="Capture d’écran 2019-04-22 à 15 57 37" src="https://user-images.githubusercontent.com/35296671/56507169-6505a500-6518-11e9-89bb-04c7fd9b41ca.png">
<img width="450" alt="Capture d’écran 2019-04-22 à 15 58 21" src="https://user-images.githubusercontent.com/35296671/56507180-6afb8600-6518-11e9-97a5-4dcff8f0a069.png">


* The **title** of your object will be the title displayed to your candidates. Use an intellegible title for your user.
  > NB: you can always edit it in the *Admin*

* The **type** depends on the nature of your object:
  * **Campus** is used to declare a school.
    * Examples: *Alem*, *Madeira*, etc.
    * Campus can contains cursus: *Alem* contains for example *01-classical* and *Piscine Go*.
  * **Cursus** is used to declare a course. 
    * Examples: *01-classical*, *Piscine Go*, etc.
    * Cursuses can contains cursuses: the main cursus *01-classical*, for example, contains cursuses like *Piscine Go*, but also all the branches that the student have access to, as *Web*, *Security*, *Algorythm*, *Design*, etc.
    * Cursuses can contains quests: *Piscine Go* of *01-classical* contains quests like *Quest 1* or *Quest 2*.
  * **Quest** is used to declare a project.
    * Examples: *Quest 1*, *Quest 2*, etc.
    * Quest contains exercises: *Quest 1* of *Piscine Go* contains exercises like *printalphabet* or *printcomb*. 
  * Exercise is used to declare exercises
    * Examples: *printalphabet*, *printcomb*, *atoi*, etc.
    * Exercises doesn't contains any children.
  * Signup is used to declare steps of the registration.
    * Examples: *Using our services*, *Tell us more about you*, etc.
    * One major object *Sign up* contains all the sign up's modular steps : *Using our services*, *Tell us more about you*, etc.
  * Onbaording is used to declare steps of the onbaording.
    * Examples: *Toad*, *Administration*, *Additional Informations*, *Chart 01*, etc.
    * Three main objects define the major steps of the onboarding : *Toad*, *Administration*, *Piscine*.
    * *Administration* contains modular steps: *Additional Informations*, *Chart 01*, etc.

> The child object is then available in the *Admin*. It can be found in the section of its type or thanks to the search bar of the cursus object's page.

> More information is available:
> * for setting attributes of an object: (soon available)
> * for setting children of an object: [Child object creation](https://github.com/01-edu/public/blob/master/doc/child-object-creation.md) 
> * for creation of modular steps in Sign up and onboarding's Administration object: [Modular step management](https://github.com/01-edu/public/blob/master/doc/modular-steps-management.md)
