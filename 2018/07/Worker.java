class Worker {
  int id;
  private Step step;

  Worker(int id) {
    this.id = id;
  }

  public Step getStep() {
    return step;
  }

  public Step setStep(Step s) {
    step = s;

    if (s != null) {
      s.setWorker(this);
    }

    return s;
  }

  public boolean isReady() {
    if (step == null) {
      return true;
    }

    return step.isComplete();
  }

  public String toString() {
    return String.format("Worker # %d", id);
  }

  public boolean work() {
    if (step == null) {
      return false;
    }

    return step.perform();
  }
}