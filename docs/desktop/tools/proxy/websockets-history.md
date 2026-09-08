> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history

ProfessionalCommunity Edition

# WebSockets history

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 You can use the **WebSockets history** to see a record of any WebSocket messages Burp's browser exchanges with web servers. You can use it to view, intercept, and modify the communication between Burp's browser and web servers. This enables you to:


-
 Study the behavior of a target website.

-
 Look for vulnerabilities in WebSockets handshakes and messages.

-
 Send interesting messages to other tools in Burp Suite for further testing.


 The WebSockets history contains the following information:


-  **#** - The request index number.

-  **URL** - The URL of the WebSocket connection.

-  **Direction** - The direction of the message (outgoing versus incoming).

-  **Edited** - Flag whether the message was modified by the user.

-  **Length** - The length of the response in bytes.

-  **Notes** - Any user-applied note.

-  **TLS** - Flag whether TLS is used.

-  **Time** - The time the message was received.

-  **Listener port** - The [listener port](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy#binding) on which the message was received.

-  **WebSocket ID** - Burp's internal ID for the WebSocket that was used for the message.


 The WebSockets history is always updated, even if [Intercept is off](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages#controls). This enables you to browse without interruption while you monitor key details about application traffic.


Right-click any item in the table to access further options, such as sending requests to other Burp tools.

## Managing the WebSockets history

 You can manage the WebSockets history in the following ways:

-

**Manage the table** - Sort and customize the table, and copy column data to your clipboard. For more information, see [Customizing Burp's tables](https://portswigger.net/burp/documentation/desktop/burps-layout/customize-tables).

-

**Filter the data** - Click the **Filter settings** bar, then choose from the following:


  -

**Settings mode** - Use predefined checkboxes and fields to set your criteria. For more information, see [Filtering the WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/filter-settings).

  -

**Script mode** - Apply a Java-based script to define your custom filter. For more information, see [Filtering the WebSockets history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/scripts).


-  **Search the data** - Enter a search term in the **Search** box to find matching items across all columns. The search applies to the underlying data, regardless of any filters you have applied.

-  **Toggle the filter** - Click **Filter on** or **Filter off** to enable or disable both the filter and the search bar. This enables you to compare filtered and unfiltered traffic without resetting your filter or search term.

-

Professional **Add custom columns** - Click the options menu **  > Add custom column** to create a personalized column that displays the data you want to see. For more information, see [Adding custom columns to the WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/custom-columns).


## Viewing a request

 If you select an item from the WebSockets history, the lower pane shows the relevant message. Any modified messages are shown separately. The message may have been modified through:


-  [User interception](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages).

-  [Automatic response modification](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy#request-and-response-interception-rules).

-  [Match and replace rules](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace).


 In addition to the main history view, you can also:


-
 Double-click an item to open it in a pop-up window.

-
 Right-click a message and select **Show new history window** to open a new history window with its own display filter.

-
 Access the [Inspector](https://portswigger.net/burp/documentation/desktop/tools/inspector), to easily view and edit interesting items.

-
 View and edit notes. To do this, click ** Notes**.


#### Related pages

- [Filtering the WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/filter-settings)
- [Filtering the WebSockets history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/scripts)
