> Source: https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample/live-capture

ProfessionalCommunity Edition

# Burp Sequencer live capture

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 When you start a live capture, Burp Sequencer repeatedly issues the request and extracts the relevant token from the application's responses. This occurs in a new results window.


 The results window contains a progress bar, and real-time details of the:


- Number of requests made.
- Number of tokens captured.
- Number of errors found.

![Sequencer live capture](https://portswigger.net/burp/documentation/desktop/images/sequencer-live-capture.png)

 The buttons in the results window enable you to control aspects of the capture:


- **Pause** or **Resume** - Temporarily pause and resume the capture.
- **Stop** - Permanently stop the capture.
- **Copy tokens** - Copy the captured tokens to the clipboard.
- **Save tokens** - Save the captured tokens as a file.
- **Analyze now** - Analyze the current sample of tokens. You need a minimum of 100 captured tokens to perform the analysis.
- **Auto-analyze** - Perform automatic analysis periodically during the live capture.

 You can use the copied or saved tokens as payloads for Burp Intruder or Burp Scanner. You can also manually load them in a future Sequencer analysis.


#### Related pages

- For instructions on how to set up a live capture, see [Obtaining a token sample](https://portswigger.net/burp/documentation/desktop/tools/sequencer/sample).
- For information on live capture settings, see [Sequencer settings](https://portswigger.net/burp/documentation/desktop/settings/tools/sequencer#live-capture).
- To learn how the results are displayed, see [Burp Sequencer results](https://portswigger.net/burp/documentation/desktop/tools/sequencer/results).
