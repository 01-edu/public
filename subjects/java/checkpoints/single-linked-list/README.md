## Single Linked List

### Instructions

Create a single linked list data structure that implements the following methods:

- `at(int index)`: to access an element by its index, if the index is out of bound return -1.
- `add(int value)`: to add an element at the end of the list.
- `remove(int index)`: to remove an element by its index.
- `size()`: to get the size of the list.

Define these methods in an interface called `LinkedList`, and implement this interface in a class called `SingleLinkedList`. Additionally, add a private method `next(Node node)` in this class that will be used to traversing the list in the other methods. This method should print the message "Go to next node\n" each time it is called.

### Explanation

A linked list is a linear data structure where each element is a separate object called a node. Each node contains two fields:

- `value`: stores the data.
- `next`: stores a reference to the next node in the list.

The first node is called the head of the list and the last node is always null. The list allows for efficient insertion and deletion of elements. However, accessing an element by its index requires traversing the list from the head to the desired position.

### Given Interface

```java
public interface LinkedList {
    int at(int index);
    void add(int value);
    void remove(int index);
    int size();
}
```

### Expected Class

```java
public class SingleLinkedList implements LinkedList {
    private Node head;
    private int size;

    private class Node {
        int value;
        Node next;

        Node(int value) {
            this.value = value;
            this.next = null;
        }
    }

    @Override
    public int at(int index) {
        // Implementation for accessing an element by its index
    }

    @Override
    public void add(int value) {
        // Implementation for adding an element at the end of the list
    }

    @Override
    public void remove(int index) {
        // Implementation for removing an element by its index
    }

    @Override
    public int size() {
        // Implementation for accessing an element by its index
    }

    private Node next(Node node) {
        // Print the message "Go to next node"
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        LinkedList list = new SingleLinkedList();

        // Add elements to the list
        list.add(1);
        list.add(2);
        list.add(3);

        // Access elements by index
        System.out.println("Element at index 0: " + list.at(0)); // Expected output: 1
        System.out.println("Element at index 1: " + list.at(1)); // Expected output: 2
        System.out.println("Element at index 2: " + list.at(2)); // Expected output: 3

        // Remove an element by index
        list.remove(1);
        System.out.println("Element at index 1 after removal: " + list.at(1)); // Expected output: 3
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Go to next node
Element at index 0: 1
Go to next node
Element at index 1: 2
Go to next node
Go to next node
Element at index 2: 3
Go to next node
Go to next node
Element at index 1 after removal: 3
$
```
