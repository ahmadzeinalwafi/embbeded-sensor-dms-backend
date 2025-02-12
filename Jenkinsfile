pipeline {
    agent {
        any
    }

    environment {
        IMAGE_NAME = 'dms-be'
        IMAGE_TAG = 'latest'
        GHCR_USERNAME = 'ahmadzeinalwafi'
        GHCR_REPO = "ghcr.io/${GHCR_USERNAME}/${IMAGE_NAME}"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    checkout scm
                }
            }
        }

        stage('Login to GHCR') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'ghcr-token-id', variable: 'GHCR_TOKEN')]) {
                        sh "echo $GHCR_TOKEN | docker login ghcr.io -u $GHCR_USERNAME --password-stdin"
                    }
                }
            }
        }

        stage('Build Image') {
            steps {
                script {
                    docker.build("${GHCR_REPO}:${IMAGE_TAG}")
                }
            }
        }

        stage('Push to GHCR') {
            steps {
                script {
                    docker.withRegistry('https://ghcr.io', 'ghcr-token-id') {
                        docker.image("${GHCR_REPO}:${IMAGE_TAG}").push()
                    }
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    sh "docker rmi ${GHCR_REPO}:${IMAGE_TAG} || true"
                }
            }
        }
    }
}