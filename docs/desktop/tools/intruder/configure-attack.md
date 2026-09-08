> Source: https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack

ProfessionalCommunity Edition

# Configuring Burp Intruder attacks

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 When you send an HTTP request to Burp Intruder, it opens in a new attack tab. Burp Intruder enables you to insert payloads into defined positions in an HTTP request, then send each version of the request to the target server. You can configure various aspects of the attack:


-  [Payload positions](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/positions) - The locations in the base request where payloads are placed.

-  [Attack type](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/attack-types) - The algorithm for placing payloads into your defined payload positions.

-  [Payload type](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-types) - The type of payload that you want to inject into the base request. You can use a simple wordlist, but Burp Suite also provides a range of options for auto-generating payloads. Burp Suite Professional includes a range of [predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists) for use with compatible payload types.

-  [Payload processing](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/processing) - Rules to manipulate each payload before it is used.

-  [Resource pool](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/resource-pool) - The allocation of resources to the attack.

-  [Attack settings](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings) - Burp Intruder attack settings.


 You can use the top-level **Intruder** menu to save the attack configuration, or load it in a future attack. Alternatively you can copy the attack configuration into any open tab. For each function you can choose whether to include the payload positions.


 Once you have configured the attack, click  **Start attack** to send the request to the target server.


#### Related pages

- For more information on using attack tabs, see [Managing attack tabs](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/managing-tabs).
