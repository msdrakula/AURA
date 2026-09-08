> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/viewing-issues

DAST

# Viewing issue details

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 This page explains how you can see more detailed information about the issues detected by Burp Suite DAST. You can also:

-  [Set an issue's status](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#setting-an-issue-status)
-  [Edit the severity of an issue](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#editing-issue-severity)
-  [Raise a ticket to remediate an issue](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets)

## View details for a specific issue

 To view a detailed description of an issue:


1.
 From the top menu, select **Scans**.

1.
 Select the scan you want to view.

1.
 Select the **Issues** tab.

1.
 Expand the issue and select the URL from the list.

1.
 Make sure that the **Advisory** tab is selected. This panel contains detailed information about the issue.

1.
 To see more information about the issue, expand the collapsible headings. Note that the headings vary depending on the type of issue detected.

1.
 To view other issues for this scan, select them from the list on the left of the window.


## Viewing the timeline for an issue

 The **Timeline** tab shows the full history of an issue.


 To view an issue's timeline:


1.
 Open the issue, as described above.

1.
 Select the **Timeline** tab.

1.
 (Optional) Filter the timeline by event type.


 For more information about the events shown in the timeline, see [Issue details](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details).


#### Note

 The timeline is only available on an issue record, which you can open at the global, folder, or site level. An issue viewed from a scan is a snapshot of that issue as it was found during the scan, so it does not have a timeline.


#### Related pages

-  [Viewing AI-enhanced scan results](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/ai-enhanced-results).

-  [Raising tickets](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets).

-  [Managing issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues).

-  [Issue details](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details).
