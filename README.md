# ioc-bench

Benchmark for common indicators of compromise

## Implemented

* `cnc-resolve-random`: Simulates C&C discovery via randomized hostname lookups (Aquatic Panda)
* `cnc-dns-over-https`: Simulates C&C discovery via DNS over HTTPS (Godlua)
* `creds-browser-cookies`: Simulates theft of web session cookies [T1539](https://attack.mitre.org/techniques/T1539/)
* `creds-gcp-exfil`: Simulates theft of GCP credentials [1552.001](https://attack.mitre.org/techniques/T1552/001/), [T15060.002](https://attack.mitre.org/versions/v10/techniques/T1560/002/), [Local Data Staging](https://attack.mitre.org/versions/v10/techniques/T1074/001/)
* `evade-bash-history`: Simulates cleanup via bash_history truncation (T1070.003)[<https://attack.mitre.org/techniques/T1070/003/>

## Planned

* `access-chrome-breakout`: Simulates an overflow where Google Chrome spawns a shell [T1189](https://attack.mitre.org/techniques/T1189/)

* `creds-sniff`: Simulates theft of credentials via network sniffing [T1040](https://attack.mitre.org/techniques/T1040/)
* `creds-keylogger`: Simulate theft of credentials via key logging [T1056](https://attack.mitre.org/techniques/T1056/001/)

* `cnc-ingress-tool-xfer`: Simulates ingress tool transfer using curl [T1105](https://attack.mitre.org/versions/v10/techniques/T1105/)

* `evade-masquerade-kthreads`: Simulates process masquerading as a kernel thread [T1036.004](https://attack.mitre.org/versions/v10/techniques/T1036/004/)
* `evade-usr-bin`: Simulates program installing itself into /usr/bin [T1036.005](https://attack.mitre.org/versions/v10/techniques/T1036/005/)
* `evade-var-tmp`: Simulates program running from /var/tmp/. [T1036.005](https://attack.mitre.org/versions/v10/techniques/T1036/005/)
* `evade-usr-bin`: Simulates program installing itself into /. [T1036.005](https://attack.mitre.org/versions/v10/techniques/T1036/005/)
* `persist-iptables`: Simulates iptables changes

* `exfil-web-service`: <https://attack.mitre.org/techniques/T1567/><https://attack.mitre.org/versions/v10/techniques/T1560/002/>
* `exfil-non-web-port`: Simulates data exfil on a non-HTTPS port
* `exec-cron-reboot`: <https://attack.mitre.org/techniques/T1053/003/>

* `exec-bash-listen`: <https://attack.mitre.org/techniques/T1059/004/>
* `exec-python-listen`: <https://attack.mitre.org/techniques/T1059/006/>
* `exec-netcat-listen`: <https://attack.mitre.org/techniques/T1059/004/>
