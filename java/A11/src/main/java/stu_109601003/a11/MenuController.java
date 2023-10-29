package stu_109601003.a11;

import javafx.fxml.FXML;

import java.io.IOException;

public class MenuController {
  @FXML
  public void onStartPressed() throws IOException {
    A11.greedyScene.getRoot().requestFocus();
    A11.currentStage.setScene(A11.greedyScene);
  }

  @FXML
  public void onExitPressed() throws IOException {
    A11.currentStage.close();
  }
}