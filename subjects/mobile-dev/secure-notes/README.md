## Secure Notes

Secure Notes is an app that allows you to create, modify, and delete notes, with an emphasis on security. Only the app's user should be able to access the notes, and only after biometric authentication. The app needs to use the `sqflite` package to store notes in a database.

### Instructions

The Secure Notes app has several objectives, including:

- Working with text fields and the keyboard
- Allowing users to edit, delete, and reorder notes
- Implementing biometric authentication
- Localizing the application for international use

### Main Screen

The main screen of the app displays a list of notes, including the note's title and description. Users can reorder notes using the `ReorderableListView` widget, `delete` notes by swiping, and `add` notes using the "add" button which opens a screen with at least 2 text fields.

<img src="./resources/secureNotes.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

<img src="./resources/secureNotes.02.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

### Add Note

This screen allows users to create new notes by entering information into the text fields. If any of the fields are empty, an error message is displayed.

### Edit Note

This screen allows users to modify existing notes by displaying the note's text in the corresponding text fields.

### Data Management

The Secure Notes app must use the `sqflite` package to store notes in a database. The Database class should also have `CRUD` methods, including `getAllNotes`, `deleteAllNotes`, `addNote`, `deleteNote`, and `updateNote`. Notes are saved in the database and are accessible even if the app is closed and reopened.

```jsx
class Database {
Database _db;

    Future create() async {
        Directory path = await getApplicationDocumentsDirectory();
        String dbPath = join(path.path, "database.db");

        _db = await openDatabase(dbPath, version: 1);
    }
}
```

### Authentication

The app must contain a biometric authentication before users can access their notes. You may use the `local_auth` package to authenticate the user.

### Localization

The app can be localized for use in multiple countries and languages. Localizing the app allows it to reach a wider audience and make it easier for non-native speakers to use.

To localize the app, follow the instructions provided by [Flutter](https://docs.flutter.dev/accessibility-and-localization/internationalization).
