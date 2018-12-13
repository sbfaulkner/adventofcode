import java.net.*;
import java.io.*;
import java.nio.file.*;
import java.util.*;

public class Parts {
  private static String getSourceFile() {
    return Parts.class.getProtectionDomain().getCodeSource().getLocation().getPath();
  }

  public static void main(String[] args) throws Exception {
    String input = args.length > 0 ? args[0] : Paths.get(getSourceFile(), "input").toString();
    File inputFile = new File(input);
    StepReader sr = new StepReader(new FileReader(inputFile));

    while (sr.readStep() != null)
      ;

    sr.close();

    String order = "";
    Step s = Step.available();

    while (s != null) {
      order = order + s.complete();

      s = Step.available();
    }

    System.out.println("Order: " + order);
  }
}
