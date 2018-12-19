package com.karl.ymir;

// handles generation for a world, no visuals
// expect only pure math here, lol

public class World {

    private Map m;

    private double[][] elevationGrid;

    public World(Map m, int width, int height) {
        this.m = m;

        elevationGrid = new double[this.m.getWidth()][this.m.getHeight()];

        double col = 0.03125;
        for(int x = 0; x < width; x++) {
            for(int y = 0; y < height; y++){
/*                if(val < 1.0) {
                    elevationGrid[x][y] = val;
                    val += 0.001;
                } else {
                    val = -1.0;
                    elevationGrid[x][y] = val;
                }*/
                elevationGrid[x][y] = col;
            }
            if(col < 1.0) {
                col += 0.03125;
            } else {
                col = 0.03125;
            }
        }
    }

    public double getElevation(int x, int y){
            return elevationGrid[x][y];
    }

    public void saveWorld() {
        System.out.println("World saved");
    }

    public void openWorld() {
        System.out.println("Save open world?");
        System.out.println("Opened world");
    }

}
