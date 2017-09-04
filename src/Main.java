//Karl Ramberg

//Ymir is a 2D procedural world generator.

import com.miolean.random.*;

//The main control class.
public class Main {

    public int mapWidth = 800;
    public int mapHeight = 600;

    public String title = "Ymir";

    public Window w;
    public Sidebar s;
    public static Main m;

    public static void main(String[] args){
        m = new Main();
    }

    public Main(){
        //Create a new Window
        s = new Sidebar(mapHeight, this);
        w = new Window(title, mapWidth, mapHeight, s);
        w.setVisible(true);

        //Initializes the loop
        loop();
    }

    public void loop(){
        while(true){

        }
    }

    //Create a new map
    public void createMap(String name){
        if(name.equals("")){
            name = getRandomName();
        }

        s.worldName.setText(name);
        w.setTitle("Ymir - " + name);
    }

    //TODO Finish method
    //Creates a random name
    public String getRandomName(){
        String result;

        WordRandom random = new WordRandom();
        double ran = Math.random();
        if(ran < .1) result = random.nextWord(1);
        else if(ran < .2) result = random.nextWord(3);
        else result = random.nextWord(2);

        return result.substring(0, 1).toUpperCase() + result.substring(1, result.length());
    }

}
