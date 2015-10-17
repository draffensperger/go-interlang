import java.io.File;

public class JavaToGo {
  public native int GoAdd(int x, int y);

  static {
    String os = System.getProperty("os.name").toLowerCase();
    String libExtension;
    if (os.contains("mac os")) {
      libExtension = "dylib";
    } else if (os.contains("windows")) {
      libExtension = "dll";
    } else {
      libExtension = "so";
    }

    String pwd = System.getProperty("user.dir");
    String previousDir = new File(pwd).getParentFile().getAbsolutePath();
    System.load(previousDir + "/go_adder/libadder.so");
    // System.load(previousDir + "/go_adder/libadder." + libExtension);
    System.load(pwd + "/java_to_go." + libExtension);
  }

  public static void main(String[] args) {
    System.out.println("Java says: about to call Go ..");
    int total = new JavaToGo().GoAdd(30, 12);
    System.out.println("Java says: result is " + total);
  }
}
