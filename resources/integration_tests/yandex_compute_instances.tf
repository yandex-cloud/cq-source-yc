resource "yandex_vpc_network" "cq-instance-test-net" {
  name = "cq-instance-test-net"
}

resource "yandex_vpc_subnet" "cq-instance-test-subnet" {
  network_id     = yandex_vpc_network.cq-instance-test-net.id
  v4_cidr_blocks = ["10.2.0.0/16"]
  name           = "cq-instance-test-subnet"
}

resource "yandex_compute_instance" "cq-instance-test-instance" {
  name = "cq-instance-test-instance"
  boot_disk {
    initialize_params {
      image_id = "fd8vmcue7aajpmeo39kk"
    }
  }
  network_interface {
    subnet_id = yandex_vpc_subnet.cq-instance-test-subnet.id
  }
  resources {
    cores = 2
    memory = 4
  }
}