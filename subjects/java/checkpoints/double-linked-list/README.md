## Double Linked List

### Instructions

Create a double linked list data structure that implements the following methods:

- `at(int index)`: to access an element by its index. If the index is out of bound return -1.
- `add(int value)`: to add an element at the end of the list.
- `remove(int index)`: to remove an element by its index. If the index is out of bound nothing happen.
- `size()`: to get the size of the list.

Define these methods in an interface called `LinkedList`, and implement this interface in a class that you will create. Additionally, add private methods `next(Node node)` and `prev(Node node)` in your class that will be used in the other methods. These methods should print the messages "Go to next node\n" and "Go to previous node\n" each time they are called, respectively.

### Explanation

A double linked list is a linear data structure where each element is a separate object called a node. Each node contains three fields:

- `value`: stores the data.
- `next`: stores a reference to the next node in the list.
- `prev`: stores a reference to the previous node in the list.

The first node is called the head of the list and the last node always has next pointing to null. The list allows for efficient insertion and deletion of elements. However, accessing an element by its index requires traversing the list from the head to the desired position.

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
public class DoubleLinkedList implements LinkedList {
    private Node head;
    private Node tail;

    private class Node {
        int value;
        Node next;
        Node prev;
        ...
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
        // Implementation for getting the size of the list
    }

    private Node next(Node node) {
        // Implementation for going to the next
    }

    private Node prev(Node node) {
        // Implementation for going to the prev
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        LinkedList list = new DoubleLinkedList();

        // Add elements to the list
        list.add(1);
        list.add(2);
        list.add(3);

        // Access elements by index
        System.out.println("Element at index 0: " + list.at(0)); // Expected output: 1
        System.out.println("Element at index 1: " + list.at(1)); // Expected output: 2
        System.out.println("Element at index 2: " + list.at(2)); // Expected output: 3

        list.add(4);

        // Access elements by index
        System.out.println("Element at index 0: " + list.at(0)); // Expected output: 1
        System.out.println("Element at index 1: " + list.at(1)); // Expected output: 2
        System.out.println("Element at index 2: " + list.at(2)); // Expected output: 3
        System.out.println("Element at index 3: " + list.at(3)); // Expected output: 4

        // Remove an element by index
        list.remove(1);
        list.remove(2);
        System.out.println("List size : " + list.size()); // Expected output: 2
        System.out.println("Element at index 1 after removal: " + list.at(1)); // Expected output: 3
    }
}
```

### Expected Output

```java
$ javac *.java -d build
$ java -cp build ExerciseRunner
Element at index 0: 1
Go to next node
Element at index 1: 2
Element at index 2: 3
Element at index 0: 1
Go to next node
Element at index 1: 2
Go to previous node
Element at index 2: 3
Element at index 3: 4
Go to next node
List size : 2
Element at index 1 after removal: 3
$
```
