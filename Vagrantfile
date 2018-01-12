VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
    config.vm.box = "azure"
    config.vm.box_url="https://github.com/azure/vagrant-azure/raw/v2.0/dummy.box"
    config.ssh.username = "usuario"
    config.ssh.private_key_path = '~/.ssh/id_rsa_deploying'
    config.vm.hostname = "middleware"
    #config.vm.network "forwarded_port", guest: 8080, host: 80
    #config.vm.network "private_network", type: "dhcp"
    #config.vm.network "private_network", ip: "192.168.50.4", auto_config: false
    config.vm.communicator = "ssh"
    config.vm.post_up_message = "La maquina esta desplegada"
    config.vm.provider "azure" do |azure, override|
        azure.tenant_id = ENV['AZURE_TENANT']
        azure.client_id = ENV['AZURE_APPID']
        azure.client_secret = ENV['AZURE_PASSWORD']
        azure.subscription_id = ENV['AZURE_SUBSCRIPTION']
        azure.admin_username = "usuario"
        azure.vm_password = ENV['AZURE_VM_PASS']
        azure.vm_size = "Basic_A0"
        #azure.vm_image_urn = "cognosys:secured-ngnix-on-centos-7-3:secured-ngnix-on-centos-7-3:1.0.2"
        azure.vm_name = "middleware"
        azure.tcp_endpoints = "80"
        azure.resource_group_name= "IV"
        azure.location = "westeurope"
    end

    # Provisionamiento con ansible
	config.vm.provision :ansible do |ansible|
        ansible.playbook = "./provision/playbook.yml"
    end
end
