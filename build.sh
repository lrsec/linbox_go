# varaibles
project_name=linbox_go

cd ../..
project_root=`pwd`


# env
export GOPATH=$project_root:$GOPATH

#dependencies
#go get github.com/go-sql-driver/mysql
#go get github.com/go-errors/errors

# clean
cd $project_root/bin
rm -rf ./*

# build
cd $project_root/src/$project_name
go install

# copy
cp $project_root/src/$project_name/config/* $project_root/bin/