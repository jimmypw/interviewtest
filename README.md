Instructions
============

Instructions for interviewer
----------------------------

To start:
- go build killer.go
- Use a sensible window manager
- vagrant up

to stop:
- vagrant destroy

Instructions for interviewee
----------------------------

You will be given a virtual machine with a selection of issues wrong with it. Your task is to perform the instructions we give you then locate and rectify any issues you encounter.
You are allowed to use the internet and or man pages to help you solve the issue. Please talk us through what you are thinking as you troubleshoot the issue.
We are interested to see if you are able to diagnose and rectify these issues and we are not rating you based on 'best practice' or 'security'

- Change the password for user 'gilbert'.
- Switch user to 'gilbert'
- There is a file called '-' in the home directory, delete it.
- Create an SSH key for the user 'gilbert', add the public key to authorized_keys and test
- Install httpd
- start httpd and verify functionality
- there is a second disk attached to the system. format it with ext4 filesystem, create a persistant mount, mount it and create a file
- run the mystery application /usr/local/bin/pbcak
- identify why your mount no longer works and attempt to recover it.

Validation
==========

Change the password for user 'gilbert'
------------------------------------
- sudo passwd gilbert
- <a password><CR>

Switch users to 'gilbert'
--------------------------
The users shell is set to nologin so before the su can take place the shell must be changed with
- sudo usermod -s /bin/bash gilbert

Then switch users with
- sudo -u gilbert -i
- or similar

Delete the file called '-'
--------------------------
The files in the directory are owned by root and must be changed before you are able to delete it

- chown -R gilbert: /home/gilbert
- rm ./-

Create an SSH key for gilbert and install it in to the authorized keys file
---------------------------------------------------------------------------
There are two issues here
1) The remaining file ownership needs to be changed if not done above.
- chown -R gilbert: /home/gilbert

2) The execute bit is not set on the .ssh directory
- chmod +x /home/gilbert/.ssh

Then create the ssh key with
- ssh-keygen

Install httpd
-------------
several issues here too
1) iptables is set to deny outbound all traffic on port 53 for DNS
- flush the iptables rules with
- iptables -F

2) nsswitch is set not to use DNS
- in /etc/nsswitch 'dns' must be added to the line begining hosts:
- hosts:      files dns myhostname

then install httpd with 
- yum install httpd

Start httpd and verify functionality
------------------------------------
There is a process which is locking port 80 which needs to be killed before you are able to start httpd
reveal the process

- netstat -anpt | grep 80
- kill the process
- ----- spoiler, it got restarted.

find the parent process
- ps waxf
- kill the parent process

restart apache
- systemctl restart apache

verify it works
- curl localhost


Format second disk
------------------
No tricks here just format the device mount it and create a file on the partition


Run the mystery application
---------------------------
User runs sudo /usr/local/bin/pbcak


Work out what has happened to the partition and try to repair it
----------------------------------------------------------------
A load of rubbish has been written to the filesystem and the filesystem destroyed. This can be seen with 
- xxd /dev/sdb
- fsck.ext4 /dev/sdb
