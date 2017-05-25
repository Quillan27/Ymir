//Karl Ramberg

import javax.swing.*;
import java.awt.*;

//The builder class for the applications window.
public class Window extends JFrame {

    //TODO Change dynamically based on OS.
    public final String LOOKANDFEEL = "com.sun.java.swing.plaf.gtk.GTKLookAndFeel";

    //Panel to the right with controls.
    public Sidebar s;

    public Container c;

    public Window(String title, int mapWidth, int mapHeight){

        //Window Settings
        setTitle(title);
        setLookAndFeel();
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setSize(mapWidth + 200, mapHeight);
        setLocationRelativeTo(null);
        setResizable(false);

        //Init Sidebar
        s = new Sidebar(mapHeight);

        //Add Sidebar aligned to the left.
        c = getContentPane();
        c.add(s, BorderLayout.EAST);

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
