pipeline {
    agent none

    stages {
//        parallel {
        stage('Product') {
            agent any

            stages {    
                stage('Get source') {
                    steps {
                        git url: 'https://github.com/cams7/intensivao-microservicos.git', branch: 'main'
                    }
                }

                stage('Docker build') {
                    steps {
                        script {
                            dockerapp = docker.build(
                                "cams7/intensivao-product:${env.BUILD_ID}",
                                '-f ./product/Dockerfile ./product'
                            )
                        }
                    }
                }

                stage('Docker push') {
                    steps {
                        script {
                            docker.withRegistry(
                                'https://registry.hub.docker.com',
                                'dockerhub'
                            ) {
                                dockerapp.push('lastest')
                                dockerapp.push("${env.BUILD_ID}")
                            }
                        }
                    }
                }

                stage('Kubernetes deploy') {
                    agent {
                        kubernetes {
                            cloud 'kubernetes'
                        }                        
                    }

                    steps {
                        kubernetesDeploy(
                            configs: './_k8s/product.yaml',
                            kubeconfigId: 'kubeconfig'
                        )        
                    }
                }
            }
        }
//        }
    }
}