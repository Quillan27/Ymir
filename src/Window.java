//Karl Ramberg

import javax.swing.*;
import java.awt.*;
import java.awt.font.FontRenderContext;
import java.awt.font.GlyphVector;
import java.awt.geom.AffineTransform;
import java.awt.image.BufferedImage;
import java.awt.image.BufferedImageOp;
import java.awt.image.ImageObserver;
import java.awt.image.RenderedImage;
import java.awt.image.renderable.RenderableImage;
import java.nio.Buffer;
import java.text.AttributedCharacterIterator;

//The builder class for the applications window.
public class Window extends JFrame {

    //TODO Change dynamically based on OS.
    private final String LOOKANDFEEL = "com.sun.java.swing.plaf.gtk.GTKLookAndFeel";

    //Panel to the right with controls.
    private Sidebar sidebar;
    private Map map;
    private Container container;

    public Window(String title, int mapWidth, int mapHeight, Sidebar sidebar, Map map){

        this.sidebar = sidebar;
        this.map = map;

        //Window Settings
        setTitle(title);
        setLookAndFeel();
        setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
        setSize(mapWidth + 200, mapHeight); //200 extra for the sidebar
        setLocationRelativeTo(null);
        setResizable(false);

        //Add Sidebar aligned to the left and Map to the right.
        container = getContentPane();
        container.add(sidebar, BorderLayout.EAST);
        container.add(map, BorderLayout.WEST);

    }



    public void setLookAndFeel(){

        try {
            UIManager.setLookAndFeel(LOOKANDFEEL);
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InstantiationException e) {
            e.printStackTrace();
        } catch (UnsupportedLookAndFeelException e) {
            e.printStackTrace();
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        }

    }

}
