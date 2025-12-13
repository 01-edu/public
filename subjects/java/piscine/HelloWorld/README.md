## HelloWorld

### Setup

Install the [Java 17 JDK](https://www.oracle.com/java/technologies/javase/jdk17-archive-downloads.html)

The files for each exercise will be submitted in a single directory. For example, the directory for this exercise will be `HelloWorld`, and the path to the file will be `HelloWorld/HelloWorld.java`.

You will need to compile your code if you intend to test it:

```shell
javac *.java -d build
```

Then run the following command :

```shell
java -cp build ExerciseRunner
```

To edit your code, you can use any IDE or text editor, though IDEA IntelliJ or JetBrains are specially dedicated and recommended.

### Instructions

Create a file `HelloWorld.java` in a directory named `HelloWorld`.

Write a function `helloWorld` that returns the string 'Hello World !'.

### Expected Functions

```java
public class HelloWorld {
    public static String helloWorld() {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(HelloWorld.helloWorld());
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Hello World !
$
```

### Submit with Git

Create in your [Gitea](<https://((DOMAIN))/git>) account the repository named `((ROOT))`.

This repository will be the folder where all the exercises must be uploaded. Each of your exercises will have its own directory, with a name matching the exercise. That directory will contain all of the files for that exercise.

Once created, clone that repository on your desktop.

To do so, open a Unix shell (e.g. Git Bash on Windows), you are going to type commands in it.

First, tell Git to remember your password (like a web browser would):

```
git config --global credential.helper store
```

If your username was `01-user` this is the command that will need to be used:

```
git clone https://((DOMAIN))/git/01-user/((ROOT)).git
```

This command needs to be adapted with **your own username**.

Now you'll need to "commit" your code using the following commands:

1. `git add HelloWorld/HelloWorld.java`
2. `git commit -m "add HelloWorld exercise solution"`
3. `git push`

Once these steps are applied, the file can now be submitted for grading on the platform by clicking on the "RUN HELLOWORLD TEST" button.

This action will run the tests on your submitted `HelloWorld.java` file.
