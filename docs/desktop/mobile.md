> Source: https://portswigger.net/burp/documentation/desktop/mobile

ProfessionalCommunity Edition

# Mobile testing

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can use Burp Suite to perform security tests for mobile applications. To do this, you need to configure the mobile device to proxy its traffic via Burp Proxy. This enables you to intercept, view, and modify all the HTTP/S requests and responses processed by the mobile app, and carry out [penetration testing using Burp](https://portswigger.net/burp/documentation/desktop/testing-workflow) in the normal way.


 Successfully intercepting HTTPS traffic from mobile applications can be complex. You may run into difficulties with the proxy configuration, or TLS certificate pinning. For more information, see the [troubleshooting](https://portswigger.net/burp/documentation/desktop/mobile/troubleshooting) section.


## Configuring a mobile device

-  [Configuring an iOS device to work with Burp Suite Professional](https://portswigger.net/burp/documentation/desktop/mobile/config-ios-device).

-  [Configuring an Android device to work with Burp Suite Professional](https://portswigger.net/burp/documentation/desktop/mobile/config-android-device).
