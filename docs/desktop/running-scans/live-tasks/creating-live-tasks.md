> Source: https://portswigger.net/burp/documentation/desktop/running-scans/live-tasks/creating-live-tasks

Professional

# Creating live tasks

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Suite's **Live tasks** feature enables you to perform some scanning operations automatically. You can use live tasks to audit for vulnerabilities, or add resources to Burp's **Target** site map.


#### Related pages

[Live tasks](https://portswigger.net/burp/documentation/desktop/running-scans/live-tasks) - Gives further information around how live tasks work.

 To create a new live task:


1.

From the Dashboard, click **New live task** to display a dialog.
1.

Select a **Task type**:

  - **Live audit**.
  - **Live passive crawl**.

1.

Select the **Tools scope**. You can set the task to inspect the traffic from the following tools:

  - **Proxy**.
  - **Repeater**.
  - **Intruder**.

1.

Select the **URL scope**. You can set the task to process the following items for the selected tools:

  - **Everything** - Includes all URLs.
  - **Suite scope** - Includes all URLs covered by the current suite-wide scope.
  - **Custom scope** - Enables you to specify your own URLs for the task to match. Live tasks use Burp Suite's standard URL matching and advanced scope control rules. See [URL matching](https://portswigger.net/burp/documentation/desktop/tools/target/scope#url-matching-rules) for more details.

1.

If required, select **Ignore duplicate items based on URL and parameter names** to reduce the number of items processed by the task.
1.

If required, click the **Scan configuration** tab and select a scan configuration for the task. For more information, see
 [Configuring scans](https://portswigger.net/burp/documentation/desktop/running-scans/configuring-scans).
1.

If required, and if you are creating an audit task, click the **Resource pool** tab and configure the resource pool that the task runs in. For more information on configuring resource pools, see [Managing resource pools for scans](https://portswigger.net/burp/documentation/desktop/running-scans/managing-resource-pools).
1.

Click **OK** to start the task.

 You can also add a pre-configured live task. Choose a task from the **Choose predefined task** drop-down in the **Scan details** tab. The available options are:


- **Passively scan all traffic through Proxy**.
- **Actively scan all in-scope traffic through Proxy**.
- **Add all items requested through Proxy to site map**.
- **Add all links observed in traffic through Proxy to site map**.
