package com.karl.ymir;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Handler implements ActionListener {

    private Sidebar s;
    public Handler(Sidebar s) {
        this.s = s;
    }

    @Override
    public void actionPerformed(ActionEvent e) {

        System.out.println("actions are performed");
        if(e.getSource() == s.getElevationButton()){

        }
    }
}