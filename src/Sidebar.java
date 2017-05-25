import com.sun.scenario.effect.impl.sw.java.JSWBlend_COLOR_BURNPeer;

import java.awt.*;
import javax.swing.*;

//A builder for the control panel (sidebar) in the main application window.
public class Sidebar extends JPanel{

    public Dimension d = new Dimension();
    public JLabel worldName;
    public JButton switchToElevation;
    public JButton switchToPolitcal;
    public JButton switchToClimate;
    public JButton switchToBiome;
    public JButton newMap;
    public JButton settings;

    public Sidebar(int mapHeight){

        //Panel fill the window's height and the last 200 pixels on the width.
        d.width = 200;
        d.height = mapHeight;
        setPreferredSize(d);

        //dark grey border
        setBorder(BorderFactory.createLineBorder(Color.DARK_GRAY));

        //configure buttons
        worldName = new JLabel("World Name");
        switchToElevation = new JButton("Elevation");
        switchToPolitcal = new JButton("Poltical");
        switchToClimate = new JButton("Climate");
        switchToBiome = new JButton("Biome");
        newMap = new JButton("New Map");
        settings = new JButton("Settings");
        switchToElevation.setHorizontalTextPosition(JButton.LEFT);

        //set layout type
        setLayout(new GridBagLayout());
        GridBagConstraints gc = new GridBagConstraints();

        //Buttons snap to the center.
        gc.anchor = GridBagConstraints.CENTER;

        //padding between components
        gc.weightx = 0.5;
        gc.weighty = 0.5;

        gc.gridwidth = 2; //Cell size is 2.

        //World Name
        gc.gridx = 0;
        gc.gridy = 0;
        add(worldName, gc);

        //Elevation
        gc.ipadx = 5;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 1;
        add(switchToElevation, gc);

        //Political
        gc.ipadx = 14;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 2;
        add(switchToPolitcal, gc);

        //Climate
        gc.ipadx = 15;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 3;
        add(switchToClimate, gc);

        //Biome
        gc.ipadx = 22;
        gc.ipady = 5;
        gc.gridx = 0;
        gc.gridy = 4;
        add(switchToBiome, gc);

        gc.gridwidth = 1; //Cell size is 1 to accommodate 2 components in the last row.

        //New Map
        gc.ipadx = 20;
        gc.ipady = 20;
        gc.gridx = 0;
        gc.gridy = 5;
        add(newMap, gc);

        //Settings
        gc.ipadx = 20;
        gc.ipady = 5;
        gc.gridx = 1;
        gc.gridy = 5;
        add(settings, gc);

    }
}
