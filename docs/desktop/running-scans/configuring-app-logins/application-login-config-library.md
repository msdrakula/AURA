> Source: https://portswigger.net/burp/documentation/desktop/running-scans/configuring-app-logins/application-login-config-library

Professional

# Managing application logins using the configuration library

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Burp Suite's configuration library enables you to store sets of login credentials and recorded login sequences so that you can use them in later scans. You can save and load application logins from the scan launcher.


 To load application logins from the library, click **Select from library** and select the required login. Burp Suite adds the selected credentials or recorded login sequence to the list.


 To save all of the scan's application logins to the configuration library:


1. Click **Save to library** to display the **Save configuration to library** dialog.
1. Enter a **Configuration name**.
1. Click **Save**.

 Burp Suite saves all of the credentials or recorded login sequences in the list to the configuration library. If you select this configuration for a different scan then Burp Suite adds the application logins to that scan.


#### Note

 You can also add application logins directly to the configuration library from the **Settings** dialog. For more information, see [Configuration library](https://portswigger.net/burp/documentation/desktop/settings/library).
