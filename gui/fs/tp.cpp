#include "tp.h"
#include "ui_tp.h"

Tp::Tp(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::Tp)
{
    ui->setupUi(this);
}

Tp::~Tp()
{
    delete ui;
}

