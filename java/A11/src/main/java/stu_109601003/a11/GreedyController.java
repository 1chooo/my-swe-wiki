package stu_109601003.a11;

import javafx.event.EventHandler;
import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.input.KeyCode;
import javafx.scene.input.KeyEvent;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.Pane;

public class GreedyController implements EventHandler<KeyEvent> {
  int rowIndexOfHead = 0;
  int columnIndexOfHead = 0;
  int rowIndexOfTail = 0;
  int columnIndexOfTail = 0;
  KeyCode tempKey;
  boolean legalToMove = false;
  boolean win = false;
  boolean lose = false;

  @FXML
  GridPane maze;
  @FXML
  Pane man1;
  @FXML
  Pane man2;
  @FXML
  Label successLabel;
  @FXML
  Label failLabel;
  @FXML
  Button backButton;

  @FXML
  public void onBackPressed() {
    A11.currentStage.setScene(A11.menuScene);
  }

  @Override
  public void handle(KeyEvent event) {
    legalToMove = false;
    System.out.println(event.getCode());

    if (!(rowIndexOfHead == rowIndexOfTail) || !(columnIndexOfHead == columnIndexOfTail)) {
      rowIndexOfTail = rowIndexOfHead;
      columnIndexOfTail = columnIndexOfHead;
    }

    if (event.getCode() == KeyCode.UP) {
      if (tempKey != KeyCode.DOWN) {
        rowIndexOfHead--;
        legalToMove = true;
      }
    }

    if (event.getCode() == KeyCode.DOWN) {
      if (tempKey != KeyCode.UP) {
        rowIndexOfHead++;
        legalToMove = true;
      }
    }

    if (event.getCode() == KeyCode.LEFT) {
      if (tempKey != KeyCode.RIGHT) {
        columnIndexOfHead--;
        legalToMove = true;
      }
    }

    if (event.getCode() == KeyCode.RIGHT) {
      if (tempKey != KeyCode.LEFT) {
        columnIndexOfHead++;
        legalToMove = true;
      }
    }

    if (event.getCode() == KeyCode.SPACE) {
      rowIndexOfHead = 0;
      columnIndexOfHead = 0;
      rowIndexOfTail = 0;
      columnIndexOfTail = 0;
      successLabel.setVisible(false);
      failLabel.setVisible(false);
      legalToMove = true;
      win = false;
      lose = false;
    }

    if (rowIndexOfHead < 0 || columnIndexOfHead < 0 || columnIndexOfHead > 4 || rowIndexOfHead > 4) {
      failLabel.setVisible(true);
      legalToMove = false;
      lose = true;
    }

    if (legalToMove && !win && !lose) {
      tempKey = event.getCode();
      walkToNewPosition();
    }

    if (rowIndexOfHead == 4 && columnIndexOfHead == 4) {
      successLabel.setVisible(true);
      win = true;
    }
  }

  private void walkToNewPosition() {
    GridPane.setRowIndex(man1, rowIndexOfHead);
    GridPane.setColumnIndex(man1, columnIndexOfHead);
    GridPane.setRowIndex(man2, rowIndexOfTail);
    GridPane.setColumnIndex(man2, columnIndexOfTail);
  }
}