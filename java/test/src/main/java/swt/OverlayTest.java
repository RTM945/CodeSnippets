package swt;

import org.eclipse.swt.*;
import org.eclipse.swt.events.SelectionAdapter;
import org.eclipse.swt.events.SelectionEvent;
import org.eclipse.swt.graphics.*;
import org.eclipse.swt.internal.win32.OS;
import org.eclipse.swt.internal.win32.TCHAR;
import org.eclipse.swt.layout.*;
import org.eclipse.swt.widgets.*;

public class OverlayTest {

    public static void main(String[] args) {
        
        Display display = new Display();
        Shell shell = new Shell(display);
        shell.setLayout(new FillLayout(SWT.VERTICAL));
        shell.setSize(250, 250);

        // create the composite
        Composite composite = new Composite(shell, SWT.NONE);
        composite.setLayout(new FillLayout(SWT.VERTICAL));

        // add stuff to the composite
        for (int i = 0; i < 5; i++) {
            new Text(composite, SWT.BORDER).setText("Text " + i);
        }

        // create the overlay over the composite
        Overlay overlay = new Overlay(composite);
        overlay.setText("No data available");

        // create the button to show/hide the overlay
        Button button = new Button(shell, SWT.PUSH);
        button.setText("Show/hide overlay");
        button.addSelectionListener(new SelectionAdapter() {
            @Override
            public void widgetSelected(SelectionEvent arg0) {
                // if the overlay is showing we hide it, otherwise we show it
                if (overlay.isShowing()) {
                    overlay.remove();
                }
                else {
                    overlay.show();
                }
            }
        });

        shell.open();
        while (shell != null && !shell.isDisposed()) {
            if (!display.readAndDispatch()) {
                display.sleep();
            }
        }
    }

}
