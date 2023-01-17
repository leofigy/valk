use gtk::prelude::*;
use gtk::{self, glib, Application, ApplicationWindow, Button, Orientation, Box};
use glib::clone;
use std::rc::Rc;
use std::cell::Cell;

const APP_ID: &str =  "xyz.leofigy.helper";

fn main() {
    println!("Hello, world!");
    let app = Application::builder().application_id(APP_ID).build();
    app.connect_activate(build_ui);
    app.run();
}

fn build_ui(app: &Application){

    // pieces
    let button_increase = Button::builder()
        .label("increase")
        .margin_top(12)
        .margin_bottom(12)
        .margin_start(12)
        .margin_end(12)
        .build();

    let button_decrease = Button::builder()
        .label("decrease")
        .margin_top(12)
        .margin_bottom(12)
        .margin_start(12)
        .margin_end(12)
        .build();

    
    let number = Rc::new(Cell::new(0));

    button_increase.connect_clicked(clone!(@weak number, @weak button_decrease =>
        move |_| {
            number.set(number.get() + 1);
            button_decrease.set_label(&number.get().to_string());
    }));
    button_decrease.connect_clicked(clone!(@weak button_increase =>
        move |_| {
            number.set(number.get() - 1);
            button_increase.set_label(&number.get().to_string());
    }));

    let boxed = Box::builder()
    .orientation(Orientation::Vertical)
    .build();

    boxed.append(&button_increase);
    boxed.append(&button_decrease);

    let window = ApplicationWindow::builder()
    .application(app)
    .title("Helper app")
    .child(&boxed)
    .build();

    window.present();
}