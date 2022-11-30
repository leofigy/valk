#include <gtk/gtk.h>



static void click_cb(GtkButton *btn, gpointer user_data){
    g_print ("Clicked.\n");
}


static void app_activate(GApplication *app, gpointer *user_data){
    g_print("Gtk app is activated\n");
    GtkWidget *win;
    GtkWidget *btn;

    win = gtk_application_window_new(
      GTK_APPLICATION(app)
    );

    gtk_window_set_title (GTK_WINDOW (win), "Wizard");
    gtk_window_set_default_size(GTK_WINDOW(win), 600, 600);


    btn = gtk_button_new_with_label("Hello pal");
    gtk_window_set_child(GTK_WINDOW(win), btn);

    g_signal_connect(
      btn,
      "clicked",
      G_CALLBACK(click_cb),
      NULL
    );

    gtk_widget_show(win);
}

int main (int argc, char **argv) {
  GtkApplication *app;
  int stat;

  app = gtk_application_new ("com.github.leofigy.config", G_APPLICATION_FLAGS_NONE);
  g_signal_connect(app, "activate", G_CALLBACK(app_activate), NULL);
  stat =g_application_run (G_APPLICATION (app), argc, argv);
  g_object_unref (app);
  return stat;
}
