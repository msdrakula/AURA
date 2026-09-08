> Source: https://portswigger.net/burp/documentation/desktop/tools/target/site-map/workflow-tools

ProfessionalCommunity Edition

# Site map workflow tools

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 Once you populate the [site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map) with information about your target, you can use the context menu to drive your workflow. To view the context menu, select an item from anywhere in the site map and right-click it.


 The context menu only shows actions that are relevant to the selected item. This section describes all the possible context menu actions.


## Add to scope / Remove from scope

 You can add or remove URLs from the target scope. Any changes you make also apply to any child branches.


 This is useful if you're testing an application that includes some sensitive URLs:


1.
 Select the whole application path (domain or directory).

1.
 Right-click and select **Add to scope**.

1.
 Right-click the sensitive items and select **Remove from scope**.


## Send to

 You can send items to other Burp tools, such as [Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder) or [Organizer](https://portswigger.net/burp/documentation/desktop/tools/organizer). This enables you to perform further attacks or analysis and use Burp to drive your workflow.


## Scan

 You can use the scanner to [scan for content or vulnerabilities](https://portswigger.net/burp/documentation/desktop/running-scans). You can scan an entire branch of a tree if you select this action from the tree view.


## Show response in browser

 You can render responses in Burp's browser, to avoid the limitations of Burp's built-in HTML renderer. To render the response, paste the unique URL that Burp generates into Burp's browser.


 Burp serves the resulting browser request with the exact response that you select: the request is not forwarded to the original web server. Burp's browser processes the response in the context of the originally requested URL. This means that relative links within the response are handled properly.


 When Burp's browser renders the response it may make additional requests, for example for images or CSS. These are handled by Burp in the usual way.


## Record an issue

Professional Manually record an issue for the selected request / response pair:

-

**Create an issue** - Add a new issue.

-

**Add to manually created issue** - Add a request / response pair to a pre-existing manually created issue.


 The issue is saved to your project and can be included when you generate a report.

 For more information, see [Manually creating issues for reports](https://portswigger.net/burp/documentation/desktop/running-scans/reporting/manual-issues).

## Request in browser

 You can resend requests in Burp's browser:


-  **In original session** - Resend the request using the cookie header that appeared in the original request.

-

**In current browser session** - Resend the request using the cookies supplied by the browser. You can use this feature to test [access controls](https://portswigger.net/web-security/access-control):


  -
 Select requests within Burp that are generated within one user context (for example, an administrator).

  -
 Resend the requests when you log in with a different user context (for example, an ordinary user).


 This method makes it much easier to deal with complex, multi-stage processes. You can simply paste a series of URLs from Burp into Burp's browser. The alternative is to repeat complicated procedures many times, and manually modify cookies with the Proxy.


## Engagement tools

Professional The **Engagement tools** submenu contains useful functions that enable you to perform engagement-related tasks:


-  **Search** - Select branches from the site map and use the [Search function](https://portswigger.net/burp/documentation/desktop/tools/search) to find items that match a specific expression.

-  **Find notes / Find scripts** - Select branches from the site map and use the [Find comments and scripts](https://portswigger.net/burp/documentation/desktop/tools/search#find-comments-and-scripts) functions to find comments or scripts.

-  **Find references** - Use the [Find references](https://portswigger.net/burp/documentation/desktop/tools/search#find-references) function to search all of Burp's tools for HTTP responses that link to your selected item.

-  **Analyze target** - The [Target Analyzer](https://portswigger.net/burp/documentation/desktop/tools/engagement-tools/target-analyzer) function shows you how many static and dynamic URLs are contained in your selected branch. It also shows you how many parameters each URL takes.

-  **Discover content** - Use the [Discover content](https://portswigger.net/burp/documentation/desktop/tools/engagement-tools/content-discovery) function to discover content and functionality that is not linked from visible content. This enables you to browse or analyze this content.

-  **Schedule task** - The [Schedule task](https://portswigger.net/burp/documentation/desktop/settings/project/tasks#schedule-tasks) function enables you to create tasks that run automatically at defined times and intervals.

-  **Generate CSRF PoC** - Use the [Generate CSRF PoC](https://portswigger.net/burp/documentation/desktop/tools/engagement-tools/generate-csrf-poc) function to create HTML that causes the request to be issued when it's viewed in a browser.

-  **Simulate manual testing** - You can use the [Manual testing simulator](https://portswigger.net/burp/documentation/desktop/tools/engagement-tools/manual-simulator) to generate HTTP traffic that's similar to the traffic caused by manual penetration testing.


## Compare site maps

 The **Compare site maps** function enables you to identify differences between two site maps. To learn more, see [Comparing site maps](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/comparing).


## Add notes / Highlight

 You can use these functions to add notes or highlights to items. To learn more, see [Annotations](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-filter#site-map-annotations).


## Expand / collapse branch / requested items

 Use these functions in the tree view to quickly expand or collapse whole branches of the tree.


## Delete items

 This function removes the selected items permanently. By default, the site map displays all the content that Burp identifies based on HTTP responses. This means that the map often includes a large amount of third-party content that the application links to. You can manage this in two ways:


-
 Configure a suitable [target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope) and [display filter](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-filter).

-
 Manually remove the irrelevant branches of the tree.


## Copy URLs

 This function copies the URLs of the selected items to the clipboard.


## Copy as curl command

 This function copies a curl command to the clipboard. You can use it to generate the selected request.


## Copy links

 You can use this function to parse items for links and copy the links to the clipboard.


## Save items

 This function lets you save the details of your selected items as an XML file. The file includes full requests and responses, and relevant metadata such as response length, HTTP status code, and MIME type.


## Show new site map window

 You can use this function to open another site map window. You can open multiple windows and configure different filters for each window.


#### Related pages

-  [Site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map).

-  [Comparing site maps](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/comparing).

-  [Editing the site map layout](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-layout).
