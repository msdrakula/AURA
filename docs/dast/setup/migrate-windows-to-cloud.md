> Source: https://portswigger.net/burp/documentation/dast/setup/migrate-windows-to-cloud

DAST

# Migrating from Windows to Cloud

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

 This guide explains how to use our migration tool to move data from your Windows self-hosted instance to PortSwigger's secure cloud.


 You can use our migration tool to transfer:


-

Sites and folders
-

Custom scan configurations
-

Extensions
-

BChecks
-

Groups and roles (but not users)
-

Scheduled recurring scans
-

Tags

## Prerequisites

 Make sure your self-hosted instance is using the latest software version.


-

To check your version, click the  icon in the upper-right corner of the screen and select **About**.
-

If you need to update your instance, download the installer from the [release page](https://portswigger.net/burp/releases.html#dast).

## Export your data

 To export your data:


1.

Run the following command from the Burp Suite DAST installation folder: `exportDataForSaasMigration`
1.

You will see a warning message. Enter `y` to continue.
1.

Make a note of where the output file is saved.

 The data is exported as a ZIP file.


## Import your data

 Use the GraphQL API to import your data:


1.

Create an API user and API key in the Burp Suite DAST UI. For more information, see [Creating API users](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user).
1.
 Save the API user's API key somewhere safe.

1.
 Open **Command Prompt** on the source instance.

1.
 Navigate to the Burp Suite DAST installation directory.

1.

Run the following command, replacing the required fields in the brackets: `importDataForSaasMigration.exe ^
   --cloud-url=<cloud-url> ^
   --api-key=<api-key-from-cloud-instance> ^
   --data-file="<path-to-export-data-file.zip>" ^
   --have-permission-to-scan=<y-or-n> ^
   --resume-previous-import=<y-or-n>`

Enter the required fields as follows:

-

`cloud-url` The URL of the instance of Burp Suite DAST that is importing the data. For example, `https://my-company.portswigger.cloud`.
-

`api-key` The API key of the user that is importing the data.
-

`data-file` The path and filename of the export zip file, created by the export tool.

#### Note

 The import only works on clean cloud instances. If your cloud instance already contains data, the tool may stop with errors (for example, due to duplicate scan configurations or extensions).


## Scan schedule behavior after migration

 Once you migrate to PortSwigger's secure cloud, your scan schedules behave as follows:


-

One-off scans are not exported.
-

Recurring scans continue, based on their schedule.
-

If a recurring scan was set to repeat a certain number of times, this behavior will not change.
-

Custom scan configurations are marked as created by the API user who runs the import.

#### Related pages

-

[Migrating from Linux to Cloud](https://portswigger.net/burp/documentation/dast/setup/migrate-linux-to-cloud)
-

[Setting up a Cloud instance](https://portswigger.net/burp/documentation/dast/setup/cloud)
