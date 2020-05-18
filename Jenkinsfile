pipeline{

  agent any
  
  stages{
    
    stage('Build') {
      steps {
        echo 'Building...'
      }
    }

    stage('Test') {
      steps {
        echo 'Testing...'
      }
    }

    stage('Push To Docker Registry') {
      steps {
        echo 'Pushing to docker registry...'
      }
    }

  }
}