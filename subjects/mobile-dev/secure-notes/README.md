# Secure Notes

### Introduction

The app where you can add your own notes. It should behave as normal notes app where you can add, modify, remove old notes, but the app should be secure. Being secure means that the notes should only be accessible through the app, and only after biometric authentication.

Packages: _sqflite_ package.

### Objective

- Work with textfields
- Work with keyboard
- Editing
- Deleting
- Reordering

### First Part

- Main Screen displayd list of notes with title, and description.
- Implement ReordableListView to reorder items in list
- Add ability to delete item by swiping
- Add "add" button, which opens screen with 3 text fields

<img src="./resources/secureNotes.01.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

<img src="./resources/secureNotes.02.png?raw=true" style = "width: 210px !important; height: 420px !important;"/>

### Second Part

- Screen to add new notes.
- If any of the fields is empty, show error.

### Third Part

- Edit Screen.
- Show text in textfields to edit.

### Fourth Part

Add sqflite package, so your notes will be saved, even when app reloads.
When app loads it should get notes from database.

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

Database class should also have CRUD methods like getAllNotes, deleteAllNotes, addNote, deleteNote, updateNote.

- getAllNotes
- deleteAllNotes
- addNote
- deleteNote
- updateNote

### Fifth Part

Make a biometric authentication. Make sure that notes are not accessible without firstly authenticating user.

- Note: you may use local_auth

### Sixth Part

[Localize](https://flutter.dev/docs/development/accessibility-and-localization/internationalization) your apllication.
