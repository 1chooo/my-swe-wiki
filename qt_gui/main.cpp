#include <QtGui>
#include "Demo2.h"
 
int main(int argv, char **args)
{
    QApplication app(argv, args);
 
    Demo2 demo;
    demo.show();
 
    return app.exec();
}