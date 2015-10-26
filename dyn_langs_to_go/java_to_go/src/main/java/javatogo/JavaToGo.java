package javatogo;

import java.io.File;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;

public class JavaToGo {
  static GoAdder GO_ADDER;
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
    String lib = pwd + "/go_adder/libadder." + libExtension;
    GO_ADDER = (GoAdder) Native.loadLibrary(lib, GoAdder.class);
  }

  public interface GoAdder extends Library {
    long Add(long x, long y);
  }

  public static void main(String[] args) {
    System.out.println("Java says: about to call Go ..");
    long total = GO_ADDER.Add(30, 12);
    System.out.println("Java says: result is " + total);
  }
}
