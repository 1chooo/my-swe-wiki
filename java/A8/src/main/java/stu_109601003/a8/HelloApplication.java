package stu_109601003.a8;

import javafx.application.Application;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import javafx.scene.input.Clipboard;
import javafx.scene.input.ClipboardContent;
import javafx.scene.input.MouseEvent;
import javafx.stage.Stage;

import java.io.*;
import java.util.ArrayList;
import java.io.IOException;


public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) {
    Group root = new Group();
    ArrayList<String> myArray = new ArrayList<>();
    Clipboard clipboard = Clipboard.getSystemClipboard();
    ClipboardContent content = new ClipboardContent();

    Label input = new Label("輸入：");
    input.setLayoutX(20);
    input.setLayoutY(32.5);

    Label output = new Label("輸出：");
    output.setLayoutX(20);
    output.setLayoutY(72.5);

    TextField inText = new TextField();
    inText.setLayoutX(70);
    inText.setLayoutY(30);
    inText.setPrefWidth(400);

    TextField outText = new TextField();
    outText.setLayoutX(70);
    outText.setLayoutY(70);
    outText.setPrefWidth(400);

    String s = "";
    Label status = new Label(s);
    status.setLayoutX(25);
    status.setLayoutY(160);

    Button create = new Button("新建");
    create.setLayoutX(60);
    create.setLayoutY(120);
    create.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {
      File file = new File("./109601003.txt");
      try {
        if (!file.exists()) {
          file.createNewFile();
          file.canWrite();
          status.setText("已建立");
        } else {
          status.setText("新建失敗");
        }
      } catch (IOException e) {
        e.printStackTrace();
      }
    });

    Button decode = new Button("編碼");
    decode.setLayoutX(220);
    decode.setLayoutY(120);
    decode.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {
      String inStr = inText.getText();

      if (inStr.equals("")) {
        status.setText("編碼失敗");
      } else {
        char[] temp;
        temp = inStr.toCharArray();

        for (int i = 0; i < inStr.length(); i++) {
          int a = (int) temp[i];
          String s1 = String.valueOf(a);
          myArray.add(s1);
        }
        String result = "";
        for (String s1 : myArray) {

          result = result + s1;
        }
        outText.setText(result);
        status.setText("已編碼");
        myArray.clear();
      }
    });

    Button duplicate = new Button("複製");
    duplicate.setLayoutX(300);
    duplicate.setLayoutY(120);
    duplicate.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {
      String dup = outText.getText();

      if (dup.equals("")) {
        status.setText("複製失敗");
      } else {
        content.putString(dup);
        clipboard.setContent(content);
        status.setText("已複製");
      }
    });

    Button save = new Button("存檔");
    save.setLayoutX(140);
    save.setLayoutY(120);
    save.addEventHandler(MouseEvent.MOUSE_CLICKED, MouseEvent -> {
      String sav = outText.getText();

      try {
        File file = new File("./109601003.txt");
        FileWriter fw = new FileWriter(file, true);
        BufferedWriter bufw = new BufferedWriter(fw);
        bufw.write(sav);
//        bufw.flush();
        bufw.newLine();
        bufw.close();
      } catch (IOException e) {
        e.printStackTrace();
      }
      status.setText("已存檔");
    });

    Button clear = new Button("清空");
    clear.setLayoutX(380);
    clear.setLayoutY(120);
    clear.addEventHandler(MouseEvent.MOUSE_CLICKED, mouseEvent -> {
      inText.clear();
      outText.clear();
      status.setText("已清空");
    });

    root.getChildren().addAll(input, output, inText, outText, status);
    root.getChildren().addAll(create, save, decode, duplicate, clear);

    Scene scene = new Scene(root, 500, 200);
    stage.setTitle("編密碼小工具");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) throws IOException {
    launch();
  }
}