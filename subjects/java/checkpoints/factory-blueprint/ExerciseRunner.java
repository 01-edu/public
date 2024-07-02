public class ExerciseRunner {
    public static void main(String[] args) {
        Factory factory = new Factory();

        // Handle invalid product type
        Object invalidProduct = factory.createProduct("C");
        if (invalidProduct != null) {
            invalidProduct.showDetails();
        } else {
            System.out.println("Invalid product type");
        }
    }
}