use gtk::prelude::*;
use gtk::{Application, ApplicationWindow, Button};

const APP_ID: &str =  "xyz.leofigy.helper";

fn main() {
    println!("Hello, world!");
    let app = Application::builder().application_id(APP_ID).build();
    app.connect_activate(build_ui);
    app.run();
}

fn build_ui(app: &Application){

    // pieces
    let button = Button::builder()
        .label("Click on me")
        .margin_top(12)
        .margin_bottom(12)
        .margin_start(12)
        .margin_end(12)
        .build();

    
    button.connect_clicked(
        |button| {
            button.set_label("Welcome|Wilkommen|Bienvenido");
        }
    );


    let window = ApplicationWindow::builder()
    .application(app)
    .title("Helper app")
    .child(&button)
    .build();

    window.present();
}