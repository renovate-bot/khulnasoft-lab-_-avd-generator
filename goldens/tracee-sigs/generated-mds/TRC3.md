---
title: Code Injection
id: TRC-3
aliases: [
    "/tracker/trc3"
]
source: Tracker
icon: khulnasoft
shortName: Code Injection
severity: high
draft: false
version: 0.1.0
keywords: "TRC-3"

category: runsec
date: 2021-04-15T20:55:39Z

remediations:

breadcrumbs: 
  - name: Tracker
    path: /tracker
  - name: Defense Evasion
    path: /tracker/defense-evasion

cvedb_page_type: cvedb_page
---

### Code Injection
Possible code injection into another process

### MITRE ATT&CK
Defense Evasion: Process Injection


### Rego Policy
```
package tracker.TRC_3

import data.tracker.helpers

__rego_metadoc__ := {
    "id": "TRC-3",
    "version": "0.1.0",
    "name": "Code injection",
    "description": "Possible code injection into another process",
    "tags": ["linux", "container"],
    "properties": {
        "Severity": 3,
        "MITRE ATT&CK": "Defense Evasion: Process Injection",
    }
}

eventSelectors := [
    {
        "source": "tracker",
        "name": "ptrace"
    },
    {
        "source": "tracker",
        "name": "security_file_open"
    },
    {
        "source": "tracker",
        "name": "process_vm_writev"
    }
]

tracker_selected_events[eventSelector] {
	eventSelector := eventSelectors[_]
}


tracker_match {
    input.eventName == "ptrace"
    arg_value = helpers.get_tracker_argument("request")
    arg_value == "PTRACE_POKETEXT"
}

tracker_match = res {
    input.eventName == "security_file_open"
    flags = helpers.get_tracker_argument("flags")

    helpers.is_file_write(flags)

    pathname := helpers.get_tracker_argument("pathname")

    regex.match(`/proc/(?:\d.+|self)/mem`, pathname)

    res := {
        "file flags": flags,
        "file path": pathname,
    }
}

tracker_match {
    input.eventName == "process_vm_writev"
    dst_pid = helpers.get_tracker_argument("pid")
    dst_pid != input.processId
}
```
