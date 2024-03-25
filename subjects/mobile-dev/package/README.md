# Package

Package can be used to organize and share a set of functions in Dart. It is simply a sharable library or modules.
Package is similar to Dart Application except that Dart Package does not have application entry point - `main`.
A minimal package consists of the following:

- `pubspec.yaml`: a metadata file that declares the package name, version, author, etc.
- `lib/`: a directory contains the public code of the package, at least one dart file.

### Instructions:

Create a Flutter package for your [Secure Notes app](../secure-notes/README.md).
You should write your own package which will work with [`sqflite`](https://pub.dev/packages/sqflite) and have CRUD functionality.
Your package should consist of `Database` class and `Note` class which will allow easy access to SQLite database.

In the end you should be able to import it like:

```dart
import 'package:note/note.dart';

```

### Objective:

- Dealing with existing, large apps
- Reusing packages across multiple apps
- Local vs remote (git) packages

### Database.dart

```dart
class Database {
Database _db;

    Future create() async {
        Directory path = await getApplicationDocumentsDirectory();
        String dbPath = join(path.path, "database.db");

        _db = await openDatabase(dbPath, version: 1);
    }
}

```

Example of Database class, where you should create table `Note` with 4 parameters:

- `id`
- `title`
- `body`
- `date`

Database class should also have CRUD methods like `getAllNotes`, `deleteAllNotes`, `addNote`, `deleteNote`, `updateNote`.

- `getAllNotes()`
- `deleteAllNotes()`
- `addNote(note: Note)`
- `deleteNote(note: Note)`
- `updateNote(oldNote: Note, newNote: Note)`

### Note.dart

Model class for Note object.

```dart
class Note {
  int id;
  String title;
  String body;
  String date;

  Note({
    this.id,
    this.title,
    this.body,
    this.date;
  });

```

### Notions

[Developing packages with Dart](https://flutter.dev/docs/development/packages-and-plugins/developing-packages)
