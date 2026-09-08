> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/browser-config-safari

ProfessionalCommunity Edition

# Configuring Safari to work with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You need to configure Safari so that you can use it for testing with Burp Suite.


#### Note

 These steps are only necessary if you want to use an external browser for manual testing with Burp. If you prefer, you can just use [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser), which is preconfigured to work with Burp Proxy already. To access Burp's browser, go to the **Proxy > Intercept** tab, and click **Open Browser**.


 To configure Safari, follow these steps:


1.

 In Safari, go to the **Safari** menu and click **Preferences**.


![Safari preferences](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_safari_1.jpg)

1.

 Click the **Advanced** tab and, under **Proxies**, click the **Change Settings** button. This will open the network configuration settings for your current network adapter.


![Safari advanced settings](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_safari_1.5.jpg)

1.
 In the **Proxies** tab, check the **Web Proxy (HTTP)** box and enter your Burp Proxy listener address in the **Web Proxy Server** field (by default, `127.0.0.1`), and your Burp Proxy listener port in the (unlabelled) port field (by default, `8080`).

1.

 Repeat these steps for the **Secure Web Proxy (HTTPS)** checkbox.


![Safari Proxy settings](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_safari_2.jpg)

1.
 Ensure the **Bypass proxy settings for these Hosts & Domains** box is empty.

1.

 Click **OK** and **Apply** and close the open dialogs.


![Safari clear bypass settings](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_safari_3.jpg)

#### Next step

- [Check your browser proxy configuration](https://portswigger.net/burp/documentation/desktop/external-browser-config/check-browser-configuration).
