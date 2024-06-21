public class ExerciseRunner {
    public static void main(String[] args) {
        DateFormatter df = new DateFormatter(1656374400, "DD/MM/YYYY");
        System.out.println(df.getFormattedDate());

        df.setFormat("dd.MM.yyyy");
        System.out.println(df.getFormattedDate());

        df.setDate(1672531200);
        System.out.println(df.getFormattedDate());

        df.setFormat("DD month yyyy");
        System.out.println(df.getFormattedDate());
    }
}