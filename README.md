# ioc-bench

A crude benchmark for malware detection and intrusion detection systems.

ioc-bench simulates a number of popular indicators of compromise from the MITRE ATT&CK framework, biasing toward those seen in more recent attacks.

How many of these simulations does your vendor detect?

## Usage

`go run .`

For the few checks that require root, you will be prompted for a password.

## Implemented

* `access-chrome-breakout`: Simulates an overflow where Google Chrome spawns a shell [T1189](https://attack.mitre.org/techniques/T1189/)

* `cnc-resolve-random`: Simulates C&C discovery via randomized hostname lookups (Aquatic Panda)
* `cnc-dns-over-https`: Simulates C&C discovery via DNS over HTTPS (Godlua)

* `creds-browser-cookies`: Simulates theft of web session cookies [T1539](https://attack.mitre.org/techniques/T1539/)
* `creds-gcp-exfil`: Simulates theft of GCP credentials [1552.001](https://attack.mitre.org/techniques/T1552/001/), [T15060.002](https://attack.mitre.org/versions/v10/techniques/T1560/002/), [Local Data Staging](https://attack.mitre.org/versions/v10/techniques/T1074/001/)
* `creds-keylogger`: Simulate theft of credentials via key logging [T1056](https://attack.mitre.org/techniques/T1056/001/)
* `creds-sniff`: Simulates theft of credentials via network sniffing [T1040](https://attack.mitre.org/techniques/T1040/)

* `evade-bash-history`: Simulates cleanup via bash_history truncation [T1070.003](https://attack.mitre.org/techniques/T1070/003/)
* `evade-deleted-service`: Simulates a service running by a binary which no longer exists
* `evade-masquerade-kthreadd`: Simulates process masquerading as a kernel thread [T1036.004](https://attack.mitre.org/versions/v10/techniques/T1036/004/)
* `evade-var-tmp`: Simulates tool transfer using curl & running from /var/tmp/. [T1036.005](https://attack.mitre.org/versions/v10/techniques/T1036/005/), [T1105](https://attack.mitre.org/versions/v10/techniques/T1105/)
* `evade-usr-bin`: Simulates program installing itself into /usr/bin [T1036.005](https://attack.mitre.org/versions/v10/techniques/T1036/005/)

* `exec-bash-reverse-shell`
* `exec-python-reverse-shell`
* `exec-netcat-reverse-shell`
* `exec-netcat-listen`: <https://attack.mitre.org/techniques/T1059/004/>

* `persist-iptables`: Simulates iptables changes

## Planned

* `exec-cron-reboot`: <https://attack.mitre.org/techniques/T1053/003/>
* `exfil-web-service`: <https://attack.mitre.org/techniques/T1567/><https://attack.mitre.org/versions/v10/techniques/T1560/002/>
* `exfil-non-web-port`: Simulates data exfil on a non-HTTPS port
