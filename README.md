# ttp-bench

![logo](./images/logo.png)

ttp-bench simulates 30 popular tactics from both the [MITRE ATT&CK framework](https://attack.mitre.org/) and published defense research. All of the simulations behave at least vaguely suspicious, such as stealing GCP credentials, sniffing your keyboard, accessing unusual DNS servers, or pretending to be a kernel process. Most simulations have multiple suspicious characteristics that lend themselves toward alerting, such as being unsigned binaries that just magically appeared on disk. How many of these simulations will your intrusion detection system detect?

A similar open-source project is [Atomic Red Team](https://github.com/redcanaryco/atomic-red-team):  which did not support Linux or macOS when ttp-bench was created. It's a bit complicated to setup and there isn't much overlap between the techniques used, so both projects remain useful in 2024.

## Screenshots

![choices](./images/ioc-choices.png)
![running](./images/ioc-running.png)

## Requirements

* The Go Programming language

Most of the checks available today mimic IoC found on UNIX-like operating systems. This is however not an intentional design goal. ttp-bench is actively tested on Linux and macOS

## Usage

To jump in, run the following to access the interactive menu of checks to execute:

```shell
go run .
```

ttp-bench supports some flags for automation:

```shell
 -all: execute all possible checks
  -checks: comma-separated list of checks to execute
  -list: list possible checks
```

For the few checks that require root, you will be prompted for a password.

## Available checks

* cnc-dns-over-https: Simulates C&C discovery via DNS over HTTPS (ala Godlua)
* cnc-resolve-random: Simulates C&C discovery via randomized hostname lookups (ala Aquatic Panda)
* creds-browser-cookies: Simulates theft of web session cookies [T1539]
* creds-gcp-exfil: Simulates theft of GCP credentials [1552.001, T15060.002]
* creds-keylogger-root: Simulate theft of credentials via key logging [T1056]
* creds-packet-sniffer-root: Simulates theft of credentials via network sniffing [T1040]
* creds-ssh-exfil: Simulates theft of GCP credentials [1552.001, T15060.002]
* evade-deleted-service: Simulates a service running by a binary which no longer exists
* evade-masquerade-kernel-thread-root: Simulates process masquerading as a kernel thread [T1036.004]
* evade-masquerade-user: Simulates process masquerading as another user process [T1036.004]
* evade-shell-history: Simulates attack cleanup via bash_history truncation [T1070.003]
* evade-tools-in-var-tmp-hidden: Simulates tool transfer using curl & running from /var/tmp/. [T1036.005]
* evade-usr-bin-exec-root: Simulates malicious program installing itself into /usr/bin [T1036.005]
* exec-bash-reverse-shell: Launches a temporary reverse shell using bash
* exec-curl-to-hidden-url: Simulate
* exec-curl-to-hidden-url: Simulates tool transfer using curl to a hidden directory [T1036.005]
* exec-drop-eicar: Simulates droppping a known virus signature (EICAR) onto filesystem
* exec-linpeas: Downloads and launches LinPEAS
* exec-netcat-listen: Launches netcat to listen on a port [T1059.004]
* exec-python-reverse-shell: Launches a temporary reverse shell using Python
* exec-traitor-vuln-probe: Simulates probing system for privilege escalation vulns
* exec-upx-listener-root: New unsigned obfuscated binary listening from a hidden directory as root
* hidden-listener: New unsigned binary listening from a hidden directory
* persist-iptables-root: Simulates attacker making iptables changes to allow incoming traffic
* persist-launchd-com-apple-root: Simulates persistance via a fake unsigned Apple launchd service
* persist-user-crontab-reboot: Simulates a command inserting itself into the user crontab for persistence
* privesc-traitor-dirty-pipe: Simulate CVE-2022-0847 (Dirty pipe) to escalate user privileges to root
* privesc-traitor-docker-socket: Simulates using Docker sockets to escalate user privileges to root
* pypi-supply-chain: Simulates a PyPI supply chain attack using a modified real-world sample
