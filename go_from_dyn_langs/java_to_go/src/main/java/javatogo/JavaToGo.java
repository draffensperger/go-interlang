package javatogo;

import java.io.File;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Platform;

public class JavaToGo {
  static GoAdder GO_ADDER;
  static {
    String lib = "/Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-inter-lang-calls/go_from_dyn_langs/go_adder/libadder.dylib";
    GO_ADDER = (GoAdder) Native.loadLibrary(lib, GoAdder.class);
  }

  public interface GoAdder extends Library {
    int Add(int x, int y);
  }

  public static void main(String[] args) {
    System.out.println("Java says: about to call Go ..");
    int total = GO_ADDER.Add(30, 12);
    System.out.println("Java says: result is " + total);
  }

  // public static void main(String[] args) {
  //   System.out.println("Java says: about to call Go ..");
  //   int total = new JavaToGo().GoAdd(30, 12);
  //   System.out.println("Java says: result is " + total);
  // }
}
