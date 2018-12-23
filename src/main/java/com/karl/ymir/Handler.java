package com.karl.ymir;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Handler implements ActionListener {

    private Map map;
    private Sidebar sidebar;

    public Handler(Map map, Sidebar sidebar) {
        this.map = map;
        this.sidebar = sidebar;

        this.sidebar.getElevationButton().addActionListener(this);
        this.sidebar.getPoliticalButton().addActionListener(this);
        this.sidebar.getClimateButton().addActionListener(this);
        this.sidebar.getBiomeButton().addActionListener(this);
        this.sidebar.getNewMapButton().addActionListener(this);
        this.sidebar.getSaveButton().addActionListener(this);
        this.sidebar.getSettingsButton().addActionListener(this);
    }

    @Override
    public void actionPerformed(ActionEvent e) {
        if(e.getSource() == sidebar.getElevationButton()) {
            map.updateMap(1);
            System.out.println("Map switched to elevation");
        } else if(e.getSource() == sidebar.getPoliticalButton()) {
            map.updateMap(2);
            System.out.println("Map switched to political");
        } else if(e.getSource() == sidebar.getClimateButton()) {
            map.updateMap(3);
            System.out.println("Map switched to climate");
        } else if(e.getSource() == sidebar.getBiomeButton()) {
            map.updateMap(4);
            System.out.println("Map switched to biome");
        } else if(e.getSource() == sidebar.getNewMapButton()) {
            map.newWorld();
        }
    }
}