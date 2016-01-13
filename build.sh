# varaibles
project_name=linbox_go

cd ../..
project_root=`pwd`:$GOPATH


# child project path
project_path_connector_tcp=connector/tcp

# env
export GOPATH=$project_root:$GOPATH

#dependencies
#go get github.com/go-sql-driver/mysql
#go get github.com/go-errors/errors

# clean
cd $project_root/bin
rm -rf ./*

# build
cd $project_root/src/$project_name/$project_path_connector_tcp
go install

# copy
cp $project_root/src/$project_name/$project_path_connector_tcp/config/* $project_root/bin/