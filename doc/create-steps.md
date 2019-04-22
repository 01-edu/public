# How to create onboarding steps

## Usage
This documentation explains how to create steps in your onboarding. 2 kind of steps can be generated in a modular way:
* forms
* documents to sign

These steps are used for the "sign up" phase, and for the administrative part of the onboarding.

## Description & examples
The sign up phase and the administration part of the onboarding are related to major objects in the admin. all the steps for each of these two sections are children of these two big objects, that you can create, update, or delete, and order as you wish. 
The objects contained in the sign up phase are object of type "sign up".
The objects contained in the administration part of the onboarding phase are object of type "onboarding".

### Relate a step to a major object
Every step is related to the major object you want to costum as a child. To create a child, you must:
1. Create a new object for your step in the admin (in Admin > "Add new object")
   * The title of your object will be the title displayed for the step you're creating. Use a intellegible title for your user.
   * This object must have the same type as your major object ("signup" for the sign up phase, "onboarding" for the administration part of the onboarding)
2. Add this new object as a child of your major object
  1. Edit the major object
  2. Go to "Children" > "Add a child"
  3. Set up the new child:
     * Enter its name in the input "Add a child name"
     * Select you child in the select input
     * Click on "ADD"

### Settings for a form step
In the child object you've created, 2 attributes must be filled:
1. subtype
2. form

To set up the child object you've created with these elements:
1. Edit you child object
2. Go to "Object attributes"
3. Add the attributes:
   * Add a new key "subtype" of type `String` with the value 'onb-adm-form-generator'
   * Add a new key "form" of type `Object` 
     * The form you're creating can have several sections. Each section is displayed with a title, and its inputs. 
       * The submission of the form will check the required inputs of all the sections you've created. 
     * To create a section, add a new key to your form object, of type `Object`, that contains :
       * A key "title" of type `String`. The value of this property will be the title displayed in the top of your form section. If you have only one section in your form step, and you don't need a section title, you don't have to set up this property.
       * An "inputs" key of type `Object`, which will contain all the inputs of the section. For each input you want to add, you must set up the input as this :
         * Add a new `Object` element in the "inputs" object. The key of this object will be used as the "name" attribute of your input. The values will be considered as the props of your input.

Set up of an input props:
* You must declare a "type" key of type `String`, wich define the type of your input : `tel`, `text`, `date`, `select`, `radio`, `switch`, `checkbox`, `textarea`. 
* Then you can fill the props you need for your input, according to its type : `placeholder`, `id`, `required`, `label`, `items`, `emptyItems`, `index`, etc.

Here are examples of the structure the form attribute could have, with two sections and an example of set up for every kind of input type:

```javascript
const formStepAttributes = {
  subtype: "onb-adm-form",
  form: {
    section1: {
      title: "My section title",
      inputs: {
        firstName: {
          index: 0,
          placeholder: 'First name',
          type: 'text',
          required: true,
        },
        tel: {
          index: 1,
          required: true,
          type: 'tel',
          label: 'Phone number',
          placeholder: '+333 33 33 33 33',
          pattern: '[+][3][0-9]{2}[0-9]{2}[0-9]{2}[0-9]{2}[0-9]{2}',
        },
        medicalInfo: {
          label: 'Medical informations',
          index: 6,
          placeholder: 'Medical Informations',
          type: 'textarea',
        },
        dateOfBirth: {
          index: 2,
          required: true,
          type: 'date',
          label: 'Date of birth',
          value: '2000-01-01',
        },
        gender: {
          index: 3,
          type: 'select',
          id: 'genders',
          required: true,
          emptyItem: { label: 'Select your gender' },
          items: [
            { label: 'Male', data: 'male' },
            { label: 'Female', data: 'female' },
          ],
        },
        environment: {
          index: 4,
          type: 'radio',
          required: true,
          label: 'Which environment do you live in ?',
          inlineBlock: true,
          items: [
            { label: 'City', data: 'city' },
            { label: 'Countryside', data: 'countryside' },
          ],
        },
        programmingAbilities: {
          index: 5,
          type: 'switch',
          label: 'I am new in programming',
          value: true,
        },
        generalContions: {
          index: 7,
          type: 'checkbox',
          label: 'I have read and I accept the general conditions',
          value: false,
        },
      },
    },
    section2: {
      title: "My second section title",
      inputs: {
        firstName: {
          index: 0,
          placeholder: 'First name',
          type: 'text',
          required: true,
        }
      }
    }
  }
}
```

This form step would look like this:
![form step example](https://user-images.githubusercontent.com/35296671/56503976-012aae80-650f-11e9-82c8-dd7d026b6eb1.png)

NB: 
* The index property is used to order your inputs. It won't be spread in the input.
* The type property is resquired. It will be used to determine the kind of input we must generate. It is spread in the input only if the input type attribute is required (type 'tel' or 'text' for example, but not for type 'select' - in this case, we will generate a select element)
* Each input type has its own attributes. You can check how to set up your inputs hey : 
  * [textInput documentation](https://alem.01-edu.org/design/Components/FormInputs/TextInput) - used for inputs type 'text', 'tel', and 'date'
  * [textArea documentation](https://alem.01-edu.org/design/Components/FormInputs/TextArea)
  * [select documentation](https://alem.01-edu.org/design/Components/FormControls/Select)
  * [radio button documentation](https://alem.01-edu.org/design/Components/FormControls/Radio)
  * [switch documentation](https://alem.01-edu.org/design/Components/FormControls/Switch)
  * [checkbox documentation](https://alem.01-edu.org/design/Components/FormControls/Checkbox)
* For switch and checkbox inputs type, you have to set the default value as a boolean property named "value".

### Settings for a document to sign step
In the child object you've created, 2 attributes must be filled:
1. subtype
2. text
3. buttonText (facultative)
4. checkbox (facultative)

To set up the child object you've created with these elements:
1. Edit you child object
2. Go to "Object attributes"
3. Add the attributes:
   * Add a new key "subtype" of type `String` with the value 'onb-adm-sign'
   * Add a new key "text" of type `String` with all the text of your document to sign as value
   * Add a new key "buttonText" of type `String` with the text you want to display in the submit button of your step
   * Add a new key "checkbox" of type `Object` if you want to force your user to click on a checkbox (ex: 'I have read and accepted the conditions') before validating his document. In the checkbox object, you should define:
     * A "label" key of type `String`, with the text you want to associate to the checkbox
     * A "required" key of type `Boolean`, set at true if you want to force the user to check it
     * A "name" key of type `String`
     * Then you can add all the properties you want to your checkbox.

Here is an example of the structure a 'document to sign' step could have:

```javascript
const signStepExample = {
  subtype: "onb-adm-sign",
  text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ornare non sem eu pretium. Integer porttitor risus eget nibh iaculis, ac lacinia orci dictum. Nunc ullamcorper consequat enim in posuere. Aliquam volutpat est odio, vel maximus arcu maximus sit amet. Donec ultricies faucibus magna id luctus. Duis et dapibus elit. In vestibulum ipsum erat, at commodo tortor convallis vel. Nunc ut ultrices nulla. Etiam lorem justo, consequat a consectetur a, porttitor non turpis. Mauris eu mollis nisl, id dignissim quam. Curabitur condimentum sollicitudin rutrum. Aenean blandit, arcu nec ullamcorper rhoncus, lectus sem lacinia lorem, venenatis dignissim velit mi et sapien. Nullam posuere augue ut magna ullamcorper dignissim. Ut rhoncus sapien vel nulla commodo finibus. Cras non leo vel urna finibus volutpat. Praesent et ex eget diam tincidunt suscipit. Phasellus bibendum neque vel placerat iaculis. Vestibulum bibendum ultrices ipsum, non sodales lectus. Cras eget orci eget elit blandit scelerisque at ut nulla. Integer ligula eros, eleifend quis sodales a, porttitor sit amet neque. Fusce mollis magna at lectus varius, quis suscipit mi cursus. Etiam id imperdiet metus, in malesuada quam. Aliquam facilisis nunc non sapien condimentum, quis iaculis nisl auctor. Nunc lorem sapien, interdum vel efficitur ac, dapibus a diam. Ut ante urna, sodales in bibendum vel, lacinia ut mauris. In vel placerat leo. In libero dui, tincidunt at sem id, faucibus sollicitudin elit.",
  buttonText: "Sign chart",
  checkbox: {
    name:"acceptChart",
    label:"I have read and accepted the Chart 01",
    required:true
  },
}
```

This 'document to sign' step would look like this:
![document to sign step example](https://user-images.githubusercontent.com/35296671/56504782-8f079900-6511-11e9-9a0e-bb638b6d7d03.png)

