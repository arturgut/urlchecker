pipeline{

  agent any
  
  stages{
    
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

    stage('Docker Push') {
      steps {
        echo 'Not implemented yet'
      }
    }

    stage('Cleanup') {
      steps {
        echo 'Time to cleanup!'
      }
    }

  }
}