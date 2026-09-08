> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/scripts

ProfessionalCommunity Edition

# Filtering the HTTP history with scripts

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 You can apply Java-based scripts to create powerful custom filters for your HTTP history. You can do this in two ways:

-

**[Load existing scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/scripts#loading-scripts-from-your-library)** - Load scripts from your Bambda library. This is your personal collection of reusable scripts. It includes any scripts you've created and saved, or ones you've imported, for example, from our GitHub repo. For more information, see [Importing scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/importing).
-

**[Create new scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/scripts#creating-custom-scripts)** - Write your own custom scripts. To get started quickly, use one of the built-in templates. These work out of the box and are easy to customize.

#### Note

 To quickly toggle the filter on or off, click **Filter on** or **Filter off**. This enables you to compare filtered and unfiltered traffic without resetting your filter.


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

 You can load and apply scripts that are stored in your library to filter the HTTP history.

 To load a script from your Bambda library:

1.

In **Proxy > HTTP history**, click the filter bar to open the **HTTP history filter** window.
1.

In the **HTTP history filter** window, click **Script mode**.
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

Click **Apply & close**.

 Burp compiles your script and applies it to every item already logged in your HTTP history, as well as any future HTTP traffic generated in this project.

## Creating custom scripts

 You can write your own scripts directly in the **HTTP history filter** window, using built-in templates or from a blank definition.

#### Note

 Before you begin writing, we recommend exploring our [Bambdas GitHub repository](https://github.com/PortSwigger/bambdas). There may be an existing script that meets your needs or provides inspiration for creating your own.


### Converting filter settings to scripts

 You can convert filter settings to a script as a starting point for further customization:

1.

In **Proxy > HTTP history**, click the filter bar to open the **HTTP
 history filter** window.
1.

Make changes to the filter settings as necessary.
1.

At the bottom of the **HTTP history filter** window, click **Convert to script**.

 Your filter is converted into a script, enabling you to customize it further using Java.

### Creating your script

 Two objects of the Montoya API are available to help you write your script:

-

`ProxyHttpRequestResponse`
-

`Utilities`

 To create a script to filter your HTTP history:

1.

In **Proxy > HTTP history**, click the filter bar to open the **HTTP history
 filter** window.
1.

In the **HTTP history filter** window, click **Script mode**.
1.

If you want to create your script from a built-in template, select **New > From template**. Select a template from the list, then click
 **Create using this template**.
1.

Write your script using Java.
1.

Click **Apply** to compile and test the script. Fix any errors shown in the **Compilation errors** panel. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).
1.

[Optional] Click **Save to library > Save**. The script is saved to your Bambda library for future use across Burp.
1.

Click **Apply & close**.

 Burp compiles your script and applies it to every item already logged in your HTTP history, as well as any future HTTP traffic generated in this project.

#### Warning

 Using slow running or resource-intensive scripts can slow down Burp. Write your script carefully to minimize performance impact.


### Example script

 In the example below, we'll create a script that filters the HTTP history to show only items that meet the following criteria:

-

The request must have a response.
-

The response must have a `3XX` status code.
-

The response must have a cookie set with the name `session`.

In this example, our script is:`if (!requestResponse.hasResponse()) {
    return false;
    }

    var response = requestResponse.response();
    return response.isStatusCodeClass(StatusCodeClass.CLASS_3XX_REDIRECTION) && response.hasCookie("session");
`

#### Related pages

-
 To get feedback, showcase your work, and connect with other developers, share your script on our [PortSwigger
 Discord](https://discord.com/invite/portswigger) #bambdas channel.

-
 To share your scripts with the community, add them to our ever-growing GitHub repository. For more information, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).
