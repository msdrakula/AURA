> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/managing

Professional

# Managing custom scan checks in your library

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 You can store and manage custom scan checks in the **Extensions > Custom scan checks** tab. This enables you to reuse your checks across your Burp projects. To add checks to the library you can create and save your own or import them from our GitHub repositories.

 Custom scan checks are listed in a table with the following information:

-

**Name** - The name of the check.
-

**Author** - The author of the check.
-

**Tags** - Any tags that are applied to the custom scan check.

#### Note

 The **Tags** and **Author** columns are only populated for custom scan checks written in the BChecks language. They are automatically populated from the definition. To modify them, edit the check's definition directly.


 To search the table, enter text in the search bar. You can also customize and sort the table, and copy column data to your clipboard. For more information, see [Customizing Burp's tables](https://portswigger.net/burp/documentation/desktop/burps-layout/customize-tables).

#### Related pages

 For instructions on how to use custom scan checks in your scans, see [Adding custom scan checks to scans](https://portswigger.net/burp/documentation/desktop/running-scans/custom-checks).


## Importing custom scan checks

 You can import custom scan checks that have been shared with you or downloaded from our [Bambda scripts](https://github.com/PortSwigger/bambdas/tree/main/CustomScanChecks) repository or [BChecks](https://github.com/PortSwigger/BChecks) repository. For more information, see [Importing custom scan checks into your library](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/importing).

#### Warning

 Custom scan checks can run arbitrary code. For security reasons, please be cautious when importing and using scan checks from unverified sources.


## Creating custom scan checks

 You can create checks directly from the custom scan checks library. For more information, see [Creating custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating).

## Exporting custom scan checks

 You can export custom scan checks to import them into other instances of Burp.

 To export checks:

1.

In the library, select the checks you want to export.
1.

Click **Export** .
1.

Select a directory.
1.

Save the checks to the directory:

  -

Single check: Enter a filename then click **Save**.
  -

Multiple checks: Click **Save** to export all selected checks into a folder. Each check is saved with its current name.

 The checks are exported with the correct file extension:

-

`.bcheck` files for BCheck-based checks
-

`.bambda` files for script-based checks

 Each `.bambda` file includes metadata (unique ID, name, function, and location). Burp uses the unique ID to match checks, so always preserve it when sharing to avoid conflicts.

 Burp matches `.bcheck` files by filename, so keep filenames consistent when sharing.

#### Related pages

 For more information about how the unique ID is used when re-importing checks, see [Importing custom scan checks - Updating your checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/importing#updating-your-custom-scan-checks).


 You can share your checks with the community by adding them to our ever-growing GitHub repositories.


 For information on submitting script custom scan checks to our Bambda scripts GitHub repository, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).


 For information on submitting BCheck custom scan checks to our BChecks GitHub repository, see [Submitting BChecks to the community](https://portswigger.net/burp/documentation/scanner/bchecks/contribute-bchecks).


 For more information about our custom BCheck language, see [BCheck definitions](https://portswigger.net/burp/documentation/scanner/bchecks).


## Managing custom scan checks

 You can perform the following actions on the checks in your library:

-

**Edit** - Select the check that you want to edit, then click .
-

**Copy** - Select the check that you want to copy, then click .
-

**Delete** - Select the check that you want to delete, then click .
