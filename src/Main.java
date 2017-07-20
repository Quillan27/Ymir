//Karl Ramberg

//Ymir is a 2D procedural world generator.

//The main control class.
public class Main {

    public int mapWidth = 800;
    public int mapHeight = 600;

    public String title = "Ymir";

    public Window w;
    public Sidebar s;

    public static void main(String[] args){
        new Main();
    }

    public Main(){
        //Create a new Window
        s = new Sidebar(mapHeight);
        w = new Window(title, mapWidth, mapHeight, s);
        w.setVisible(true);

        //Initializes the loop
        loop();
    }

    public void loop(){

        while(27==27){
            if(s.newMapButton.getModel().isPressed()){
                System.out.println("Woot!");
                createMap("");
            }
        }
    }

    public void createMap(String name){
        if(name.equals("")){
            name = getRandomName();
        }

        s.worldName.setText(name);
    }

    public String getRandomName(){
        return "Random";
    }

}
