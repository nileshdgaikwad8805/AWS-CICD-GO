pipeline {
    agent any

    options {
        timestamps()
        disableConcurrentBuilds()
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main',
                     url: 'https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git'
            }
        }

        stage('Go Version') {
            steps {
                sh 'go version'
                sh 'go env GOPATH GOMOD'
            }
        }


        stage('Build') {
            steps {
                sh 'mkdir -p bin'
                sh 'go build -o bin/application main.go'
                sh 'chmod +x bin/application'
                sh 'ls -lh bin/application'
            }
        }

        stage('Package') {
            steps {
                sh 'tar -czf cloudpulse-${BUILD_NUMBER}.tar.gz bin/application Procfile static'
                sh 'ls -lh cloudpulse-${BUILD_NUMBER}.tar.gz'
            }
        }
    }

    post {
        success {
            echo 'AWS-CICD-GO CI pipeline completed successfully.'
        }
        failure {
            echo 'Pipeline failed. Check the first meaningful error in Console Output.'
        }
        always {
            archiveArtifacts artifacts: 'cloudpulse-*.tar.gz', allowEmptyArchive: true
        }
    }
}
