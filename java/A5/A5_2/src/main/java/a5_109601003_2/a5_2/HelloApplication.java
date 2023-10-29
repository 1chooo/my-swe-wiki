/**
 * Assignment 5
 * Student Number: 109601003
 * Course: CE1002
 */

package a5_109601003_2.a5_2;

import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.fxml.FXMLLoader;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.input.MouseButton;
import javafx.scene.input.MouseEvent;
import javafx.stage.Stage;

import java.io.IOException;

public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) throws IOException {
    Group root = new Group();
    ToggleGroup groupRadioButton = new ToggleGroup();
    RadioButton radioButton1 = new RadioButton("A");
    radioButton1.setToggleGroup(groupRadioButton);
    radioButton1.setLayoutX(75);
    radioButton1.setLayoutY(100);
    RadioButton radioButton2 = new RadioButton("B");
    radioButton2.setLayoutX(175);
    radioButton2.setLayoutY(100);
    radioButton2.setToggleGroup(groupRadioButton);
    Button btPress2 = new Button("Press2");
    btPress2.setLayoutX(115);
    btPress2.setLayoutY(150);
    btPress2.addEventHandler(MouseEvent.MOUSE_CLICKED, new EventHandler<MouseEvent>() {
      @Override
      public void handle(MouseEvent mouseEvent) {
        if (radioButton1.isSelected()) {
          System.out.println("A");
        } else if (radioButton2.isSelected()) {
          System.out.println("B");
        } else {
          System.out.println("Not choose any button.");
        }
      }
    });


    CheckBox checkBox = new CheckBox("CheckBox");
    checkBox.setLayoutX(100);
    checkBox.setLayoutY(300);
    Button btPress3 = new Button("Press3");
    btPress3.setLayoutX(115);
    btPress3.setLayoutY(350);

    btPress3.addEventHandler(MouseEvent.MOUSE_CLICKED, new EventHandler<MouseEvent>() {
      @Override
      public void handle(MouseEvent mouseEvent) {
        if (checkBox.isSelected()) {
          System.out.println("true");
        } else {
          System.out.println("false");
        }
      }
    });

    TextField textField = new TextField();
    textField.setLayoutX(350);
    textField.setLayoutY(225);
    Button btPress1 = new Button("Press1");
    btPress1.setLayoutX(550);
    btPress1.setLayoutY(225);
    btPress1.addEventHandler(MouseEvent.MOUSE_CLICKED, new EventHandler<MouseEvent>() {
      @Override
      public void handle(MouseEvent mouseEvent) {
        System.out.println(textField.getText());
      }
    });


    root.getChildren().add(radioButton1);
    root.getChildren().add(radioButton2);
    root.getChildren().add(btPress2);
    root.getChildren().add(checkBox);
    root.getChildren().add(btPress3);
    root.getChildren().add(textField);
    root.getChildren().add(btPress1);

    Scene scene = new Scene(root, 750, 500);
    stage.setTitle("Hello World");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) {
    launch();
  }
}