> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/retest-issues

DAST

# Re-testing issues

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can re-test issues to confirm that they are fixed, without waiting for a full scan. Burp Suite DAST replays the relevant requests and updates the status of issues, depending on the outcome.


 To re-test an issue:


1.
 Go to the **Issues** tab.

1.
 Select the issue you want to re-test.

1.
 Click **Re-test issue**.


 When the re-test finishes, Burp Suite DAST updates the issue's scanner status and records the result in the issue's [timeline](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details).


## Automatic resolution confirmation

 You do not have to re-test issues manually to confirm a fix. When a later scan covers a location where an issue was previously found, Burp Suite DAST automatically re-tests that issue. If the issue is no longer present, its scanner status changes to **Fixed (confirmed)**, with evidence.


 Issues are only confirmed as fixed when the scanner has positively re-tested them. If issues don't appear in a scan they are not treated as resolved.


## Re-test results

 Each re-test is recorded in the issue's timeline. The record includes:


- The date.
- Who triggered the re-test.
- The outcome.

 A re-test has one of the following outcomes:


- **Fixed (confirmed)** - The issue is no longer present at its location.
- **Present** - The issue is still present.

#### Related pages

- [Issue statuses](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses)
- [Managing issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues)
- [Issue details](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details)
