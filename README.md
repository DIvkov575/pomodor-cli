pomodoro timer
---

configurable pomodoro-cli timer written in golang


inorder to install:
1) download apropriate binary from Githubreleases
2) execute install.sh
3) run 'pomodoro-cli run'
4) if error occurs in step 3, add a pomodoro-config.yaml containing text below to your user root dir (~/) 

"""cycles_lengths:

  - "25m"

  - "5m"

cycle_names:

  - "Work"

  - "Break"

confirm_new: false"""


---

use `pomodoro run` inorder to execute

-c flag to specify number of cycles

-s to change dir source for config file from default (~/)

---
press 'q' or 'ctrl+c' to quit during execution

press 's' to skip current cycle

---

credit to github.com/caarlos0/timer for timer animation


