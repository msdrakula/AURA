> Source: https://portswigger.net/burp/documentation/desktop/tools/repeater/groups

ProfessionalCommunity Edition

# Managing tab groups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can organize tabs into groups to manage large numbers of open tabs. This also enables you to send requests from multiple tabs in sequence.


#### Note

For more information on sending grouped requests as a sequence, see [Sending requests in sequence](https://portswigger.net/burp/documentation/desktop/tools/repeater/send-group).

## Creating a new tab group

 To create a tab group:


1. Click the add tab button  and select  **Create tab group** to display the create group context menu.
1. Enter a **Group name**.
1. Click **Select all** to add all tabs to the group, or use the checkboxes to pick specific ones. You can shift-click to select multiple tabs.
1. Select a group color. The group's tabs are highlighted with this color on the tab bar.
1. Click **Create** to set up the group. An icon for the group is added to the tab bar. Click the icon to display the individual tabs within the group.

 Any tab groups you create remain open even if you restart Burp Suite.


## Editing existing groups

 You can edit existing groups of tabs:


- Edit the group - Right-click the group and select **Edit group** to display the context menu. The fields on this menu are the same as those used when you create a group.
- Add a tab to a group - Right-click the tab and select a group from the **Add tab to group** menu.
- Remove a tab from a group - Right-click the tab and select **Remove tab from group**. Note that if you remove the last tab from a group, the group is automatically closed.
- Close all tabs in the group except the selected tab - Right-click the tab and select **Close other tabs in group**.
- Close a group but keep its tabs open - Right-click the group icon and select **Ungroup tabs**.

## Closing tab groups

 To close a group and all of its tabs, right-click the group icon and select **Delete entire group**.
