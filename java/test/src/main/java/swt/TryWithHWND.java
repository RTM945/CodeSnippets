package swt;

import com.sun.jna.Native;
import com.sun.jna.Pointer;
import com.sun.jna.platform.win32.WinDef.HWND;
import com.sun.jna.platform.win32.WinUser;
import com.sun.jna.platform.win32.User32;
import com.sun.jna.platform.win32.WinUser.WNDENUMPROC;
import com.sun.jna.win32.StdCallLibrary;

public class TryWithHWND {

   public static void main(String[] args) {
      final User32 user32 = User32.INSTANCE;
      HWND hwnd = user32.FindWindow(null, "Untitled - Notepad");
      System.out.println(hwnd);
      
    //   user32.EnumWindows(new WNDENUMPROC() {
    //      int count = 0;
    //      @Override
    //      public boolean callback(HWND hWnd, Pointer arg1) {
    //         byte[] windowText = new byte[512];
    //         user32.GetWindowTextA(hWnd, windowText, 512);
    //         String wText = Native.toString(windowText);

    //         // get rid of this if block if you want all windows regardless of whether
    //         // or not they have text
    //         if (wText.isEmpty()) {
    //            return true;
    //         }

    //         System.out.println("Found window with text " + hWnd + ", total " + ++count
    //               + " Text: " + wText);
    //         return true;
    //      }
    //   }, null);
   }
}
    