package dialog

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <stdlib.h>
// static GtkWidget* msg_dialog_new(GtkWindow *parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, char *msg) {
// 	return gtk_message_dialog_new(parent, flags, type, buttons, "%s", msg);
// }
import "C"

import (
	"fmt"
	"unsafe"
)

type DlgBuilder struct {
	title   string
	message string
	msgType C.GtkMessageType
	buttons []buttonDef
}

type buttonDef struct {
	label    string
	response C.GtkResponseType
}

func init() {
	if C.gtk_init_check(nil, nil) != C.TRUE {
		panic("GTK initialization failed")
	}
}

func Question(format string, args ...any) *DlgBuilder {
	return &DlgBuilder{
		message: fmt.Sprintf(format, args...),
		msgType: C.GTK_MESSAGE_QUESTION,
	}
}

func (b *DlgBuilder) Title(format string, args ...any) *DlgBuilder {
	b.title = fmt.Sprintf(format, args...)
	return b
}

func (b *DlgBuilder) OKButton(label string) *DlgBuilder {
	b.buttons = append(b.buttons, buttonDef{label, C.GTK_RESPONSE_OK})
	return b
}

func (b *DlgBuilder) CancelButton(label string) *DlgBuilder {
	b.buttons = append(b.buttons, buttonDef{label, C.GTK_RESPONSE_CANCEL})
	return b
}

func (b *DlgBuilder) Run() bool {
	hasOK := false
	hasCancel := false

	for _, button := range b.buttons {
		if button.response == C.GTK_RESPONSE_OK {
			hasOK = true
		} else if button.response == C.GTK_RESPONSE_CANCEL {
			hasCancel = true
		}
	}

	if !hasOK {
		b.OKButton("OK")
	}
	if !hasCancel {
		b.CancelButton("Cancel")
	}
	return b.run() == C.GTK_RESPONSE_OK
}

func (b *DlgBuilder) run() C.gint {
	title := C.CString(b.title)
	defer C.free(unsafe.Pointer(title))
	message := C.CString(b.message)
	defer C.free(unsafe.Pointer(message))

	dialog := C.msg_dialog_new(nil,
		C.GtkDialogFlags(0),
		b.msgType,
		C.GTK_BUTTONS_NONE,
		message)

	C.gtk_window_set_title((*C.GtkWindow)(unsafe.Pointer(dialog)), title)

	dlgPtr := (*C.GtkDialog)(unsafe.Pointer(dialog))
	for _, buttonDef := range b.buttons {
		label := C.CString(buttonDef.label)
		defer C.free(unsafe.Pointer(label))
		button := C.gtk_dialog_add_button(dlgPtr, label, C.gint(buttonDef.response))
		if buttonDef.response == C.GTK_RESPONSE_CANCEL {
			C.gtk_widget_grab_focus(button)
		}
	}
	defer closeDialog(dialog)

	return C.gtk_dialog_run(dlgPtr)
}

func closeDialog(dialog *C.GtkWidget) {
	C.gtk_widget_destroy(dialog)

	for C.gtk_events_pending() != 0 {
		C.gtk_main_iteration()
	}
}
