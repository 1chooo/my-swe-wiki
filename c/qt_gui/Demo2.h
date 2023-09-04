#ifndef DEMO2_H
#define DEMO2_H
 
#include <QWidget>
 
class QLabel;
class QPushButton;
class QCheckBox;
class QRadioButton;
class QLineEdit;
 
// Demo2 繼承自 QWidget
class Demo2 : public QWidget {
    Q_OBJECT
 
public:
    Demo2(QWidget *parent = 0);
 
private:
    QLabel *b1;
    QPushButton *b2;
    QCheckBox *b3;
    QRadioButton *b4;
    QLineEdit *b5;
};
 
#endif