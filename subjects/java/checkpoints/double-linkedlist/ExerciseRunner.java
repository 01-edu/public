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