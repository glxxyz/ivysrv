pipeline {
  agent any
  triggers {
    pollSCM('* * * * *')
  }
  options {
    skipDefaultCheckout(true)
  }
  stages {
    stage('build') {
      steps {
          cleanWs()
          checkout scm
//          sh 'XDG_CACHE_HOME=/tmp/.cache make docker'
          sh 'docker tag jmalloc/ivysrv:dev us-east1-docker.pkg.dev/ccd-starter-jenkins-gke/ivysrv-jenkins'
          sh 'docker push us-east1-docker.pkg.dev/ccd-starter-jenkins-gke/ivysrv-jenkins'
          sh 'skaffold deploy --images=us-east1-docker.pkg.dev/ccd-starter-jenkins-gke/ivysrv-jenkins:latest --profile=beta'
          sh 'skaffold deploy --images=us-east1-docker.pkg.dev/ccd-starter-jenkins-gke/ivysrv-jenkins:latest --profile=prod'
      }
    }
  }
  post {
    always {
      cleanWs(cleanWhenNotBuilt: false,
              deleteDirs: true,
              disableDeferredWipeout: true,
              notFailBuild: true,
              patterns: [[pattern: '.gitignore', type: 'INCLUDE'],
                         [pattern: '.propsfile', type: 'EXCLUDE']])
    }
  }
}