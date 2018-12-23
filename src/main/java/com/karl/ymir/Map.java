// created by Karl Ramberg - Mar. 21 2018

// handles all visual sides of the world, no generation

package com.karl.ymir;

import javax.swing.*;
import java.awt.image.BufferedImage;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class Map extends JLabel {

    private World world;
    private BufferedImage image;
    private int width;
    private int height;

    public Map(int width, int height) {
        this.width = width;
        this.height = height;

        image = new BufferedImage(this.width, this.height, BufferedImage.TYPE_INT_ARGB);
        world = new World(this.width, this.height);

        updateMap(1);

        setIcon(new ImageIcon(image));
        setText("No Map Loaded");
    }

    // redraws map image using a map type and the open world
    public void updateMap(int type) {
        ArrayList<String> palette = createPalette(type);

        if (type == 1) { // elevation
            for (int x = 0; x < width; x++) {
                for (int y = 0; y < height; y++) {
                    double elevation = world.getElevation(x, y);
                    //System.out.println(elevation);
                    //System.out.println((int)elevation * 31.0);
                    String color = palette.get((int)(elevation * 31.0)); // TODO change based on final double range, for now 0.0 to 1.0
                    //System.out.println(color);
                    image.setRGB(x, y, convertHexToRGBA(color));
                    //(int)(elevation * 32.0)
                }
            }
        }
    }

    public void newWorld(){
        System.out.println("Generating New World...");
        world = new World(width, height);
        updateMap(1);
        setIcon(new ImageIcon(image));
    }

    // creates color array(palette) from a text file
    private ArrayList<String> createPalette(int type) {
        String path;
        if(type == 1) {
            path = "res/text/elevation.pal";
        } else if(type == 2) {
            path = "res/text/political.pal";
        } else if(type == 3) {
            path = "res/text/climate.pal";
        } else {
            path = "res/text/biome.pal";
        }

        ArrayList<String> colors = new ArrayList<String>();
        try {
            colors = splitFile(path);
        } catch (IOException e) {
            e.printStackTrace();
        }
        return colors;
    }

    // splits newline seperated file into an array
    private ArrayList<String> splitFile(String path) throws IOException{
        BufferedReader br = new BufferedReader(new FileReader(path));
        ArrayList<String> lines = new ArrayList<String>();
        String line = br.readLine();
        while(line != null) {
            lines.add(line);
            //System.out.println(line);
            line = br.readLine();
        }

        return lines;
    }

    // converts hexidecimal color notion to an int that Image.setRGB() can use
    private int convertHexToRGBA(String hex) {
        int r = Integer.parseInt(hex.substring(1,3),16);
        int g = Integer.parseInt(hex.substring(3,5),16);
        int b = Integer.parseInt(hex.substring(5,7),16);

        return (255<<24) + (r<<16) + (g<<8) + b;
    }

    public int getWidth() { return width; }
    public void setWidth(int width) { this.width = width; }
    public int getHeight() { return height; }
    public void setHeight(int height) { this.height = height; }
}
