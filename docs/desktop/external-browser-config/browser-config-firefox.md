> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/browser-config-firefox

ProfessionalCommunity Edition

# Configuring Firefox to work with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You need to configure Firefox so that you can use it for testing with Burp Suite.


#### Note

 These steps are only necessary if you want to use an external browser for manual testing with Burp. If you prefer, you can just use [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser), which is preconfigured to work with Burp Proxy already. To access Burp's browser, go to the **Proxy > Intercept** tab, and click **Open Browser**.


 To configure Firefox, follow these steps:


1.

 In Firefox, go to the Firefox Menu and select **Preferences > Options**.


![Firefox browser preferences](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_ff_1.jpg)

1.

 Select the **General** tab and scroll to the **Network Proxy** settings. Click the **Settings** button.


![Firefox network settings](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_ff_2.jpg)

1.

 Select the **Manual proxy configuration** option.

1.
 Enter your Burp Proxy listener address in the **HTTP Proxy** field (by default this is set to `127.0.0.1`).

1.

 Enter your Burp Proxy listener port in the **Port** field (by default, `8080`). Make sure the **Use this proxy server for all protocols** box is checked.


![Firefox proxy configuration](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_ff_3.jpg)

1.

 Delete anything that appears in the **No proxy for** field. Click **OK** to close all the options dialogs.


![Firefox clear "no proxy for" field](https://portswigger.net/burp/documentation/desktop/images/support/installingandconfiguring_configuringyourbrowser_ff_4.jpg)

#### Next step

- [Check your browser proxy configuration](https://portswigger.net/burp/documentation/desktop/external-browser-config/check-browser-configuration).
