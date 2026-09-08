> Source: https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/testing

ProfessionalCommunity Edition

# Testing HTTP match and replace rules

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 When adding or editing a HTTP match and replace rule, you can test your rule using the built-in test function. This enables you to confirm that the rule correctly matches and replaces the intended text.

 To test a HTTP match and replace rule in the match/replace rule editor:

1. Review the sample message under **Original request** or **Original response**. Optionally, replace this sample message with the specific request or response you'd like to test the rule against.
1. Click **Test**. Burp applies the rule to the original message, creating a modified request or response.
1. Review the modified request or response under **Auto-modified request** or **Auto-modified response**.
1. Adjust the rule as necessary.

To restore the sample request or response, click .
