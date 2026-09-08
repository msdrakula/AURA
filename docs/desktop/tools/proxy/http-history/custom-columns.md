> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/custom-columns

Professional

# Adding custom columns in the HTTP history

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

You can use Java-based scripts to add powerful custom columns to the HTTP history table. Custom columns enable you to see more detail about the items in your HTTP history for a more focused analysis.

You can apply custom column scripts in two ways:

-

**[Load existing scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/custom-columns#loading-scripts-from-your-library)** - Load scripts from your Bambda library. This is your personal collection of reusable scripts. It includes any scripts you've created and saved, or ones you've imported, for example, from our GitHub repo. For more information, see [Importing scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/importing).
-

**[Create new scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/custom-columns#creating-custom-scripts)** - Write your own custom scripts. To get started quickly, use one of the built-in templates. These work out of the box and are easy to customize.

## Keyboard shortcuts

 To speed up your workflow when creating or loading scripts, you can use the following keyboard shortcuts:

-

**Save** - `Ctrl + S` or `Cmd + S`
-

**Save as** - `Ctrl + Shift + S` or `Cmd + Shift + S`
-

**Create new script** - `Ctrl + N` or `Cmd + N`
-

**Load recent script** - `Ctrl + O` or `Cmd + O`

## Loading scripts from your library

 You can load and apply scripts that are stored in your library to add custom columns to the HTTP history table.

 To load a script from your Bambda library:

1.

In **Proxy > HTTP history**, click the options menu ** > Add custom column**. The **Add custom column** window opens.
1.

Click  **Load**.
1.

Select a recent script from the list.
1.

If the script you want to load isn't in the list, click **View all** to view all scripts stored
 in your library.

  1. Select a script.
  1. Click  Load.

1.

[Optional] If required, edit the script:

  1.

Make your changes.
  1.

Click **Apply** to compile and test the script. Fix any errors shown in the **Compilation errors** panel. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).
  1.

Save your changes:

    -

To overwrite the existing script, click **Save to library > Save**.
    -

To save a new version, click **Save to library > Save as**.

1.

Enter a name for your column in the **Column header** field.
1.

Click **Apply & close**.

## Creating custom scripts

 You can write your own scripts directly in the **Add custom column** window, using built-in templates or from a blank definition.

#### Note

 Before you begin writing, we recommend exploring our [Bambdas GitHub repository](https://github.com/PortSwigger/bambdas). There may be an existing script that meets your needs or provides inspiration for creating your own.


 Two objects of the Montoya API are available to help you write your script:

-

`ProxyHttpRequestResponse`
-

`Utilities`

 To create a custom column for your HTTP history table:

1.

In **Proxy > HTTP history**, click the options menu ** > Add custom column**. The **Add custom column** window opens.
1.

If you want to create your script from a built-in template, select **New > From template**. Select a template from the list, then click
 **Create using this template**.
1.

Enter a name for your column in the **Column header** field.
1.

Write a script using Java to specify the data that the custom column displays.
1.

Click **Apply** to compile and test the script. Fix any errors shown in the **Compilation errors** panel. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).
1.

[Optional] Click **Save to library > Save**. The script is saved to your Bambda library for future use across Burp.
1.

Click **Apply & close**.

#### Warning

 Using slow running or resource-intensive scripts can slow down Burp. Write your script carefully to minimize performance impact.


### Example script

In the example below, we'll write a script to create a custom column containing the **Server** header value of the response.`
    if (!requestResponse.hasResponse()) {
    return "";
    }

    var response = requestResponse.response();

    return response.hasHeader("Server")
    ? response.headerValue("Server")
    : "";
`

#### Related pages

-
 To get feedback, showcase your work, and connect with other developers, share your script on our [PortSwigger
 Discord](https://discord.com/invite/portswigger) #bambdas channel.

-
 To share your scripts with the community, add them to our ever-growing GitHub repository. For more information, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).
