import java.net.*;
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.regex.*;

public class Build {
  private static final Pattern pattern = Pattern.compile(
    "^Step (?<dependency>[A-Z]) must be finished before step (?<id>[A-Z]) can begin.$");

  private static String getSourceFile() {
    return Build.class.getProtectionDomain().getCodeSource().getLocation().getPath();
  }

  public static void main(String[] args) throws Exception {
    String input = args.length > 0 ? args[0] : Paths.get(getSourceFile(), "input").toString();
    Build build = new Build();

    build.readSteps(input);

    String order = build.assemble();

    System.out.printf("Order: %s\n", order);
    System.out.printf("Time: %d\n", build.getAssemblyTime());

    build.reset();

    order = build.assemble(4);

    System.out.printf("Order: %s\n", order);
    System.out.printf("Time: %d\n", build.getAssemblyTime());
  }

  private HashMap<String, Step> steps;
  private int assemblyTime;

  Build() {
    steps = new HashMap<String, Step>();
    assemblyTime = 0;
  }

  public String assemble() {
    return assemble(1);
  }

  public String assemble(int numberOfWorkers) {
    String order = "";
    Worker workers[] = new Worker[numberOfWorkers];

    System.out.print("Second");

    for (int w = 0; w < numberOfWorkers; w++) {
      System.out.printf("   Worker %d", w+1);
      workers[w] = new Worker(w + 1);
    }

    System.out.println("   Done");

    while (order.length() < steps.size()) {
      System.out.printf("  %2d  ", assemblyTime);

      for (Worker w : workers) {
        if (w.isReady()) {
          w.setStep(nextStep());
        }
      }

      for (Worker w : workers) {
        Step s = w.getStep();

        System.out.printf("   %4s    ", s == null ? "." : s.getId());

        if (w.work()) {
          order = order + s.getId();
        }
      }

      System.out.printf("   %s\n", order);

      assemblyTime++;
    }

    return order;
  }

  public Step getStep(String id) {
    Step step = steps.get(id);

    if (step == null) {
      step = new Step(id);
      steps.put(id, step);
    }

    return step;
  }

  public int getAssemblyTime() {
    return assemblyTime;
  }

  public Step nextStep() {
    for (Step s : steps.values()) {
      if (s.isAvailable()) {
        return s;
      }
    }

    return null;
  }

  public void readSteps(String input) throws IOException {
    File inputFile = new File(input);
    BufferedReader br = new BufferedReader(new FileReader(inputFile));

    try {
      String line;

      while ((line = br.readLine()) != null) {
        Matcher m = pattern.matcher(line);

        if (!m.matches()) {
          throw new IllegalArgumentException("Unrecognized step string [" + line + "]");
        }

        Step step = getStep(m.group("id"));
        Step dependency = getStep(m.group("dependency"));

        step.addDependency(dependency);
      }
    }

    finally {
      br.close();
    }
  }

  public void reset() {
    for (Step s: steps.values()) {
      s.reset();
    }

    assemblyTime = 0;
  }
}
