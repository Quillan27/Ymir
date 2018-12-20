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

    private Handler handler;
    private Container container;
    private Sidebar sidebar;
    private Map map;

    private Dimension screenSize;
    private int width;
    private int height;

    public Window() {

        setTitle("Ymir");
        setDefaultCloseOperation(JFrame.DISPOSE_ON_CLOSE);
        setLocationRelativeTo(null);
        setResizable(false);

        // get screen resolution to scale window/map size
        screenSize = Toolkit.getDefaultToolkit().getScreenSize();
        System.out.println("Screen Resolution: " + screenSize.getWidth() + " " + screenSize.getHeight());
        width = (int)screenSize.getWidth() / 2;
        height = (int)(screenSize.getHeight() / 1.5);
        setSize(width + (width / 4), height); // set window to map size plus extra for the sidebar

        sidebar = new Sidebar(width, height);
        map = new Map(width, height);
        handler = new Handler(map, sidebar);

        // add sidebar aligned to the left and map to the right
        container = getContentPane();
        container.add(sidebar, BorderLayout.EAST);
        container.add(map, BorderLayout.WEST);

        setVisible(true);
    }
}
