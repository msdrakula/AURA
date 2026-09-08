> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/evaluating-inputs

ProfessionalCommunity Edition

# Evaluating inputs with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Once you have discovered functionality that is worth investigating further, you can use a range of Burp's tools to evaluate the user controllable inputs. This enables you to determine which inputs are most important to the application's function. For example, these may be inputs that impact the session or state of the application.


 To evaluate inputs with Burp Suite, you can:


- Manually evaluate individual inputs.
- Scan the inputs.
- Fuzz the inputs.

## Before you start

 We recommend that you complete the following steps before starting this tutorial:


- Set the test scope. For more information, see [Setting the initial test scope in Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/test-scope).
- Map the target application. For more information, see [Mapping the target website with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping).
- Identify requests that you want to investigate further. For more information, see [Identifying
 high-risk functionality](https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/high-risk-functionality).

## Steps

 You can follow along with the process below using a lab with a SQL injection vulnerability. For example,
 [SQL injection in WHERE clause allowing retrieval of hidden data](https://portswigger.net/web-security/sql-injection/lab-retrieve-hidden-data).


### Manually evaluating individual inputs

 You can manually evaluate how individual inputs impact the application:


1. Send a request to Burp Repeater.
1.

Go to the **Repeater** tab and modify each input in turn. For example, you can:

  1. Remove the input.
  1. Give the input an empty value.
  1. Insert a Collaborator payload in the input. Highlight the input, then right-click and select **Insert
 Collaborator payload**.

1. Click **Send** to send the request.
1. Review the responses for noteworthy behavior, such as input reflections or differences in response times.
1.

To identify subtle changes between responses, send the responses to Burp Comparer. In Burp Comparer:

  1. Select the two responses to compare.
  1. Click **Words** or **Bytes** to compare the responses. A new window opens with the results. Comparer
 highlights any differences between the responses.

### Scanning inputs

 If you're using Burp Suite Professional, you can scan inputs to identify potential vulnerabilities, which Burp Scanner flags as issues.


 To scan a single input:


1. Highlight the input. You can do this from any message editor in Burp, for example from **Proxy** >
 **HTTP history**.
1. Right-click and select **Scan selected insertion point**. The scan launcher opens.
1. Click **OK** to start the scan using the default configuration.

 To scan multiple inputs at the same time:


1. Send the request to Burp Intruder.
1.

Go to **Intruder**. Add a payload position for each input you want to scan:

  1. Select the input.
  1.

Click **Add §**.

1. Right-click the request and select **Scan defined insertion points**. The scan launcher opens.
1. Click **OK** to start the scan using the default configuration.

 Burp Scanner audits the request using only the selected inputs. To view the issues identified by the scan:


1. Go to the **Dashboard** tab.
1. From the **Tasks** list, select the relevant scan task.
1. Go to the **Issues** tab. This contains a record of the issues found in the scan.

### Fuzzing inputs

 You can fuzz an input to identify potential vulnerabilities:


1. Highlight the input in the request, then right-click and select **Send to Intruder**.
1.

Go to **Intruder**. Notice that the input is
 automatically marked as a payload position.

![Set payload position for evaluating inputs](https://portswigger.net/burp/documentation/desktop/images/tutorials/evaluating-inputs-1.png)

1. In the **Payloads** side panel, under **Payload configuration**, add a list of fuzz strings. If you're using Burp Suite Professional, open the
 **Add from list**
 dropdown menu and select the built-in **Fuzzing - full** list. Otherwise, add your own list.
1. Click ** Start attack**. The attack starts in a new dialog. Burp Intruder sends a request for each fuzz payload.
1. When the attack is finished, study the responses to look for any noteworthy behavior.

#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater)
- [Burp Comparer](https://portswigger.net/burp/documentation/desktop/tools/comparer)
- [Burp Collaborator](https://portswigger.net/burp/documentation/desktop/tools/collaborator)
