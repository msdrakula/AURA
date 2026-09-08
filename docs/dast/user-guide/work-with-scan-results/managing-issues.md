> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues

DAST

# Managing issues in Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 After a scan, you can manage the issues that Burp Suite DAST finds. You can:


- [Set an issue's status](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#setting-an-issue-status).
- [Add comments to issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#adding-a-comment-to-an-issue).
- [Edit the severity of an issue](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#editing-issue-severity).
- [Update many issues at once](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#updating-issues-in-bulk).
- [Download issues as a CSV file](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#downloading-issues-as-a-csv-file).

## Setting an issue status

 After reviewing an issue you can set a status for it, such as **False positive** or **Accepted risk**. If you and the scanner set different statuses, Burp Suite DAST shows the status you set, unless the scanner confirms the issue is fixed. For more information, see [Issue statuses](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses).


 To set an issue's status:


1.
 Go to the **Issues** tab.

1.
 Select an issue to open it.

1.

From the **Status** drop-down, select one of the following:

  - **False positive**
  - **Accepted risk**
  - **Verified risk**
  - **Fixed (unconfirmed)**
  - **Use scanner status** - clears a status you have set and returns the issue to its scanner status.

1.
 (Optional) Add a note.

1.
 Click **Change issue status**.


 Burp Suite DAST records the change in the issue's [timeline](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/viewing-issues#viewing-the-timeline-for-an-issue), along with the time, date, and your username.


 The status stays with the issue in every later scan of the site. It does not apply to the same vulnerability found at another location, or on another site. For more information, see [Issue statuses](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses).


#### Note

 Issues marked as **False positive** are moved to the bottom of the **Issues** tab and excluded from dashboard statistics and charts. Issues marked as **False positive** or **Accepted risk** do not cause CI/CD pipeline builds to fail.


## Adding a comment to an issue

 Add comments to record your decisions and share context with your team. Comments appear in the issue's timeline.


 To add a comment:


1.
 Select an issue to open it.

1.
 Select the **Timeline** tab.

1.
 Enter your comment.

1.
 Click **Add comment**.


## Editing issue severity

 Editing the severity of an issue can be useful if the real-world impact it poses is different from how Burp has rated it.


 To edit the severity of an issue:


1.
 Go to the **Issues** tab.

1.
 Select an issue to open it.

1.
 From the **Severity** drop-down, select a severity.

1.
 (Optional) Add a note.

1.
 Click **Change issue severity**.


 Burp Suite DAST records the change in the issue's [timeline](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/viewing-issues#viewing-the-timeline-for-an-issue), along with the time, date, and your username.


## Updating issues in bulk

 You can update many issues in a single action, across different sites and applications.


 To update issues in bulk:


1.
 Go to the **Issues** tab.

1.
 Select the issues you want to update.

1.
 Click **Change status** or **Change severity**, then select the new value.


 You cannot add a note to a bulk change.


 You can also change an issue's status directly from the issue list, without opening the issue.


## Downloading issues as a CSV file

 You can download the issue list as a CSV file. The file contains the issues that match your current filters.


 To download issues as a CSV file:


1.
 Go to the **Issues** tab on the **Home** page, or on a site or folder.

1.
 Filter the issue list, if necessary.

1.
 Click **Download as CSV**, above the issue list.


 The file lists the following details for each issue:


- Issue type.
- Origin.
- Path.
- Severity.
- Confidence.
- Site name.
- Status.
- Last seen date.

 The file can contain up to 100,000 issues.


#### Related pages

- [Issue statuses](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses)
- [Tracking issues over time](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/track-issues)
- [Best practices for managing false positives](https://portswigger.net/burp/documentation/dast/user-guide/working-with-scans/false-positives-best-practice)
- [MCP server](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/mcp-server)
