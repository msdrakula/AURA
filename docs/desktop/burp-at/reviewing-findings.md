> Source: https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings

Professional

# Working with Burp AT's results

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Burp AT records vulnerabilities it finds as issues in Burp, logging its actions so you can verify its work. This page explains how to access and use the findings from a Burp AT task.


#### More information

 For more information on what Burp issues contain and how they work, see [Issues](https://portswigger.net/burp/documentation/desktop/running-scans/results/issues).


## Reviewing issues

 You can review issues raised by Burp AT wherever Burp shows issues:


-

**In the conversation** - Burp AT reports each finding as it works, so you can see it in context straight away.
-

**In the Dashboard** - Issues Burp AT raises when scanning appear on the **Issues** tab of the scan task. To see every issue in the project, go to the **Dashboard** tab and select **All issues** from the bottom dock.
-

**On the site map** - Each issue is linked to the corresponding site map entry, so it shows against the relevant node in **Target > Site map**.

 Burp AT sets the severity and confidence of issues itself. We recommend you confirm these before acting on them.


 For more information on where Burp AT's work appears across Burp, see [Picking up Burp AT's work in other Burp tools](https://portswigger.net/burp/documentation/desktop/burp-at/tasks#picking-up-burp-at-s-work-in-other-burp-tools).


## Verifying issues

 As with "regular" Burp issues, issues reported as a result of a Burp AT task include the requests and responses used to find them. This enables you to easily reproduce issues, confirm they're exploitable, and report them with confidence. Select an issue and click **Request** or **Response** to see this information.


 You can send requests from Burp AT to other Burp tools to do further work on manually. For example, you might send a request to Repeater to reproduce an issue, or to Intruder to fuzz the request. For more information, see [Sending a request to another Burp tool](https://portswigger.net/burp/documentation/desktop/burp-at/tasks#sending-a-request-to-another-burp-tool).


 Burp AT does not have prior knowledge of how your application is meant to work, so it may report intended behavior as a vulnerability. For example, it may flag a search feature as an injection point.


## Editing and deleting issues

 You can review, edit, and delete issues raised by Burp AT.


 To change or remove an issue, right-click it in the Dashboard or site map and select:


-

**Set severity** - Reassign the issue's severity level. You can also flag the issue as a false positive.
-

**Set confidence** - Reassign the issue's confidence level.
-

**Delete issue** - Delete selected issues from the table.

 To restore severity or confidence values you've changed, right-click an issue and select **Restore original value**.


#### Note

 Burp AT can update its own issues as it works, for example by lowering severity or deleting an issue it no longer believes to be true.


 For more information on managing issues manually, see [Issues](https://portswigger.net/burp/documentation/desktop/running-scans/results/issues#managing-issues).


## Including issues in your report

 You can include issues raised by Burp AT in a report in the same way you would with issues discovered during a standard scan.


 To report an issue, right-click it in the Dashboard or site map and select **Report issue**.


 Burp AT presents the issues it found as a table in the task conversation. You can export this to use in your own reporting:


-

To copy the table to your clipboard, click the copy icon  and select **Markdown**, **CSV**, or **TSV**.
-

To save the table as a file, click the download icon  and select **CSV** or **Markdown**.

#### Related pages

-

[Reporting scan results](https://portswigger.net/burp/documentation/desktop/running-scans/reporting) - Gives detailed information on generating a report and exporting issue data.
-

[Generating a report](https://portswigger.net/burp/documentation/desktop/getting-started/generate-reports) - Walks through the reporting wizard.
-

[Manually creating issues for reports](https://portswigger.net/burp/documentation/desktop/running-scans/reporting/manual-issues) - Explains how to add issues you found yourself.

#### Next step

-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust)
