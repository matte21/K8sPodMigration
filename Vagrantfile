IMAGE_NAME = "ubuntu/bionic64"
NODES = 2

Vagrant.configure("2") do |config|
    config.ssh.insert_key = false

    config.vm.provider "virtualbox" do |v|
        v.memory = "2048"
        v.cpus = 2
    end

    config.vm.define "k8s-master" do |master|
        master.vm.box = IMAGE_NAME
        master.vm.network "private_network", ip: "192.168.50.10"
        master.vm.hostname = "k8s-master"
        master.vm.provision "ansible" do |ansible|
            ansible.playbook = "cluster-setup/master-playbook.yml"
            ansible.extra_vars = {
                node_ip: "192.168.50.10",
            }
        end
    end

    (1..NODES).each do |i|
        config.vm.define "node-#{i}" do |node|
            node.vm.box = IMAGE_NAME
            node.vm.network "private_network", ip: "192.168.50.#{i + 10}"
            node.vm.hostname = "node-#{i}"
            node.vm.provision "ansible" do |ansible|
                ansible.playbook = "cluster-setup/node-playbook.yml"
                ansible.extra_vars = {
                    node_ip: "192.168.50.#{i + 10}",
                }
            end
        end
    end

end

# FYI: this file and the ansible playbooks it uses were taken from
# https://kubernetes.io/blog/2019/03/15/kubernetes-setup-using-ansible-and-vagrant/ .
# To make things work, it was necessary to perform many tweaks, including (but
# not limited to) those described in https://github.com/ceph/ceph-ansible/issues/1154 ,
# https://stackoverflow.com/questions/56998761/using-vagrant-ansible-to-spin-up-a-multi-node-kubernetes-cluster-fails-detecting
# and https://medium.com/@joatmon08/playing-with-kubeadm-in-vagrant-machines-part-2-bac431095706 .
