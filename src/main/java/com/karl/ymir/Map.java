// created by Karl Ramberg - Mar. 21 2018

// handles all visual sides of the world, no generation

package com.karl.ymir;

import javax.swing.*;
import java.awt.image.BufferedImage;

public class Map extends JLabel {
    private static final long serialVersionUID = 1L;

    private BufferedImage img;
    private World w;

    public Map() {
        w = new World(this, getWidth(), getHeight());
        img = new BufferedImage(800, 600, BufferedImage.TYPE_INT_ARGB);
        fillImg("#001f0f");

        setIcon(new ImageIcon(img));
    }

    public void fillImg(String col) { // TEMP?
        for(int i = 0; i < getWidth(); i++) {
            for(int j = 0; j < getHeight(); j++) {
                img.setRGB(i, j, convertHexToRGBA(col));
            }
        }
        System.out.println("Map filled with " + col);
    }

    // converts hexidecimal color notion to an int that Image.setRGB() can use
    private int convertHexToRGBA(String hex) {

        // get substring and parse an int from the hexidecimal, 0-255
        int r = Integer.parseInt(hex.substring(1,3),16);
        int g = Integer.parseInt(hex.substring(3,5),16);
        int b = Integer.parseInt(hex.substring(5,7),16);
        int a = 255;

        // combine into one integer using bit manipulation
        int c = 0;
        c += a<<24;
        c += r<<16;
        c += g<<8;
        c += b;

        return c;

    }

    public int getWidth() { return img.getWidth(); }
    public int getHeight() { return img.getHeight(); }
}
