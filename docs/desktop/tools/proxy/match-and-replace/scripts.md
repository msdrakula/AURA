> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/scripts

Professional

# Creating HTTP match and replace rules with scripts

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 You can create powerful match and replace rules using Java-based scripts. These enable you to perform custom modifications to HTTP messages as they pass through Burp's proxy.

 You can apply scripts in the following ways:

-

**[Load existing scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/scripts#loading-scripts-from-your-library)** - Load scripts from your Bambda library. This is your personal collection of reusable scripts. It includes any scripts you've created and saved, or ones you've imported, for example, from our GitHub repo. For more information, see [Importing scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/importing).
-

**[Create new scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/scripts#creating-custom-scripts)** - Write your own custom scripts. To get started quickly, use one of the built-in templates, which work without modification.

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

 You can load and apply scripts that are stored in your library to add a new rule.

 To load a script from your Bambda library:

1.

In **Proxy > Match and replace**, click **Add** to open the **Add match/replace rule** window.
1.

In the **Add match/replace rule** window, click **Script mode**.
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

Test the rule using the built-in test function. For more information, see [Testing HTTP match and replace rules](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/testing).
  1.

Save your changes:

    -

To overwrite the existing script, click **Save to library > Save**.
    -

To save a new version, click **Save to library > Save as**.

1.

Click **OK**.

 If the script is error-free, it's added to the **HTTP match and replace rules** table and automatically enabled for the current project.

 If errors exist, they appear in the **Compilation errors** panel. You'll need to fix these before you can add the script to the table. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).

## Creating custom scripts

 You can write your own scripts directly in the **Add match/replace rule** window, using built-in templates or from a blank definition.

#### Note

 Before you begin writing, we recommend exploring our [Bambdas GitHub repository](https://github.com/PortSwigger/bambdas). There may be an existing script that meets your needs or provides inspiration for creating your own.


 Two objects of the Montoya API are available to help you write your script:

-

`ProxyHttpRequestResponse`
-

`Utilities`

 The script must return either the `HttpRequest` or `HttpResponse` object.

 For advanced use cases, you can also access a subset of the Montoya API functionality. This enables you to create more complex script.

#### Warning

 Use the Montoya API functionality carefully when creating match and replace scripts. While we've restricted access to known dangerous functionality, certain methods may still potentially impact Burp's performance or cause memory leaks.


 To create a script to add a new rule:

1.

In **Proxy > Match and replace**, click **Add** to open the **Add match/replace rule** window.
1.

In the **Add match/replace rule** window, click **Script mode**.
1.

If you want to create your script from a built-in template, select **New > From template**. Select a template from the list, then click
 **Create using this template**.
1.

Write your script using Java.
1.

Test the rule using the built-in test function. For more information, see [Testing HTTP match and replace rules](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/testing).
1.

[Optional] Click **Save to library > Save**. The script is saved to your Bambda library for future use across Burp.
1.

Click **Apply & close**.

 If the script is error-free, it's added to the **HTTP match and replace rules** table and automatically enabled for the current project.

 If errors exist, they appear in the **Compilation errors** panel. You'll need to fix these before you can add the script to the table. For more information, see [Troubleshooting scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/troubleshooting).

#### Warning

 Using slow running or resource-intensive scripts can slow down Burp. Write your script carefully to minimize performance impact.


### Example scripts

In the example below, we'll create a **Request script** that forces all HTTP requests to `https://ginandjuice.shop` and adds a `User: Admin` header.

 In this example, our script is:
`return requestResponse.request()
    .withService(HttpService.httpService("https://ginandjuice.shop"))
    .withAddedHeader("User", "Admin")
    .withUpdatedHeader("Host", "ginandjuice.shop");
`

In the example below, we'll create a **Response script** that uses the `MontoyaAPI` functionality to send items to Organizer with the note "Cached response" when they meet the following criteria:

- The response has an `X-Cache` header with a value of `Hit`.

 In this example, our script is:
`if(requestResponse.response().headerValue("X-Cache").contains("Hit")) {
    api().organizer().sendToOrganizer(HttpRequestResponse.httpRequestResponse(requestResponse.request(), requestResponse.response(), Annotations.annotations("Cached response")));
}
return requestResponse.response();
`

#### Related pages

-
 To get feedback, showcase your work, and connect with other developers, share your script on our [PortSwigger
 Discord](https://discord.com/invite/portswigger) #bambdas channel.

-
 To share your scripts with the community, add them to our ever-growing GitHub repository. For more information, see [Submitting scripts to our GitHub repository](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/contribute-scripts).
