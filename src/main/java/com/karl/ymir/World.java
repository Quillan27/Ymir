package com.karl.ymir;

// handles generation for a world, no visuals
// expect only pure math here, lol

import java.util.Random;

public class World {

    private double[][] elevationGrid;

    private int width;
    private int height;

    public World(int width, int height) {
        this.width = width;
        this.height = height;

        elevationGrid = new double[width][height];

        initElevation();
        addPeaks(20);
        addNoise();
    }

    private void initElevation() { // the earth is flat
        for(int x = 0; x < width; x++) {
            for (int y = 0; y < height; y++) {
                elevationGrid[x][y] = 0.0;
            }
        }
    }

    private void addPeaks(int count) { //TODO so many magic numbers, i might actually be merlin
        Random r = new Random(System.currentTimeMillis());
        int[][] peaks = new int[count][2];

        for(int i = 0; i < count; i++) {
            int randX = r.nextInt(width);
            int randY =  r.nextInt(height);
            //System.out.println(randX + " " + randY);
            elevationGrid[randX][randY] = 0.75 + (1.0 - 0.75) * r.nextDouble();
            peaks[i][0] = randX;
            peaks[i][1] = randY;
        }

        for(int x = 0; x < width; x++) {
            for (int y = 0; y < height; y++) {
                double elevScore = 0.0;
                for(int i = 0; i < count; i++) {
                    int dis = getDis(x, y, peaks[i][0], peaks[i][1]);
                    if(dis < 350) {
                        elevScore += 350 - dis;
                    }
                }
                elevationGrid[x][y] += elevScore / 1200;
                elevationGrid[x][y] = chomp(elevationGrid[x][y], 0.0, 1.0);
            }
        }
    }

    private void addNoise() { // shitty random noise with no pattern, perlin would spit on this
        Random r = new Random(System.currentTimeMillis());
        for(int x = 0; x < width; x++) {
            for(int y = 0; y < height; y++) {
                double randChange = 0.0 + (0.01 - 0.0) * r.nextDouble();
                elevationGrid[x][y] += randChange;
                elevationGrid[x][y] = chomp(elevationGrid[x][y], 0.0, 1.0);
            }
        }
    }

    private double chomp(double num, double min, double max) {
        if(num < min) {
            return min;
        } else if (num > max) {
            return max;
        } else {
            return num;
        }
    }

    private int getDis(int x1, int y1, int x2, int y2) {
        return (int)Math.sqrt((x1 - x2)*(x1 - x2) + (y1 - y2) * (y1 - y2));
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
