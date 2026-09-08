> Source: https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/generate-logs

DAST

# Downloading logs and debug packs

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This section describes how to download event logs and debug packs.


## Downloading the event log

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 The event log can be useful for debugging. To download the event log in CSV format:


1.
 Open the **Scans** tab and select a scan.

1.
 Select the **Logging** tab.

1.
 Click **Download event log**.


## Downloading the scan debug pack

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you require assistance, our Support team may ask you for the scan debug pack.


 To download the scan debug pack:


1.
 Open the **Scans** tab and select a scan.

1.
 Go to **Logging > Scan debug pack**.

1.
 Click **Download**.


## Downloading the verbose debug pack

 In some circumstances, our Support team may ask you for the verbose debug pack. It contains a Burp Suite DAST project file that holds a scan's data and configuration settings.


 This is only available to download for scans that were performed with verbose logging enabled.


 To download the verbose debug pack:


1. Open the **Scans** tab and select a scan.
1. Go to the **Logging** tab.
1. Click **Download verbose debug pack**.

#### Note

You can enable user activity logging to record user actions throughout Burp Suite. You can then download this log as a CSV file that contains a list of timestamped actions and successful logins.


For more information, see [User activity log](https://portswigger.net/burp/documentation/dast/user-guide/reference/user-activity-logs).
