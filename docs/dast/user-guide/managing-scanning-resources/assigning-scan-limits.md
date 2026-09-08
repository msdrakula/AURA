> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/assigning-scan-limits

DAST

# Assigning scan limits

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 To avoid placing too much load on your scanning machines, you can limit the maximum number of concurrent scans that they can run. For newly added scanning machines, the default scan limit is 1. If the scanning machine has enough RAM and CPU cores, you can increase the scan limit. For more information, refer to the [System requirements](https://portswigger.net/burp/documentation/dast/setup/self-hosted/system-req-overview).


#### Note

This page explains how to configure fixed scanning machines for a standard instance of Burp Suite DAST. For information on how to use cloud-based auto-scaling scanning, see the [Managing auto-scaling scan resources](https://portswigger.net/burp/documentation/dast/user-guide/managing-scanning-resources/self-hosted/kubernetes) page.

 If you decide to change the scan limits, increase the number slowly and monitor the effect this has on performance.


 To assign scan limits, do the following steps:


1.
 From the settings menu  select **Scanning resources**.

1.
 Under **Scanning machines**, click **Manage scanning machines**.

1.
 Make sure the **Scanning machines** tab is selected.

1.
 To change the **Concurrent scan limit**, click the plus  or minus  icons.


 If you have a Classic license for a specific number of concurrent scans, go to the **Licensing** tab to view this total number, and the remaining number of concurrent scans that you are licensed to run.


## Additional scans

 If you have a Classic license for a specific number of concurrent scans, you can increase the number of concurrent scans at any time from [your account page](https://portswigger.net/users/youraccount) on `portswigger.net`.
