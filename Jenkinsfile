pipeline {
  agent {
    docker {
      image 'golang:1.14.2'
      args '-e GOPROXY=https://goproxy.io -e GO111MODULE=on -v ${pwd}:go/src/app'
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