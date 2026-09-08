> Source: https://portswigger.net/burp/documentation/desktop/burp-at/tasks

Professional

# Tasks

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 Burp AT organizes work into tasks. A task is a conversation focused on one area of testing, with its own messages and results. Tasks can split into parallel sub-tasks if needed.


 Task data is stored in the Burp project you are working in.


## Starting a new task

 To start a new task, click the **+** icon at the top of the **Tasks** panel.


 You can run multiple tasks at once, with each running independently. The **Tasks** panel shows the current project's tasks, with the most recently-updated at the top. To return to an earlier task, select it in the **Tasks** panel.


#### Note

 Each task runs in one Burp instance at a time. If you open a task in a new instance, control of the task moves to that instance.


## Managing running tasks

 A task is an ongoing conversation, meaning that you can send follow-up prompts to dig into a result, adjust what you asked for, or point Burp AT at something new.


 Burp AT shows its rationale as it works. To see more information about an action Burp AT took, click **Show details** in the conversation next to the action. Expanding this drop-down shows the tool and inputs used, and, for actions that send HTTP traffic, the request and response themselves. It also shows how the action was approved: run automatically, approved manually, or handled by **Smart** mode.


 When Burp AT needs your approval for an action, it shows a dialog in the conversation. For more information on how approvals and tool permissions work in Burp AT, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


 To stop Burp AT's current action, click **Stop**. The task stays open, but any actions that are currently running end.


#### Note

 If an action fails, Burp AT explains what went wrong and suggests how to continue. If it stalls or drifts off track, consider stopping it and starting a new task with a refined goal.


## Picking up Burp AT's work in other Burp tools

 Burp AT works in the Burp project you have open. This means you can follow what it's doing in many of Burp's tools, and conduct further manual testing at any point.


### Where Burp AT's activity appears

 As Burp AT works, its activity appears in:


-

**Logger** - every request Burp AT sends, with the accompanying response.
-

**Dashboard** - any scans Burp AT starts. You can pause, resume, and inspect these from the **Tasks** panel on the Dashboard.
-

**Target > Site map** - the hosts and endpoints it discovers.
-

**Issues** - the vulnerabilities it confirms. For more information on how Burp AT reports issues, see [Working with Burp AT's results](https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings).

### Sending a request to another Burp tool

 You can send requests created by Burp AT directly to another Burp tool to use them in your manual testing. To do this:


1.

In the conversation, click **Show details** next to the action that sent the request to display the **HTTP exchange** panel.
1.

Click **Send to**.
1.

Select the tool you want to send the request to.

## Managing tasks

 To rename a task, right-click it in the **Tasks** panel and select **Rename**, or double-click its name.


 To remove a task from the list, right-click it in the **Tasks** panel and select **Archive**.


 To restore an archived task:


1.

Go to **Settings > Archived tasks**.
1.

Click **Restore** on the task you want to bring back.

## Attaching resources

 You can give Burp AT extra context by attaching resources to a task. When you attach a resource, Burp AT can use its information straight away instead of having to rediscover information. Attaching relevant resources usually produces a more focused, accurate result.


 You can send resources to Burp AT from:


-

**Proxy** - a message from **Proxy > HTTP history**.
-

**Target** - a node from your site map.
-

**Issues** - issues discovered by the scanner.

 To share a resource, right-click it in the relevant Burp tool and select **Send to Burp AT**. You can also press `Ctrl+G` (Windows/Linux) or `Cmd+G` (Mac), or select **Send to Burp AT** from the [command palette](https://portswigger.net/burp/documentation/desktop/tools/command-palette).


 Shared resources are available to every task in the open project. Sending more resources adds to this shared pool rather than replacing what's there.


 To attach a shared resource to a task, click the **+** icon in the message box to open the **Burp Resources** menu and select the item from the list. Attached resources appear as a tile above the message box.


## Sub-tasks

 Burp AT can ask for permission to create sub-tasks as it works. These are separate processes that work on different parts of the goal at the same time. Each sub-task has its own conversation window, enabling you to see results and direct the sub-task if required. Splitting a goal into sub-tasks enables Burp AT to cover larger goals faster and more accurately than working through everything in one thread.


 Sub-tasks do not communicate with each other directly, but each can see the contents of the current Burp project.


 Burp AT lists sub-tasks in the **Tasks** panel. When you select a task, Burp AT displays any related sub-tasks in a drop-down menu. It also shows them in the **Sub-tasks** list on the main task conversation. Each sub-task shows a name and status indicating whether it needs human intervention, for example to approve a tool call or resolve an error.


#### Next step

-

[Prompting Burp AT effectively](https://portswigger.net/burp/documentation/desktop/burp-at/prompting)
