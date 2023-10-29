/**
 * Assignment: 10
 * Student Number: 109601003
 * Course: CE1002
 */

package stu_109601003.a10;

import javafx.animation.KeyFrame;
import javafx.animation.Timeline;
import javafx.application.Application;
import javafx.beans.property.DoubleProperty;
import javafx.scene.Scene;
import javafx.scene.input.KeyCode;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.shape.Circle;
import javafx.stage.Stage;
import javafx.util.Duration;

import java.io.IOException;

public class HelloApplication extends Application {
  @Override
  public void start(Stage stage) throws IOException {
    BallPane ballPane = new BallPane();

    Scene scene = new Scene(ballPane, 750, 500);
    scene.setFill(Color.rgb(151, 173, 172));
    stage.setTitle("Bounding Ball Animation");
    stage.setScene(scene);
    stage.show();
  }

  public static void main(String[] args) {
    launch();
  }
}

class BallPane extends Pane {
  public final double radius = 20;
  private double x = radius;
  private double y = radius;
  private double dx = 1;
  private double dy = 1;
  private Circle circle = new Circle(x, y, radius);
  private Timeline animation;

  public BallPane() {
    circle.setFill(Color.rgb(222, 125, 44));
    getChildren().add(circle);
    animation = new Timeline(new KeyFrame(Duration.millis(8), e -> moveBall()));
    animation.setCycleCount(Timeline.INDEFINITE);
    animation.play();
  }

  public void play() {
    animation.play();
  }

  public void pause() {
    animation.pause();
  }

  public void increaseSpeed() {
    animation.setRate(animation.getRate() + 0.1);
  }

  public void decreaseSpeed() {
    animation.setRate(animation.getRate() > 0 ? animation.getRate() - 0.1 : 0);
  }

  public DoubleProperty rateProperty() {
    return animation.rateProperty();
  }

  protected void moveBall() {
    if (x < radius || x > getWidth() - radius) {
      dx *= -1;
    }
    if (y < radius || y > getHeight() - radius) {
      dy *= -1;
    }
    x += dx;
    y += dy;
    circle.setCenterX(x);
    circle.setCenterY(y);
  }
}