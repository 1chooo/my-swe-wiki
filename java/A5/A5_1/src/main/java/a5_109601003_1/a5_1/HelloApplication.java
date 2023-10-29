/**
 * Assignment 5
 * Student Number: 109601003
 * Course: CE1002
 */

package a5_109601003_1.a5_1;

import javafx.application.Application;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.FontPosture;
import javafx.scene.text.FontWeight;
import javafx.stage.Stage;

import java.io.IOException;

public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) throws IOException {
    Group root = new Group();
    Button btSend = new Button("Send");
    btSend.setLayoutX(350);
    btSend.setLayoutY(425);
    Label label = new Label("Duffy Bear");
    label.setFont(Font.font("Duffy Bear", FontWeight.BOLD, FontPosture.ITALIC, 20));
    label.setLayoutX(425);
    label.setLayoutY(125);
    Image image = new Image("File:src/Duffy.png");
    CheckBox checkBox = new CheckBox("Duffy Duffy Duffy ~~~");
    checkBox.setLayoutX(425);
    checkBox.setLayoutY(200);
    Label label2 = new Label("Is Duffy cute?");
    label2.setFont(Font.font("Duffy Duffy Duffy ~~~", FontWeight.BOLD, FontPosture.ITALIC, 20));
    label2.setLayoutX(425);
    label2.setLayoutY(250);

    ToggleGroup groupRadioButton = new ToggleGroup();
    RadioButton radioButton1 = new RadioButton("Yes");
    radioButton1.setToggleGroup(groupRadioButton);
    radioButton1.setTextFill(Color.RED);
    radioButton1.setLayoutX(425);
    radioButton1.setLayoutY(300);
    RadioButton radioButton2 = new RadioButton("Sure");
    radioButton2.setTextFill(Color.GREEN);
    radioButton2.setLayoutX(500);
    radioButton2.setLayoutY(300);
    radioButton2.setToggleGroup(groupRadioButton);

    root.getChildren().add(btSend);
    root.getChildren().add(label);
    root.getChildren().add(checkBox);
    root.getChildren().add(label2);
    root.getChildren().add(radioButton1);
    root.getChildren().add(radioButton2);

    ImageView imageView2 = new ImageView(image);
    imageView2.setFitWidth(320);
    imageView2.setFitHeight(180);
    imageView2.setLayoutX(50);
    imageView2.setLayoutY(150);
    root.getChildren().add(imageView2);

    Scene scene = new Scene(root, 750, 500);
    stage.setTitle("Hello World");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) {
    launch();
  }
}