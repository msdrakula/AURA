> Source: https://portswigger.net/burp/documentation/desktop/burp-at/permissions

Professional

# Configuring autonomy

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Burp AT keeps you in control of what it can do independently, from checking with you before every action to running autonomously. You can set permissions at a high level using autonomy modes, or configure permissions for individual tools if you want a more granular approach. These controls enable you to leave Burp AT to work knowing it won't take an action you've restricted.


## Enabling and disabling tools

 Burp AT never uses a tool you've disabled. Tool settings work on a per-project basis, so you can disable different tools in different projects.


 To enable or disable tools:


1. Click **Settings**, then select **Tools** to open **Tool settings**.
1. On the **Enabled tools** tab, turn a tool's toggle on or off. To change a whole group at once, use the group dropdown.

 Use the **Search tools** box to find a specific tool. To restore default permissions, select **Reset to Burp's default**.


## Configuring autonomy modes

 Each Burp AT task has an autonomy mode that defines when Burp AT needs to ask for permission before using a tool:


-

**Manual** - Burp AT uses the permission set for each tool in the **Tool settings** dialog.
-

**Smart** - Burp AT approves routine tool use on its own and stops to ask you about anything that it judges to be potentially risky.
-

**Autonomous** - Burp AT runs without asking for approval, apart from a few high-impact tools that always require permission to run.

 To change the autonomy mode for a task, select an option from the **Mode** drop-down in the message box. By default, new tasks use **Smart** mode.


#### Note

**Smart** mode uses AI to judge each action. It may occasionally ask about something harmless, or act on something you'd rather it checked. If you want to approve every action yourself, use **Manual** mode and set the tools you want to supervise to **Ask**.


 Burp AT never uses tools that you've disabled in any autonomy mode.


## Configuring Manual mode

 In **Manual** mode, Burp AT uses the permission you set for each tool. To set these, select **Settings > Tools** and click the **Manual mode** tab. You can set each tool to:


-

**Ask** - Burp AT always asks for your permission before it uses the tool.
-

**Act** - Burp AT uses the tool without asking.

 The default setting for each tool varies depending on its potential impact. You can change these to suit your workflow.


#### Note

 A few high-impact tools always require your approval, whatever mode you're in - for example, running a script, adding a target to scope, or creating a custom scan check. You can't set these to **Act**, and **Always approve** isn't offered for them, so Burp AT asks every time it uses one of these tools.


 To set permissions for a whole group of tools at once, use the group dropdown.


## Approving an action

 When Burp AT needs your permission to use a tool, it pauses and explains what it wants to do. To see the action it plans to take before you decide, click **Show detail**. You can select:


-

**Approve** - Burp AT uses the tool once.
-

**Always approve** - Burp AT uses the tool and sets it to **Act** in **Tool settings**. This change applies to all tasks in your current project. This option isn't always available: it appears only in **Manual** mode when the tool is set to **Ask**, and never for tools that always require approval.
-

**Reject** - Burp AT doesn't use the tool, and finds another way to continue.

#### Note

 To make Burp AT ask for permission to use the tool again after selecting **Always approve**, set the tool back to **Ask**.


 If an approval request expires before you respond, Burp AT does not use the tool. The request shows a status of **Approval expired**, and you can continue the conversation.


## Burp AT and scope

 Burp AT respects the scope you define in Burp, and checks whether a target is in scope before it acts.


 Burp AT can edit scope, but it always asks for approval first. To prevent it from asking to edit scope, open **Tool settings** (**Settings > Tools**) and, on the **Enabled tools** tab, disable the **Add to scope** and **Remove from scope** tools.


#### More information

 For more information on defining scope, see [Target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope).


#### Next step

-

[Working with Burp AT's results](https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings)
