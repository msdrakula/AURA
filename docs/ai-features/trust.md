> Source: https://portswigger.net/burp/documentation/ai-features/trust

DASTProfessional

# AI trust and data handling

-

**Last updated: ** September 3, 2026
-

**Read time: ** 11 Minutes

 This page answers common questions about how Burp AT and Burp AI keep your testing safe, how you stay in control when using AI features, and how PortSwigger handles the data you send.

 Both Burp AT and Burp AI run on PortSwigger's AI infrastructure. The first section covers topics that apply to both. The sections that follow give specific information on each product.

## Shared infrastructure and data

#### How does Burp send and process your data?

 When you use an AI feature, Burp sends your data to PortSwigger's AI infrastructure, which manages communication with the AI provider:


1.

**Request:** Your data, such as an HTTP request or vulnerability description, is sent to PortSwigger's AI infrastructure. In Burp Suite Professional, this is sent from your desktop client. In Burp Suite DAST, it is sent from a dedicated scanning resource created specifically to handle the investigation.
1.

**Authentication:** PortSwigger verifies your license. Only licensed users can access the AI service.
1.

**Task processing:** A temporary processing agent is created to handle your request. It is decommissioned immediately after the task is complete, and is never reused across requests or users.
1.

**Provider transmission:** The agent passes your data to the AI provider.
1.

**Provider processing:** The AI provider processes your data and returns a response to Burp. It does not retain any of your data once processing is complete.
1.

**Storage:** PortSwigger stores your prompt, the AI response, and related metadata to support troubleshooting, auditing, and billing.

 All communication between Burp, PortSwigger, and the AI providers is encrypted using TLS 1.2 or later.


#### What data does PortSwigger store?

 PortSwigger stores your prompts, AI responses, and associated metadata to support troubleshooting, auditing, billing, and account management. This includes:


-

Any data you send as context, such as highlighted HTTP headers or code snippets.
-

The full conversation for a Burp AT task, including the agent's messages, its tool calls, and the results those tools return.
-

Any resources you send to Burp AT, such as HTTP messages, site map nodes, and scanner issues.
-

Task metadata, such as which feature was used, issue details, and target URLs.
-

The timestamp at which the request was started.
-

Your license details, token counts, credit consumption, and request counts.

 All stored data is encrypted using AES-256. Only authorized PortSwigger staff can access it, and only when necessary to support troubleshooting or resolve an issue. Access is managed through role-based permissions, and every internal interaction with your data is individually logged.


#### How long is my data retained?

 Retention depends on the type of data:


-

Conversation and audit data (the record of your prompts, AI responses, and associated task details) is retained indefinitely to support troubleshooting and long-term feature development.
-

Operational logs, such as error and infrastructure logs, are deleted after 31 days.

#### Do AI providers retain or train on my data?

 No. Our contracts with our AI providers prohibit them from monitoring your data, storing it, or using it for model training. Your data passes through provider infrastructure while a request is being processed, but the provider does not retain it once the request is complete, and our contracts prohibit any secondary use.


#### Does PortSwigger use my data to improve its AI?

 PortSwigger reserves the right to use anonymized data to improve its AI features and diagnose issues.


#### Which AI models are used, and can I choose?

 Not currently. PortSwigger uses a multi-model approach to get the best results across a wide range of security testing tasks, and selects the model for each task based on extensive performance testing, matching each feature to the model that delivers the most accurate and reliable results. Burp AI uses models from OpenAI and Anthropic. Burp AT uses Anthropic models.


 We only use models that meet our strict data privacy standards, and providers that can guarantee your data is not used for training. Our infrastructure manages the connection to these models, so we can upgrade model versions after assessing them. These updates happen automatically, giving you access to current AI versions without the need to update Burp.


#### Where is data processed?

 PortSwigger's AI infrastructure runs on AWS in the US-East (Virginia) and EU-West (Ireland) regions. Our AI providers process requests in US data centers. Routing between the US and EU regions is handled automatically. You cannot currently choose a specific geographic location.


#### What security controls protect the AI infrastructure?

 Our AI infrastructure uses server- and client-side input validation, and token-based authentication that limits access to licensed users.


 Burp AT checks each tool call against your testing scope and the permissions you've set before it runs. Burp's own tooling enforces these limits, rather than leaving them to the AI. For more information, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


#### What happens if there's a security incident involving an AI provider?

 In this case, PortSwigger would respond according to its ISO 27001-aligned incident response process and its contractual obligations with the affected provider. AI usage is logged internally to support our investigations.


#### What happens if there's an outage?

 We operate failover systems to keep the AI features available. If the infrastructure serving a request becomes unavailable, requests are automatically routed to other infrastructure to maintain service.


#### Is Burp's AI covered by PortSwigger's compliance and Data Processing Agreement?

 Yes. PortSwigger's AI is covered by PortSwigger's ISO 27001 certification, and PortSwigger's [Data Processing Agreement](https://portswigger.net/data-processing-agreement) covers all personal data processed on your behalf. For full details, see the [PortSwigger Trust Center](https://trust.portswigger.net/).


## Burp AT

 This section covers trust and safety questions specific to Burp AT.

#### Does Burp AT run automatically?

 No. Burp AT only runs when you open it and give it a prompt. Otherwise it takes no action, and sends no data to PortSwigger or our AI providers. Once running, it works within the autonomy levels you set, and asks for approval on anything you haven't configured it to be able to do on its own.


#### Can Burp AT take actions without my approval?

 Yes, but only if you configure it to. You have full control over what each tool can do without approval. When a tool call needs approval, Burp AT pauses and waits for your decision. For more information, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


#### Can Burp AT test out of scope targets?

 Burp AT works within the scope you set, and checks whether a target is in scope before it acts.


 Burp AT can edit scope if you allow it, using its **Add to scope** and **Remove from scope** tools. You can restrict or disable these individually in **Settings > Tools**. For details on what each tool can and can't do, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


 This scope check applies to Burp AT's own tool actions. Custom code that it runs through its scripting tools can send requests directly, including to targets outside your defined scope. If you need to prevent this, disable the scripting tools in **Settings > Tools**.


#### Could Burp AT disrupt a production target?

 Burp AT sends the same kinds of requests to your target that you would when testing manually - it has no ability to affect a site beyond what Burp's tools can already do. It only acts within the scope you define, which is enforced by Burp's own tooling rather than by the AI. High-impact actions always require your approval.


 For tight control on a sensitive target, run the task in **Manual** mode and set the tools you want to supervise to **Ask**. For more information, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


#### Can Burp AT send destructive requests?

 Yes, if you allow it to. Burp AT can send any request you could send manually, including one that changes or deletes data.


 Burp AT only acts within the scope you define, which is enforced by Burp's own tooling rather than by the AI. Some high-impact tools always require your approval.


 To review every request before Burp AT sends it:


1.

Go to **Settings > Tools** to open **Tool settings**.
1.

Go to the **Manual mode** tab and set **Send request** to **Ask**.
1.

Select **Manual** from the **Mode** drop-down in the message box.

 Burp AT then asks for your approval each time it wants to send a request. Per-tool permissions apply only in **Manual** mode.


 Custom code that Burp AT runs through its scripting tools sends requests directly. Those requests don't pass through the **Send request** tool. Running a script always requires your approval, so we strongly recommend that you read the script before you approve it. To block Burp AT from running scripts, disable the scripting tools in **Settings > Tools**.


#### Can I stop Burp AT mid-task?

 Yes. While Burp AT is working, click **Stop** to stop it immediately.


#### Can I disable Burp AT?

 Yes. To turn off all of Burp's AI features, go to **Settings > AI** and select **Disable AI features**. This stops all AI communication, including Burp AT, and grays out AI options throughout Burp.


 To limit what Burp AT can do without turning it off completely, consider disabling individual tools. For more information, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


#### Can data from one project leak into another?

 No. Each task is part of the Burp project you're working in. Burp AT can only see and act on the data in that project, so work from one project can't be seen by another.


#### What requests does Burp AT send to my target?

 Burp AT sends requests to the targets you've told it to test, in the same way you would when testing manually. These requests go to your target directly from your local installation of Burp. PortSwigger's AI infrastructure does not connect to your target.


 Burp AT captures its traffic in the Logger, so it is fully visible to you. It also attaches the relevant requests and responses to every issue it raises, so you can see exactly what happened and reproduce any findings.


#### What does Burp AT send to PortSwigger and AI providers?

 Burp AT sends the information the agent needs to reason about your task:


-

Your prompts, including the testing goals you set.
-

Any resources you send to Burp AT, such as a message from **Proxy > HTTP history**, a node from your site map, or an issue found by the scanner.
-

The agent's own messages.
-

Each tool call the agent makes, and the arguments it passes to that tool.
-

The result each tool returns. For tools that send HTTP requests, this includes the full request and response.
-

Task metadata, such as token counts and the time each request was started.

 Burp AT does not upload your Burp project. Its tools run locally in Burp, and Burp returns only the results of the tool calls the agent makes.


 PortSwigger stores the conversation as your task runs, and includes it in requests to the AI provider to enable the agent to keep track of what it has already done. For details of what PortSwigger stores, see [What data does PortSwigger store?](https://portswigger.net/burp/documentation/ai-features/trust#what-data-does-portswigger-store)

#### Does Burp AT send sensitive data?

 It can, depending on what you're testing. The requests Burp AT sends to your target, the responses it gets back, and any resources you send it may contain sensitive data. Burp AT does not redact that data before it sends requests to PortSwigger and AI providers.


 You can configure exactly what Burp AT can do without approval. If you need to check what leaves your machine, configure it to require manual approval before it sends any request. For more information, see [Can Burp AT send destructive requests?](https://portswigger.net/burp/documentation/ai-features/trust#can-burp-at-send-destructive-requests)

## Burp AI

 This section covers trust and safety questions specific to Burp AI, the suite of assistive AI tools built into Burp Suite Professional and Burp Suite DAST.

#### Does Burp AI run automatically?

No. Burp AI requires explicit setup before it can run:

- **Burp Suite Professional:** a feature only runs when you actively use it. No data is sent to PortSwigger or our AI providers in the background.
- **Burp Suite DAST:** AI features are not available by default. An administrator must enable them for your installation before they appear in the UI. Once enabled, they only run against sites you have configured.

#### Can I disable Burp AI?

Yes, in both Burp Suite Professional and Burp Suite DAST:

- **Burp Suite Professional:** go to **Settings > AI** and select **Disable AI features**. Burp then cannot access PortSwigger's AI infrastructure, and all AI-related features are grayed out in the UI.
- **Burp Suite DAST:** access is managed at the organization level by an administrator. Go to **Settings > Burp AI** and toggle **Enable Burp AI**. You need the **Enable/Disable Burp AI for the installation** permission to change this setting. When disabled, AI features are removed from the UI and any AI results from previous scans are hidden.

#### Can I stop a Burp AI feature mid-task?

Yes, in both Burp Suite Professional and Burp Suite DAST:

- **Burp Suite Professional:** for multi-step processes such as Explore Issue or Repeater tasks, pause or stop execution at any time from the **Tasks** panel.
- **Burp Suite DAST:** to stop AI-enhanced scanning, cancel the scan.

#### Can I control which targets are analyzed?

 Yes. You can use Burp's scope settings to restrict Burp AI to specific hosts and URLs. In DAST, you can also control which issue severity and confidence combinations are investigated on a per-site basis.


- For scope in Burp Suite Professional, see [Target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope).
- For scope in Burp Suite DAST, see [Configuring AI-enhanced scanning](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/configuring-ai-enhanced-scanning).

#### Does Burp AI send sensitive data?

 It can. Burp AI does not automatically redact or mask the data sent to AI providers. Depending on the features you use and the traffic you're testing, this may include sensitive data.


 As an exception, AI-generated recorded logins replace actual usernames and passwords with placeholders before they are sent to PortSwigger or the AI providers. This is the same behavior as manually-recorded logins.


#### What does each Burp AI feature send?

All Burp AI features use the same infrastructure, but the data sent varies by feature:

- **Explainer:** the highlighted text or code snippet, and the context type (for example, `REQUEST_HEADERS` or `RESPONSE_BODY`).
- **Explore Issue / AI-enhanced scanning:** the full HTTP request and response pair in which the issue was found, including headers; the issue name, target URL, and insertion point; and, in Burp Suite DAST only, the origin scan ID and site ID.
- **Burp AI in Repeater:** the full HTTP request and response pair from the Repeater tab you ran the task from, including headers; your prompt; and any notes attached to the tab.
- **False positive analysis:** a screen capture of the vulnerability finding, and the associated HTTP request and response.
- **Broken access control:** a screenshot of the page content returned during the access check.
- **Recorded logins:** the page text, HTML attributes of interactive elements, and a screen capture of the page.

#### How does Burp AI stay focused on security testing?

 Burp AI uses structured prompts to keep the AI focused on security testing, reducing the risk of hallucinations or unintended actions. There are also built-in limits on the number of steps the AI can take during a task.


 In Burp Suite DAST, AI-enhanced scanning has additional guardrails that block destructive HTTP methods (such as DELETE and PUT) and SQL operations that could cause data loss. These run automatically and cannot be disabled by users.
