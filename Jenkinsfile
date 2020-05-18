pipeline{

  agent any
  
  stages{
    
    stage('Build') {
      steps {
        echo 'Building...'
        sh '''
          make docker-build-dev 
        '''
      }
    }

    stage('Run') {
      steps {
        echo 'Running docker image...'
        sh '''
          make docker-run-dev
        '''
      }
    }

    stage('Test') {
      steps {
        echo 'Testing...'
        sh '''
          make test 
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