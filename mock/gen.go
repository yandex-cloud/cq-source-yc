package mock

//go:generate mockgen -destination=compute_disk_server_mock.go -package=mock github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1 DiskServiceServer
