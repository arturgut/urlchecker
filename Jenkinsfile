pipeline{

  agent any
  
  stages{
    
    stage('Build Dev') {
      steps {
        echo 'Building...'
        sh '''
          echo env.GIT_COMMIT
          echo env.GIT_BRANCH
          echo env.GIT_REVISION
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
          make docker-build-final
        '''
      }
    }

    stage('Docker Push') {
      steps {
        echo 'Not implemented yet'
      }
    }

    stage('Cleanup') {
      steps {
        echo 'Time to cleanup! Currently not implemented.'
      }
    }

  }
}