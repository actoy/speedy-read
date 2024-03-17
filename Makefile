goverter:
	cd ./biz/conversion/api & go run github.com/jmattheis/goverter/cmd/goverter@v1.2.0 gen ./

kitex_gen:
	kitex -module speedy/read -service speedy-read thrift/speedy_read.thrift