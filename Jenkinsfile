pipeline {
    agent any
    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/Aboubakar1234/projectDevOps'  
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build('custom-go-app:latest')
                }
            }
        }
        stage('Run Docker Container') {
            steps {
                script {
                    docker.image('custom-go-app:latest').run('-d -p 8085:8085')
                }
            }
        }
    }
}

