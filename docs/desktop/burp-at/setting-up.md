> Source: https://portswigger.net/burp/documentation/desktop/burp-at/setting-up

Professional

# Getting started with Burp AT

-

**Last updated: ** September 3, 2026
-

**Read time: ** 5 Minutes

 Burp AT pursues the goals you give it, choosing the right actions for each step and adapting as it learns about your target.


 This page walks you through setting up your first task, following Burp AT's work as it runs, and verifying what it finds.


## What you need

 To use Burp AT, you need:


-

Burp Suite Professional 2026.7 or later.
-

A PortSwigger account with a valid Burp Suite Professional license associated.
-

AI credits on your account. Burp AT uses credits as it works. For more information, see [AI credits](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits).
-

Outbound HTTPS access to `ai.portswigger.net` and `login.portswigger.net` on port 443. For help with connectivity to PortSwigger's AI services, see [Troubleshooting AI connectivity](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-connectivity-troubleshooting).
-

A target you're authorized to test.

## Running your first task

### Step 1: Set your target scope

 Burp AT respects the scope you define in Burp, and checks whether a target is in scope before it acts. Set your scope before you start.


 For information on how to set scope in Burp, see [Target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope).


 Keep the scope narrow for your first task. This makes it easier to follow what Burp AT does.


 Burp AT can edit scope, but it always asks for approval first. To prevent it from asking, open **Tool settings** (**Settings > Tools**) and, on the **Enabled tools** tab, disable the **Add to scope** and **Remove from scope** tools.


### Step 2: Open Burp AT

 To open Burp AT, click **Burp AT** in the top-right corner of Burp. You can also press `Ctrl+Shift+G` (Windows/Linux) or `Cmd+Shift+G` (Mac), or select **Open Burp AT** from the command palette.


 The first time you open it, Burp prompts you to sign in to your PortSwigger account. Sign in with the account your Burp Suite Professional license is registered to. You only need to sign in once, though logins expire after two weeks of inactivity.


#### Why you need to sign in to Burp AT

 Burp AT keeps a persistent, private workspace for your data. Your tasks, goals, and findings are saved, enabling you to pick up where you left off. Because that data is sensitive and specific to you, we need you to authenticate when you access Burp AT.


 When it opens, Burp AT automatically creates a new task. A task is a conversation focused on one area of testing, and its data is stored in the Burp project you're working in.


#### Note

 For more information on tasks in Burp AT, see [Tasks](https://portswigger.net/burp/documentation/desktop/burp-at/tasks).


### Step 3: Set Burp AT's autonomy

 Each task has an autonomy mode that defines when Burp AT needs to ask for permission before using a tool. By default, tasks use **Smart** mode, in which Burp AT approves routine tool use on its own and stops to ask you about anything it judges to be potentially risky.


 To change the task's autonomy mode, use the **Mode** drop-down in the message box.


#### Note

**Smart** mode uses AI to judge each action. It may occasionally ask about something harmless, or act on something you'd rather it checked. If you want to approve every action yourself, use **Manual** mode and set the tools you want to supervise to **Ask**.


 For more information on configuring Burp AT's autonomy, including setting per-tool permissions, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


### Step 4: Give Burp AT context

 You can give Burp AT context by attaching resources from your Burp project. This usually produces a more focused, accurate result.


 To add a resource to your task:


1.

Right-click a message in **Proxy > HTTP history**, a node in your site map, or an issue, and select **Send to Burp AT**.
1.

In Burp AT, click the **+** icon in the message box to open the **Burp Resources** menu, and select the resource.

#### Note

 The requests Burp AT sends to your target, the responses it gets back, and any resources you send it may contain sensitive data. Burp AT does not redact that data before it sends requests to PortSwigger and AI providers. Only send resources that you're permitted to share.


-

For more information on how Burp AT handles your data, see [AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust#burp-at).
-

For more information on using resources, see [Attaching resources](https://portswigger.net/burp/documentation/desktop/burp-at/tasks#attaching-resources).

### Step 5: Send your first prompt

 Enter your prompt in the message box and press the return key.


 Say what you want Burp AT to work on and what you want to find out. The clearer you are, the more focused its work. Aim for an outcome you can verify yourself, so that you can confirm what Burp AT finds and report it with confidence.


 For example: `Test https://example.com for broken access control on the /admin endpoints. I can sign in as a low-privileged user with these credentials.`

 This is only an example. Burp AT is designed to handle work of any size and scope, from wide-ranging goals to focused testing tasks.


#### More information

-

[Prompting Burp AT effectively](https://portswigger.net/burp/documentation/desktop/burp-at/prompting)

### Step 6: Watch Burp AT's progress

 Burp AT plans an approach and carries it out using Burp's tools, drawing on the resources you've attached and adapting as it learns about your target. It shows its rationale as it goes.


 In **Smart** mode, Burp AT approves routine tool use on its own. If it needs your approval to take an action, it pauses and explains what it wants to do. Click **Show details** to see the action it plans to take.


#### Note

 A few high-impact tools, such as running a script, adding a target to scope, or creating a custom scan check, always require your approval.


### Step 7: Verify what it found

 Burp AT records vulnerabilities it finds as issues in Burp, and logs its actions so you can verify its work.


 When the task finishes:


-

Read back through the conversation to see which tools Burp AT used, and click **Show details** on any action to see further information.
-

Review any issues Burp AT raised. Select an issue and click **Request** or **Response** to see the traffic used to find that issue.

 Burp AT is built to extend your testing, not replace your judgment. Treat each finding as a lead to verify before you act on it.


 A task is an ongoing conversation, so you can send follow-up prompts to dig into a result, adjust what you asked for, or point Burp AT at something new.


#### More information

-

[Working with Burp AT's results](https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings)
