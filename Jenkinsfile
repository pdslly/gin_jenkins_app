pipeline {
  agent none
  stages {
    stage('Build') {
      agent {
        docker {
          image 'golang:1.14.2'
          args '-e GOPROXY=https://goproxy.io -e GO111MODULE=on -v /var/jenkins_home/workspace/gin_demo:/go/src/app'
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