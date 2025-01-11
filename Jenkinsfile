pipeline {
    agent {
        label 'jenkins-slave'
    }
    environment {
        DOCKER_IMAGE = 'custom-go-app:latest'
        KUBERNETES_NAMESPACE_DEV = 'development'
        KUBERNETES_NAMESPACE_PROD = 'production'
    }
    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', credentialsId: 'GitHubCredentials', url: 'https://github.com/Aboubakar1234/projectDevOps'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build(env.DOCKER_IMAGE, './webapi')
                }
            }
        }
        stage('Load Image into Minikube') {
            steps {
                script {
                    sh "minikube image load ${env.DOCKER_IMAGE}"
                }
            }
        }
        stage('Deploy to Development') {
            steps {
                script {
                    sh """
                    kubectl apply -f kubernetes/deployment-dev.yaml
                    kubectl apply -f kubernetes/service-dev.yaml
                    """
                }
            }
        }
        stage('Test in Development') {
            steps {
                script {
                    def response = sh(script: 'curl -s -o /dev/null -w "%{http_code}" http://$(minikube ip):8085/health', returnStdout: true).trim()
                    if (response != '200') {
                        error "Health check failed with response code: ${response}"
                    }
                }
            }
        }
        stage('Deploy to Production') {
            steps {
                script {
                    sh """
                    kubectl apply -f kubernetes/deployment-prod.yaml
                    kubectl apply -f kubernetes/service-prod.yaml
                    """
                }
            }
        }
    }
}

