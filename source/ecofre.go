package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew("com.test", glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err := gtk.BuilderNewFromFile("ecofre-main.glade")
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_window_main_destroy":    onMainWindowDestroy,
			"button_key_add_clicked_cb": clickedTestButton,
		}
		builder.ConnectSignals(signals)

		// Get the object with the id of "window_main".
		obj, err := builder.GetObject("window_main")
		fmt.Println(reflect.TypeOf(obj))
		errorCheck(err)

		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		win, err := isWindow(obj)
		fmt.Println(reflect.TypeOf(win))
		errorCheck(err)

		// Show the Window and all of its components.
		win.Show()
		application.AddWindow(win)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_window_main_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}

func clickedTestButton() {
	showMessageInfo("Test!")
	showMessageError("Test Error!")
}

func showMessageInfo(message string) {
	fmt.Println(message)

	dialog := gtk.MessageDialogNew(nil,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		message)
	dialog.Run()
	dialog.Destroy()
}

func showMessageError(message string) {
	fmt.Println(message)

	dialog := gtk.MessageDialogNew(nil,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_OK,
		message)
	dialog.Run()
	dialog.Destroy()
}
