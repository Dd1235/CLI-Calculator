pipeline {
  agent any

  environment {
    DOCKERHUB_REPO = 'calculator-imt2023006'
    IMAGE_TAG = "${env.BUILD_NUMBER}"
  }

  options {
    skipDefaultCheckout(true)
    timestamps()
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        sh 'go version'
        sh 'go mod download'
        sh 'go build -o bin/calculator ./cmd/calculator'
      }
    }

    stage('Test') {
      steps {
        sh 'go test -v ./... | tee test.out'
      }
      post {
        always {
          archiveArtifacts artifacts: 'test.out', onlyIfSuccessful: false
        }
      }
    }

    stage('Docker Build') {
      when { expression { currentBuild.currentResult == null || currentBuild.currentResult == 'SUCCESS' } }
      steps {
        script {
          sh """
            docker build -t ${DOCKERHUB_USER ?: ''}/${DOCKERHUB_REPO}:${IMAGE_TAG} .
            docker images | grep ${DOCKERHUB_REPO} || true
          """
        }
      }
    }

    stage('Docker Push') {
      when { expression { currentBuild.currentResult == null || currentBuild.currentResult == 'SUCCESS' } }
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKERHUB_USER', passwordVariable: 'DOCKERHUB_PASS')]) {
          sh """
            echo "${DOCKERHUB_PASS}" | docker login -u "${DOCKERHUB_USER}" --password-stdin
            docker tag ${DOCKERHUB_USER}/${DOCKERHUB_REPO}:${IMAGE_TAG} ${DOCKERHUB_USER}/${DOCKERHUB_REPO}:latest
            docker push ${DOCKERHUB_USER}/${DOCKERHUB_REPO}:${IMAGE_TAG}
            docker push ${DOCKERHUB_USER}/${DOCKERHUB_REPO}:latest
          """
        }
      }
    }
  }

  post {
    success {
      echo "Pushed image: ${env.DOCKERHUB_USER}/${env.DOCKERHUB_REPO}:${env.IMAGE_TAG}"
    }
  }
}
