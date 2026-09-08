> Source: https://portswigger.net/burp/documentation/desktop/running-scans/extension-checks

Professional

# Adding extension scan checks to scans

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can use scan checks created by extensions for web applications or API-only scans.

 When you load an extension, any scan checks it creates are automatically enabled in the scan launcher. They run alongside Burp Scanner's built-in checks during the audit.

#### Related pages

[Installing extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing)

 You can disable or enable extension scan checks in bulk:

1.

In the scan launcher, go to the **Scan configuration** tab.
1.

From the dropdown, select the type of scan configuration you want to use.
1.

Under **Audit configuration**, select **Scan checks**.
1.

In the settings panel, go to the **Extension** tab.
1.

Use the **Enabled** toggle to disable or enable every extension scan check at once.

 Burp Scanner runs all the enabled custom checks when auditing.

#### Note

 To enable or disable scan checks from a specific extension, you must disable the extension itself. For more information, see [Managing extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/managing-extensions).


## Managing the extension scan checks table

 Extension scan checks are listed in a table with the following information:

-

**Name** - The name of the scan check.
-

**Extension** - The extension that created the scan check.
-

**Type** - The type of scan check, either **Active**, **Passive**, or **Unspecified**.
-

**Check runs** - How often the check runs, either **Per insertion point**, **Per
 request**, or **Per host**.

 You can customize and sort the table contents. For more information, see [Customizing Burp's tables](https://portswigger.net/burp/documentation/desktop/burps-layout/customize-tables).
