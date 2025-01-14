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
                // Récupère le code du repo (contenant le Dockerfile et le code de l'application)
                git branch: 'main', credentialsId: 'GitHubCredentials', url: 'https://github.com/Aboubakar1234/projectDevOps'
            }
        }
        
        stage('Start Minikube & Build Docker Image') {
            steps {
                script {
                    // 1) Démarre Minikube en driver Docker
                    // 2) "eval $(minikube docker-env)" redirige les commandes Docker
                    //    vers le daemon Docker interne de Minikube
                    // 3) Construit l'image directement dans Minikube
                    sh """
                      minikube start --driver=docker
                      eval \$(minikube docker-env)
                      docker build -t ${env.DOCKER_IMAGE} ./webapi
                    """
                }
            }
        }
        
        stage('Deploy to Development') {
            steps {
                script {
                    // Déploie dans le namespace "development"
                    // Notez l'imagePullPolicy: Never dans votre YAML
                    // car l'image est déjà dans Minikube
                    sh """
                    #cree le nameSpace et le || true ne le cree s'il existe deja 
                    kubectl create namespace development || true

                    #deployer
                    kubectl apply -f kubernetes/deployment-dev.yaml
                    kubectl apply -f kubernetes/service-dev.yaml
                    """
                }
            }
        }
        
        stage('Test in Development') {
            steps {
                script {
                    // Vérifie l’endpoint /health sur le NodePort exposé (8085)
                    def response = sh(
                      script: 'curl -s -o /dev/null -w "%{http_code}" http://$(minikube ip):30085/health',
                      returnStdout: true
                    ).trim()
                    
                    if (response != '200') {
                        error "Health check failed with response code: ${response}"
                    }
                }
            }
        }
        
        stage('Deploy to Production') {
            steps {
                script {
                    // Déploie dans le namespace "production"
                    // Même logique, l’image locale est déjà connue de Minikube
                    sh """
                    #cree le nameSpace et le || true ne le cree s'il existe deja 
                    kubectl create namespace production || true

                    #deployer
                    kubectl apply -f kubernetes/deployment-prod.yaml
                    kubectl apply -f kubernetes/service-prod.yaml
                    """
                }
            }
        }
    }
}
