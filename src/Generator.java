//Karl Ramberg

// class for generating elevation values and others based on elevation
public class Generator {

    // elevation, political, climate, biome map values
    private int[][] elevation;
    private int[][] poltical;
    private int[][] climate;
    private int[][] biome;

    private int width, height;

    public Generator(int width, int height){
        this.width = width;
        this.height = height;
    }

    //elevation is default
    public int[][] generateNewWorld(){
        elevation = new int[width][height];

        //TODO elevation gen
        for(int i = 0; i < elevation.length; i++){
            for(int j = 0; j < elevation[i].length; j++){
                elevation[i][j] = 14; // temp val
            }
        }

        // TODO pol derives from elev

        // TODO clim derives from elev

        // TODO bio derives from elev

        return elevation;
    }
}
