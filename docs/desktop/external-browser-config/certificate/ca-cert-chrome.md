> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate/ca-cert-chrome

ProfessionalCommunity Edition

# Installing Burp's CA certificate in Chrome

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 To test applications in your own browser over HTTPS, you need to install Burp Suite's CA certificate.


#### Note

 These steps are only necessary if you want to use your own external browser for manual testing with Burp. If you prefer, you can just use [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser), which is preconfigured to work with Burp Proxy already. To access Burp's browser, go to the **Proxy > Intercept** tab, and click **Open Browser**.


 Before you install Burp's CA certificate:


-
 Make sure that the [proxy listener](https://portswigger.net/burp/documentation/desktop/external-browser-config/check-listener) is active.

-  [Configure your browser to work with Burp](https://portswigger.net/burp/documentation/desktop/external-browser-config).


 The process to install Burp's CA certificate for use with Chrome is different for each operating system. Follow the relevant process to install the CA certificate:


- [MacOS](https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate/ca-cert-chrome-macos).
- [Windows](https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate/ca-cert-chrome-windows).
- [Linux](https://portswigger.net/burp/documentation/desktop/external-browser-config/certificate/ca-cert-chrome-linux).
