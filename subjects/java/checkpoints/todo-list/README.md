## TodoList

### Instructions

Write a class `Task` that has two attributes: `description` and `status`, with the corresponding getters and setters.

Write a class called `TodoList` that represents a to-do list. The class should be able to hold a specified number of tasks, each with a description and a status. Include methods to add tasks, set a task's status, update task descriptions, and display all tasks in a neat format showing their descriptions and statuses.

### Expected Functions

```java
enum TaskStatus {
    NEW, IN_PROGRESS, COMPLETED
}

public class Task {
    private String description;
    private TaskStatus status;

    public void setDescription(String description);
    // Getters and Setters
    // ...
}

public class TodoList {
    private Task[] tasks;
    private int capacity;

    public TodoList(int capacity);

    public void addTask(String description);

    public void setStatus(int index, TaskStatus status);

    public void setDescription(int index, String newDescription);

    public void displayTasks();
}
```

### Usage

Here is a possible TaskManager.java to test your TodoList class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        TodoList myList = new TodoList(3); // List can hold up to 3 tasks
        myList.addTask("Go grocery shopping");
        myList.addTask("Pay electricity bill");
        myList.setStatus(0, TaskStatus.COMPLETED); // Mark the first task as completed
        myList.setDescription(1, "Pay all utility bills"); // Update the description of the second task
        myList.displayTasks(); // Display the list of tasks
    }
}
```

and its output:

```shell
$ javac TaskList.java TaskManager.java
$ java TaskManager
Tasks:
Go grocery shopping               | COMPLETED
Pay all utility bills             | NEW
```
