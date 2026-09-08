> Source: https://portswigger.net/burp/documentation/dast/user-guide/integrating-splunk

DAST

# Integrating Burp Suite DAST with Splunk

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you or your teams use Splunk for your Security Information and Event Management (SIEM), you may like to integrate this with Burp Suite DAST.


 Once configured, this enables you to stream issues directly to Splunk for advanced analysis, enabling real-time monitoring and event management.


#### Note

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud If you're a Cloud user, you need to use Splunk Cloud Platform and host on `*.splunkcloud.com`.


## Prerequisites

-
 You have access to Burp Suite DAST as an administrator.

-
 You have access to Splunk as an administrator.


## Configuring a connection to Splunk

 To configure a connection to Splunk:


1.

 In Splunk:


  -
 Make a note of the URL where Splunk is hosted.

  -
 Create a new **HTTP Event Collector** token. Copy the **Token Value**.

  -
 Make a note of the **HTTP Port Number**. This is typically `8088`.


1.
 In Burp Suite DAST, go to **Settings**  and select **Integrations**.

1.
 Find the tile for Splunk and click **Configure**.

1.
 Enter the Splunk URL and HTTP Port Number in the **Splunk URL** field. For example, `http://10.100.1.100:8088`.

1.
 Enter the Splunk **Token Value** in **Splunk token value**.

1.
 Click **Connect**, and make a note of the name of the integration token that Burp Suite DAST sends to Splunk.

1.
 In Splunk, go to **Search & Reporting** and search for the integration token in the **Event** list.


## Creating a new event type in Splunk

 If you want to use the **Vulnerabilities data model** in the Splunk Common Information Model (CIM) add-on, you need to configure a new **Event type** in Splunk:

1.
 In your Splunk settings, create a new **Event type**.

1.
 Specify a search string for the **Event type**, in order to use the **HTTP Event Collector** token to filter events from Burp Suite DAST.

1.
 Enter a name for the **Event type** that enables you to identify issues sent by Burp Suite DAST.

1.
 Add the tags `report` and `vulnerability` to the **Event type**.

1.
 Save the **Event type**.


## Disconnecting from Splunk

 To disconnect from Splunk:

1.
 In Burp Suite DAST, Go to **Settings**  and select **Integrations**.

1.
 Find the tile for Splunk and click **Edit**.

1.
 Click **Disconnect Splunk** and then click **OK**.
