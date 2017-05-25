//Karl Ramberg

//Ymir is a 2D procedural world generator.

//The main control class.
public class Main {

    public int mapWidth = 600;
    public int mapHeight = 500;

    public String title = "Ymir";

    public Window w;

    public static void main(String[] args){
        new Main();
    }

    public Main(){

        //Create a new Window
        w = new Window(title, mapWidth, mapHeight);
        w.setVisible(true);

    }

}
