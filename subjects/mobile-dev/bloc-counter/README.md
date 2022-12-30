# Bloc Counter

### Introduction

Usually, when projects get bigger and more complex, managing all the widges, their states, and updating their children's states will become troublesome. To avoid possible complexity, we recommend using **Patterns**.
Patterns are built in a way that lets developers control the hierarchy of widgets much easier, and a popular one in the Flutter is - BLoC.

Implement simple counter app using Bloc Pattern.  
When you start writing your own application, you will need to structure your app first.

<img src="./resources/blocCounter.02.png?raw=true" width="200"/>

### Objective

- Observe state changes with `BlocObserver`.
- `BlocProvider`, Flutter widget which provides a bloc to its children.
- `BlocBuilder`, Flutter widget that handles building the widget in response to new states.

In this subject you will implement Bloc pattern, which is created by Google.
BLoC pattern uses Reactive Programming to handle the flow of data within an app.

Bloc consist of 2 concepts :

- `Streams`
- `Sinks`
  , which are provided by `StreamController`.

<pre>
<code>
💡 See this article <a href="https://medium.com/flutterpub/architecting-your-flutter-project-bd04e144a8f1"> Architect your Flutter project using BLoC pattern</a>, by Sagar Suri.
</code> 
</pre>

<img src="./resources/blocCounter.01.png?raw=true"/>

<pre>
<code>
💡 Documentation <a href="https://bloclibrary.dev/#/gettingstarted">https://bloclibrary.dev/#/gettingstarted</a>.
</code> 
</pre>

### Part 1

- Create new flutter app, so it will generate sample Counter App
- Add flutter_bloc as a dependecy of your app
- App structure should be similar to:

```bash
  —lib
      —bloc
          —counter_bloc.dart
          —counter_event.dart
      -widgets
          —yourwidgets.dart
      —..
  —main.dart
```

### Part 2

- add enum to bloc/couter_event.dart

```jsx
enum CounterEvent {
  increment
}
```

- go to bloc/counter_bloc.dart and create `CounterBloc` class which extends from `CounterBloc`

> 💡 note: you can generate bloc files using vscode extension

- implement override of `mapEventToState` function in `CounterBloc` class, so it will be switching between different events, and return value as an int. In our case we only have case with increment state

### Part 3 (UI and getting data)

- Inside MyApp class wrap your home page with BlocProvider class.

```jsx
return MaterialApp(
  home: BlocProvider<CounterBloc>(
    create: (context) => CounterBloc(),
    child: Home(),
  ),
);
```

- Inside Home class create instance of CounterBloc class
- Use BlocBuilder to state from CounterBloc
- Add "+" button to call CounterEvent.increment

### Bonus

- add button and event to handle decrementing counter
