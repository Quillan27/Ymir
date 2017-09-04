//Karl Ramberg

import javax.swing.*;
import java.awt.*;
import java.awt.image.BufferedImage;

public class Map extends JLabel {

    private int width;
    private int height;
    public BufferedImage mapImg;
    private Dimension d;

    private int color;

    public Map(int width, int height) {

        this.width = width;
        this.height = height;

        d = new Dimension(width, height);
        setPreferredSize(d);

        //default to black
        color = toRGBAInt(119,148,198,255);

        mapImg = new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);

        newMap(getRandomWorldName());

    }

    public void newMap(String name) {

        //map changes, solid for now
        for (int i = 0; i < width; i++) {
            for (int j = 0; j < height; j++) {
                mapImg.setRGB(i, j, color);
                System.out.println(color);
            }
        }

        //display map
        setIcon(new ImageIcon(mapImg));

    }

    //TODO finish method
    public String getRandomWorldName(){

        return "New World";

    }

    private int toRGBAInt(int r, int g, int b, int a){

        int c = 0;
        c += a<<24;
        c += r<<16;
        c += g<<8;
        c += b;
        System.out.println("Color: "+c);
        return c;

    }
}

