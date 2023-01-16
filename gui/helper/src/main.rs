use gtk::prelude::*;
use gtk::{Application, ApplicationWindow};

const APP_ID: &str =  "xyz.leofigy.helper";

fn main() {
    println!("Hello, world!");
    let app = Application::builder().application_id(APP_ID).build();
    app.connect_activate(build_ui);
    app.run();
}

fn build_ui(app: &Application){
    let window = ApplicationWindow::builder()
    .application(app)
    .title("Helper app")
    .build();

    window.present();
}