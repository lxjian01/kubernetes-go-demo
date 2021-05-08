# kubernetes-go-demo
    install client-go v1.15.0
    go get k8s.io/client-go@kubernetes-1.15.0
    
# Distributed scheduling framework machinery, like python celery
    github.com/RichardKnop/machinery/v2
   
# start gin server
    cd /opt/code/kubernetes-go-demo
    go run main.go   # run gin server
    go run main.go version  # see kubernetes-go-demo version
    go run main.go worker  # run machinery worker
    
# test
    curl -H "Content-Type: application/json" -X POST "http://127.0.0.1:8018/kubernetes/deployment"
    curl -H "Content-Type: application/json" -X PUT "http://127.0.0.1:8018/kubernetes/deployment"