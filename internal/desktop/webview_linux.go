//go:build linux && cgo

package desktop

/*
#cgo pkg-config: gtk+-3.0 webkit2gtk-4.1
#include <gtk/gtk.h>
#include <webkit2/webkit2.h>
#include <stdlib.h>

static GtkWidget *meb_window = NULL;

static void meb_on_destroy(GtkWidget *w, gpointer data) {
	(void)w;
	(void)data;
	gtk_main_quit();
}

static gboolean meb_quit_idle(gpointer data) {
	(void)data;
	if (meb_window != NULL) {
		gtk_widget_destroy(meb_window);
		meb_window = NULL;
	} else {
		gtk_main_quit();
	}
	return G_SOURCE_REMOVE;
}

static GtkWidget *meb_on_create(WebKitWebView *web, WebKitNavigationAction *action, gpointer data) {
	(void)action;
	(void)data;
	webkit_web_view_load_uri(web, webkit_uri_request_get_uri(webkit_navigation_action_get_request(action)));
	return NULL;
}

void meb_quit_window(void) {
	g_idle_add(meb_quit_idle, NULL);
}

int meb_run_window(const char *url, const char *title, int width, int height) {
	if (!gtk_init_check(NULL, NULL)) {
		return 1;
	}
	g_set_prgname("aura");
	g_set_application_name("AURA");
	gdk_set_program_class("aura");

	meb_window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
	gtk_window_set_title(GTK_WINDOW(meb_window), title);
	gtk_window_set_default_size(GTK_WINDOW(meb_window), width, height);
	gtk_window_set_position(GTK_WINDOW(meb_window), GTK_WIN_POS_CENTER);
	g_signal_connect(meb_window, "destroy", G_CALLBACK(meb_on_destroy), NULL);

	WebKitSettings *settings = webkit_settings_new();
	webkit_settings_set_enable_javascript(settings, TRUE);
	webkit_settings_set_enable_developer_extras(settings, TRUE);
	webkit_settings_set_javascript_can_access_clipboard(settings, TRUE);

	WebKitWebView *web = WEBKIT_WEB_VIEW(webkit_web_view_new_with_settings(settings));
	g_signal_connect(web, "create", G_CALLBACK(meb_on_create), NULL);
	gtk_container_add(GTK_CONTAINER(meb_window), GTK_WIDGET(web));
	webkit_web_view_load_uri(web, url);

	gtk_widget_show_all(meb_window);
	gtk_window_present(GTK_WINDOW(meb_window));
	gtk_main();
	meb_window = NULL;
	return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func nativeSupported() bool { return true }

func runNative(uiURL, title string, width, height int) error {
	cURL := C.CString(uiURL)
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cURL))
	defer C.free(unsafe.Pointer(cTitle))
	if C.meb_run_window(cURL, cTitle, C.int(width), C.int(height)) != 0 {
		return fmt.Errorf("gtk display is not available")
	}
	return nil
}

func quitNative() {
	C.meb_quit_window()
}
