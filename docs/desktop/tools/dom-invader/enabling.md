> Source: https://portswigger.net/burp/documentation/desktop/tools/dom-invader/enabling

ProfessionalCommunity Edition

# Enabling DOM Invader

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 DOM Invader is preinstalled in Burp's browser, but is disabled by default as some of its features may interfere with your other testing activities.


 To enable DOM Invader:


1.

 Go to the **Proxy > Intercept** tab and open Burp's browser.

1.

 In the upper-right corner of the browser window, click the Burp Suite logo. If you can't see this logo, click the jigsaw icon first. A panel opens containing tabs for the Burp Suite Navigation Recorder and DOM Invader settings menu.


![Enabling DOM invader](https://portswigger.net/burp/documentation/desktop/images/dom-invader-settings-main.png)

1.

 From the DOM Invader settings, toggle the switch so that **DOM Invader is on**.

1.

 Click **Reload** to refresh the browser. This is necessary for your changes to take effect.

1.

 Right-click anywhere in the main browser window and select **Inspect** to open the browser's DevTools panel. Note that this now contains the **DOM Invader** tab. For the best experience, we recommend docking the panel to the bottom of the browser window.


![Docking the DevTools panel to the bottom of the browser window](https://portswigger.net/burp/documentation/desktop/images/dom-invader-innerhtml-sink.png)

#### Note

 By default, DOM Invader remembers your previous settings, including whether it was on or off. Keep this in mind if you close Burp's browser while DOM Invader is still enabled. To disable this behavior, go to **Settings > Tools > Burp's browser** and deselect **Store settings and history after closing**.
