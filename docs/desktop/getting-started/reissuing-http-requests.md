> Source: https://portswigger.net/burp/documentation/desktop/getting-started/reissuing-http-requests

ProfessionalCommunity Edition

# Reissue requests with Burp Repeater

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 In this tutorial, you'll use Burp Repeater to send an interesting request over and over again. This lets you study the target website's response to different input without having to intercept the request each time. This makes it much simpler to probe for vulnerabilities, or confirm ones that were identified by Burp Scanner, for example.


#### Web Security Academy

 To follow along, you'll need an account on `portswigger.net`. If you don't have one already, [registration is free](https://portswigger.net/users/register) and it grants you full access to the Web Security Academy.


 If you haven't completed our previous tutorial on setting the target scope, you'll need to do so before continuing. See [Set the target scope](https://portswigger.net/burp/documentation/desktop/getting-started/setting-target-scope).


## Sending a request to Burp Repeater

 The most common way of using Burp Repeater is to send it a request from another of Burp's tools. In this example, we'll send a request from the HTTP history in Burp Proxy.


### Step 1: Identify an interesting request

 In the previous tutorial, you browsed a fake shopping website. Notice that each time you accessed a product page, the browser sent a `GET /product` request with a `productId` query parameter.


![Identifying an interesting request](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-product-page-request.png)

 Let's use Burp Repeater to look at this behavior more closely.


### Step 2: Send the request to Burp Repeater

 Right-click on any of the `GET /product?productId=[...]` requests and select **Send to Repeater**.


![Sending a request to Burp Repeater](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-send-to-repeater.png)

 Go to the **Repeater** tab to see that your request is waiting for you in its own numbered tab.


### Step 3: Send the request and view the response

 Click **Send** and view the response from the server. You can resend this request as many times as you like and the response will be updated each time.


![Reissuing a request in Burp Repeater](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater.png)

---

## Testing different input with Burp Repeater

 By resending the same request with different input each time, you can identify and confirm a variety of input-based vulnerabilities. This is one of the most common tasks you
 will perform during manual testing with Burp Suite.


### Step 1: Resend the request with different input

 Change the number in the `productId` parameter and resend the request. Try this with a few arbitrary numbers, including a couple of larger ones.


![Experimenting with different input](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater-invalid-productid.png)

### Step 2: View the request history

 Use the arrows to step back and forth through the history of requests that you've sent, along with their matching responses. The drop-down menu next to each arrow also lets you jump
 to a specific request in the history.


![Stepping through the request history in Burp Repeater](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater-history.png)

 This is useful for returning to previous requests that you've sent in order to investigate a particular input further.


 Compare the content of the responses, notice that you can successfully request different product pages by entering their ID, but receive a `Not Found` response if the server was unable to find a product with the given ID. Now we know how this page is supposed to work, we can use Burp Repeater to see how it responds to unexpected input.


### Step 3: Try sending unexpected input

 The server seemingly expects to receive an integer value via this `productId` parameter. Let's see what happens if we send a different data type.


 Send another request where the `productId` is a string of characters.


![Submitting unexpected input](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater-string-productid.png)

### Step 4: Study the response

 Observe that sending a non-integer `productId` has caused an exception. The server has sent a verbose error response containing a stack trace.


![Studying the error message in the response](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater-error.png)

 Notice that the response tells you that the website is using the Apache Struts framework - it even reveals which version.


![Discovering the leaked framework version](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-repeater-apache-struts.png)

 In a real scenario, this kind of information could be useful to an attacker, especially if the named version is known to contain additional vulnerabilities.


 Go back to the lab in Burp's browser and click the **Submit solution** button. Enter the Apache Struts version number that you discovered in the response (2 2.3.31).


![Submitting the solution](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-solution.png)

 Congratulations, that's another lab under your belt! You've used Burp Repeater to audit part of a website and successfully discovered an information disclosure vulnerability.


**Next step - **Running your first scan (Pro users only) [CONTINUE](https://portswigger.net/burp/documentation/desktop/getting-started/running-your-first-scan)

#### In this tutorial

1. [Downloading and installing Burp Suite.](https://portswigger.net/burp/documentation/desktop/getting-started/download-and-install)
1. [Intercepting HTTP traffic with Burp Proxy.](https://portswigger.net/burp/documentation/desktop/getting-started/intercepting-http-traffic)
1. [Modifying requests in Burp Proxy.](https://portswigger.net/burp/documentation/desktop/getting-started/modifying-http-requests)
1. [Setting the target scope.](https://portswigger.net/burp/documentation/desktop/getting-started/setting-target-scope)
1. Manually reissuing requests with Burp Repeater.
1. [Running your first scan.](https://portswigger.net/burp/documentation/desktop/getting-started/running-your-first-scan)
1. [Generating a report.](https://portswigger.net/burp/documentation/desktop/getting-started/generate-reports)
1. [What next?](https://portswigger.net/burp/documentation/desktop/getting-started/what-next)
