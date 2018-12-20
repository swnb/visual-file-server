if [ ! -d './bin/' ]; then
    mkdir "./bin/"
fi

echo "you are building tools for gat \n"

go build -o ./bin/new ./.tools/create-controller.go

echo "you can use ./bin/new to create new controller \n"

echo "to get help ; use -h option to see what you can do"