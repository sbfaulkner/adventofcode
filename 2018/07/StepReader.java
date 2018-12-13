import java.io.*;
import java.util.regex.*;

public class StepReader extends BufferedReader {
  private static final Pattern pattern = Pattern.compile(
    "^Step (?<dependency>[A-Z]) must be finished before step (?<id>[A-Z]) can begin.$");

  StepReader(FileReader fileReader) {
    super(fileReader);
  }

  public Step readStep() throws IOException {
    String line = readLine();

    if (line == null) {
      return null;
    }

    Matcher m = pattern.matcher(line);

    if (!m.matches()) {
      throw new IllegalArgumentException("Unrecognized step string [" + line + "]");
    }

    String id = m.group("id");
    Step step = Step.get(id);

    if (step == null) {
      step = new Step(id);
    }

    String dependencyId = m.group("dependency");
    Step dependency = Step.get(dependencyId);

    if (dependency == null) {
      dependency = new Step(dependencyId);
    }

    step.addDependency(dependency);

    return step;
  }
}
