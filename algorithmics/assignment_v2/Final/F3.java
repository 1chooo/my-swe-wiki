/**
 * Math have to import
 */

package Final;

import java.util.ArrayList;
import static java.lang.Math.abs;
import java.util.Scanner;

public class F3 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);

    int runTimes = myObjects.nextInt();
    ArrayList<Integer> xPoint = new ArrayList<>();
    ArrayList<Integer> yPoint = new ArrayList<>();

    for ( int i = 0; i < runTimes; i++) {
      int times = myObjects.nextInt();
      for ( int j = 0; j < times; j++) {
        int tempX = myObjects.nextInt();
        int tempY = myObjects.nextInt();
        xPoint.add(tempX);
        yPoint.add(tempY);
      }

      ArrayList<Integer> selectedPointX = new ArrayList<>();
      ArrayList<Integer> selectedPointY = new ArrayList<>();
      ArrayList<Integer> otherPointX = new ArrayList<>();
      ArrayList<Integer> otherPointY = new ArrayList<>();

      MST mst = new MST(times, xPoint, yPoint,
              selectedPointX, selectedPointY,
              otherPointX, otherPointY);
      mst.pickFirstPoint();
      mst.countDistance();
      mst.showAns();

      xPoint.clear();
      yPoint.clear();
    }
  }
}

class MST {
  private int times;
  private ArrayList<Integer> xPoint;
  private ArrayList<Integer> yPoint;
  private ArrayList<Integer> selectedPointX;
  private ArrayList<Integer> selectedPointY;
  private ArrayList<Integer> otherPointX;
  private ArrayList<Integer> otherPointY;

  private int distance;

  public MST(int times, ArrayList<Integer> xPoint, ArrayList<Integer> yPoint,
             ArrayList<Integer> selectedPointX, ArrayList<Integer> selectedPointY,
             ArrayList<Integer> otherPointX, ArrayList<Integer> otherPointY) {
    this.times = times;
    this.xPoint = xPoint;
    this.yPoint = yPoint;
    this.selectedPointX = selectedPointX;
    this.selectedPointY = selectedPointY;
    this.otherPointX = otherPointX;
    this.otherPointY = otherPointY;
  }

  public void pickFirstPoint() {
    selectedPointX.add(xPoint.get(0));
    selectedPointY.add(yPoint.get(0));

    for ( int i = 1; i < xPoint.size(); i++) {
      otherPointX.add(xPoint.get(i));
      otherPointY.add(yPoint.get(i));
    }
//    System.out.println(selectedPointX);
//    System.out.println(otherPointX);
  }

  public void countDistance() {
    while (otherPointX.size() >= 1) {
//      System.out.println(otherPointX);
//      System.out.println(otherPointY);
//      System.out.println(selectedPointX);
//      System.out.println(selectedPointY);
      ArrayList<Integer> allManhattanDistance = new ArrayList<>();
      for ( int i = 0; i < selectedPointX.size(); i++) {
        for ( int j = 0; j < otherPointX.size(); j++) {
          int temp = abs(selectedPointX.get(i) - otherPointX.get(j)) +
                  abs(selectedPointY.get(i) - otherPointY.get(j));
          allManhattanDistance.add(temp);
//          System.out.println(temp);
        }
      }

//      System.out.println(allManhattanDistance);

//      int minDistance = Collections.min(allManhattanDistance);
//      int minDistanceIndex = allManhattanDistance.indexOf(minDistance);


      int minDistance = allManhattanDistance.get(0);
      int minDistanceIndex = 0;
      for ( int i = 1; i < allManhattanDistance.size(); i++ ) {
        if ( allManhattanDistance.get(i) < minDistance ) {
          minDistance = allManhattanDistance.get(i);
          minDistanceIndex = i;
        }
      }

      minDistanceIndex = (minDistanceIndex + 1) % otherPointX.size() - 1;
      if ( minDistanceIndex == -1) {
        minDistanceIndex = otherPointX.size() - 1;
      }

      selectedPointX.add(otherPointX.get(minDistanceIndex));
      selectedPointY.add(otherPointY.get(minDistanceIndex));
      otherPointX.remove(minDistanceIndex);
      otherPointY.remove(minDistanceIndex);

      distance += minDistance;
//      System.out.println(minDistance);
    }
  }

  public void showAns() {
    System.out.println(distance);
  }
}