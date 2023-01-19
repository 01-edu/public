## Student

### Instructions

Create a `class` named `Student`, that extends the `Person` `class` that you created earlier.

Attributes:

- `batch`: `int`
- `level`: `int`
- `secretKey`: private `string`. Defaults to "01".

Constructor:
  - `name`: `string`
  - `cityOfOrigin`: `string`
  - `age`: `int`
  - `height`: `int`
  - `batch`: `int`
  - `level`: `int`

### Inheritance

A class can inherit the fields and methods from another class. By doing so, every public field and method will be visible to the child class. It can be useful in cases where one class is a superset of the other, and you feel that it is not useful to copy all of the code from one class to another. If necessary, the behavior of certain methods can be changed in the child class by using the `@override` keyword.

Let's say we have a class `TV` with a lot of code. Then the class `SmartTV` can have all of the features of `TV` plus its own smart features.

```dart
class TV {
  bool _hasPower = false;

  int _channel = 1;
  int _minChannel = 1;
  int _maxChannel = 99;

  bool _isMute = false;
  int _volume = 50;
  int _minVolume = 0;
  int _maxVolume = 100;

  void connectPower() {
    this._hasPower = true;
  }

  void disconnectPower() {
    this._hasPower = false;
    this._channel = 1;
    this._volume = 50;
    this._isMute = false;
  }

  void toggleMute() {
    this._isMute = !this._isMute;
  }

  bool get mute => this._isMute;

  int get volume => (this._isMute) ? 0 : this._volume;

  int get channel => this._channel;

  bool get hasPower => this._hasPower;

  set channel(int channelNum) {
    this._channel =
        (channelNum >= this._maxChannel || channelNum <= this._minChannel)
            ? this._channel
            : channelNum;
  }

  void increaseVolume() {
    this._isMute = false;
    if (this._volume <= this._maxVolume) {
      this._volume++;
    }
  }

  void decrementVolume() {
    this._isMute = false;

    this._volume =
        (this._volume <= this._minVolume) ? this._minVolume : this._volume - 1;
  }

  void incrementVolume() {
    this._isMute = false;

    this._volume =
        (this._volume >= this._maxVolume) ? this._maxVolume : this._volume + 1;
  }

  void decrementChannel() {
    this._isMute = false;

    this._channel = (this._channel <= this._minChannel)
        ? this._minChannel
        : this._channel - 1;
  }

  void incrementChannel() {
    this._isMute = false;

    this._channel = (this._channel >= this._maxChannel)
        ? this._maxChannel
        : this._channel + 1;
  }

  TV({required bool hasPower}) {
    this._hasPower = hasPower;
  }
}
```

`SmartTV` can be created with **much** less code, and still have all of the base functionality of `TV`.

Notice that `@override` changes the behavior of the `disconnectPower` method.

> If you override a method, you can still use the method from the parent class by writing `super.` before invoking the method: `super.disconnectPower()`.

```dart
class SmartTV extends TV {
  String? _app;

  void startApp(String app) => this._app = app;
  void quitApp() => this._app = null;

  bool get tvMode => this._app == null;
  String? get app => this._app;

  @override
  void disconnectPower() {
    super.disconnectPower();
    this._app = null;
  }

  SmartTV({required bool hasPower}) : super(hasPower: hasPower);
}
```

```dart
void main() {
  var tv = new SmartTV(hasPower: true);

  print(tv.volume); // 50
  tv.toggleMute();
  print(tv.volume); // 0
  tv.increaseVolume();
  print(tv.volume); // 51

  tv.startApp('Chrome');
  print(tv.tvMode); // true
  tv.quitApp();
  print(tv.tvMode); // false
}

```

> Don't forget to use the constructor of the parent class.
