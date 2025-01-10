pipeline {
    agent {
    	label 'jenkins-slave'
    }
    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', credentialsId: 'GitHubCredentials', url :'https://github.com/Aboubakar1234/projectDevOps'  
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build('custom-go-app:latest', './webapi')
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

