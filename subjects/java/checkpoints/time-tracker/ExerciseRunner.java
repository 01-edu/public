public class ExerciseRunner {
    public static void main(String[] args) {
        // Example of a very short project
        ProjectTime shortProject = new ProjectTime("2023-05-14 09:00", "2023-05-14 09:30");
        System.out.println("Short Project Total Logged Time: " + shortProject.getHoursLogged());  // Should be "30 m"

        // Example of an overnight project
        ProjectTime overnightProject = new ProjectTime("2023-05-14 20:00", "2023-05-15 08:00");
        System.out.println("Overnight Project Total Logged Time: " + overnightProject.getHoursLogged());  // Should be "12 h"

        // Example of a full day project
        ProjectTime fullDayProject = new ProjectTime("2023-05-14 00:00", "2023-05-15 00:00");
        System.out.println("Full Day Project Total Logged Time: " + fullDayProject.getHoursLogged());  // Should be "24 h"

        // Example with incorrect format
        ProjectTime errorProject = new ProjectTime("2023-05-14", "2023-05-15 08:00");
        System.out.println("Error Project Total Logged Time: " + errorProject.getHoursLogged());  // Should handle error or indicate incorrect format
    }
}