# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network :private_network, ip: "192.168.33.10"

  #config.vm.provision "docker" do |d|
  #  d.pull_images "tutum/hello-world"
  #end
  #config.vm.provision "shell" do |s|
  #  s.inline = <<-EOF
  #   docker ps -q | xargs docker kill
  #   docker run -n app-0 -p 8080:80 -d tutum/hello-world
  #   docker run -n app-1 -p 8081:80 -d tutum/hello-world
  #   docker run -n app-2 -p 8082:80 -d tutum/hello-world
  #   docker run -n app-3 -p 8083:80 -d tutum/hello-world
  #  EOF
  #end
end
