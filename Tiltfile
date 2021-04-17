# Load external restart proces
load('ext://restart_process', 'docker_build_with_restart')

# File service deployment and live development
k8s_yaml(helm('deployments/file-service', name='file-service'))

docker_build_with_restart(
  'k3d-mr-registry:5000/file-service', 
  '.',
  entrypoint='/start_app',
  target='dev',
  dockerfile='./services/file-service/Dockerfile',
  live_update=[
    sync('./services/file-service', '/usr/src/app'),
  ],
)