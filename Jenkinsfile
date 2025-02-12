pipeline {
    agent any  // Uses any available agent with Docker installed

    environment {
        IMAGE_NAME = 'dms-be'   // Change this
        IMAGE_TAG = 'latest'
        GHCR_USERNAME = 'ahmadzeinalwafi'  // Change this
        GHCR_REPO = "ghcr.io/${GHCR_USERNAME}/${IMAGE_NAME}"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Image') {
            steps {
                script {
                    image = docker.build("${GHCR_REPO}:${IMAGE_TAG}")
                }
            }
        }

        stage('Push to GHCR') {
            steps {
                script {
                    docker.withRegistry('https://ghcr.io', 'ghcr-token-id') {
                        image.push()
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    docker.image("${GHCR_REPO}:${IMAGE_TAG}").remove()
                }
            }
        }
    }
}
