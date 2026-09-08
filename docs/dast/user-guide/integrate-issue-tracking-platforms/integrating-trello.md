> Source: https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-trello

DAST

# Integrating Burp Suite DAST with Trello

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 If you or your teams use Trello, you may like to integrate this with Burp Suite DAST. Once configured, this enables you to create Trello cards from within Burp Suite DAST for any security issues found by your scans.


## Prerequisite

-
 You must have access to Burp Suite DAST as an administrator.


## (Recommended) Create a new Trello user for the integration

 To integrate with Trello, Burp Suite DAST must be linked to a specific Trello user.


We recommend creating a new Trello user specifically for the integration. This allows you to control which boards are available for use in Burp Suite DAST - simply by giving your user access to those boards.


## Connect Burp Suite DAST to Trello

To connect Burp Suite DAST to Trello:


1.
 Log in to Trello as the user you want to use for the integration.

1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu, select **Integrations**.

1.
 On the Trello tile, select **Configure**.

1.
 Click **Get my API key**.

1.
 Review and accept Trello's Terms and Conditions, if relevant.

1.
 Copy your Trello API key to your clipboard.


![Getting your Trello API key](https://portswigger.net/burp/documentation/dast/images/trello/trello-api-key.png)

1.
 In Burp Suite DAST, paste your Trello API key into the field provided.

1.
 Click **Continue**.


![Entering your Trello API key into Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/trello/enter-trello-api-key.png)

1.
 Under **Authorization URL**, copy the URL shown and paste it into your browser.


![Getting your Trello API token](https://portswigger.net/burp/documentation/dast/images/trello/getting-trello-api-token.png)

1.
 To allow `BurpSuiteEnterpriseEdition` access to your Trello account, click **Allow**.


![Giving Burp Suite DAST access to Trello](https://portswigger.net/burp/documentation/dast/images/trello/giving-burp-suite-dast-access-to-your-trello-account.png)

1.
 Copy the token to your clipboard.


![Getting your Trello access token](https://portswigger.net/burp/documentation/dast/images/trello/trello-access-token.png)

1.
 In Burp Suite DAST, paste your token into the field provided.

1.
 Click **Connect**.


![Managing your Trello connection in Burp Suite DAST](https://portswigger.net/burp/documentation/dast/images/trello/managing-trello-connection-in-burp-suite-dast.png)

 If Burp Suite DAST successfully connects to Trello, you'll be presented with options to configure both manual and automatic card creation.


#### Note

 You must enable at least one of these in order to complete the Trello configuration.


### Enable manual Trello card creation

 To enable users to [create Trello cards manually](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/trello-cards) from within Burp Suite DAST, you need to configure the list of Trello boards and lists that they can choose from:


1.
 Select a board from the **Board** drop-down list.

1.
 Select a list type from the **List** drop-down list.

1.
 Click the **+** symbol.

1.

 If necessary, repeat these steps to add more boards and lists.


#### Note

 You need to add separate entries for each list, even when adding multiple lists from the same board.

1.
 Click **Save**.


![Enabling manual Trello card creation](https://portswigger.net/burp/documentation/dast/images/trello/manual-trello-card-permit.png)

### Enable automatic Trello card creation

 You can configure Burp Suite DAST to create Trello cards automatically. Cards are created for any issues that meet the minimum severity and confidence levels that you specify.


#### Note

 To avoid inadvertently flooding your Trello backlog with an overwhelming number of cards, we recommend setting high severity and confidence levels initially. You can then lower these once you have a better understanding of how many cards are created as a result of your scans.


1.
 Click **Enable**.

1.
 Select a board from the **Board** drop-down list.

1.
 Select a list from the **List** drop-down list.

1.
 Use the sliders to set the minimum issue severity and confidence levels that trigger Trello card
 creation.

1.
 Click **Save**.


![Enable automatic card creation for Trello](https://portswigger.net/burp/documentation/dast/images/trello/auto-trello-card-create.png)

## Raising Trello cards from within Burp Suite DAST

 For information on how users can manually raise Trello cards, refer to [Raise Trello cards from within Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/trello-cards).
