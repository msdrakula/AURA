> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/authentication-mechanisms/credential-stuffing

ProfessionalCommunity Edition

# Credential stuffing with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Credential stuffing is a form of brute-force attack in which you attempt to log into a website using known username and password combinations from other websites. These sets of credentials are usually collated from earlier data breaches.


 These attacks rely on the fact that users often reuse the same credentials across multiple different sites. Crucially, as you're only attempting to access each account once, defense mechanisms such as account locking are effectively powerless against this kind of attack, although you may still need to bypass any rate limiting that's in place. For some ideas on how to do this, see the
 [Authentication](https://portswigger.net/web-security/authentication) topic on our Web Security Academy.


## Before you start

 Obtain a list of known username and password pairs. For the example below, you can use the following wordlists. They are already sorted into correct pairs:


#### Username list
`carlos

                        root

                        admin

                        test

                        guest

                        info

                        adm

                        mysql

                        wiener

                        user

                        administrator

                        oracle

                        ftp

                        pi

                        puppet

                        ansible

                        ec2-user

                        vagrant

                        azureuser

                        academico

                        acceso

                        access

                        accounting

                        accounts

                        acid

                        activestat

                        ad

                        adam

                        adkit

                        admin

                        administracion

                        administrador

                        administrator

                        administrators

                        admins

                        ads

                        adserver

                        adsl

                        ae

                        af

                        affiliate

                        affiliates

                        afiliados

                        ag

                        agenda

                        agent

                        ai

                        aix

                        ajax

                        ak

                        akamai

                        al

                        alabama

                        alaska

                        albuquerque

                        alerts

                        alpha

                        alterwind

                        am

                        amarillo

                        americas

                        an

                        anaheim

                        analyzer

                        announce

                        announcements

                        antivirus

                        ao

                        ap

                        apache

                        apollo

                        app

                        app01

                        app1

                        apple

                        application

                        applications

                        apps

                        appserver

                        aq

                        ar

                        archie

                        arcsight

                        argentina

                        arizona

                        arkansas

                        arlington

                        as

                        as400

                        asia

                        asterix

                        at

                        athena

                        atlanta

                        atlas

                        att

                        au

                        auction

                        austin

                        auth

                        auto

                        autodiscover`

#### Password list
`123456

                        password

                        12345678

                        qwerty

                        123456789

                        12345

                        1234

                        111111

                        peter

                        1234567

                        dragon

                        123123

                        baseball

                        abc123

                        football

                        monkey

                        letmein

                        shadow

                        master

                        666666

                        qwertyuiop

                        123321

                        mustang

                        1234567890

                        michael

                        654321

                        superman

                        1qaz2wsx

                        7777777

                        121212

                        0

                        qazwsx

                        123qwe

                        killer

                        trustno1

                        jordan

                        jennifer

                        zxcvbnm

                        asdfgh

                        hunter

                        buster

                        soccer

                        harley

                        batman

                        andrew

                        tigger

                        sunshine

                        iloveyou

                        2000

                        charlie

                        robert

                        thomas

                        hockey

                        ranger

                        daniel

                        starwars

                        klaster

                        112233

                        george

                        computer

                        michelle

                        jessica

                        pepper

                        1111

                        zxcvbn

                        555555

                        11111111

                        131313

                        freedom

                        777777

                        pass

                        maggie

                        159753

                        aaaaaa

                        ginger

                        princess

                        joshua

                        cheese

                        amanda

                        summer

                        love

                        ashley

                        nicole

                        chelsea

                        biteme

                        matthew

                        access

                        yankees

                        987654321

                        dallas

                        austin

                        thunder

                        taylor

                        matrix

                        mobilemail

                        mom

                        monitor

                        monitoring

                        montana

                        moon

                        moscow

                        daniel`

## Steps

 You can follow along with the process below using the [Excessive trust in client-side controls](https://portswigger.net/web-security/logic-flaws/examples/lab-logic-flaws-excessive-trust-in-client-side-controls) lab from our Web Security Academy.


1.
 Send the request for submitting the login form to Burp Intruder.

1.
 Go to **Intruder** and select **Pitchfork attack** from the attack type drop-down menu.

1.

 In the request, highlight the username value and click **Add §** to mark it as a payload position. Do the same for the password.


![Setting payload positions for a credential stuffing attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/credential-stuffing/positions.png)

1.
 In the **Payloads** side panel, select position `1` from the **Payload position** drop-down list.

1.

 Under **Payload configuration**, paste the list of usernames.


![Adding payloads for a credential stuffing attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/credential-stuffing/payloads-one.png)

1.
 Select position `2` from the **Payload position** drop-down list, and paste the list of passwords.

1.
 Click ** Start attack**. The attack starts running in the new dialog. Intruder sends a request for each pair of usernames and passwords in the list.

1.

 When the attack is finished, study the responses to look for any behavior that may indicate a valid login. For example, look for any anomalous error messages, response times, or status codes. In the example below, one of the requests has received a
 `302` response.


![Viewing the results of a credential stuffing attack](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/credential-stuffing/results.png)

1.
 To investigate the contents of a response in detail, right-click and select **Send to Comparer (response)**. Do the same for the original response.

1.
 Go to the **Comparer** tab. Select the two responses and click **Words** or **Bytes** to compare the responses. Any differences are highlighted.


#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists)
- [Grep - match](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-match)
- [Grep - extract](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/settings#grep-extract)
- [Burp Comparer](https://portswigger.net/burp/documentation/desktop/tools/comparer)
- Academy : [Vulnerabilities in password-based login mechanisms](https://portswigger.net/web-security/authentication/password-based)
