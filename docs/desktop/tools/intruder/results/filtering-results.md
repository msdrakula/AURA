> Source: https://portswigger.net/burp/documentation/desktop/tools/intruder/results/filtering-results

ProfessionalCommunity Edition

# Filtering attack results

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 You can use the following filters to manage the results of a Burp Intruder attack:

-

**Capture filter** - Choose which items Intruder captures during the attack to control resource usage. For more information, see [Capture filter](https://portswigger.net/burp/documentation/desktop/tools/intruder/results/filtering-results#capture-filter).

-

**View filter** - Choose which items are displayed in the results table. This helps you to focus on relevant results by hiding unnecessary data without deleting items. For more information, see [View filter](https://portswigger.net/burp/documentation/desktop/tools/intruder/results/filtering-results#view-filter).


 The current filters are described in the filter bars above the results table.

![Intruder filter bars](https://portswigger.net/burp/documentation/desktop/images/intruder-filter.png)

## Capture filter

 To filter the items that Intruder captures during the attack:

1.

 In the **Intruder results** window, click the **Capture filter** bar. The **Intruder capture filter** window opens.

1.

Configure the capture filter settings.
1.

Click **Save settings**. The **Intruder capture filter** window closes.
1.

To apply the settings to your attack, select **Apply capture filter**.

#### Note

 The capture filter only impacts items generated after it's applied. To prevent unwanted items being added to the results table while configuring the filter, pause the attack by clicking .


### Capture filter settings

 In the **Intruder capture filter** window you can configure the following settings:

-

**Discard items without responses** - Exclude requests that don't receive a response.
-

**Capture by search term** - Enter a term. Only responses containing the term are captured. Options for applying the search term include:


  -

**Regex** - Specify whether the search term is a literal string or a regular expression.
  -

**Case sensitive** - Specify whether the search term is case-sensitive.
  -

**Negative search** - Show or hide items that do not match the search term.

-

**Capture by status code** - Capture responses based on their HTTP status code.
-

**Capture by annotation** - Only capture items with comments or highlights.

### Managing the capture filter settings

 The **Intruder capture filter** window includes options for quickly managing the filter settings, making it easier to refine and reset your changes. The available options are:

-

**Show all** - Capture all items.
-

**Hide all** - Stop capturing items.
-

**Revert changes** - Undo any changes you made since opening the current filter window.

## View filter

 To filter the results that are displayed in the results window:

1.

 In the **Intruder results** window, click the **View filter** bar. The **Intruder view filter** window opens.

1.

Configure the view filter settings.
1.

Click **Apply**. The results window updates to display only the items that match your filter criteria.

 The view filter only controls what is displayed. If you hide items, they are not deleted: they reappear if you reset the filter.

### View filter settings

 In the **Intruder view filter** window you can configure the following settings:

-

**Filter by search term** - Show or hide responses containing a specified term. You have the following options:


  -

**Regex** - Specify whether the search term is a literal string or a regular expression.
  -

**Case sensitive** - Specify whether the search term is case-sensitive.
  -

**Negative search** - Show or hide items that do not match the search term.

-

**Filter by status code** - Show or hide responses based on their HTTP status code.
-

**Filter by annotation** - Only show items with comments or highlights.

### Managing the view filter settings

 The **Intruder view filter** window includes options for quickly managing the filter settings, making it easier to refine and reset your changes. The available options are:

-

**Show all** - Show all items.
-

**Hide all** - Hide any items.
-

**Revert changes** - Undo any changes you made since opening the current filter window.
