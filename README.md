# CloudPulse (Go) - AWS CodePipeline, no YAML

Source -> CodeBuild (console build commands) -> Elastic Beanstalk (Go platform)

Build commands:
    GOOS=linux GOARCH=amd64 go build -o application main.go
    chmod +x application

Output files: application, Procfile
Test: GET /  and  GET /health
