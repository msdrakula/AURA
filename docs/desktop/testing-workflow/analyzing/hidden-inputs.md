> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/hidden-inputs

ProfessionalCommunity Edition

# Checking for hidden inputs with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Hidden inputs are HTTP headers, cookies, or parameters that affect how a site responds to requests but are not used during normal interaction with the site. Identifying hidden inputs can provide you with additional ways to trigger unintended behavior from your target.


 You can discover potential hidden inputs in Burp Suite using the Param Miner extension. Param Miner uses a built-in wordlist and information taken from the scope to guess at potential hidden inputs. It then sends a request containing each input, followed by another request with this input omitted. If the responses differ, this suggests that the tested header, cookie, or parameter is supported by the website.


## Before you start

 We recommend that you complete the following steps before starting this tutorial:


- Install the Param Miner BApp. For more information, see [Installing extensions from the BApp Store](https://portswigger.net/burp/documentation/desktop/extend-burp/extensions/installing/bapp-store).
- Set a test scope. For more information, see [Setting the initial test scope in Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/test-scope).
- Map the visible attack surface of your target site. For more information, see [Mapping the visible attack surface with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/visible-attack-surface).

## Steps

 You can follow along with the process below using [Gin and Juice Shop](https://ginandjuice.shop), our deliberately vulnerable demonstration site. This site has hidden inputs that Param Miner can discover.


 To check for hidden inputs with Param Miner:


1. In Burp Suite, open **Target** > **Site map**.
1. Select the request you want to run Param Miner against. You can select multiple requests if required.
1. Select **Extensions** > **Param Miner** > **Guess params**.
1.

Select the type of hidden inputs you want to Param Miner to guess:

  - **Guess GET parameters**
  - **Guess cookie parameters**
  - **Guess headers**
  - **Guess everything!**

1. On the **Attack Config** dialog, click **OK**. If you selected **Guess everything!** you need to click
 **OK** a few times to close the dialog. Param Miner sends a series of requests to the target.
1. To view the results of the test, select **Extensions** > **Installed** >
 **Param Miner** > **Output**. This tab display a log of Param Miner's run, including any hidden inputs identified.

 In Burp Suite Professional, hidden inputs also appear as a **Secret input** issue in the **Dashboard** tab. Select these issues to view the relevant requests and responses.


#### Related pages

- [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser)
- [Site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map)
- [Practical Web Cache Poisoning](https://portswigger.net/research/practical-web-cache-poisoning) - a blog post that explains how to use Param Miner to find web cache poisoning vulnerabilities.
