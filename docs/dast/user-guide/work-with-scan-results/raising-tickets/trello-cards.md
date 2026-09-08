> Source: https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/trello-cards

DAST

# Raising Trello cards from within Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If your admin user has [configured a Trello integration](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-trello), you can raise Trello cards for issues directly from Burp Suite DAST.


1.
 From the top menu, select **Scans**.

1.
 Select the scan you want to view.

1.
 Select the **Issues** tab.

1.
 Expand the issue and select the URL from the list.

1.

 In the upper-right corner of the page, click the **Raise Trello card** button.


#### Note

 If you have also integrated Burp Suite DAST with other issue-tracking platforms, you may need to select this from the **Raise ticket** drop-down.

1.
 You can create a new Trello card, or link to an existing Trello card:


  -
 To create a new card, select your board and list from the drop-down menu, then click **Create**.

  -
 To link to an existing card, select **Link to existing card**, enter the short URL, and then click **Link**.


![Creating a Trello card from within Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/trello/creating-new-trello-card-burp-suite-dast.png)

## Raising Trello cards for multiple issues

 You can raise Trello cards for more than one issue at a time. For more information, refer to [Raising tickets for multiple issues](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/raise-tickets-multiple).
