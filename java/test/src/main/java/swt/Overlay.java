package swt;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

import org.eclipse.swt.SWT;
import org.eclipse.swt.events.ControlEvent;
import org.eclipse.swt.events.ControlListener;
import org.eclipse.swt.events.DisposeEvent;
import org.eclipse.swt.events.DisposeListener;
import org.eclipse.swt.events.PaintEvent;
import org.eclipse.swt.events.PaintListener;
import org.eclipse.swt.graphics.Color;
import org.eclipse.swt.graphics.Point;
import org.eclipse.swt.graphics.Rectangle;
import org.eclipse.swt.layout.GridData;
import org.eclipse.swt.layout.GridLayout;
import org.eclipse.swt.widgets.Composite;
import org.eclipse.swt.widgets.Control;
import org.eclipse.swt.widgets.Label;
import org.eclipse.swt.widgets.Scrollable;
import org.eclipse.swt.widgets.Shell;

/**
 *  A customizable overlay over a control.
 *  
 *  @author Loris Securo
 */
public class Overlay {

    private List<Composite> parents;
    private Control objectToOverlay;
    private Shell overlay;
    private Label label;
    private ControlListener controlListener;
    private DisposeListener disposeListener;
    private PaintListener paintListener;
    private boolean showing;
    private boolean hasClientArea;
    private Scrollable scrollableToOverlay;

    public Overlay(Control objectToOverlay) {

        Objects.requireNonNull(objectToOverlay);

        this.objectToOverlay = objectToOverlay;

        // if the object to overlay is an instance of Scrollable (e.g. Shell) then it has 
        // the getClientArea method, which is preferable over Control.getSize
        if (objectToOverlay instanceof Scrollable) {
            hasClientArea = true;
            scrollableToOverlay = (Scrollable) objectToOverlay;
        }
        else {
            hasClientArea = false;
            scrollableToOverlay = null;
        }

        // save the parents of the object, so we can add/remove listeners to them
        parents = new ArrayList<Composite>();
        Composite parent = objectToOverlay.getParent();
        while (parent != null) {
            parents.add(parent);
            parent = parent.getParent();
        }

        // listener to track position and size changes in order to modify the overlay bounds as well
        controlListener = new ControlListener() {
            @Override
            public void controlMoved(ControlEvent e) {
                reposition();
            }

            @Override
            public void controlResized(ControlEvent e) {
                reposition();
            }
        };

        // listener to track paint changes, like when the object or its parents become not visible (for example changing tab in a TabFolder)
        paintListener = new PaintListener() {
            @Override
            public void paintControl(PaintEvent arg0) {
                reposition();
            }
        };

        // listener to remove the overlay if the object to overlay is disposed
        disposeListener = new DisposeListener() {
            @Override
            public void widgetDisposed(DisposeEvent e) {
                remove();
            }
        };

        // create the overlay shell
        overlay = new Shell(objectToOverlay.getShell(), SWT.NO_TRIM);

        // default values of the overlay
        overlay.setBackground(objectToOverlay.getDisplay().getSystemColor(SWT.COLOR_GRAY));
        overlay.setAlpha(200);

        // so the label can inherit the background of the overlay
        overlay.setBackgroundMode(SWT.INHERIT_DEFAULT);

        // label to display a text
        // style WRAP so if it is too long the text get wrapped
        label = new Label(overlay, SWT.WRAP);

        // to center the label
        overlay.setLayout(new GridLayout());
        label.setLayoutData(new GridData(SWT.CENTER, SWT.CENTER, true, true));

        showing = false;
        overlay.open();
        overlay.setVisible(showing);
    }

    public void show() {

        // if it's already visible we just exit
        if (showing) {
            return;
        }

        // set the overlay position over the object
        reposition();

        // show the overlay
        overlay.setVisible(true);

        // add listeners to the object to overlay
        objectToOverlay.addControlListener(controlListener);
        objectToOverlay.addDisposeListener(disposeListener);
        objectToOverlay.addPaintListener(paintListener);

        // add listeners also to the parents because if they change then also the visibility of our object could change
        for (Composite parent : parents) {
            parent.addControlListener(controlListener);
            parent.addPaintListener(paintListener);
        }

        showing = true;
    }

    public void remove() {

        // if it's already not visible we just exit
        if (!showing) {
            return;
        }

        // remove the listeners
        if (!objectToOverlay.isDisposed()) {
            objectToOverlay.removeControlListener(controlListener);
            objectToOverlay.removeDisposeListener(disposeListener);
            objectToOverlay.removePaintListener(paintListener);
        }

        // remove the parents listeners
        for (Composite parent : parents) {
            if (!parent.isDisposed()) {
                parent.removeControlListener(controlListener);
                parent.removePaintListener(paintListener);
            }
        }

        // remove the overlay shell
        if (!overlay.isDisposed()) {
            overlay.setVisible(false);
        }

        showing = false;
    }

    public void setBackground(Color background) {
        overlay.setBackground(background);
    }

    public Color getBackground() {
        return overlay.getBackground();
    }

    public void setAlpha(int alpha) {
        overlay.setAlpha(alpha);
    }

    public int getAlpha() {
        return overlay.getAlpha();
    }

    public boolean isShowing() {
        return showing;
    }

    public void setText(String text) {
        label.setText(text);

        // to adjust the label size accordingly
        overlay.layout();
    }

    public String getText() {
        return label.getText();
    }

    private void reposition() {

        // if the object is not visible, we hide the overlay and exit
        if (!objectToOverlay.isVisible()) {
            overlay.setBounds(new Rectangle(0, 0, 0, 0));
            return;
        }

        // if the object is visible we need to find the visible region in order to correctly place the overlay

        // get the display bounds of the object to overlay
        Point objectToOverlayDisplayLocation = objectToOverlay.toDisplay(0, 0);

        Point objectToOverlaySize;

        // if it has a client area, we prefer that instead of the size 
        if (hasClientArea) {
            Rectangle clientArea = scrollableToOverlay.getClientArea();
            objectToOverlaySize = new Point(clientArea.width, clientArea.height);
        }
        else {
            objectToOverlaySize = objectToOverlay.getSize();
        }

        Rectangle objectToOverlayBounds = new Rectangle(objectToOverlayDisplayLocation.x, objectToOverlayDisplayLocation.y, objectToOverlaySize.x,
                objectToOverlaySize.y);

        Rectangle intersection = objectToOverlayBounds;

        // intersect the bounds of the object with its parents bounds so we get only the visible bounds
        for (Composite parent : parents) {

            Rectangle parentClientArea = parent.getClientArea();
            Point parentLocation = parent.toDisplay(parentClientArea.x, parentClientArea.y);
            Rectangle parentBounds = new Rectangle(parentLocation.x, parentLocation.y, parentClientArea.width, parentClientArea.height);

            intersection = intersection.intersection(parentBounds);

            // if intersection has no size then it would be a waste of time to continue
            if (intersection.width == 0 || intersection.height == 0) {
                break;
            }
        }

        overlay.setBounds(intersection);
    }

}
