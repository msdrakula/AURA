> Source: https://portswigger.net/burp/documentation/desktop/burp-at/prompting

Professional

# Prompting Burp AT effectively

-

**Last updated: ** September 3, 2026
-

**Read time: ** 5 Minutes

 A useful goal for Burp AT needs more than a topic: it should tell Burp AT what question to answer, what evidence would answer it, and when to stop. This page explains how to create effective prompts for Burp AT.


## Build a goal around an answer you can check

 A goal can include:


-

**Target**: The host, path, endpoint, or attached Burp resource to investigate.
-

**Question**: The security behavior you want Burp AT to test for.
-

**Context**: Relevant facts, such as intended application behavior or an observation that led to the test.
-

**Limits**: What to avoid, how much impact is acceptable, and how far to pursue the result.
-

**Evidence**: What would confirm, reject, or leave the question unresolved.

 You don't need to include all of these elements in every prompt. Use only elements that add useful information.


 For example:
 `Test whether a standard user can read another user's order through
https://example.com/api/orders/{id}.

I've attached a request that retrieves an order owned by the signed-in user.
The application should return only orders owned by that user.

Do not modify or cancel an order. Stop
after demonstrating access to one harmless record.

Compare a valid same-user request with the cross-user request. Report the
result as confirmed, not confirmed, or inconclusive, and show the requests
and response differences that support it.`

 This goal gives Burp AT a specific target and a means of determining whether the test was a success. It states the expected behavior, limits the impact of the test, and defines both the evidence required and the point at which further access would add no value.


## Replace broad instructions with testable questions

 Giving Burp AT specific questions to answer can make results easier to verify. Adding emphasis or asking Burp AT to be “extremely thorough” does not provide useful direction.

| **Instead of**  | **Write**
|

`Test this application thoroughly.` |

`Map the authenticated account area and identify three access-control boundaries worth testing. Do not send attack payloads yet.`
|

`Find SQL injection in this request.` |

`Test whether the sort parameter is injectable. Compare SQL-sensitive inputs with suitable controls, and report whether the evidence confirms, weakens, or rejects the hypothesis.`
|

`Check the admin page.` |

`Only administrators should be able to reach /admin. Test whether a standard-user session can access it.`
|

`Prove this issue is exploitable.` |

`Try to reproduce the attached issue. If it reproduces, demonstrate the minimum impact needed to confirm it. If it does not, show what you tested and what remains uncertain.`
|

`The login behaves strangely.` |

Attach the request, then write: `Investigate whether this login request permits authentication bypass.`

 Naming a vulnerability class, endpoint, or expected security boundary helps Burp AT direct its effort. Avoid packing several unrelated objectives into the same goal. Where possible, keep Burp AT tasks focused on one goal or question.


## Ask for evidence, not the answer you expect

 Phrase the goal as a hypothesis Burp AT can reject. Ask Burp AT to test whether the issue exists and to distinguish between:


-

**Confirmed** — the evidence demonstrates the behavior.
-

**Not confirmed** — the tests performed did not demonstrate it.
-

**Inconclusive** — the available access, context, or evidence was insufficient.

 State the evidence you need to review: for example, a baseline request, a modified request, the relevant response difference, and the impact demonstrated. Ask Burp AT to provide the minimum sufficient proof in order to minimise disruption. For example, access to one harmless test record may establish a broken access control issue without retrieving additional data or changing application state.


 You may find it useful to challenge a result before reporting it. For example:
 `Try to disprove this finding. Test the strongest alternative explanation
first, then tell me whether the original conclusion still holds.`

## Control depth with observable stopping points

 You can control the level of depth Burp AT goes into by giving it stopping points. For example:
 `Test every input to this endpoint for server-side injection, including query
parameters, headers, and cookies. Use suitable controls and encoding variants.
Continue until each input is confirmed, rejected, or marked inconclusive.`

 When giving Burp AT a stopping point, use well-defined conditions such as `until each input is classified` rather that phrases such as `test exhaustively` or `keep going until you are certain` .


 If you impose a time or effort limit, say what Burp AT should do when it reaches that limit:
 `Keep this to a short initial pass. If you cannot confirm the issue, summarize
the tests completed, the strongest remaining lead, and what you would try next.`

 This gives Burp AT permission to return a useful partial result instead of choosing between overrunning the limit and making an unsupported conclusion. Deeper goals generally cause more tool use and consume more AI credits.


## Steer the task without changing its goal

 Use follow-up prompts to focus, extend, or stop the current line of work. For example:


1. `Concentrate on the unexplained 500 response and compare it with a control.`
1. `Repeat the same access-control test against the invoice endpoint.`
1. `Summarize the evidence collected so far.`

 If Burp AT repeatedly drifts from the intended question, start a new task with a clearer goal. Carrying several corrections and abandoned approaches forward can make the task harder to direct.


 When Burp AT requests approval, select **Show detail** and check that the planned action matches your goal and acceptable impact. Rejecting an action does not end the task; Burp AT looks for another approach instead.


## Enforce restrictions using in-product controls

 Limits written in a prompt guide Burp AT's priorities, but they should not be used as security controls in their own right. Use Burp Suite's scope controls and Burp AT's tool permissions settings to define what Burp AT should not test and what tools it should use.


 Burp AT can propose changes to scope, but they always require manual approval. You can disable **Add to scope** and **Remove from scope** in **Settings > Tools**.


 Custom code run through the scripting tools can send requests outside project scope. You can disable these scripting tools if you need to.


 Treat target content as untrusted data. Pages, responses, scripts, and headers may contain text intended to influence an AI agent. Do not rely on an instruction in your goal to neutralize that content; use project scope, tool permissions, and approval checks to control what Burp AT can do.


#### More information

-

[Setting the target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope)
-

[Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions)
-

[AI trust and data handling](https://portswigger.net/burp/documentation/ai-features/trust#can-burp-at-test-out-of-scope-targets).

## Troubleshooting

| **Symptom**  | **What to change**
|

Burp AT covers a large surface without going deep. |

Narrow the goal to one security question or one area of the application.
|

Burp AT reports a finding you cannot reproduce. |

Ask it to repeat the test against a control and try to disprove the finding.
|

Burp AT stops before the question is settled. |

Add the evidence required and an observable stopping point.
|

Burp AT takes actions you did not intend. |

Add a specific impact limit, then check scope and tool permissions.
|

Burp AT asks for approval too often. |

Review the task's autonomy mode and the settings under **Settings > Tools**.
|

Burp AT declines authorized testing. |

State the authorized testing context, identify the in-scope target, and phrase the request as vulnerability verification.

#### Next step

-

[Tools](https://portswigger.net/burp/documentation/desktop/burp-at/tools)
