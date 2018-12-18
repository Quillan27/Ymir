// created by Karl Ramberg - Mar. 21 2018
package com.karl.ymir;

import javax.imageio.ImageIO;
import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.FileInputStream;
import java.io.IOException;

public class Window extends JFrame {

    private static final long serialVersionUID = 1L;

    private Container c;
    private Sidebar s;
    private Map m;

    public Window() {

        // window settings
        setTitle("Ymir");
        setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
        setLocationRelativeTo(null);
        setResizable(false);
        
        // create sidebar and default world
        m = new Map();
        s = new Sidebar(m);

        updateSize();

        // add sidebar aligned to the left and world to the right.
        c = getContentPane();
        c.add(s, BorderLayout.EAST);
        c.add(m, BorderLayout.WEST);

        setVisible(true);
    }

    public void updateSize() {
        setSize(m.getWidth() + 200, m.getHeight());
    }
}
