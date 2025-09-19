public class ExerciseRunner {
    public static void main(String[] args) {
        Factory factory = new Factory();

        ConcreteProductA productA = factory.createProduct("A");
        if (productA != null) {
            productA.showDetails();
        } else {
            System.out.println("Invalid product type");
        }

        ConcreteProductB productB = factory.createProduct("B");
        if (productB != null) {
            productB.showDetails();
        } else {
            System.out.println("Invalid product type");
        }

        Object invalidProduct = factory.createProduct("C");
        if (invalidProduct != null) {
            System.out.println("Invalid product type should return null");
        } else {
            System.out.println("Invalid product type");
        }
    }
}