public class ExerciseRunner {
    public static void main(String[] args) {
        RegexDirector director = new RegexDirector();
        RegexBuilder builder = new ConcreteRegexBuilder();

        director.setBuilder(builder);
        Regex regex = director.construct();

        System.out.println(regex);
    }
}