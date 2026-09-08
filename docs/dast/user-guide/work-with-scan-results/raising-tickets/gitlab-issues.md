> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/gitlab-issues

DAST

# Raising GitLab issues from within Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If your admin user has [configured a GitLab integration](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-gitlab), you can raise GitLab issues from within Burp Suite DAST for any issues found by a scan.


1.
 From the top menu, select **Scans**.

1.
 Select the scan you want to view.

1.
 Select the **Issues** tab.

1.
 Expand the issue and select the URL from the list.

1.
 In the upper-right corner of the page, click the **Raise GitLab issue** button.


#### Note

 If you have also integrated Burp Suite DAST with other issue-tracking platforms, you may need to select this from the **Raise ticket** drop-down.

1.
 You can create a new GitLab issue, or link to an existing GitLab issue:


  -
 To create a new issue, select your project and ticket type from the drop-down menu, then click **Create**.

  -
 To link to an existing GitLab issue, select **Link to existing issue**, choose the relevant project, and enter the relevant GitLab issue number. Then click **Link**.


## Raising GitLab issues for multiple issues

 You can raise GitLab issues for more than one issue at a time. For more information, refer to [Raising tickets for multiple issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/raise-tickets-multiple).
