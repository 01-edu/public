## Bloc Counter

As projects get bigger and more complex, managing all the widgets, their states, and updating their children's states can become troublesome. To avoid potential complexity, it is recommended to use **Patterns**. Patterns are built to help developers control the hierarchy of widgets more easily, and one popular pattern in Flutter is `BLoC`.

In this project, you will implement a simple `counter app`,like the example bellow, using the `BLoC` pattern. When you start writing your own application, structuring it correctly is important.

<img src="./resources/blocCounter.02.png?raw=true" width="200"/>

### Instructions

- Observe state changes with `BlocObserver`.
- Use `BlocProvider`, a Flutter widget which provides a bloc to its children.
- Use `BlocBuilder`, a Flutter widget that handles building the widget in response to new states.

`BLoC` pattern uses reactive programming to handle the flow of data within an app.

A Bloc consists of 2 concepts, `Streams` and `Sinks`, which are provided by `StreamController`.

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

### Project Setup

- Create a new Flutter app, so it generates a sample Counter App.
- Add `flutter_bloc` as a dependency to your app.

Your app structure should be similar to:

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

### Bloc Implementation

- Add an enum to `bloc/couter_event.dart`.

```jsx
enum CounterEvent {
  increment
}
```

- Go to `bloc/counter_bloc.dart` and create a `CounterBloc` class that extends `CounterBloc`.

> Note: you can generate bloc files using vscode extension.

- Implement an override of the `mapEventToState` function in the `CounterBloc` class, so it switches between different events and returns the value as an `int`. In our case, we only have a case with an increment state.

### UI and getting data

- Inside the `MyApp` class, wrap your home page with a `BlocProvider` class.

```jsx
return MaterialApp(
  home: BlocProvider<CounterBloc>(
    create: (context) => CounterBloc(),
    child: Home(),
  ),
);
```

- Inside the `Home` class, create an instance of the `CounterBloc` class.
- Use `BlocBuilder` to state from `CounterBloc`.
- Add a "+" button to call `CounterEvent.increment`.

### Bonus

- Add button and event to handle decrementing the counter.
