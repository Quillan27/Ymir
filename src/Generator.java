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
        int col = 0;
        int count = 0;
        for(int i = 0; i < elevation.length; i++){
            for(int j = 0; j < elevation[i].length; j++){
                elevation[i][j] = col; // temp val
                if(count>16){
                    count=0;
                    col++;
                }
                if(col>31){
                    col = 31;
                }
                count++;
            }
            col = 0;
            count = 0;
        }

        // TODO pol derives from elev

        // TODO clim derives from elev

        // TODO bio derives from elev

        return elevation;
    }

    public int[][] switchType(int type){
        return elevation;
    }
}
