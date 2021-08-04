resource "yandex_vpc_network" "cq-subnet-test-net" {
  name = "cq-subnet-test-net"
}

resource "yandex_vpc_subnet" "cq-subnet-test-subnet" {
  network_id     = yandex_vpc_network.cq-subnet-test-net.id
  v4_cidr_blocks = ["10.2.0.0/16"]
  name           = "cq-subnet-test-subnet"
}