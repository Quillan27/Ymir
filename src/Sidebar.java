import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.IOException;
import javax.imageio.ImageIO;
import javax.swing.*;

//A builder for the control panel (sidebar) in the main application window.
public class Sidebar extends JPanel {

    public Dimension d;

    public JLabel worldName;

    private Handler handler;
    private Map map;

    //buttons
    private JButton elevationButton;
    private JButton politicalButton;
    private JButton climateButton;
    private JButton biomeButton;
    private JButton newMapButton;
    private JButton saveButton;
    private JButton settingsButton;

    //icons
    private Image elevationIcon;
    private Image politicalIcon;
    private Image climateIcon;
    private Image biomeIcon;
    private Image newMapIcon;
    private Image saveIcon;
    private Image settingsIcon;

    public Sidebar(int mapHeight, Map map) {

        this.map = map;
        handler = new Handler();

        //Panel fill the window's height and the last 200 pixels on the width.
        d = new Dimension(200, mapHeight);
        setPreferredSize(d);

        /* CONFIGURE GUI */

        //set world name
        worldName = new JLabel("World Name");
        worldName.setFont(new Font("Roboto", Font.BOLD, 24));

        elevationButton = new JButton("Elevation");
        politicalButton = new JButton("Poltical");
        climateButton = new JButton("Climate");
        biomeButton = new JButton("Biome");
        newMapButton = new JButton("");
        settingsButton = new JButton("");
        saveButton = new JButton("");

        //get resources for icons
        try {

            elevationIcon = ImageIO.read(getClass().getResource("icons/elevation.png"));
            politicalIcon = ImageIO.read(getClass().getResource("icons/political.png"));
            climateIcon = ImageIO.read(getClass().getResource("icons/climate.png"));
            biomeIcon = ImageIO.read(getClass().getResource("icons/biome.png"));
            newMapIcon = ImageIO.read(getClass().getResource("icons/newMap.png"));
            settingsIcon = ImageIO.read(getClass().getResource("icons/settings.png"));
            saveIcon = ImageIO.read(getClass().getResource("icons/save.png"));

        } catch (IOException e) {

            e.printStackTrace();

        }

        //set icons
        elevationButton.setIcon(new ImageIcon(elevationIcon));
        politicalButton.setIcon(new ImageIcon(politicalIcon));
        climateButton.setIcon(new ImageIcon(climateIcon));
        biomeButton.setIcon(new ImageIcon(biomeIcon));
        newMapButton.setIcon(new ImageIcon(newMapIcon));
        saveButton.setIcon(new ImageIcon(saveIcon));
        settingsButton.setIcon(new ImageIcon(settingsIcon));

        //nifty tooltips
        elevationButton.setToolTipText("Switch to Elevation View");
        politicalButton.setToolTipText("Switch to Political View");
        climateButton.setToolTipText("Switch to Climate View");
        biomeButton.setToolTipText("Switch to Biome View");
        newMapButton.setToolTipText("Generate New Map");
        saveButton.setToolTipText("Save World");
        settingsButton.setToolTipText("Change Settings");

        //consistent fonts
        elevationButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        politicalButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        climateButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        biomeButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        newMapButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        saveButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        settingsButton.setFont(new Font("Roboto", Font.PLAIN, 12));

        //Add listeners
        elevationButton.addActionListener(handler);
        politicalButton.addActionListener(handler);
        climateButton.addActionListener(handler);
        biomeButton.addActionListener(handler);
        newMapButton.addActionListener(handler);
        saveButton.addActionListener(handler);
        settingsButton.addActionListener(handler);

        /* LAYOUT */

        //set layout type
        setLayout(new GridBagLayout());
        GridBagConstraints gc = new GridBagConstraints();

        //Buttons snap to the center.
        gc.anchor = GridBagConstraints.CENTER;

        //padding between components
        gc.weightx = 0.5;
        gc.weighty = 0.5;

        //Cell size is 3 for wider buttons.
        gc.gridwidth = 3;

        //World Name
        gc.gridx = 0;
        gc.gridy = 0;
        add(worldName, gc);

        //Elevation
        gc.ipadx = 10;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 1;
        add(elevationButton, gc);

        //Political
        gc.ipadx = 19;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 2;
        add(politicalButton, gc);

        //Climate
        gc.ipadx = 20;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 3;
        add(climateButton, gc);

        //Biome
        gc.ipadx = 27;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 4;
        add(biomeButton, gc);

        //Cell size is now 1 to accommodate 3 components in the last row.
        gc.gridwidth = 1;

        //New Map
        gc.ipadx = 0;
        gc.ipady = 0;
        gc.gridx = 0;
        gc.gridy = 5;
        add(newMapButton, gc);

        //Save
        gc.ipadx = 0;
        gc.ipady = 0;
        gc.gridx = 1;
        gc.gridy = 5;
        add(saveButton, gc);

        //Settings
        gc.ipadx = 0;
        gc.ipady = 0;
        gc.gridx = 2;
        gc.gridy = 5;
        add(settingsButton, gc);

    }

    private class Handler implements ActionListener {

        private Handler(){}

        @Override   //Button functions
        public void actionPerformed(ActionEvent e) {

            //Elevation Button Functions
            if(e.getSource() == elevationButton){}

            //Political Button Functions
            if(e.getSource() == politicalButton){}

            //Climate Button Functions
            if(e.getSource() == climateButton){}

            //Biome Button Functions
            if(e.getSource() == biomeButton){}

            //New Map Button, creates a new map. see Map.java
            if(e.getSource() == newMapButton) {

                String name  = "";
                if(name.equals("")){
                    name = map.getRandomWorldName();
                }
                worldName.setText(name);
                map.newMap(name);

            }

            //Save Button Functions
            if(e.getSource() == saveButton){}

            //Settings Button Functions
            if(e.getSource() == settingsButton){}
        }
    }
}