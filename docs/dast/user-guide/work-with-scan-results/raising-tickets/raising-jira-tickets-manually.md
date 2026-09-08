> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/raising-jira-tickets-manually

DAST

# Raising Jira tickets manually

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If your admin user has [configured a Jira integration](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-jira), you can raise Jira tickets for issues directly from your scan results.


1.
 From the top menu, select **Scans**.

1.
 Select the scan you want to view.

1.
 Select the **Issues** tab.

1.
 Expand the issue and select the URL from the list.

1.

 In the upper-right corner of the page, click the **Raise Jira Ticket** button.


#### Note

 If you have also integrated Burp Suite DAST with other issue-tracking platforms, you may need to select this from the **Raise ticket** drop-down.

1.
 You can create a new Jira ticket, or link to an existing Jira ticket:


  -
 To create a new Jira ticket, select a rule from the drop-down menu and click **Create**.

  -
 To link to an existing Jira ticket, select **Link to existing ticket**, enter the Jira ticket number,
 then click **Link**.


 A ticket containing a link to the issue is added to the Jira project backlog. This contains the issue details from Burp Suite DAST.


 In Burp Suite DAST, the issue now contains a **Linked tickets** tab.


#### Note

 The HTTP requests and responses for issues are currently not included automatically in the Jira ticket. Although a link to the issue is provided, if the developer assigned to investigate the issue does not have access to a Burp Suite DAST account, you may need to download the HTML report and attach it to the Jira ticket manually.


## Unlink Jira tickets

 To unlink a Jira ticket from an issue in Burp Suite DAST:


1.
 Select the issue.

1.
 Go to the **Linked tickets** tab.

1.
 Click the **Unlink** button next to the relevant ticket.


#### Note

 When you unlink a ticket in Burp Suite DAST, the ticket itself is not deleted. You need to close this manually in Jira.


#### Related pages

-
 To learn how Burp Suite DAST manages duplicate tickets, see [Managing Jira ticket duplication](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/managing-jira-ticket-duplication).

-
 To learn how to raise tickets for multiple issues, see [Raising tickets for multiple issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/raise-tickets-multiple).

-
 To learn how to raise tickets automatically, see [Raising Jira tickets automatically](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/raising-jira-tickets-automatically).
