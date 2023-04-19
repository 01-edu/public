## Quizz App

### Instructions

You need to develop a fully functional [**stateful Quizz app**](https://docs.flutter.dev/development/ui/interactive) in Flutter that allows users to select a **category** and answer **true or false questions** related to the selected category.

**Features:**

- The app needs to feature at least 5 different categories that the users can select from the main page, i.e, history, culture, math, geography, etc.
- The app needs to feature at least 10 questions to each `quizz`, so you need to make sure each category has at least 10 questions as well.
- Appropriate name and images must be added to all the selected categories. (follow the examples bellow)
- Upon selecting a category, the app needs to display the corresponding image to that category, along with the question and true/false buttons.
- After the user answers the question, the app needs to display whether the answer was correct or incorrect.
- At the end of the quiz, the user must be presented with a view of the score he got and a button to go back to the main page.

### Categories page

The `Categories` page needs to feature a **grid list** view of categories with the appropriate images and names.

The app will have models for `Question` and `Categories`, as shown below:

```jsx
//models/Question.dart
class Question {
  String question;
  bool answer;

  Question(this.question, this.answer);
}

//models/Category.dart

class Category {
  String name;
  String imageUrl;
  List<Question> questions;

  Category(this.name, this.imageUrl, this.questions);
}
```

<center>
<img src="./resources/quizApp.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

### Detailed View

After selecting a category, the app will navigate to a second route called `DetailedView`. Here, the corresponding `image` and `question` related to the selected category will be displayed, along with `true/false` buttons that allow the user to answer the question.

<center>
<img src="./resources/quizApp.02.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

When answering the question, the app needs to display whether the answer was correct or incorrect. For example you can change the color of the button or the background, so that the user knows which answer is right or wrong.

### Score View

After the user completes the quiz, they should be presented with a score view, which shows their overall score which is the number of correct and incorrect answers.

The score view should have a button that allows the user to return to the main page and start a new quiz.

Here's an example of what the score view could look like:

<center>
<img src="./resources/quizApp.03.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>

> Note: only standard Dart package, package:flutter are allowed.

> You are free to apply any styling of your choice to the app.

### Bonus

- Add specific images related to each question.
- Add a timer to each question, and if the user doesn't answer within the time limit, the app moves on to the next question automatically.
- Add difficulty Levels that allow users to choose between easy, medium, and hard difficulty levels for each category. The questions displayed should be relevant to the chosen difficulty level.
- Leaderboard, create a leaderboard that displays the top scores achieved by different users of the app.
- Shuffle the questions within each category so that the order of the questions is different each time the user plays the game.

### Notions

- [Widgets](https://docs.flutter.dev/development/ui/widgets)
- [Stateful/stateless widget](https://docs.flutter.dev/development/ui/interactive#stateful-stateless-widgets)
- [Routing](https://docs.flutter.dev/development/ui/navigation)
