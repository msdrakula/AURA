> Source: https://portswigger.net/burp/documentation/desktop/external-browser-config/check-listener

ProfessionalCommunity Edition

# Check that Burp's proxy listener is active

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp's proxy listener is a local HTTP proxy server that listens for incoming connections from your browser. It allows you to monitor and intercept all HTTP requests and responses sent and received by your browser. This lies at the heart of Burp's user-driven workflow.


 By default, Burp creates a single listener on port 8080 of the loopback interface. The first time you start Burp, you need to check that this listener is active and running.


 In Burp, go to the **Tools > Proxy** tab in the **Settings** dialog.


 In the **Proxy listeners** panel, you should see an entry for the interface `127.0.0.1:8080` with the **Running** checkbox selected, indicating that the listener is active and running. If so, everything is fine and you can move on to [configuring your browser](https://portswigger.net/burp/documentation/desktop/external-browser-config).


![Check the proxy listener is running](https://portswigger.net/burp/documentation/desktop/images/check-proxy-listener-running-940.png)

 If not, click the  icon in the **Proxy listeners** field and select **Restore defaults**. Look at the **Running** checkbox again to see if the proxy listener is now running. If the checkbox is now selected, the listener is active. Otherwise, continue with the following steps.


![Proxy listener restore defaults](https://portswigger.net/burp/documentation/desktop/images/ca-check-proxy-restore-defaults-940.png)

 If it is still not running, the default port 8080 might not be available, for example, because it is already being used by another application. In this case, you need to try using a different port. Select the entry `127.0.0.1:8080` and click the **Edit** button. The **Edit proxy listener** dialog opens.


 In the **Bind to port** field, enter the number of a new port that you think is free and click **OK**.


![Bind to port](https://portswigger.net/burp/documentation/desktop/images/check-proxy-bind-port-940.png)

 Try to activate the listener by selecting the **Running** checkbox. If you still can't activate it, then the new port that you selected is probably also blocked. Repeat this process to try a different port.


![Activate listener with a new port](https://portswigger.net/burp/documentation/desktop/images/check-proxy-port-8081-940.png)
