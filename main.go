package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)



func main() {
// Initialize the router with the routes for the application.

	router := gin.Default()


	router.GET("/etcdownership", fixEtcdOwnership)
	router.GET("/kubecertauth", fixKubeCertAuth)
	router.GET("/recommendations", getRecommendation)
	router.Run(":8080")

}

// R1.1.12 Ensure that the etcd data directory ownership is set to etcd:etcd 
// Change the ownership of the etcd data directory to etcd:etcd
func fixEtcdOwnership(c *gin.Context) {
	cmd := exec.Command("chown", "etcd:etcd", "/var/lib/etcd")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message": "Error changing ownership of etcd data directory"})
	} else {
		c.JSON(http.StatusOK,gin.H{"message": "Etcd ownership fixed Successfully"})
	}
}

// R1.2.6 Ensure the --authorization-mode=RBAC argument in  /etc/kubernetes/manifests/kube-apiserver.yaml is set to RBAC
func fixKubeCertAuth (c *gin.Context) {
	input, err := ioutil.ReadFile("/etc/kubernetes/manifests/kube-apiserver.yaml")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message": "Error reading the kube-apiserver.yaml file"})
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("--authorization-mode=Node,RBAC"), []byte("--authorization-mode=RBAC"), -1)
	if err = ioutil.WriteFile("/tmp/kube-apiserver.yaml", output, 0666); err != nil {
	if err = ioutil.WriteFile("/etc/kubernetes/manifests/kube-apiserver.yaml", output, 0666); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message": "Error writting the kube-apiserver.yaml file"})
		os.Exit(1)
	} else {
		c.JSON(http.StatusOK,gin.H{"message": "Kubelet certificate authority fixed Successfully"})
	}
}
}

// Get the recommendations from the output of kube-bench in the node
func getRecommendation (c *gin.Context) {
	input, err := os.Open("/opt/kube-bench/output/recommendations.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message": "Error reading the recommendations.txt file"})
		os.Exit(1)
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		c.JSON(http.StatusOK,gin.H{"message": scanner.Text()})
	}
}

// R1.2.18	Ensure that the --audit-log-path argument is set
// To be implemented 

// R1.2.19	Ensure that the --audit-log-maxage argument is set to 30 or as appropriate
// To be implemented 

// R1.2.20	Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate
// To be implemented 

// R1.2.22	Ensure that the --request-timeout argument is set as appropriate
// To be implemented 

// R1.3.2	Ensure that the --profiling argument is set to false Controller Manager
// To be implemented 

// R1.4.1	Ensure that the --profiling argument is set to false - Scheduler
// To be implemented 

