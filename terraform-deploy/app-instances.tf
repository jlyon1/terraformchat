provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region     = "${var.region}"
}

resource "aws_instance" "chat" {
  ami           = "ami-1853ac65"
  instance_type = "t2.micro"

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install -y docker",
      "sudo yum install -y docker",
      "sudo service docker start",
      "sudo curl -L https://github.com/docker/compose/releases/download/1.5.1/docker-compose-`uname -s`-`uname -m` > docker-compose",
      "sudo chown root docker-compose",
      "sudo mv docker-compose /usr/local/bin",
      "sudo chmod +x /usr/local/bin/docker-compose"
    ]
  }

  provisioner "local-exec" {
    command = "echo ${aws_instance.chat.public_ip} > ip_address.txt"
  }
}
