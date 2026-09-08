> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/troubleshooting

ProfessionalCommunity Edition

# Troubleshooting Burp extensions

-

**Last updated: ** September 3, 2026
-

**Read time: ** 9 Minutes

 This guide helps you resolve common issues when using Burp extensions.

 If the troubleshooting steps in this guide don't resolve your issue, please contact our support team at `support@portswigger.net`.

## Installing extensions

#### You can't access the BApp Store in Burp

 To access extensions from the BApp Store in Burp, your device must be able to access `portswigger.net`. You might not be able to access extensions because:


- You're offline.
- Your network requires an upstream proxy.
- An intercepting proxy is intercepting and resigning traffic with its own certificate. Burp doesn't trust self-signed certificates by default, so it blocks the connection.

 If you're working offline in Burp, you can use a separate browser to download extensions from our website, then install them manually. For more information, see [Installing extensions manually](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/manual-install).


### Step 1: Check your internet connection

 Make sure your computer is connected to the internet. To verify this, open an external browser and try visiting `https://portswigger.net`. If you can't access the site, check your network settings or contact your administrator.


### Step 2: Configure an upstream proxy

 Some networks require an upstream proxy for internet access. If your network requires one, configure Burp to use it:


1.

In Burp, click ** Settings**. The Settings dialog opens.
1.

Go to **Network > Connections**.
1.

Under **Upstream proxy servers**, click **Add**. The Add upstream proxy rule dialog opens.
1.

Enter the details of the upstream proxy. For more information, see [Connections settings - Upstream proxy servers](https://portswigger.net/burp/documentation/desktop/settings/network/connections#upstream-proxy-servers).
1.

Click **OK**.
1.

The upstream proxy rule is added to the table.

### Step 3: Identify and resolve an intercepting proxy

 Some networks use intercepting proxies, such as ZScaler, to inspect and decrypt encrypted traffic. These proxies intercept HTTPS connections and re-sign certificates, which means Burp won't trust the connection.


 To check if an intercepting proxy is impacting your connection:


1.

In Burp's browser, go to `https://portswigger.net/bappstore`.
1.

In Burp, go to **Settings > Network > TLS**.
1.

Under **Server TLS certificates**, find the entry for `portswigger.net`.
1.

Check the issuer of the certificate:

  -

If the issuer is a well-known CA provider such as Amazon, it's unlikely that an intercepting proxy is interfering with the connection.
  -

If the certificate is issued by an intercepting proxy (such as ZScaler), or your company's security system, then your traffic is being intercepted.

 If your traffic is being intercepted by a proxy, Burp might not trust the proxy's certificate. This can prevent access to the BApp Store.


 To fix this, add the proxy's certificate authority (CA) to Burp's trusted certificates:


1.

Find the path to the proxy's certificate. If you're unsure, check your system settings or ask a network administrator.
1.

In Burp, go to **Settings > Network > TLS**.
1.

Under **Custom CA certificates**, click **Add**.
1.

Select the certificate file.
1.

Restart Burp, then check whether you can access the BApp Store.

#### You can't install extensions from the BApp Store in Burp

 If you can't install an extension from the BApp Store, the **Install** button is grayed out. This may be because:


-

You need to update Burp.
-

You need to install Jython or JRuby.
-

You need to upgrade to Burp Suite Professional.

### You need to update Burp

 Some extensions require features or API methods that were introduced in newer versions of Burp. To update Burp:


1.

Click the **Help** top-level menu.
1.

Select **Check for updates**. A dialog opens with details of the latest Burp version.
1.

Click **Update now**.
1.

Wait for Burp to prepare the update. When it's ready, you'll be prompted to restart. Click **Update and restart** to complete the process.
1.

When Burp restarts, return to **Extensions > BApp store** to download your extension.

### You need to configure Jython or JRuby

 To use an extension that is written in Python or Ruby, you'll need to configure Jython or JRuby. Burp is a Java application and requires Java-compatible implementations of these languages to run the extension code.


 To configure Jython or JRuby:


1.

Download the [Jython standalone JAR file](https://www.jython.org/download.html) or the
 [Ruby JAR file](https://www.jruby.org/download).
1.

In Burp Suite, click ** Settings** to open the
 **Settings** dialog.
1.

Go to **Extensions**.
1.

Under **Python Environment** or **Ruby Environment**, click **Select file**.
1.

Select the downloaded JAR file and click **Open**.

Once configured, return to **Extensions > BApp Store**, click , select **Refresh list**, then download your extension.

#### Related pages

[Extensions settings](https://portswigger.net/burp/documentation/desktop/settings/extensions).


### You need to upgrade to Burp Suite Professional

 Some extensions require Burp Suite Professional features. If you're using Burp Suite Community edition, you must upgrade to install these extensions.


 For more information, see [Upgrade to Burp Suite Professional](https://portswigger.net/burp/upgrade-community-to-pro).


#### An extension isn't working as expected

 If an extension is installed but not working properly, try the following steps to investigate the issue.


### Step 1: Update Burp

 It's best practice to run extensions on the latest version of Burp. If an extension requires features that were introduced in a newer version than the one you're using, it may not function correctly.


 To update Burp:


1.

Click the **Help** top-level menu.
1.

Select **Check for updates**. A dialog opens with details of the latest Burp version.
1.

Click **Update now**.
1.

Wait for Burp to prepare the update. When prompted, click **Update and restart** to finish.
1.

When Burp restarts, check that your extension is working as expected.

### Step 2: Visit the extension's repository

 All extensions in the BApp Store include a link to PortSwigger's fork of the author's GitHub repository. This may include additional documentation and troubleshooting information.


 To access the author's GitHub repository:


1.

In Burp, go to **Extensions > BApp Store**.
1.

Select the extension you want to investigate.
1.

In the description panel, scroll to **Source** and click the repository link.
1.

Review the `README` file and any other resources in the repository.

### Step 3: Use AI to review the extension's code

 You can use an LLM to help you understand how the extension works and identify any bugs. To support this, we provide a `CLAUDE.md` file that includes essential context on how Burp extensions are structured for the model.


 Before you start, you'll need to install or access your preferred LLM. These instructions use Claude Code, but you can adapt them for use with other LLMs. For more information on Claude Code, see the
 [Anthropic documentation](https://www.anthropic.com/claude-code).


#### Note

 While we review all extensions submitted to the BApp Store, they are written by third parties and can run arbitrary code. We therefore can't guarantee their quality or safety.


#### Step 1: Access the extension code file

 To access the extension's code:


1.

In Burp, go to **Extensions > BApp Store**.
1.

Select the extension you want to review.
1.

In the description panel, scroll to **Source** and click the repository link.
1.

In GitHub, if the repository was forked, click the **Forked from** link to access the original repository.
1.

Fork the original repository to your own **GitHub** account.
1.

Clone your forked repository to your local machine.
1.

Open or import the local clone into your **IDE**.

#### Step 2: Run your LLM to analyze the extension

To use Claude Code to review the extension:

1.

Download our [CLAUDE.md](https://github.com/PortSwigger/extension-template-project/blob/main/CLAUDE.md) file and
 [supporting documentation](https://github.com/PortSwigger/extension-template-project/tree/main/docs) from **GitHub**.
1.

Add the `CLAUDE.md` file to the root of the extension's folder.
1.

Create a `docs` folder in the extension's folder and add the supporting documentation inside it.
1.

Open a terminal and navigate to the extension's folder.
1.

Run Claude Code using the following command: `claude`.
1.

Prompt Claude to review the code, explain how to use the extension, and identify any bugs.

 Claude should automatically read the contents of the `ExtensionTemplateProject` folder, including the `CLAUDE.md` file, then explain how the extension works. If you think it hasn't read the `CLAUDE.md` file, directly prompt it to do so before continuing.


#### Note

 To use the `CLAUDE.md` file with an LLM other than Claude Code, prompt the LLM to read the file and supporting documentation, or provide their contents as part of your context window.


 If you identify any issues with the extension's code, you can report this to the extension's author to fix.


### Step 4: Report an issue or suggest a fix to the extension's author

 PortSwigger doesn't maintain the community-created extensions in the BApp Store. If you believe you've found a bug or need support for a particular extension, please contact the extension's author through their repository.


 To do this, you need to access the author's original extension repository:


1.

In Burp, go to **Extensions > BApp Store**.
1.

Select the extension you want to investigate.
1.

In the description panel, scroll to **Source** and click the repository link.
1.

If the repository was forked, click the **Forked from** link to access the original repository.

#### Option 1: Report an issue

 If you don't have a proposed solution, you can still report the issue to make the extension's author aware of the problem:


1.

In the extension author's original repository, click the **Issues** tab.
1.

Review existing issues to see if the problem has already been reported.
1.

If not, click **New issue** and describe the problem in detail, including:


  -

The version of Burp Suite you're using.
  -

The operating system.
  -

Any relevant steps to reproduce the issue.
  -

Screenshots or logs, if applicable.

#### Option 2: Suggest a fix

 If you've identified the cause of the issue and have a proposed solution, you can suggest a fix. This may speed up the resolution process.


#### Note

 You can optionally use an LLM to help identify and implement appropriate code changes. Make sure to review and test any code suggested by the LLM before submitting it, as it may not always produce correct or secure output. For LLM setup instructions, see
 [Step 3: Use AI to review the extension's code](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/troubleshooting#step-3-use-ai-to-review-the-extension-s-code).


 To suggest a fix:


1.

Fork the extension author's original repository.
1.

Make changes to the extension's code.
1.

Create a pull request. In the description, clearly describe the problem and proposed fix, including:

  -

The version of Burp Suite you're using.
  -

The operating system.
  -

Any relevant steps to reproduce the issue.
  -

A summary of the changes introduced by your fix.
  -

Screenshots or logs that support your solution, if applicable.

## Performance issues

#### Burp has performance issues while using extensions

 Extensions can cause Burp to run slowly. Extensions on the BApp Store aren't tested by PortSwigger for performance.


 If you notice performance issues while using extensions, try the following troubleshooting steps:


### Step 1: Disable unused extensions

 Running too many extensions can slow Burp down. To minimize resource usage:


1.

In Burp, go to **Extensions > Installed**.
1.

In the **Loaded** column, uncheck any extensions you're not currently using.
1.

Restart Burp.
1.

Test Burp's performance.

#### Note

 To estimate the total impact of installed extensions, go to the **Extensions > Installed** tab and check the **Total estimated system impact** indicator.


### Step 2: View the performance impact estimate for particular extensions

 Before testing an extension manually, you can get a general idea of its impact by checking the **Estimated system impact** ratings in the BApp Store.


#### Warning

 These ratings are estimates only and may not fully reflect real-world performance. For example, we are unable to fully test extensions that add custom tabs or context menu options. The most reliable way to assess performance impact is to test extensions manually in Burp.


 To check the expected system impact of a specific BApp Store extension:


1.

Go to **Extensions > BApp Store**.
1.

Select the extension and scroll down the right-hand panel to **Estimated system impact**.
1.

Review the estimated impact across the following categories:


  -

**Memory** - Potential impact on Burp Suite's memory usage.
  -

**CPU** - Additional processing load on the CPU.
  -

**Time** - Impact on Burp's overall speed. This includes the responsiveness of the interface and how long tools take to complete tasks.
  -

**Scanner** - Potential increase in scan duration.
  -

**Overall** - The highest impact rating among all categories.

### Step 3: Manually test extensions

 Extensions can impact performance individually, or the combination of extensions can cause a performance issue. Follow these steps to pinpoint the cause:


1.

Disable all extensions:

  1.

In Burp, go to **Extensions > Installed**.
  1.

In the Burp extensions table, click anywhere in the list and press `Ctrl + A` or `Cmd + A` to select all extensions.
  1.

Right-click then select **Unload**. This unloads all your installed extensions.
  1.

Test Burp's performance without any extensions enabled.

1.

In **Extensions > Installed**, re-enable extensions one at a time, starting with the ones you want to use immediately.
1.

Test Burp's performance after enabling each extension. If Burp slows down, the recently enabled extensions may be causing the issue.
1.

If Burp slows down after enabling multiple extensions, disable them selectively to identify conflicts.

#### I get an error message saying `java.lang.OutOfMemoryError: Metaspace`

 You may see this error message if you load several Python or Ruby extensions, or if you unload and reload extensions multiple times.


 To avoid this issue, configure Java to allocate more `Metaspace` storage:


1.

Open a terminal or command prompt.
1.

Add `-XX:MaxMetaspaceSize=1G` to the command you use to launch Burp, as follows: `java -XX:MaxMetaspaceSize=1G -jar FILE_PATH.jar`

#### Related pages

[Launching Burp Suite from the command line](https://portswigger.net/burp/documentation/desktop/troubleshooting/launch-from-command-line)
