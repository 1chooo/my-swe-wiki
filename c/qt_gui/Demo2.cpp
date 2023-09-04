#include <QtGui>
#include "Demo2.h"
 
Demo2::Demo2(QWidget *parent) : QWidget(parent) {
    // 將成員變數屬性初始化
    b1 = new QLabel(tr("QLabel"));
    b2 = new QPushButton(tr("QPushButton"));
    b3 = new QCheckBox(tr("QCheckBox"));
    b4 = new QRadioButton(tr("QRadioButton"));
    b5 = new QLineEdit;
     
    // 建立版面樣式物件
    QVBoxLayout *layout = new QVBoxLayout;
    layout->addWidget(b1);
    layout->addWidget(b2);
    layout->addWidget(b3);
    layout->addWidget(b4);
    layout->addWidget(b5);
     
    // 設定版面樣式與視窗標題
    setLayout(layout);
    setWindowTitle(tr("Demo2"));
}