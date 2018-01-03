//Karl Ramberg

//The main control class.
public class Engine {

    private int mapWidth = 800;
    private int mapHeight = 600;

    private String title = "Ymir";

    private Window window;
    private Sidebar sidebar;
    private Map map;

    public static void main(String[] args){
        new Engine();
    }

    public Engine(){
        map = new Map(mapWidth, mapHeight);
        sidebar = new Sidebar(mapHeight, map);
        window = new Window(title, mapWidth, mapHeight, sidebar, map);
        window.setVisible(true);
    }
}