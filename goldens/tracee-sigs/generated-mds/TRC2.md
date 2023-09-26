---
title: Anti Debugging
id: TRC-2
aliases: [
    "/tracker/trc2"
]
source: Tracker
icon: khulnasoft
shortName: Anti Debugging
severity: high
draft: false
version: 0.1.0
keywords: "TRC-2"

category: runsec
date: 2021-04-15T20:55:39Z

remediations:

breadcrumbs: 
  - name: Tracker
    path: /tracker
  - name: Defense Evasion
    path: /tracker/defense-evasion

avd_page_type: avd_page
---

### Anti Debugging
Process uses anti-debugging technique to block debugger

### MITRE ATT&CK
Defense Evasion: Execution Guardrails


### Rego Policy
```
package tracker.TRC_2

__rego_metadoc__ := {
    "id": "TRC-2",
    "version": "0.1.0",
    "name": "Anti-Debugging",
    "description": "Process uses anti-debugging technique to block debugger",
    "tags": ["linux", "container"],
    "properties": {
        "Severity": 3,
        "MITRE ATT&CK": "Defense Evasion: Execution Guardrails",
    }
}

tracker_selected_events[eventSelector] {
	eventSelector := {
		"source": "tracker",
		"name": "ptrace"
	}
}

tracker_match {
    input.eventName == "ptrace"
    arg := input.args[_]
    arg.name == "request"
    arg.value == "PTRACE_TRACEME"
}
```
