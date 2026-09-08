> Source: https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/monitor-scan-progress

DAST

# Monitoring scan progress

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 The **Timeline** shows you the progress of a scan, so that you can see what stage it's at and how much time remains.


 To monitor the scan progress:


1.
 Click **Scans** on the top menu to display the list of scans.

1.
 To see which scans are running, set the **Status** column header to . If a scan is running, a scanning icon appears in this column.

1.
 Select a running scan.

1.

Select the **Timeline** tab. Notice that the following information is shown:

  -
 The current scan **Status**.

  -  **Start time**, **Duration**, and an estimate of the **Time remaining**.

  -
 The progress bar, which indicates the status of the scan.

  -
 Any error messages, if applicable.


#### Note

 The time remaining for your scan depends on several factors, including your chosen [scan configuration](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scan-configurations/custom-configs).


![Progress bar](https://portswigger.net/burp/documentation/dast/images/scan-progress-bar.png)

 If you scan a web app, Burp Scanner performs the crawl and audit phases at the same time. The crawl phase isn't required for API scans.

#### Note

 You can also review the timeline for completed scans. This can help you to troubleshoot any errors that may have occurred.


## Reviewing scan errors

 The timeline shows error messages when applicable. The error messages include information about the cause of the issue and possible remedies.


 To review scan errors:


1.
 Open the **Timeline** tab.

1.
 Notice that errors are indicated on the progress bar by a red cross .

1.
 To review an error, scroll down to the stage where the error occurred.

1.
 Follow the remediation advice for the error and run the scan again.


 For further help, refer to the [Troubleshooting documentation](https://portswigger.net/burp/documentation/dast/user-guide/troubleshooting).
