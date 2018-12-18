package com.karl.ymir;

// handles generation for a world, no visuals
// expect only pure math here, lol

public class World {

    private Map map;

    private double[][] elevGrid;
    private double[][] poliGrid;
    private double[][] climGrid;
    private double[][] biomGrid;

    public World(Map map, int width, int height) {
        this.map = map;

        elevGrid = new double[map.getWidth()][map.getHeight()];
        poliGrid = new double[map.getWidth()][map.getHeight()];
        climGrid = new double[map.getWidth()][map.getHeight()];
        biomGrid = new double[map.getWidth()][map.getHeight()];
    }

    public void saveWorld() {
        System.out.println("World saved");
    }

    public void openWorld() {
        System.out.println("Save open world?");
        System.out.println("Opened world");
    }

    public void addPerlinNoise(double[][] grid) {

    }
}
