> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating

ProfessionalCommunity Edition

# Creating scripts

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Bambdas are lightweight, reusable Java-based scripts that enable you to fine-tune and extend Burp Suite's functionality. They can be used for tasks such as creating custom match-and-replace rules, table columns, and filters.

You can create scripts from the following locations:

-

**[Bambda library](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating#creating-scripts-in-the-bambda-library)** - Create any type of script, starting from a built-in template or a blank definition.
-

**[Specific Burp tools](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating#creating-scripts-from-specific-burp-tools)** - Create scripts directly in certain tools.

 Before you begin, we recommend exploring our [Bambdas GitHub repository](https://github.com/PortSwigger/bambdas). There may be an existing script that meets your needs or provides inspiration for creating your own.

#### Warning

 Slow running or resource-intensive scripts can slow down Burp. Write your script carefully to minimize performance impact.


## Creating scripts in the Bambda library

 In the Bambda library you can create new scripts using built-in templates or from a blank definition. After saving scripts to your library you can load and apply them across Burp.

To create a new script in your library:

1.

Go to **Extensions > Bambda library**.
1.

Click  **New** and select either **Blank** or **From template**.
1.

If you selected **From template**, select a template from the list, then click  **Create using this template**.
1.

Click the name field and enter a unique name.
1.

Click the **Function** drop-down menu and select the task that the script will perform.
1.

Click the **Location** drop-down menu and select the Burp tool where you want to use the script.
1.

Write the script in Java.
1.

[Optional] You can test scan check scripts against real HTTP messages. For more information, see [Testing
 custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/testing).
1.

Click **Save**. The script is saved to your library. Any errors are shown in the **Compilation errors** panel. You must resolve these before you can apply your script. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).
1.

Click **Save & close**.

 The script is saved to your library. If you created a script that defines a scan check function, Burp also adds it to your custom scan checks library, under
 **Extensions > Custom scan checks**. The new scan check is immediately available from the scan launcher.

#### Note

 Press `Ctrl + S` or `Cmd + S` to quickly save your scripts.


#### Related pages

-
 To get feedback, showcase your work, and connect with other developers, share your script on our [PortSwigger
 Discord](https://discord.com/invite/portswigger) `#bambdas` channel.

-
 To share your scripts with the community, add them to our ever-growing GitHub repository. For more information, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).


## Creating scripts from specific Burp tools

 Many tools in Burp enable you to create and apply scripts directly. For more information, see the feature-specific instructions.

### Filtering tables

For instructions on how to create scripts for filtering tables, see the following pages:

-

[Filtering the HTTP history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/scripts)
-

[Filtering the WebSockets history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/scripts)
-

[Configuring the Logger view filter with scripts](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/view-with-scripts)
-

[Configuring the Logger capture filter with scripts](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/capture-with-scripts)
-

[Filtering the site map with scripts](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/scripts)

### Adding custom columns

Professional For instructions on how to create scripts for adding custom columns to tables, see the following pages:

-

[Adding custom columns in the HTTP history](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/custom-columns)
-

[Adding custom columns in the WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/custom-columns)
-

[Adding custom columns in Logger](https://portswigger.net/burp/documentation/desktop/tools/logger/custom-columns)

### Adding custom scan checks

Professional For instructions on how to create custom scan checks, see
 [Creating custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/creating).

### Adding custom actions

Professional Custom actions are tasks that you can apply to HTTP messages in Burp Repeater to extract, transform, and analyze data.

 For instructions on how to create custom actions in Burp, see [Custom actions](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/custom-actions).

 For guidance on writing custom actions, see [Writing custom actions](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/writing-custom-actions).

### Adding match and replace rules

Professional Match and replace rules automatically replace parts of HTTP messages as they pass through the proxy.

For instructions on how to create HTTP match and replace rules with scripts, see [Creating HTTP match and replace rules with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/scripts).
