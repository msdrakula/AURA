> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/check-browser-configuration

ProfessionalCommunity Edition

# Checking your browser proxy configuration

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Follow the steps below to check your browser proxy configuration.


#### Note

 These steps are only necessary if you want to use an external browser for manual testing with Burp. If you prefer, you can just use [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser), which is preconfigured to work with Burp Proxy already. To access Burp's browser, go to the **Proxy > Intercept** tab, and click **Open Browser**.


1.
 Make sure you have [checked that the proxy listener is active](https://portswigger.net/burp/documentation/desktop/external-browser-config/check-listener) and have configured your chosen browser.

1.
 In Burp Suite, go to the **Proxy > Intercept** tab. To activate HTTP interception, click **Intercept is off**.

1.
 With Burp Suite running, open the browser that you configured and go to any HTTP URL. Your browser should sit waiting for the request to complete, because Burp Suite has intercepted the HTTP request that your browser is trying to send.

1.
 In Burp Suite, go to the **Proxy > Intercept** tab. On the **Intercept** tab, the intercepted HTTP request displays in the main panel.

1.
 Click **Forward** to release the request from Burp Suite.

1.
 Go back to your browser. You should now see the requested page loading as it would during normal browsing.

1.
 To deactivate HTTP interception, click **Intercept is on**.


 These steps enable you to test web applications that use HTTP. To test HTTPS URLs, you need to [install Burp's CA certificate](https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate).


#### Next step

- [Install Burp's CA certificate](https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate).
