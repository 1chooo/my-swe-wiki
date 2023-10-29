/**
 * Assignment: 11
 * Student Number: 109601003
 * Course: CE1002-A
 */

package stu_109601003.a11;

import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.stage.Stage;
import java.io.IOException;

public class A11 extends Application {
  public static Stage currentStage;
  public static Scene menuScene;
  public static Scene greedyScene;

  @Override
  public void start(Stage primaryStage) throws IOException {
    FXMLLoader fxmlLoader1 = new FXMLLoader(A11.class.getResource("menu.fxml"));
    FXMLLoader fxmlLoader2 = new FXMLLoader(A11.class.getResource("greedy.fxml"));

    currentStage = primaryStage;
    menuScene = new Scene(fxmlLoader1.load());
    greedyScene = new Scene(fxmlLoader2.load());
    currentStage.setTitle("Greedy Snake");
    currentStage.setScene(menuScene);
    currentStage.show();
  }
  public static void main(String[] args) {
    launch();
  }
}