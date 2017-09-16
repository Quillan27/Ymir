//Karl Ramberg

import javax.swing.*;
import java.awt.*;
import java.awt.image.BufferedImage;

public class Map extends JLabel {

    private int width;
    private int height;
    public BufferedImage mapImg;
    private Dimension d;
    private Generator generator;

    private int color;

    public Map(int width, int height) {
        generator = new Generator();
        this.width = width;
        this.height = height;

        d = new Dimension(width, height);
        setPreferredSize(d);

        //default to black
        color = hexToInt("#7794c6");

        mapImg = new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
        setText("  No Map Loaded");

    }

    public void newMap(String name) {
        //activate the generator!!
        int[][][] world = generator.generateNewWorld();

        //convert to image
        //TODO

        //display map
        setIcon(new ImageIcon(mapImg));

    }

    //TODO finish method
    public String getRandomWorldName(){

        return "New World";

    }

    private int hexToInt(String hex){

        System.out.println("Color: "+hex);

        //Get substring and parse an int from the hexidecimal, 0-255
        int r = Integer.parseInt(hex.substring(1,3),16);
        int g = Integer.parseInt(hex.substring(3,5),16);
        int b = Integer.parseInt(hex.substring(5,7),16);
        int a = 255;

        //combine into one integer using bit manipulation
        int c = 0;
        c += a<<24;
        c += r<<16;
        c += g<<8;
        c += b;

        return c;

    }
}

