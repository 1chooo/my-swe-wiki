/**
 * Assignment 9
 * Course: CE1002
 * Student Number: 109601003
 */

package stu_109601003.a9;

import javafx.application.Application;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.ComboBox;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import javafx.scene.input.MouseEvent;
import javafx.scene.text.Font;
import javafx.stage.Stage;

public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) {
    Group root = new Group();
    Font f1 = new Font(24);

    Currency USD = new Currency("美元");
    Currency TWD = new Currency("台幣");
    Currency JPY = new Currency("日幣");
    Currency EUR = new Currency("歐元");
    Currency CNY = new Currency("人民幣");

    TextField textField = new TextField();
    ComboBox<String> country1 = new ComboBox<>();
    ComboBox<String> country2 = new ComboBox<>();
    Button arrow = new Button("↔");
    Button convert = new Button("轉換");
    Label status = new Label();

    textField.setLayoutX(20);
    textField.setLayoutY(25);
    textField.setPrefWidth(460);

    country1.setLayoutX(20);
    country1.setLayoutY(60);
    country1.getItems().addAll(USD.getName(), TWD.getName(), EUR.getName(), JPY.getName(), CNY.getName());
    country1.setValue("選擇幣別");

    country2.setLayoutX(175);
    country2.setLayoutY(60);
    country2.getItems().addAll(USD.getName(), TWD.getName(), EUR.getName(), JPY.getName(), CNY.getName());
    country2.setValue("選擇幣別");

    arrow.setLayoutX(132.5);
    arrow.setLayoutY(60);
    arrow.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {
      String tmp = country1.getValue();
      country1.setValue(country2.getValue());
      country2.setValue(tmp);
    });

    convert.setLayoutX(20);
    convert.setLayoutY(95);

    convert.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {

      String inStr = textField.getText();
      float num = Float.parseFloat(inStr);
      String temp1 = country1.getValue();
      String temp2 = country2.getValue();

      float countNum = switch (temp1) {
        case "美元" -> (num);
        case "台幣" -> (float) (num / 29.42);
        case "日幣" -> (float) (num / 124.819687);
        case "歐元" -> (float) (num / 0.913381);
        case "人民幣" -> (float) (num / 6.347357);
        default -> 0;
      };

      float ans = switch (temp2) {
        case "美元" -> (countNum);
        case "台幣" -> (float) (countNum * 29.42);
        case "日幣" -> (float) (countNum * 124.819687);
        case "歐元" -> (float) (countNum * 0.913381);
        case "人民幣" -> (float) (countNum * 6.347357);
        default -> 0;
      };

      int flag = 0;
      String strAns = String.valueOf(ans);

      for (int i = 0; i < strAns.length(); i++) {
        if (strAns.charAt(i) == '.') {
          flag = i;
        }
      }

      if (strAns.length() > flag + 3) {
        strAns = strAns.substring(0, flag + 4);
      }

      String s = inStr + " " + temp1 + " = " + strAns + " " + temp2;
      status.setText(s);
    });

    status.setLayoutX(25);
    status.setLayoutY(145);
    status.setFont(f1);

    root.getChildren().addAll(textField, arrow, convert, status);
    root.getChildren().addAll(country1, country2);
    Scene scene = new Scene(root, 500, 200);
    stage.setTitle("匯率轉換器");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) {
    launch();
  }
}

class Currency {
  private String Name;

  public Currency(String Name) {
    this.Name = Name;
  }

  public String getName() {
    return Name;
  }
}