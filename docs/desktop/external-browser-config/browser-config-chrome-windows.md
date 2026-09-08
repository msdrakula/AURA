> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/browser-config-chrome-windows

ProfessionalCommunity Edition

# Configuring Chrome to work with Burp Suite - Windows

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 If you want to use Chrome with Burp Suite, you need to configure the proxy settings.


#### Note

 These steps are only necessary if you want to use an external browser for manual testing with Burp. If you prefer, you can just use [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser), which is preconfigured to work with Burp Proxy already. To access Burp's browser, go to the **Proxy > Intercept** tab, and click **Open Browser**.


To configure Chrome to work with Burp Suite, follow these steps:

1.
 Open Chrome and go to the **Customize** (hamburger) menu.

1.
 Select **Settings** and open the **System** menu.

1.

 Click **Open your computer's proxy settings**. The **Proxy Settings** window enables you to set up the proxy server.


![Chrome proxy settings](https://portswigger.net/burp/documentation/desktop/images/chrome-proxy-settings.png)

1.
 Make sure that **Automatically detect settings** and **Use setup script** are **Off**.

1.
 Set **Use a proxy server** to **On**.

1.
 Enter your Burp Proxy listener address in the **Address** field (by default, `127.0.0.1`).

1.
 Enter your Burp Proxy listener port in the **Port** field (by default, `8080`).

1.
 Make sure that **Don't use the proxy server for local (intranet) addresses** is unchecked.

1.

 Click **Save**.


![Chrome proxy settings - Windows](https://portswigger.net/burp/documentation/desktop/images/config-chrome-ms.png)

#### Next step

- [Check your browser proxy configuration](https://portswigger.net/burp/documentation/desktop/external-browser-config/check-browser-configuration).
