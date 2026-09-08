> Source: https://portswigger.net/burp/documentation/desktop/burp-at/skills

Professional

# Skills

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Skills are focused testing techniques written by PortSwigger's Research team that extend what Burp AT knows how to do.


 Burp AT uses skills automatically. As it works, it evaluates your goal and its testing methodology, and decides if and when to invoke skills from its library.


## What are skills?

 A skill is a combination of a prompt and accompanying code. This combination gives Burp AT context on how to approach a problem, and enables it to tackle that problem in a reliable, repeatable way. When using a skill, Burp AT follows a methodology built for that specific task, rather than assembling an approach from general knowledge. Skills give Burp AT a growing library of proven techniques to draw on, taken from PortSwigger's cutting-edge research.


 Skills are published and updated automatically. You don't need to update Burp to take advantage of new skills.


## Skills and permissions

 Skills can potentially combine multiple Burp tools. For example, a single skill might use Repeater and Intruder as part of the same technique.


 Skills are subject to the same permissions and autonomy settings as the rest of Burp AT. If a skill attempts to use a tool that requires approval, you are prompted in the same way as you would be if Burp AT was using those tools outside of a skill.


#### More information

 For more information on controlling what Burp AT can do on its own, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


## Seeing when a skill is used

 When Burp AT uses a skill, it shows a banner under its response. Expand this banner to see a list of the skills that Burp AT ran when generating the response. Select an individual skill to see an overview of what the skill did and the tools it called.


#### More information

 For more information on following and verifying Burp AT's work, see [Working with Burp AT's results](https://portswigger.net/burp/documentation/desktop/burp-at/reviewing-findings).


## Disabling a skill

 Skills are enabled by default. To disable a skill, click the options menu  next to the skill where it appears in the task, and select **Disable skill**. Disabling a skill applies across all tasks in the current project and takes effect the next time you send a message to Burp AT.


 To re-enable a skill, select **Settings > Skills** and click **Re-enable** on the skill.


#### Note

 Disabling a skill only stops Burp AT from using that technique. It doesn't change your tool permissions or autonomy settings.


#### Next step

-

[Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions)
