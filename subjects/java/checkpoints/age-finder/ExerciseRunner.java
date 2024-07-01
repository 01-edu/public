public class ExerciseRunner {
    public static void main(String[] args) {
        AgeFinder AgeFinder = new AgeFinder();

        // Test case 1
        String date1 = "2000-01-01";
        System.out.println("Age: " + AgeFinder.calculateAge(date1));

        // Test case 2
        String date2 = "1990-06-15";
        System.out.println("Age: " + AgeFinder.calculateAge(date2));

        // Test case 3
        String date3 = "2010-12-25";
        System.out.println("Age: " + AgeFinder.calculateAge(date3));
    }
}