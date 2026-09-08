> Source: https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-filter

ProfessionalCommunity Edition

# Filtering the site map

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can use the display filter to hide some of the content in the site map. This makes it easier for you to analyze the content.


 The filter bar above the site map displays the current display filter. To configure this, click the filter bar to open the **Site map filter** window. The window has two tabs - **Settings mode** and **Script mode**.


 The filters only control what is displayed. If you hide items, they are not deleted: they reappear if you reset the filter.


## Settings mode

 Settings mode enables you to filter your site map quickly, using the following settings:


-  **Request type** - Show only [in-scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope) items, only requested items, or only requests with parameters. You can also hide not-found items.

-  **MIME type** - Filter responses that contain particular MIME types, such as HTML, CSS, or images.

-  **Status code** - Filter responses based on their HTTP status code.

-  **Folders** - Hide empty folders in the tree view. This enables you to hide folders where all the child items are hidden by other display filter attributes.

-  Professional **Search term** - Show responses that contain a specific search term. You can use a literal string or a regular expression, and you can make the search case-sensitive. To only show items that don't match the search term, select **Negative search**.

-  **File extension** - Filter items by their file extension.

-  **Annotation** - Filter items with [notes or highlights](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-filter#site-map-annotations).


#### Note

 To use different display filters, you can pop up additional site map windows and apply a different display filter to each window. Select **Show new site map window** from the context menu.


## Script mode

Script mode enables you to apply scripts to define powerful custom filters for your site map. For more information, see
 [Filtering the site map with scripts](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/scripts).

## Site map annotations

 In the **Contents** view, you can annotate items by adding notes and highlights. This enables you to describe the purpose of different URLs or flag interesting items for further investigation.


 To add a highlight, click the left-hand column and select a color from the drop-down menu.


 To add a note:


- In the **Contents** panel, select the item from the list.
- Click ** Notes**.
- Enter your comment in the **Notes** panel.

 To easily find the items later, sort the content by the **Host** or **Notes** columns or use the display filter.


#### Related pages

-  [Site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map).

-  [Site map workflow tools](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/workflow-tools).

-  [Comparing site maps](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/comparing).

-  [Editing the site map layout](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/site-map-layout).
