/**
 * Assignment 7
 * Student Number: 109601003
 * Course: CE1002
 */

package stu_109601003.a7;

import javafx.application.Application;
import javafx.geometry.Pos;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.TextField;
import javafx.scene.input.MouseEvent;
import javafx.scene.text.Font;
import javafx.stage.Stage;

import java.util.ArrayList;

public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) {
    Group root = new Group();
    Font font = new Font(16);
    Font font1 = new Font(24);

    Button btBack = new Button("<-");       // first row
    btBack.setFont(font);
    btBack.setLayoutX(4);
    btBack.setLayoutY(396);
    btBack.setPrefSize(120, 50);

    Button btZero = new Button("0");
    btZero.setFont(font);
    btZero.setLayoutX(128);
    btZero.setLayoutY(396);
    btZero.setPrefSize(120, 50);

    Button btDot = new Button(".");
    btDot.setFont(font);
    btDot.setLayoutX(252);
    btDot.setLayoutY(396);
    btDot.setPrefSize(120, 50);

    Button btEqual = new Button("=");
    btEqual.setFont(font);
    btEqual.setLayoutX(376);
    btEqual.setLayoutY(396);
    btEqual.setPrefSize(120, 50);

    Button btOne = new Button("1");         // second row
    btOne.setFont(font);
    btOne.setLayoutX(4);
    btOne.setLayoutY(342);
    btOne.setPrefSize(120, 50);

    Button btTwo = new Button("2");
    btTwo.setFont(font);
    btTwo.setLayoutX(128);
    btTwo.setLayoutY(342);
    btTwo.setPrefSize(120, 50);

    Button btThree = new Button("3");
    btThree.setFont(font);
    btThree.setLayoutX(252);
    btThree.setLayoutY(342);
    btThree.setPrefSize(120, 50);

    Button btClear = new Button("C");
    btClear.setFont(font);
    btClear.setLayoutX(376);
    btClear.setLayoutY(342);
    btClear.setPrefSize(120, 50);

    Button btFour = new Button("4");         // third row
    btFour.setFont(font);
    btFour.setLayoutX(4);
    btFour.setLayoutY(288);
    btFour.setPrefSize(120, 50);

    Button btFive = new Button("5");
    btFive.setFont(font);
    btFive.setLayoutX(128);
    btFive.setLayoutY(288);
    btFive.setPrefSize(120, 50);

    Button btSix = new Button("6");
    btSix.setFont(font);
    btSix.setLayoutX(252);
    btSix.setLayoutY(288);
    btSix.setPrefSize(120, 50);

    Button btSqrt = new Button("^");
    btSqrt.setFont(font);
    btSqrt.setLayoutX(376);
    btSqrt.setLayoutY(288);
    btSqrt.setPrefSize(120, 50);

    Button btSeven = new Button("7");         // forth row
    btSeven.setFont(font);
    btSeven.setLayoutX(4);
    btSeven.setLayoutY(234);
    btSeven.setPrefSize(120, 50);

    Button btEight = new Button("8");
    btEight.setFont(font);
    btEight.setLayoutX(128);
    btEight.setLayoutY(234);
    btEight.setPrefSize(120, 50);

    Button btNine = new Button("9");
    btNine.setFont(font);
    btNine.setLayoutX(252);
    btNine.setLayoutY(234);
    btNine.setPrefSize(120, 50);

    Button btPercent = new Button("%");
    btPercent.setFont(font);
    btPercent.setLayoutX(376);
    btPercent.setLayoutY(234);
    btPercent.setPrefSize(120, 50);

    Button btPlus = new Button("+");         // fifth row
    btPlus.setFont(font);
    btPlus.setLayoutX(4);
    btPlus.setLayoutY(180);
    btPlus.setPrefSize(120, 50);

    Button btMinus = new Button("-");
    btMinus.setFont(font);
    btMinus.setLayoutX(128);
    btMinus.setLayoutY(180);
    btMinus.setPrefSize(120, 50);

    Button btMultiply = new Button("*");
    btMultiply.setFont(font);
    btMultiply.setLayoutX(252);
    btMultiply.setLayoutY(180);
    btMultiply.setPrefSize(120, 50);

    Button btDivide = new Button("/");
    btDivide.setFont(font);
    btDivide.setLayoutX(376);
    btDivide.setLayoutY(180);
    btDivide.setPrefSize(120, 50);

    TextField textField = new TextField();
    textField.setStyle("-fx-font-family: 'monospaced';");
    textField.setFont(font1);
    textField.setAlignment(Pos.CENTER_RIGHT);
    textField.setLayoutX(15);
    textField.setLayoutY(50);
    textField.setPrefSize(470, 80);

    btBack.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> {
      int len = textField.getLength();
      if (len == 0) {
        int stop = 1;
      } else {
        textField.deleteText((len - 1), len);
      }
    });
    btZero.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("0"));
    btDot.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("."));
    btEqual.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> {
      ArrayList<Integer> operatorsSite = new ArrayList<>();
      ArrayList<Character> operators = new ArrayList<>();

      String s = textField.getText();

      for (int i = 0; i < s.length(); i++) {
        char ch = s.charAt(i);
        if (ch == '^') {
          operatorsSite.add(i);
        } else if (ch == '*') {
          operatorsSite.add(i);
        } else if (ch == '/') {
          operatorsSite.add(i);
        } else if (ch == '%') {
          operatorsSite.add(i);
        } else if (ch == '+') {
          operatorsSite.add(i);
        } else if (ch == '-') {
          operatorsSite.add(i);
        }
      }

      int init = 0;
      ArrayList<Float> numList = new ArrayList<>();
      String num;

      for (int i = 0; i < s.length(); i++) {
        for (Integer operator : operatorsSite) {
          if (i == operator) {
            num = s.substring(init, i);
            init = i + 1;
            numList.add(Float.parseFloat(num));
          }
        }
        if (i == s.length() - 1) {
          num = s.substring(init);
          numList.add(Float.parseFloat(num));
        }
      }

      for (int index : operatorsSite) {
        char c = s.charAt(index);
        operators.add(c);
      }

      while (true) {
        int site;
        float temp;
        if (operators.contains('^')) {
          site = operators.indexOf('^');
          temp = (float) Math.pow(numList.get(site), numList.get(site + 1));
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else if (operators.contains('*')) {
          site = operators.indexOf('*');
          temp = numList.get(site) * numList.get(site + 1);
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else if (operators.contains('/')) {
          site = operators.indexOf('/');
          temp = numList.get(site) / numList.get(site + 1);
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else if (operators.contains('%')) {
          site = operators.indexOf('%');
          temp = numList.get(site) % numList.get(site + 1);
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else if (operators.contains('+')) {
          site = operators.indexOf('+');
          temp = numList.get(site) + numList.get(site + 1);
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else if (operators.contains('-')) {
          site = operators.indexOf('-');
          temp = numList.get(site) - numList.get(site + 1);
          numList.set(site, temp);
          numList.remove(site + 1);
          operators.remove(site);
          continue;
        } else {
          String ans = (String.valueOf(numList.get(0)));
          int site1 = 0;
          for (int i = 0; i < ans.length(); i++) {
            char ch1 = ans.charAt(i);
            if (ch1 == '.') {
              site1 = i;
            }
          }
          if (ans.length() > site1 + 5) {
            ans = ans.substring(0, site1 + 5);
          }
          textField.setText(ans);
          break;
        }
      }
    });

    btOne.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("1"));
    btTwo.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("2"));
    btThree.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("3"));
    btClear.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.setText(""));

    btFour.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("4"));
    btFive.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("5"));
    btSix.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("6"));
    btSqrt.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("^"));

    btSeven.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("7"));
    btEight.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("8"));
    btNine.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("9"));
    btPercent.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("%"));

    btPlus.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("+"));
    btMinus.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("-"));
    btMultiply.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("*"));
    btDivide.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> textField.appendText("/"));

    root.getChildren().addAll(btBack, btZero, btDot, btEqual);
    root.getChildren().addAll(btOne, btTwo, btThree, btClear);
    root.getChildren().addAll(btFour, btFive, btSix, btSqrt);
    root.getChildren().addAll(btSeven, btEight, btNine, btPercent);
    root.getChildren().addAll(btPlus, btMinus, btMultiply, btDivide);
    root.getChildren().add(textField);

    Scene scene = new Scene(root, 500, 450);
    stage.setTitle("計算機");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) {
    launch();
  }
}