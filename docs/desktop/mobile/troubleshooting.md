> Source: https://portswigger.net/burp/documentation/desktop/mobile/troubleshooting

ProfessionalCommunity Edition

# Troubleshooting for mobile devices

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 In this section, we'll provide some simple steps you can take to resolve common issues when trying to test mobile apps using Burp Suite.


#### I can't access HTTPS URLs on iOS even after installing Burp's CA certificate

 In Burp, go to the **Dashboard** tab and look at the event log. If you see a large number of `Connection reset` or `Remote host terminated the handshake` errors, the problem may be due to compatibility issues with TLS 1.3.


 You can prevent the proxy from using TLS 1.3 by disabling it in the proxy listener settings as follows:


1.

 In Burp, click on **Settings ** to open the **Settings** dialog.

1.

 Go to the **Tools > Proxy** tab.

1.

 Select the proxy listener that you use for your mobile device and click **Edit**.

1.

 In the **Edit proxy listener** dialog, go to the **TLS Protocols** tab.

1.

 Select **Use custom protocols**, then deselect **TLSv1.3** from the list.


 If this was the problem, you should now be able to access HTTPS URLs as normal and will see fewer errors in the event log. If you're still having issues, please contact our [support team](https://portswigger.net/support).
