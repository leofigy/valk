#ifndef TP_H
#define TP_H

#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui { class Tp; }
QT_END_NAMESPACE

class Tp : public QMainWindow
{
    Q_OBJECT

public:
    Tp(QWidget *parent = nullptr);
    ~Tp();

private:
    Ui::Tp *ui;
};
#endif // TP_H
