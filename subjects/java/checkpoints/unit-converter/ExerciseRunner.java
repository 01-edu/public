public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(UnitConverter.convert(new String[]{"celsius", "fahrenheit", "100"}));
        System.out.println(UnitConverter.convert(new String[]{"fahrenheit", "celsius", "212"}));
        System.out.println(UnitConverter.convert(new String[]{"kilometers", "miles", "5"}));
        System.out.println(UnitConverter.convert(new String[]{"pounds", "kilograms", "10"}));
        System.out.println(UnitConverter.convert(args));
    }
}
