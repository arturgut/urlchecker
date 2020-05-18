pipeline{

  agent any
  
  stages{
    
    stage('Get Commit Details'){
      steps('Get SHA') {
        sh '''
        export GIT_COMMIT_SHA=$(git log -n 1 --pretty=format:'%h')
        echo "Git commit SHA=$GIT_COMMIT_SHA"
        '''
      }
    }

    stage('Build Dev') {
      steps {
        echo 'Building...'
        sh '''
          make docker-build-dev 
        '''
      }
    }

    stage('Test') {
      steps {
        echo 'Testing...'
        sh '''
          make docker-run-test 
        '''
      }
    }

    stage('Build Final') {
      steps {
        echo 'Building...'
        sh '''
          make build
          make docker-build-final
        '''
      }
    }

    stage('Docker Push') {
      steps {
        echo 'Not implemented yet. $GIT_COMMIT_SHA'
      }
    }

    stage('Cleanup') {
      steps {
        echo 'Time to cleanup! Currently not implemented.'
      }
    }

  }
}