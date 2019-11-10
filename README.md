## Ansible Playbooks

Here lies all files and Ansible playbooks I use to manage my boxes.

## Playbooks

To run playbook:

```
ansible-playbook -K --ask-vault-pass --private-key=<your-private-key> -i hosts run.yml
```

**Note:** When creating a user, make sure you also add the user under `AllowUsers` in `/etc/ssh/sshd_config` file.
