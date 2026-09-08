> Source: https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/scanning-specific-http-messages

Professional

# Scanning specific HTTP messages

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Scanning specific HTTP messages makes it easy to run focused scans on a particular set of requests or responses.


 You can scan HTTP messages from most places that display HTTP traffic in Burp Suite. In tools that display lists of HTTP requests (such as the **Site map** and **HTTP history** tabs) you can select multiple entries to scan.


 To scan the selected HTTP messages, right-click and select one of the scan options from the context menu. There are three options available:


-

**Scan**. This menu item has two options:

  - **Open scan launcher**. This opens a scan launcher window from where you can configure the scan.
  - **Add to task**. This enables you to add a scan of the message to a pre-existing task.

-

**Do passive scan**. Burp Scanner analyzes the contents of the base request and response, rather than sending its own requests.
-

**Do active scan**. Burp Scanner sends its own requests to the target to probe for vulnerabilities.

## Configuring an audit of specific HTTP messages

 To configure an audit of specific HTTP messages:


1.

Right-click the messages required and select **Scan**.
1.

From the **Scan type** tab of the scan launcher, select **Audit selected items**.
1.

Select the task that you want the audit to run under:

  - To add the audit to an existing task, select **Add to task** and select the required task from the list.
  - To have the audit run under its own task, select **Create new task**.

1.

Optionally, select **Consolidate items** to remove unnecessary messages from the audit. You can consolidate items using the following criteria:

  - Duplicates (messages that have the same URLs and parameters).
  - Out-of-scope messages based on the current suite scope.
  - Messages with no parameters.
  - Messages with a specified file extension.

1.

Optionally, specify details for the remaining launcher tabs:

  - Scan configuration. For more information, see [Configuring scans](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-scans).
  - Resource pool. For more information, see [Managing resource pools for scans](https://portswigger.net/burp/documentation/desktop/running-scans/managing-resource-pools).

1.

Click **Scan** to start the audit.

#### Note

 Alternatively, you can configure a crawl or a combined crawl and audit of the selected HTTP messages. To do this, select the appropriate option in the **Scan type** tab of the scan launcher. Burp automatically fills the **Scan details > URLs to scan** field based on the selected URLs. For more information on how this restricts the scope of the crawl, see
 [Setting the scan scope](https://portswigger.net/burp/documentation/desktop/running-scans/setting-pro-scope).


 You can configure the remaining scan settings as you would for a crawl and audit. For more information, see [Running a full crawl and audit](https://portswigger.net/burp/documentation/desktop/running-scans/webapp-scans/full-crawl-and-audit).


#### Related pages

- [Setting the scope in Burp Suite Professional](https://portswigger.net/burp/documentation/desktop/running-scans/setting-pro-scope#items-to-scan) - Gives detailed information on the scope options available to you in Burp Suite Professional.
- [Configuring scans](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-scans) - Gives further information on configuring scans in Burp Suite Professional.
- [Resource pools](https://portswigger.net/burp/documentation/desktop/running-scans/managing-resource-pools) - Gives information on the use cases for resource pools and how to configure them.
