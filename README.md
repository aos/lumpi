## Raspberry Pi files and playbooks

Here lies all files and Ansible playbooks I use to manage the Raspberry Pi.

## Playbooks

To run playbook:

```
ansible-playbook -K --ask-vault-pass --private-key=<your-private-key> run.yml
```

**Note:** When creating a user, make sure you also add the user under `AllowUsers` in `/etc/ssh/sshd_config` file.
