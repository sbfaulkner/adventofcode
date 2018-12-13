import java.util.*;

public class Step {
  private ArrayList<Step> dependencies;
  private String id;

  private static HashMap<String, Step> allSteps = new HashMap<String, Step>();

  public static Collection<Step> all() {
    return allSteps.values();
  }

  public static Step available() {
    for (Step s : all()) {
      if (s.isAvailable()) {
        return s;
      }
    }

    return null;
  }

  public static int count() {
    return allSteps.size();
  }

  public static Step get(String id) {
    return allSteps.get(id);
  }

  Step(String id) {
    this.id = id;
    this.dependencies = new ArrayList<Step>();

    allSteps.put(id, this);
  }

  public boolean addDependency(Step step) {
    return dependencies.add(step);
  }

  public String complete() {
    for (Step s : all()) {
      s.removeDependency(this);
    }

    allSteps.remove(this.id);

    return this.id;
  }

  public String getId() {
    return id;
  }

  public boolean removeDependency(Step step) {
    return dependencies.remove(step);
  }

  public boolean isAvailable() {
    return dependencies.isEmpty();
  }

  public String toString() {
    return String.format("Step: %s <- %s", id, dependencies);
  }
}
