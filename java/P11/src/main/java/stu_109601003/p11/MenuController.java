package stu_109601003.p11;

import javafx.fxml.FXML;

import java.io.IOException;

public class MenuController {
  @FXML
  public void onStartPressed() throws IOException {
    P11.mazeScene.getRoot().requestFocus();
    P11.currentStage.setScene(P11.mazeScene);
  }

  @FXML
  public void onExitPressed() throws IOException {
    P11.currentStage.close();
  }
}
