# Sign up & onboarding's Administration section - Modular steps management

## Usage
> After their first authentication in the app, every candidate has to do his **sign up** and his **onboarding**. The steps that compose the **sign up** and the **administration** section of the onboarding are either:
>   * Forms (identification, medical information, etc.)
>   * Documents to sign (general conditions, charts, regulations, etc.)
>
> All the sections are modular: you can add, update, delete and order them as you wish. 
>
> This documentation explains how to manage these steps. 

## Create your step child object
### Create a new object for your step in the admin
> Information is available for object's creation: [Object creation](https://github.com/01-edu/public/blob/master/doc/object-creation.md)

* This object must have the same type as its future parent object (*signup* or *onboarding*).

> Your step is then available in the *Admin*. You can find it in the section of its type (*SignUp* or *Onboarding*) or thanks to the search bar of the cursus object's page.
 

### Add this new object as a child of your parent's object

* Edit the parent object: *Sign up* or *Administration*

> Information is available for object's creation: [Child object creation](https://github.com/01-edu/public/blob/master/doc/child-object-creation.md)

## Settings for a `form` step
> In the step object you have created, 2 attributes must be filled:
> 1. Subtype
> 2. Form

### Description

#### Edit the step object you have created :
> in *Object attributes*

<img width="1073" alt="Capture d’écran 2019-04-22 à 15 59 33" src="https://user-images.githubusercontent.com/35296671/56507445-3936ef00-6519-11e9-90c8-d85056e9330b.png">

* Add a new key **subtype** of type `String` with the exact value 'onb-adm-form-generator'
* Add a new key **form** of type `Object` 
  * The form you are creating can have several sections. Each section is displayed with a title, and its inputs. 
    > NB: The submission of the form will check the required inputs of all the sections you have created for the form. 
  * To create a section, add a new key to your form object, of type `Object`, that contains:
    * A **title** key of type `String`. The value of this property will be the title displayed in the top of your form section. 
      > If you have only one section in your form step, and you don't need a section title, you do not have to set up this property.
    * An **inputs** key of type `Object`, which will contain all the inputs of the section. For each wanted input, add a new `Object` element in the "inputs" object.
      > The key of this object will be used as the "name" attribute of your input.  
      > The values will be considered as the properties of your input.

#### Defining an input:
* You must declare a **type** key of type `String`, which define the type of your input : `tel`, `text`, `date`, `select`, `radio`, `switch`, `checkbox`, `textarea` (and soon a special type `countries`). 
* Then you can fill the props you need for your input, according to its type: `placeholder`, `id`, `required`, `label`, `items`, `emptyItems`, `index`, etc...

#### Important indication: 
* The **index** property is used to order your inputs. It will not be passed onto the input. Be mindful not to set the same index twice.
* The **type** property is required. It will be used to determine the kind of input we must generate. It is passed onto the input only if the input type attribute is required (type 'tel' or 'text' for example, but not for type 'select' - in this case, we will generate a select element)
  * At the moment, we are creating a special type 'countries' (a `Select` displaying all the countries). Documentation will be updated as soon as it is available.
* `onChange` prop are ignored as the event is handled by the app.
* For `switch` and `checkbox` input types, the default value has to be set as a boolean property named **value**.
* More information for each inputs is available in the design documentation: 
  * [textInput documentation](https://alem.01-edu.org/design/Components/FormInputs/TextInput) - used for inputs type 'text', 'tel', and 'date'
  * [textArea documentation](https://alem.01-edu.org/design/Components/FormInputs/TextArea)
  * [select documentation](https://alem.01-edu.org/design/Components/FormControls/Select)
  * [radio button documentation](https://alem.01-edu.org/design/Components/FormControls/Radio)
  * [switch documentation](https://alem.01-edu.org/design/Components/FormControls/Switch)
  * [checkbox documentation](https://alem.01-edu.org/design/Components/FormControls/Checkbox)

### Examples

Here is an example of the form step's attributes. It presents a form with two sections, and an example of each kind of input type. 
> NB : this example object will soon be provided in the admin.

```json
{
  "subtype": "onb-adm-form",
  "form": {
    "section1": {
      "title": "My section title",
      "inputs": {
        "firstName": {
          "index": 0,
          "placeholder": "First name",
          "type": "text",
          "required": true
        },
        "tel": {
          "index": 1,
          "required": true,
          "type": "tel",
          "label": "Phone number",
          "placeholder": "+333 33 33 33 33",
          "pattern": "[+][3][0-9]{2}[0-9]{2}[0-9]{2}[0-9]{2}[0-9]{2}"
        },
        "medicalInfo": {
          "label": "Medical informations",
          "index": 6,
          "placeholder": "Medical Informations",
          "type": "textarea"
        },
        "dateOfBirth": {
          "index": 2,
          "required": true,
          "type": "date",
          "label": "Date of birth",
          "value": "2000-01-01"
        },
        "gender": {
          "index": 3,
          "type": "select",
          "id": "genders",
          "required": true,
          "emptyItem": { "label": "Select your gender" },
          "items": [
            { "label": "Male", "data": "male" },
            { "label": "Female", "data": "female" }
          ]
        },
        "environment": {
          "index": 4,
          "type": "radio",
          "required": true,
          "label": "Which environment do you live in ?",
          "inlineBlock": true,
          "items": [
            { "label": "City", "data": "city" },
            { "label": "Countryside", "data": "countryside" }
          ]
        },
        "programmingAbilities": {
          "index": 5,
          "type": "switch",
          "label": "I am new in programming",
          "value": true
        },
        "generalConditions": {
          "index": 7,
          "type": "checkbox",
          "label": "I have read and I accept the general conditions",
          "value": false
        }
      }
    },
    "section2": {
      "title": "My second section title",
      "inputs": {
        "firstName": {
          "index": 0,
          "placeholder": "First name",
          "type": "text",
          "required": true
        }
      }
    }
  }
}
```

This 'form' step would look like this:

![form step example](https://user-images.githubusercontent.com/35296671/56503976-012aae80-650f-11e9-82c8-dd7d026b6eb1.png)

## Settings for a `document to sign` step
The newly created child can be customized with these attributes :

| name       | fullfillment |
| ---------- | ---------    |
| subtype    | **required** |
| text       | **required** |
| buttonText | optionnal    |
| checkbox   | optionnal    |

### Description

#### To set up the child object you have created with these elements:

1. Edit you step object
2. Go to *Object attributes*
3. Add the following attributes:
   * Add a new key **subtype** of type `String` with the exact value 'onb-adm-sign'
   * Add a new key **text** of type `String` with the text of your document to sign as value
   * Add a new key **buttonText** of type `String` with the text that you want to display in the submit button of your step. If you do not fill this property, the default value is 'Sign'.
   * Add a new key **checkbox** of type `Object`, if you want to force your user to click on a checkbox before validating his document (ex: 'I have read and accepted the conditions'). In the checkbox object, you should define:
     * A **label** key of type `String`, with the text you want to associate to the checkbox
     * A **required** key of type `Boolean`, set at true if you want to force the user to check it
     * A **name** key of type `String`
     * Then you can add all other properties you want to your checkbox.

### Examples

Here is an example of the structure a 'document to sign' step could have:

```json
{
  "subtype": "onb-adm-sign",
  "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent ornare non sem eu pretium. Integer porttitor risus eget nibh iaculis, ac lacinia orci dictum. Nunc ullamcorper consequat enim in posuere. Aliquam volutpat est odio, vel maximus arcu maximus sit amet. Donec ultricies faucibus magna id luctus. Duis et dapibus elit. In vestibulum ipsum erat, at commodo tortor convallis vel. Nunc ut ultrices nulla. Etiam lorem justo, consequat a consectetur a, porttitor non turpis. Mauris eu mollis nisl, id dignissim quam. Curabitur condimentum sollicitudin rutrum. Aenean blandit, arcu nec ullamcorper rhoncus, lectus sem lacinia lorem, venenatis dignissim velit mi et sapien. Nullam posuere augue ut magna ullamcorper dignissim. Ut rhoncus sapien vel nulla commodo finibus. Cras non leo vel urna finibus volutpat. Praesent et ex eget diam tincidunt suscipit. Phasellus bibendum neque vel placerat iaculis. Vestibulum bibendum ultrices ipsum, non sodales lectus. Cras eget orci eget elit blandit scelerisque at ut nulla. Integer ligula eros, eleifend quis sodales a, porttitor sit amet neque. Fusce mollis magna at lectus varius, quis suscipit mi cursus. Etiam id imperdiet metus, in malesuada quam. Aliquam facilisis nunc non sapien condimentum, quis iaculis nisl auctor. Nunc lorem sapien, interdum vel efficitur ac, dapibus a diam. Ut ante urna, sodales in bibendum vel, lacinia ut mauris. In vel placerat leo. In libero dui, tincidunt at sem id, faucibus sollicitudin elit.",
  "buttonText": "Sign chart",
  "checkbox": {
    "name":"acceptChart",
    "label":"I have read and accepted the Chart 01",
    "required":true
  }
}
```

This 'document to sign' step would look like this:

![document to sign step example](https://user-images.githubusercontent.com/35296671/56504782-8f079900-6511-11e9-9a0e-bb638b6d7d03.png)

