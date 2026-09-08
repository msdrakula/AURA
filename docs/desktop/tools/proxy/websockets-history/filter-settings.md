> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/filter-settings

ProfessionalCommunity Edition

# Filtering the WebSockets history

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can filter the WebSockets history to make it easier to analyze. This enables you to systematically examine a large Proxy history and understand where different kinds of interesting requests appear.

 The filter bar above the list of interactions describes the current display filter. To configure this, click the filter bar to open the **WebSockets history filter** window.

![WebSocket history filters](https://portswigger.net/burp/documentation/desktop/images/filter-settings-websockets.png)

The **WebSockets history filter** window has two tabs - **Settings mode** and **Script mode**.

The filters only control what is displayed. If you hide items, they are not deleted: they reappear if you reset the filter.

#### Note

 To quickly toggle the filter and search bar on or off, click **Filter on** or **Filter off**. This enables you to compare filtered and unfiltered traffic without resetting your filter or search term.


## Settings mode

Settings mode enables you to filter your WebSockets history quickly, using the following settings:

-

**Filter by request type** - You can choose to show:


  -
 Only the items that are [in-scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope).

  -
 Only incoming messages.

  -
 Only outgoing messages.


-

**Filter by search term** - You can:


  -
 Filter responses that contain a specified search term.

  -
 Use a literal string or a regular expression.

  -
 Make your search case-sensitive.

  -
 Select **Negative search**, so only items that don't match the search term are shown.


-  **Filter by annotation** - This enables you to only show items with notes or highlights.

-  **Filter by listener** - You can show items received on a specific [listener port](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy#binding). This can be useful when testing access controls.


## Script mode

Script mode enables you to apply scripts to define powerful custom filters. [Learn more about
 scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/scripts).

## Adding annotations

 You can add notes and highlights to history items. This enables you to describe the purpose of different items, and to flag interesting items for further investigation.

 To highlight a WebSockets history item:

1.
 In the **WebSockets history** tab, select the history item from the list.

1.
 Right-click the item and select **Highlight**.

1.
 Select a color from the list.


 To add a note:

1.
 In the **WebSockets history** tab, select the history item from the list.

1.
 Click ** Notes**.

1.
 Enter your comment in the **Notes** panel.


 You can also annotate items as they appear in the [Intercept](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages) tab. These automatically appear in the WebSockets history.
