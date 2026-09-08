> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/manual-install

ProfessionalCommunity Edition

# Installing extensions manually

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can install extensions manually from a file. This is useful when:

-

You want to use a custom extension in `.jar`, `.py`, or `.rb` format. For instructions, see
 [Manually installing a custom extension](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/manual-install#manually-installing-a-custom-extension).
-

You are working offline in Burp and can't access the BApp Store. In this situation, you can use a separate browser to download extensions from the BApp Store on the PortSwigger website and import the
 `.bapp` file into Burp. For instructions, see [Manually installing a BApp Store extension](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/manual-install#manually-installing-a-bapp-store-extension).

#### Note

 If neither of these situations apply, we recommend installing extensions directly from the BApp Store in Burp. For more information, see [Installing extensions from the BApp Store](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/bapp-store).


 If you can't access the BApp Store and you're not intentionally working offline, see our [troubleshooting guide](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/troubleshooting).


## Manually installing a custom extension

 You can install your own Java, Python, or Ruby custom extensions.

#### Note

 If you want to install an extension written in Python or Ruby, you'll need to configure Jython or JRuby before you can add the extension to Burp. For instructions, see our [troubleshooting guide](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/troubleshooting#you-need-to-configure-jython-or-jruby).


 To install a custom extension:

1.

Go to **Extensions > Installed** and click **Add**.
1.

Under **Extension Details**, use the **Extension type** dropdown to select the type of extension you want to install. You can choose from Java, Python, or Ruby.
1.

Click **Select file** and choose the extension file.
1.

[Optional] Under **Standard output** and **Standard error**, choose where to save output and error messages.
1.

Click **Next** to load the extension.
1.

Review any messages displayed in the **Output** and **Errors** tabs.
1.

Click **Close**.

 The extension is installed and automatically enabled. It's added to the bottom of the list in **Extensions > Installed**. Extensions process traffic in list order, so you may want to adjust their position. For instructions, see [Managing extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/managing-extensions).

## Manually installing a BApp Store extension

 You can download extensions from the BApp Store on the PortSwigger website, then install them in Burp.

#### Note

 We review all extensions submitted to the BApp Store, but as they are written by third parties, we can't guarantee their quality or security. As extensions can run arbitrary code, we recommend that you review them yourself before installing. To view the source code for an extension, visit our [GitHub page](https://github.com/portswigger/).


 For more information about how we protect your data when you use AI-powered extensions, see [AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust).


 To manually install a BApp Store extension:

1.

Go to the BApp Store web page.
1.

Click the name of the extension you want to install. The extension description opens.
1.

Click **Download BApp**.
1.

If you're using a closed network, move the `.bapp` file into your network.
1.

In Burp, go to **Extensions > BApp Store**, click , then select **Import BApp file**.
1.

Select the file you previously downloaded and click **Open**.

 The extension is installed and automatically enabled. It's added to the bottom of the list in **Extensions > Installed**. Extensions process traffic in list order, so you may want to adjust their position. For instructions, see [Managing extensions](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/managing-extensions).
