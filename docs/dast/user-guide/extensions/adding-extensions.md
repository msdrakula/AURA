> Source: https://portswigger.net/burp/documentation/dast/user-guide/extensions/adding-extensions

DAST

# Adding extensions to Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 When you add extensions to Burp Suite DAST, they are uploaded to your **Extension library**.
 Users can then apply extensions from this central repository on a site-by-site basis for them to be used during scans.

## Prerequisite permissions for adding extensions

 Only users with the `Manage extensions` permission can add extensions to the library. Initially, this is only assigned to the built-in `Administrator` role.

#### Warning

 Be careful when granting this permission to additional users. During a scan, extensions run on your scanning machine with the permissions of the `burpsuite` OS user. Therefore, there is a potential security risk if someone inadvertently uploads a fake extension created by a malicious third party.


## Adding BApps to Burp Suite DAST

 To add a BApp:

1.

 Download the BApp from the [BApp Store](https://portswigger.net/bappstore?products=enterprise). Make sure that it is compatible with Burp Suite DAST - you can filter the store to make this easier.

1.

 Log in to Burp Suite DAST as a user with permission to manage extensions.

1.

 From the settings menu , select **Extensions** to open the **Extension library**.

1.

 On the **BApp extensions** tab, click **Upload BApp**.

1.

 Select the `.bapp` file that you downloaded from the BApp Store.


 The extension is now in your **Extension library**. Your users can [apply the extension to specific sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scanning-with-extensions) to use it during scans.

## Adding custom extensions to Burp Suite DAST

 If you're proficient in Java, you can create your own custom extensions for Burp Suite DAST. Learn more about [Creating Burp
 extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/creating).

### Requirements for extensions

To use an extension with Burp Suite DAST, it needs to meet the following conditions:

-
 The extension is written in Java 21 or lower.

-
 The extension doesn't require user interaction, or the use of the user interface.

-
 The extension doesn't use features that are exclusive to Burp Suite Professional or Community Edition, such as Repeater, Intruder, or Proxy.


### Adding a custom extension

 To add a custom extension:

1.

 Log in to Burp Suite DAST as a user with permission to manage extensions.

1.

 From the settings menu , select **Extensions** to open the **Extension library**.

1.

 On the **Custom extensions** tab, click **Upload extension**.

1.

 Select the JAR file for the extension.

1.

 Enter a name and description for the extension, then click **Add**.


 The extension is now in your **Extension library**.
 Your users can [apply the extension to specific sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scanning-with-extensions) to use it during scans.

## Adding BChecks to Burp Suite DAST

 You can download BChecks created by PortSwigger, and by the Burp Suite community, from the [BChecks GitHub repository](https://github.com/PortSwigger/BChecks).

 If you have access to Burp Suite Professional, you can also create your own custom scan checks, enabling you to target your scans and make your testing workflow as efficient as possible.
 For more information, see [Creating custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating).

 To add a BCheck:

1.

 Log in to Burp Suite DAST as a user with permission to manage extensions.

1.

 From the settings menu , select **Extensions** to go to the **Extension library**.

1.

 On the **BChecks** tab, click **Upload BCheck**.

1.

 Select the BCheck you want to upload.


 Files that you want to import should be in plain text format with a `.bcheck` extension.


 The extension is now in your **Extension library**. Your users can [apply the extension to specific sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scanning-with-extensions) to use it during scans.
