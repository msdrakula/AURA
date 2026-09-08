> Source: https://portswigger.net/burp/documentation/desktop/burp-at

Professional

# Burp AT

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Burp AT brings agentic AI to human-led pentesting. With direct access to Burp Suite's trusted tooling and skills from PortSwigger's world-renowned research team, Burp AT pursues the testing goals you give it, choosing the right actions for each step and adapting as it learns about your target.


![Burp AT, with the Tasks panel on the left and the conversation window on the right](https://portswigger.net/burp/documentation/desktop/images/burp-at-landing.png)

 Burp AT uses Burp's proven tools rather than improvising its own. These tools are deterministic - they behave predictably every time - so Burp AT produces consistent, repeatable results you can trust.


 Burp AT works in your open Burp project, adding to the project as it works. You can attach resources as context, such as a request, a site map node, or an issue, so Burp AT can draw on information you already have instead of rediscovering it from scratch.


#### Next step

-

[Getting started with Burp AT](https://portswigger.net/burp/documentation/desktop/burp-at/setting-up)

## Staying in control

 You have control of Burp AT's autonomy at all times. Fine-grained, per-tool permissions let you decide what it can do on its own and what it must check with you first, from requiring you to approve every action to running autonomously. These boundaries are managed by Burp's own tooling rather than depending on instructions that an AI could reinterpret or ignore. Burp AT can propose any action, but it can't carry out an action your settings don't allow.


 As it works, Burp AT explains each step it took and the reasoning behind its actions. The HTTP requests and responses it sends are captured in the Logger alongside your own, so you can follow its work and verify its findings. The result is testing you can rely on: fast, consistent, and fully accountable to you.


 PortSwigger takes the privacy and security of your data seriously. Burp AT runs on infrastructure built for security, privacy, and transparency. For information on what Burp AT sends and how your data is protected, see [AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust#burp-at).


## When to use Burp AT

 You can use Burp AT at any point in a project, taking on everything from broad goals to quick, focused tasks. You might use it to map an unfamiliar target, confirm a suspected vulnerability on a specific endpoint, reproduce a known issue to capture the evidence, or test an area for a particular class of vulnerability.


 These are starting points, not a fixed list of use cases. If you can describe a testing task, it's worth trying. A task is a good fit when you can state the goal in a sentence or two, the target and scope are defined, and you can verify the result yourself.


 Burp AT is built to extend your testing, not replace your judgment. Treat each finding as a lead to verify before you act on it.


 Burp AT can also help you build your own skills. When it does something you don't recognize, you can read its reasoning in the conversation and ask follow-up questions to explore its methodology further.


#### Note

 Burp AT uses AI credits, which are deducted from your balance as it works. For more information on how credits work, see [AI credits](https://portswigger.net/burp/documentation/desktop/burp-ai/ai-credits).


#### In this section

-

[Getting started with Burp AT](https://portswigger.net/burp/documentation/desktop/burp-at/setting-up)
-

[Tasks](https://portswigger.net/burp/documentation/desktop/burp-at/tasks)
-

[Prompting Burp AT effectively](https://portswigger.net/burp/documentation/desktop/burp-at/prompting)
-

[Tools](https://portswigger.net/burp/documentation/desktop/burp-at/tools)
-

[Skills](https://portswigger.net/burp/documentation/desktop/burp-at/skills)
-

[Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions)
-

[Working with Burp AT's results](https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings)
-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust)
