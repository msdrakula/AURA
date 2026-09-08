> Source: https://portswigger.net/burp/documentation/desktop/running-scans/custom-checks

Professional

# Adding custom scan checks to scans

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can use custom scan checks for web applications or API-only scans. This enables you to tailor scans to your specific needs.

 Any scan checks saved to your custom scan checks library are automatically enabled. They run alongside Burp Scanner's built-in checks during the audit.

#### Related pages

 For more information on how to create or import custom scan checks, see [Custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks).


 To disable or enable custom scan checks in a scan configuration:

1.

In the scan launcher, go to the **Scan configuration** tab.
1.

From the dropdown, select the type of scan configuration you want to use.
1.

Under **Audit configuration**, select **Scan checks**.
1.

In the settings panel, go to the **Custom** tab.
1.

Do one of the following:

  -

Use the **Enabled** toggle switch to enable or disable every custom check at once.
  -

Use the checkboxes to disable or enable specific scan checks.

 Burp Scanner runs all the enabled custom checks when auditing.

## Managing the custom scan checks table

 Custom scan checks are listed in a table with the following information:

-

**Enabled** - Shows whether the scan check is currently enabled.
-

**Name** - The name of the scan check.
-

**Type** - The type of scan check, either **Active**, **Passive** or
 **Unspecified**.
-

**Check runs** - How often the check runs, either **Per insertion point**, **Per
 request**, or **Per host**.

 You can adjust the table contents as follows:

-

**Search the table** - Enter text in the search bar.
-

**Filter the table** - Click **Active** or **Passive** to filter the table by check type.
-

**Customize and sort the table** - For instructions, see [Customizing Burp's tables](https://portswigger.net/burp/documentation/desktop/burps-layout/customize-tables).

 Checks that aren't saved to your custom scan checks library are marked with an asterisk `*`. This may happen if you load a scan configuration that includes custom checks. To add the checks to your library, right-click and select **Save to library**.
