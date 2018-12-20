package com.karl.ymir;

// handles generation for a world, no visuals
// expect only pure math here, lol

public class World {

    private Odin odin;
    private double[][] elevationGrid;

    public World(int width, int height) {
        elevationGrid = new double[width][height];

        for(int x = 0; x < width; x++) { // start flat in the middle
            for(int y = 0; y < height; y++){
                elevationGrid[x][y] = 0.5;
            }
        }

        odin = new Odin();

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
