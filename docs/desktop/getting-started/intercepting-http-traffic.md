> Source: https://portswigger.net/burp/documentation/desktop/getting-started/intercepting-http-traffic

ProfessionalCommunity Edition

# Intercept HTTP traffic with Burp Proxy

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 In this tutorial, you'll use a live, deliberately vulnerable website to learn how to intercept requests with Burp Proxy.


## Intercepting a request

 Burp Proxy lets you intercept HTTP requests and responses sent between Burp's browser and the target server. This enables you to study how the website behaves when you perform different actions.


### Step 1: Launch Burp's browser

 Go to the **Proxy > Intercept** tab.


 Set the intercept toggle to **Intercept on**.


![Intercept on](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-intercept-on.png)

 Click **Open Browser**. This launches Burp's browser, which is preconfigured to work with Burp right out of the box.


 Position the windows so that you can see both Burp and Burp's browser.


### Step 2: Intercept a request

 Using Burp's browser, try to visit `https://portswigger.net` and observe that the site doesn't load. Burp Proxy has intercepted the HTTP request that was issued by the browser before
 it could reach the server. You can see this intercepted request on the **Proxy > Intercept** tab.


![Viewing an intercepted request in Burp Proxy](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-intercepted-request.png)

 The request is held here so that you can study it, and even modify it, before forwarding it to the target server.


### Step 3: Forward the request

 Click the **Forward** button to send the intercepted request. Click **Forward** again to send any subsequent requests that are intercepted, until the page loads in Burp's browser. The
 **Forward** button sends all the selected requests.


### Step 4: Switch off interception

 Due to the number of requests browsers typically send, you often won't want to intercept every single one of them. Set the intercept toggle to **Intercept off**.


![Proxy Intercept off](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-intercept-off.png)

 Go back to the browser and confirm that you can now interact with the site as normal.


### Step 5: View the HTTP history

 In Burp, go to the **Proxy > HTTP history** tab. Here, you can see the history of all HTTP traffic that has passed through Burp Proxy, even while intercept was switched off.


 Click on any entry in the history to view the raw HTTP request, along with the corresponding response from the server.


![Viewing the HTTP history in Burp Proxy](https://portswigger.net/burp/documentation/desktop/images/getting-started/quick-start-pro-proxy-history.png)

 This lets you explore the website as normal and study the interactions between Burp's browser and the server afterward, which is more convenient in many cases.


**Next step - **Modifying HTTP requests with Burp Proxy [CONTINUE](https://portswigger.net/burp/documentation/desktop/getting-started/modifying-http-requests)

#### In this tutorial

1. [Downloading and installing Burp Suite.](https://portswigger.net/burp/documentation/desktop/getting-started/download-and-install)
1. Intercepting HTTP traffic with Burp Proxy.
1. [Modifying requests in Burp Proxy.](https://portswigger.net/burp/documentation/desktop/getting-started/modifying-http-requests)
1. [Setting the target scope.](https://portswigger.net/burp/documentation/desktop/getting-started/setting-target-scope)
1. [Manually reissuing requests with Burp Repeater.](https://portswigger.net/burp/documentation/desktop/getting-started/reissuing-http-requests)
1. [Running your first scan.](https://portswigger.net/burp/documentation/desktop/getting-started/running-your-first-scan)
1. [Generating a report.](https://portswigger.net/burp/documentation/desktop/getting-started/generate-reports)
1. [What next?](https://portswigger.net/burp/documentation/desktop/getting-started/what-next)
