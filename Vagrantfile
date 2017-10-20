file_to_disk = './second_disk.vdi'

Vagrant.configure(2) do |config|
  config.vm.box = "bento/centos-7.3"
  config.vm.provider "virtualbox" do |v|
    v.memory = 8192
    v.cpus = 4
    v.customize ['createhd', '--filename', file_to_disk, '--size', 500 * 1024]
    v.customize ['storageattach', :id, '--storagectl', 'SATA Controller', '--port', 1, '--device', 0, '--type', 'hdd', '--medium', file_to_disk]
  end

  

  config.vm.provision "shell", inline: <<-SHELL
    yum install -y epel-release
    yum install -y nc bind-utils man man-pages iptables-utils vim nano vi lsof telnet puppet psmisc
    adduser -s /usr/sbin/nologin -d /home/gilbert -m gilbert
    echo 'trap "cd ~" SIGINT' > /etc/profile.d/traps.sh
    mkdir /home/gilbert/.ssh
    touch /home/gilbert/~
    chmod -x /home/gilbert/.ssh
    find /home/gilbert -exec chown root: {} \\;
    chown gilbert: /home/gilbert

    cp /vagrant/killer /usr/local/bin/pbcak
    chown root: /usr/local/bin/pbcak
    chmod 4755 /usr/local/bin/pbcak

    # Mess stuff up
    iptables -A OUTPUT -p udp -m udp --dport 53 -j DROP
    sed -i 's/ dns//g' /etc/nsswitch.conf
    cat << EOF > /usr/local/sbin/simpleserver.sh
#!/bin/bash
while [ 1 ]; do
  /usr/bin/nc -l -p 80 -e '/bin/echo Hello World'
done
EOF

    chmod +x /usr/local/sbin/simpleserver.sh

    cat << EOF > /etc/systemd/system/simpleserver.service
[Unit]
Description=Simple Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
WorkingDirectory=/tmp
ExecStart=/usr/local/sbin/simpleserver.sh
Restart=no
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF
    systemctl daemon-reload
    systemctl enable simpleserver
    systemctl start simpleserver

    # umount /vagrant
  SHELL
end
