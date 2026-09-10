//go:build linux && cgo

package desktop

/*
#cgo pkg-config: gtk+-3.0 webkit2gtk-4.1
#include <gtk/gtk.h>
#include <webkit2/webkit2.h>
#include <stdlib.h>

static GtkWidget *meb_window = NULL;

static gboolean meb_quit_idle(gpointer data) {
	(void)data;
	if (meb_window != NULL) {
		gtk_window_close(GTK_WINDOW(meb_window));
	} else {
		gtk_main_quit();
	}
	return G_SOURCE_REMOVE;
}

static void meb_on_destroy(GtkWidget *widget, gpointer data) {
	(void)widget;
	(void)data;
	meb_window = NULL;
	gtk_main_quit();
}

static void meb_on_title(WebKitWebView *web, GParamSpec *pspec, gpointer data) {
	(void)pspec;
	GtkWindow *win = GTK_WINDOW(data);
	const gchar *page = webkit_web_view_get_title(web);
	if (page != NULL && page[0] != '\0') {
		gtk_window_set_title(win, page);
	}
}

static gboolean meb_on_load_failed(WebKitWebView *web, WebKitLoadEvent ev, gchar *uri, GError *error, gpointer data) {
	(void)web;
	(void)ev;
	(void)data;
	g_printerr("aura: webview failed to load %s: %s\n", uri, error && error->message ? error->message : "unknown");
	return FALSE;
}

static void meb_on_load_changed(WebKitWebView *web, WebKitLoadEvent ev, gpointer data) {
	(void)web;
	(void)data;
	if (ev == WEBKIT_LOAD_FINISHED) {
		g_printerr("aura: webview loaded\n");
	}
}

int meb_ui_run(const char *title, const char *url, const char *icon_path, int width, int height) {
	g_set_prgname("aura");
	g_set_application_name("AURA");
	gdk_set_program_class("aura");

	if (!gtk_init_check(NULL, NULL)) {
		return 1;
	}
	if (icon_path != NULL && icon_path[0] != '\0') {
		gtk_window_set_default_icon_from_file(icon_path, NULL);
		gtk_window_set_default_icon_name("aura");
	}

	WebKitWebsiteDataManager *dm = webkit_website_data_manager_new_ephemeral();
	WebKitWebContext *wctx = webkit_web_context_new_with_website_data_manager(dm);
	webkit_web_context_set_sandbox_enabled(wctx, FALSE);
	webkit_web_context_clear_cache(wctx);

	meb_window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
	gtk_window_set_title(GTK_WINDOW(meb_window), title);
	gtk_window_set_default_size(GTK_WINDOW(meb_window), width, height);
	gtk_window_resize(GTK_WINDOW(meb_window), width, height);
	gtk_window_set_position(GTK_WINDOW(meb_window), GTK_WIN_POS_CENTER);
	gtk_window_set_role(GTK_WINDOW(meb_window), "aura");
	if (icon_path != NULL && icon_path[0] != '\0') {
		gtk_window_set_icon_from_file(GTK_WINDOW(meb_window), icon_path, NULL);
		gtk_window_set_icon_name(GTK_WINDOW(meb_window), "aura");
	}
	g_signal_connect(meb_window, "destroy", G_CALLBACK(meb_on_destroy), NULL);

	GtkWidget *web = webkit_web_view_new_with_context(wctx);
	gtk_widget_set_hexpand(web, TRUE);
	gtk_widget_set_vexpand(web, TRUE);
	WebKitSettings *settings = webkit_web_view_get_settings(WEBKIT_WEB_VIEW(web));
	webkit_settings_set_enable_javascript(settings, TRUE);
	webkit_settings_set_javascript_can_access_clipboard(settings, TRUE);
	webkit_settings_set_enable_developer_extras(settings, TRUE);
	webkit_settings_set_hardware_acceleration_policy(settings, WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER);
	g_signal_connect(web, "notify::title", G_CALLBACK(meb_on_title), meb_window);
	g_signal_connect(web, "load-failed", G_CALLBACK(meb_on_load_failed), NULL);
	g_signal_connect(web, "load-changed", G_CALLBACK(meb_on_load_changed), NULL);

	gtk_container_add(GTK_CONTAINER(meb_window), web);
	gtk_widget_show_all(meb_window);
	gtk_window_present(GTK_WINDOW(meb_window));
	webkit_web_view_load_uri(WEBKIT_WEB_VIEW(web), url);
	gtk_main();
	return 0;
}

void meb_ui_quit(void) {
	g_idle_add(meb_quit_idle, NULL);
}
*/
import "C"

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"unsafe"
)

func nativeAvailable() bool { return true }

// Native reports whether this build can open a GTK/WebKit window.
func Native() bool { return nativeAvailable() }

func prepareWebKitEnv() {
	// Blank WebKitGTK windows on Kali/NVIDIA/VMs are usually GPU compositing.
	if os.Getenv("WEBKIT_DISABLE_COMPOSITING_MODE") == "" {
		_ = os.Setenv("WEBKIT_DISABLE_COMPOSITING_MODE", "1")
	}
	if os.Getenv("WEBKIT_DISABLE_DMABUF_RENDERER") == "" {
		_ = os.Setenv("WEBKIT_DISABLE_DMABUF_RENDERER", "1")
	}
	if os.Getenv("WEBKIT_DISABLE_SANDBOX") == "" {
		_ = os.Setenv("WEBKIT_DISABLE_SANDBOX", "1")
	}
}

// RunNative opens a GTK/WebKit window on the current (main) thread and blocks until it closes.
func RunNative(ctx context.Context, title, uiURL string) error {
	runtime.LockOSThread()
	prepareWebKitEnv()
	cTitle := C.CString(title)
	cURL := C.CString(uiURL)
	cIcon := C.CString(findIconFile())
	defer C.free(unsafe.Pointer(cTitle))
	defer C.free(unsafe.Pointer(cURL))
	defer C.free(unsafe.Pointer(cIcon))

	done := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			C.meb_ui_quit()
		case <-done:
		}
	}()

	code := C.meb_ui_run(cTitle, cURL, cIcon, 1440, 900)
	close(done)
	if code != 0 {
		return fmt.Errorf("gtk display is not available")
	}
	return nil
}
