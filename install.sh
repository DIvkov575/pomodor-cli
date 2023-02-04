cp ./pomodoro-cli /usr/sbin/pomodoro-cli
cp ./pomodoro-cli /usr/local/bin/pomodoro-cli
touch ~/pomodoro-config.yaml
echo "cycles_lengths:
  - "25m"
  - "5m"
cycle_names:
  - "Work"
  - "Break"
confirm_new: false
" > ~/pomodoro-config.yaml
