> Source: https://portswigger.net/burp/documentation/desktop/running-scans/results/event-log

ProfessionalCommunity Edition

# Event log

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Key events are added to the event log as your tasks progress. This includes start times, alerts, and progress updates. You can use this information to troubleshoot problems with your scans.


 You can choose to view either a task-specific event log, which contains events triggered by an individual task, or a project-level event log containing entries for all events that have occurred across all tasks in your project.


 To view the project-level event log:


1. Go to the **Dashboard** tab.
1. From the bottom dock, select **Event log**.

 To view a task-specific event log:


1. Go to the **Dashboard** tab.
1. From the **Tasks** list, select the relevant task.
1. In the main panel, go to the **Event log** tab.

 Each item in the event log contains the following details:


- **Time** - The time of the event.
- **Type** - Critical, error, info, or debug.
- **Source** - The task number or tool that triggered the event.
- **Message** - A summary of the event.

 You can customize and sort the table, and copy column data to your clipboard. For more information, see [Customizing Burp's tables](https://portswigger.net/burp/documentation/desktop/burps-layout/customize-tables).


 Select an event to view further information. The **Event detail** panel includes a description of the event and any recommended remedial action.


 To filter the event log by event type, use the buttons at the top of the log. To filter the events by a specific term, for example a task number, use the search function.


 To delete all items in the event log, right-click any entry and select **Clear event log**.


#### Related pages

[Burp Scanner error reference](https://portswigger.net/burp/documentation/scanner/burp-scanner-error-reference) - Provides additional information about some of the error messages that may appear in the event log, as well as troubleshooting tips.
