public class ExerciseRunner {
    public static void main(String[] args) {
        MonthlyPeriod MonthlyPeriod = new MonthlyPeriod();

        String startDate1 = "2020-01-01";
        String endDate1 = "2023-06-15";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate1, endDate1));

        String startDate2 = "2015-05-20";
        String endDate2 = "2015-10-19";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate2, endDate2));

        String startDate3 = "2015-05-20";
        String endDate3 = "2015-10-20";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate3, endDate3));

        String startDate4 = "2018-12-25";
        String endDate4 = "2021-12-25";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate4, endDate4));

        String startDate5 = "2018-10-25";
        String endDate5 = "2019-11-30";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate5, endDate5));
    }
}