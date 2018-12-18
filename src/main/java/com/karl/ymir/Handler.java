package com.karl.ymir;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Handler implements ActionListener {

    public Handler() {}

    @Override
    public void actionPerformed(ActionEvent e) {

        System.out.println("actions are performed");

    }
}