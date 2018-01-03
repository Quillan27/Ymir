//Karl Ramberg

import javax.swing.*;
import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.*;
import java.util.ArrayList;


// class for raw integer conversion and map displaying
public class Map extends JLabel {

    public BufferedImage img;
    private Generator generator;

    public Map(int width, int height) {
        generator = new Generator(width, height);

        setPreferredSize(new Dimension(width, height));

        img = new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
        setText("No Map Loaded");
    }

    public void newMap(String name) throws IOException {
        // activate the generator!!
        int[][] vals = generator.generateNewWorld();

        // convert to vals to img
        //TODO
        valToImg(vals, 0);

        // display map
        setIcon(new ImageIcon(img));
    }

    // TODO implement random name gen
    public String getRandomWorldName(){
        return "New World";
    }

    // converts integers to colors and puts them into image
    private void valToImg(int[][] vals, int type) throws IOException {
        ArrayList<String> palette;

        switch(type){
            case 1: palette = textToArr("res/text/politicalCol.txt");
                    break;
            case 2: palette = textToArr("res/text/climateCol.txt");
                    break;
            case 3: palette = textToArr("res/text/biomeCol.txt");
                    break;
            default: palette = textToArr("res/text/elevationCol.txt");
                    break;
        }

        int col;
        for(int i = 0; i < vals.length; i++){
            for(int j = 0; j < vals[i].length; j++){
                col = hexToInt(palette.get(vals[i][j]));
                img.setRGB(i, j, col);
            }
        }
    }

    // #ffffff format to rbga integer
    private int hexToInt(String hex){
        // get substring and parse an int from the hexidecimal, 0-255
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

    // partitions a files lines to separate indices
    private ArrayList<String> textToArr(String path) throws IOException {
        ArrayList<String> lines = new ArrayList<String>();
        String line;
        BufferedReader br = new BufferedReader(new FileReader(path));

        line = br.readLine();
        while(line != null){
            lines.add(line);
            line = br.readLine();
        }
        return lines;
    }
}

