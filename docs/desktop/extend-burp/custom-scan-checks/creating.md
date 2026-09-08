> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating

Professional

# Creating custom scan checks

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Custom scan checks enable you to extend Burp Scanner with your own vulnerability detection logic. You can create two types of custom scan checks:

-

**Scripts** - Written in Java with access to our Montoya API. Best if you want to build more complex checks.
-

**BChecks** - Written in our custom BCheck language. Best for quick, lightweight checks.

 To help you get started, we provide the following:

-

Built-in starter templates in the editor.
-

Inline suggestions and error highlighting in the editor.
-

A range of community and reference resources.

#### Related pages

- [Custom scan checks writing guide](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/writing-guide)
- [Passive
 scan check worked example](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/passive-worked-example)
- [Active
 scan check worked example](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/active-worked-example)
-

[Bambda scripts GitHub repository](https://github.com/PortSwigger/bambdas/tree/main/CustomScanChecks) - Examples of custom scan checks written in Java, created by the community and our researchers.
-

[BChecks repository](https://github.com/PortSwigger/BChecks) - Examples of custom scan checks written in our BCheck language, created by the community and our researchers.

#### Warning

 Slow running or resource-intensive scripts can slow down Burp. Write your script carefully to minimize performance impact.


## Creating script-based checks

 To create a new custom scan check using Java:

1.

Go to **Extensions > Custom scan checks**.
1.

Click  **New** and select **Blank script** or
 **From template**.
1.

If you selected **From template**:

  1.

Select the **Script mode** tab.
  1.

Select a template from the list.
  1.

Click  Create using this template.

1.

Select the script **Type**. You can choose from **Active** or **Passive**.
1.

Select when the **Script runs**. You can choose from **Per insertion point**, **Per
 request**, or **Per host**.
1.

[Optional] For scan checks that require out-of-band testing, enable the **Use Collaborator** toggle to generate payloads and customize how Burp handles interaction callbacks. For more information, see [Using Collaborator in checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/writing-guide#using-collaborator-in-checks).
1.

Write the script in Java. For more information, see [Custom scan checks writing guide](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating/writing-guide).
1.

Click **Validate**. Any errors are shown in the **Errors** panel. You must resolve these before you can use your scan check. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).
1.

[Optional] Test the script against real HTTP messages. For instructions, see [Testing custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/testing).
1.

Click **Save & close**.

 The check is saved to your custom scan checks library for use in scans and across projects.

## Creating BCheck-based checks

 To create a custom scan check using our custom BChecks language:

1.

Go to **Extensions > Custom scan checks**.
1.

Click  **New** and select either **Blank BCheck** or
 **From template**.
1.

If you selected **From template**:

  1.

Select the **BCheck mode** tab.
  1.

Select a template from the list.
  1.

Click  **Create using this template**.

1.

Write the script in our BCheck language. For reference documentation, see [BCheck definitions](https://portswigger.net/burp/documentation/scanner/bchecks).
1.

Click **Validate**. Any errors are shown in the Errors panel. You must resolve these before you can use your scan check.
1.

[Optional] To standardize the indentation and whitespace, right-click the editor and select **Format BCheck**.
1.

[Optional] Test the BCheck against real HTTP messages. For more information, see [Testing custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/testing).
1.

Click **Save & close**.

 The check is saved to your custom scan checks library for use in scans and across projects.

#### Related pages

-

For instructions on how to use custom scan checks in your scans, see [Adding custom scan checks to scans](https://portswigger.net/burp/documentation/desktop/running-scans/custom-checks).
-

To get feedback, showcase your work, and connect with other developers, share your custom scan check on our
 [PortSwigger Discord](https://discord.com/invite/portswigger) #bambdas or #bchecks channel.
-

To learn how to export your custom scan checks so that you can share them with others, see [Exporting custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/managing#exporting-custom-scan-checks).
-

To share your custom scan checks with the community, add them to our ever-growing GitHub repositories:

  -

For information on submitting script-based custom scan checks to our Bambda scripts GitHub repository, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).
  -

For information on submitting BCheck-based custom scan checks to our BChecks GitHub repository, see [Submitting BChecks to the community](https://portswigger.net/burp/documentation/scanner/bchecks/contribute-bchecks).
