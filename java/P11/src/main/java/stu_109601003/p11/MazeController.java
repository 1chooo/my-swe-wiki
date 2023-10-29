package stu_109601003.p11;

import javafx.event.EventHandler;
import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.input.KeyCode;
import javafx.scene.input.KeyEvent;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.Pane;

public class MazeController implements EventHandler<KeyEvent> {
  int rowIndex = 0;
  int columnIndex = 0;

  @FXML
  GridPane maze;
  @FXML
  Pane man;
  @FXML
  Label successLabel;
  @FXML
  Button backButton;

  @FXML
  public void onBackPressed() {
    P11.currentStage.setScene(P11.menuScene);
  }

  @Override
  public void handle(KeyEvent event) {
    System.out.println(event.getCode());

    if (event.getCode() == KeyCode.UP)
      rowIndex--;
    if (event.getCode() == KeyCode.DOWN)
      rowIndex++;
    if (event.getCode() == KeyCode.LEFT)
      columnIndex--;
    if (event.getCode() == KeyCode.RIGHT)
      columnIndex++;
    if (event.getCode() == KeyCode.SPACE) {
      rowIndex = 0;
      columnIndex = 0;
      successLabel.setVisible(false);
    }

    walkToNewPosition();

    if (rowIndex == 4 && columnIndex == 4) {
      successLabel.setVisible(true);
    }
  }

  private void walkToNewPosition() {
    int maxRowIndex = maze.getRowCount() - 1;
    int maxColumnIndex = maze.getColumnCount() - 1;

    if (rowIndex < 0)
      rowIndex = 0;
    if (columnIndex < 0)
      columnIndex = 0;
    if (columnIndex > maxColumnIndex)
      columnIndex = maxColumnIndex;
    if (rowIndex > maxRowIndex)
      rowIndex = maxRowIndex;

    GridPane.setRowIndex(man, rowIndex);
    GridPane.setColumnIndex(man, columnIndex);
  }
}