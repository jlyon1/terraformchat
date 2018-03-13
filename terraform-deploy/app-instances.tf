provider "digitalocean" {
  token = "${var.do_token}"
}

resource "digitalocean_droplet" "chat" {
  image              = "docker"
  name               = "chat"
  region             = "nyc3"
  size               = "512mb"
  private_networking = true
  count = 1

  ssh_keys = [
    "${var.ssh_fingerprint}",
  ]

  connection {
    user        = "root"
    type        = "ssh"
    private_key = "${file(var.pvt_key)}"
    timeout     = "2m"
  }

  provisioner "remote-exec" {
    inline = [
      "git clone https://github.com/jlyon1/terraformchat.git",
      "cd terraformchat",
      "curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose",
      "chmod +x /usr/local/bin/docker-compose",
      "nohup docker-compose up &",

    ]
  }

  provisioner "local-exec" {
    command = "echo ${digitalocean_droplet.chat.price_hourly} > price.txt"
  }
}
