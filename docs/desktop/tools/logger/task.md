> Source: https://portswigger.net/burp/documentation/desktop/tools/logger/task

ProfessionalCommunity Edition

# Task Logger

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can examine traffic generated for a single task. This enables you to investigate a task that behaves unexpectedly, or monitor the progress of a specific task.


 To view task-specific log entries:


1. Go to **Dashboard**.
1. Select the relevant tasks from the **Tasks** list.
1. In the main panel, go to the **Logger** tab.

![Logger task view](https://portswigger.net/burp/documentation/desktop/images/getting-started/getting-started-logger-tab-5-dashboard-logger.png)

 The task-specific Logger functions in a similar way to Burp Logger, with a couple of notable differences:


- The default memory allocation for each task is only 10MB, or 20MB if you give Burp Suite access to 1GB or more of memory.
- You cannot capture or filter by tool, as the task-specific Logger only captures and displays traffic from the tool used for the task.

#### Related pages

-
 For information on how to edit what Logger captures, see [Capture filter](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/capture).

-
 For information on how to edit what Logger displays, see [View filter](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/view).
