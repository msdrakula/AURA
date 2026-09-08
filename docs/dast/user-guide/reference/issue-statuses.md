> Source: https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-statuses

DAST

# Issue statuses

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST tracks each issue across your scans and gives it a status. The status tells you where an issue is in its lifecycle, so you can see what the scanner has found and record your team's own assessment.


 There are two kinds of status:


- **Scanner statuses** - set automatically by the scanner, based on what it finds in each scan.
- **Your statuses** - set by you, to record how your team has triaged the issue.

## Scanner statuses

 The scanner sets these statuses automatically. You cannot change them.


- **New** - The scanner found this issue for the first time.
- **Present** - The scanner found this issue again in the latest scan.
- **Regressed** - The issue was previously resolved, but the scanner has found it again.
- **Fixed (confirmed)** - The scanner re-tested the issue's location and confirmed that it is no longer present.
- **Inconclusive** - The scanner could not conclusively test the issue, so its status is unconfirmed.

## Your statuses

 You can set one of the following statuses to record your team's assessment of an issue. For the steps, see [Managing issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues).


- **False positive** - The issue is not a genuine vulnerability.
- **Accepted risk** - The issue is a genuine vulnerability, but you have decided not to act on it. For example, it is mitigated by other security measures.
- **Verified risk** - You have confirmed that the issue is a genuine vulnerability that needs remediation.
- **Fixed (unconfirmed)** - You believe the issue is fixed, but the scanner has not yet confirmed it.

## How scanner and your statuses work together

 An issue has a single status at any time. Burp Suite DAST decides which status to show as follows:


- If you have set a status, the issue shows your status.
- If you have not set a status, the issue shows its scanner status from the latest scan.
- When the scanner confirms that an issue is fixed, this takes precedence. The issue shows **Fixed (confirmed)**, even if you set a different status.
- When the scanner cannot conclusively re-test an issue, the issue shows **Inconclusive**, unless you have set a different status.

 To clear a status you have set and return an issue to its scanner status, select **Use scanner status**.


## What a status applies to

 You set a status on a single issue, on one site. Burp Suite DAST keeps that issue as a lasting record, so the status you set stays with it in every later scan of the site.


 A status does not apply to any other issue:


- If the scanner finds the same type of vulnerability at a different location on the same site, that is a separate issue with its own status.
- If the scanner finds the same vulnerability on a different site, that is also a separate issue. Marking it on one site does not mark it on another.

 To give the same status to several issues at once, including issues on different sites, use bulk editing. For more information, see [Updating issues in bulk](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues#updating-issues-in-bulk).


#### Related pages

- [Managing issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/managing-issues)
- [Tracking issues over time](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/track-issues)
- [Issue details](https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details)
