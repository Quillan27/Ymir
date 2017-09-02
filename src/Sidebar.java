import com.sun.scenario.effect.impl.sw.java.JSWBlend_COLOR_BURNPeer;

import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.io.IOException;
import javax.imageio.ImageIO;
import javax.swing.*;

//A builder for the control panel (sidebar) in the main application window.
public class Sidebar extends JPanel{

    public Dimension d = new Dimension();

    public JLabel worldName;

    public Handler h;

    //buttons
    public JButton elevationButton;
    public JButton politicalButton;
    public JButton climateButton;
    public JButton biomeButton;
    public JButton newMapButton;
    public JButton settingsButton;
    public JButton saveButton;

    //icons
    public Image elevationIcon;
    public Image politicalIcon;
    public Image climateIcon;
    public Image biomeIcon;
    public Image newMapIcon;
    public Image settingsIcon;
    public Image saveIcon;
    
    public Main m;

    public Sidebar(int mapHeight, Main m){
        this.m = m;
        h = new Handler();

        //Panel fill the window's height and the last 200 pixels on the width.
        d.width = 200;
        d.height = mapHeight;
        setPreferredSize(d);

        //dark grey border
        setBorder(BorderFactory.createLineBorder(Color.decode("#052D49")));

        //configure buttons
        worldName = new JLabel("World Name");
        elevationButton = new JButton("Elevation");
        politicalButton = new JButton("Poltical");
        climateButton = new JButton("Climate");
        biomeButton = new JButton("Biome");
        newMapButton = new JButton("");
        settingsButton = new JButton("");
        saveButton = new JButton("");

        //get resources for icons
        try {
            elevationIcon = ImageIO.read(getClass().getResource("elevation.png"));
            politicalIcon = ImageIO.read(getClass().getResource("political.png"));
            climateIcon = ImageIO.read(getClass().getResource("climate.png"));
            biomeIcon = ImageIO.read(getClass().getResource("biome.png"));
            newMapIcon = ImageIO.read(getClass().getResource("newMap.png"));
            settingsIcon = ImageIO.read(getClass().getResource("settings.png"));
            saveIcon = ImageIO.read(getClass().getResource("save.png"));
        } catch (IOException e) {

            e.printStackTrace();
        }

        //set font for world name
        worldName.setFont(new Font("Roboto", Font.BOLD, 24));
        worldName.setForeground(Color.decode("#052D49"));

        //set icons
        elevationButton.setIcon(new ImageIcon(elevationIcon));
        politicalButton.setIcon(new ImageIcon(politicalIcon));
        climateButton.setIcon(new ImageIcon(climateIcon));
        biomeButton.setIcon(new ImageIcon(biomeIcon));
        newMapButton.setIcon(new ImageIcon(newMapIcon));
        settingsButton.setIcon(new ImageIcon(settingsIcon));
        saveButton.setIcon(new ImageIcon(saveIcon));

        //nifty tooltips
        elevationButton.setToolTipText("Switch to Elevation View");
        politicalButton.setToolTipText("Switch to Political View");
        climateButton.setToolTipText("Switch to Climate View");
        biomeButton.setToolTipText("Switch to Biome View");
        newMapButton.setToolTipText("Generate New Map");
        settingsButton.setToolTipText("Change Settings");
        saveButton.setToolTipText("Save World");

        //consistent fonts
        elevationButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        politicalButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        climateButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        biomeButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        newMapButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        settingsButton.setFont(new Font("Roboto", Font.PLAIN, 12));
        saveButton.setFont(new Font("Roboto", Font.PLAIN, 12));

        newMapButton.addActionListener(h);
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

    private class Handler implements ActionListener, KeyListener {

        public Handler(){
        }

        @Override
        public void actionPerformed(ActionEvent e) {

            if(e.getSource() == newMapButton){
                String name  = "";
                m.createMap(name);
            }
        }

        @Override
        public void keyTyped(KeyEvent e) {

        }

        @Override
        public void keyPressed(KeyEvent e) {

        }

        @Override
        public void keyReleased(KeyEvent e) {

        }
    }
}