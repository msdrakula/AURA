> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/prototype-pollution

ProfessionalCommunity Edition

# Testing for prototype pollution with DOM Invader

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Prototype pollution is a JavaScript vulnerability. It enables an attacker to add arbitrary properties to global object prototypes, which may then be inherited by user-defined objects. This enables attackers to control object properties that would otherwise be inaccessible.


 You can test for client-side prototype pollution vulnerabilities using DOM Invader. DOM Invader can automatically detect prototype pollution sources and scan for gadgets that you can use to craft an exploit. It can use the prototype pollution sources it discovers to pollute the `Object.prototype` as a proof of concept.


## Before you start

 Enable DOM Invader. For more information, see [Enabling DOM Invader](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/enabling.html).


## Steps

 You can follow along with this process in the [DOM XSS via client-side prototype pollution](https://portswigger.net/web-security/prototype-pollution/client-side/lab-prototype-pollution-dom-xss-via-client-side-prototype-pollution) Web Security Academy lab.


### Enabling prototype pollution detection in DOM Invader

1. In the upper-right corner of Burp's browser, click the Burp Suite logo and select the **DOM Invader** tab. The **Settings** menu is displayed.
1. Toggle the **DOM Invader** switch so that **DOM Invader is on**.
1. Click **Attack types** and toggle the switch so that **Prototype pollution is on**.
1. Click **Reload** to reload the browser and make your changes take effect.

### Finding potential sources for prototype pollution

1.

Right-click in the browser window and select **Inspect** to open the devtools panel.
1.

Click the **DOM Invader** tab.
1.

Browse around your target site to identify potential sources for prototype pollution. DOM Invader displays any sources found in the **Sources** list.

### Testing sources manually

1. While on the page in which the source was found, expand the **Sources** list and click **Test**. DOM Invader opens the same page in a new browser tab.
1. From the new tab, open the devtools panel and select the **Console** tab.
1. Expand the **Object** node to display the `Object.prototype`.
1. Confirm that the `Object.prototype` output now contains a property called `testproperty`.
1. Create a new object in the console using the command `let myObject = {};`.
1. Use the command `console.log(myObject.testproperty);` to view the new object. Confirm that this new object has inherited `testproperty`.

### Creating a proof of concept exploit

1. Select the source from the **Sources** list and click **Scan for gadgets**. DOM Invader opens a new tab and starts scanning.
1. Once the scan has finished, right-click in the new tab's browser window and select **Inspect** to open the devtools panel.
1. Click the **DOM invader** tab and check the contents of the **Sinks** list. These are sinks that DOM Invader was able to access via the identified gadgets.
1.
 Click **Exploit** next to a sink to test the sink with a proof-of-concept exploit. DOM Invader opens a new window in which it attempts to call the `alert()` function. If it is able to call the function, then an exploitable prototype pollution vulnerability is confirmed.


#### Related pages

-  [DOM Invader](https://portswigger.net/burp/documentation/desktop/tools/dom-invader) - Gives further information on how to use DOM Invader.

-  [DOM-based vulnerabilities](https://portswigger.net/web-security/dom-based) - Explains what the DOM is and how insecure processing of DOM data can introduce vulnerabilities.
