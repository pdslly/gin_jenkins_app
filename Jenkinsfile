pipeline {
  agent {
    docker {
      image 'golang:1.14.2'
      args '-e GOPROXY=https://goproxy.io -e GO111MODULE=on -v /var/jenkins_home/workspace/gin_demo:/go/src/app'
    }
  }
  stages {
    stage('Build') {
      steps {
        sh 'go build -o app'
      }
    }
  }
}