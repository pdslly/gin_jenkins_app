pipeline {
  agent none
  stages {
    stage('Build') {
      agent {
        docker {
          image 'golang:1.14.2'
          args '-e GOPROXY=https://goproxy.io -e GO111MODULE=on -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 -v /var/jenkins_home/workspace/gin_demo:/go/src/app'
        }
      }
      steps {
        sh 'go build -o app'
      }
    }
    stage('Deploy') {
      agent any
      steps {
        sh 'docker restart ginapp'
      }
    }
  }
}