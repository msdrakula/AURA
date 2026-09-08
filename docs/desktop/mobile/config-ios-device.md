> Source: https://portswigger.net/burp/documentation/desktop/mobile/config-ios-device

ProfessionalCommunity Edition

# Configuring an iOS device to work with Burp Suite Professional

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can test web applications and mobile apps using an iOS device. To do this, you need to do the following:


-
 Configure your Burp Proxy listener to accept connections on all network interfaces.

-
 Connect both your device and your computer to the same wireless network.

-
 To interact with HTTPS traffic, you need to install a CA certificate on your iOS device.


## Step 1: Configure the Burp Proxy listener

 To configure the proxy settings for Burp Suite Professional:


1.
 Open Burp Suite Professional click Settings  to open the
 **Settings** dialog.

1.
 Go to **Tools > Proxy**.

1.
 In **Proxy listeners**, click **Add**.

1.
 In the **Binding** tab, set **Bind to port** to `8082` (or another port that is not in use).

1.

 Select **All interfaces** and click **OK**.


![Add a proxy listener](https://portswigger.net/burp/documentation/desktop/images/add-proxy-listener.png)

1.
 At the prompt, click **Yes**.


## Step 2: Configure your device to use the proxy

 To configure the proxy settings for your iOS device:


1.
 In your iOS device, go to **Settings > Wi-Fi**.

1.
 Make sure that the Wi-Fi button is on and connect to your Wi-Fi network.

1.

 Select the information icon (**i**) next to your Wi-Fi network.


![iOS Wi-Fi settings](https://portswigger.net/burp/documentation/desktop/images/ios-wifi.png)

1.
 Set **Configure Proxy** to **Manual**.

1.
 Set **Server** to the IP address of the computer that is running Burp Suite Professional.

1.
 Set **Port** to the port value that you configured for the Burp Proxy listener, in this example `8082`.

1.

 Touch **Save** .


![iOS proxy configuration](https://portswigger.net/burp/documentation/desktop/images/ios-proxy-configuration.png)

## Step 3: Install a CA certificate on your iOS device

 In order to interact with HTTPS traffic, you need to install a CA certificate from your Burp Suite Professional installation on your iOS device.


 To install the CA certificate to your iOS device:


1.
 Make sure that Burp Suite Professional is running on your computer.

1.
 Use the browser on your iOS device to go to `http://burpsuite` and select **CA Certificate**.

1.
 When the CA certificate downloads, select **Profile downloaded** in the **Settings** menu.

1.

 On the **Install Profile** screen, select **Install**.


![iOS Install Profile screen](https://portswigger.net/burp/documentation/desktop/images/ios-install-profile.png)

1.
 On the **Installing Profile** screen, select **Install**.

1.
 When the profile is installed, select **Done**.

1.
 Go to **Settings > General > About > Certificate Trust Settings**.

1.

 Activate the toggle switch for `Portswigger CA`.


![iOS Certificate Trust Settings](https://portswigger.net/burp/documentation/desktop/images/ios-cert-trust.png)

## Step 4: Test the configuration

 To test the configuration:


1.
 Open Burp Suite Professional.

1.
 Go to **Proxy > Intercept** and click **Intercept is off** to switch intercept on.

1.
 Open the browser on your iOS device and go to an HTTPS web page.


 The page should load without any security warnings. You should see the corresponding requests within Burp Suite Professional.
