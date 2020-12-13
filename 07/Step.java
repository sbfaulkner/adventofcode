import java.util.*;

public class Step {
  private ArrayList<Step> dependencies;
  private String id;
  private int duration;
  private int progress;
  private Worker worker;

  Step(String stepId) {
    id = stepId;
    dependencies = new ArrayList<Step>();
    duration = 60 + id.charAt(0) - 'A' + 1;
    progress = 0;
  }

  public boolean addDependency(Step step) {
    return dependencies.add(step);
  }

  public String getId() {
    return id;
  }

  public boolean isAvailable() {
    if (worker != null) {
      return false;
    }

    for (Step d : dependencies) {
      if (!d.isComplete()) {
        return false;
      }
    }

    return true;
  }

  public boolean isComplete() {
    return progress == duration;
  }

  public boolean perform() {
    progress++;
    return isComplete();
  }

  public void reset() {
    progress = 0;
    worker = null;
  }

  public Worker setWorker(Worker w) {
    worker = w;

    return w;
  }

  public String toString() {
    return String.format("Step %s", id);
  }
}
