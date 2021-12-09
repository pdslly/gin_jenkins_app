pipeline {
  agent {
    docker {
      image 'golang:1.17'
      args '-p 3000:3000'
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