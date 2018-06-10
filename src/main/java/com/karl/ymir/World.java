// created by Karl Ramberg - Mar. 21 2018
package com.karl.ymir;

import javax.swing.*;
import java.awt.image.BufferedImage;

public class World extends JLabel {
    private static final long serialVersionUID = 1L;

	private int width, height;

    private double[][] grid;

    private int type;
    private String foobar = "foobar";
    private BufferedImage img;

    public World(int type) {
        this.type = type;

        if(this.type==0){ //default
            width = 800;
            height = 600;
            img = new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
            fillImg(hexToRGBA("#0000ff"));
        }

        setIcon(new ImageIcon(img));
    }

    public void fillImg(int col){
        for(int i = 0; i < width; i++){
            for(int j = 0; j < height; j++){
                img.setRGB(i, j, col);
            }
        }
        System.out.println("Image Filled");
    }

    private int hexToRGBA(String hex) {
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

    public int getHeight(){
        return height;
    }

    public int getWidth(){
        return width;
    }
}
