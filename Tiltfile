default_registry('hub.ziemba.net')
registry = 'hub.ziemba.net'
website = 'ip.ziemba.net'
version = 'V0.0.2'
timestamp = str(local("powershell get-date -format 'yyyyMMdd_HHmm'")).strip()
################################################################################
docker_build(registry + '/website', 
            './Docker',
            dockerfile='./Docker/Dockerfile',
            extra_tag= [registry + '/' + website + ':' + timestamp, registry + '/' + website + ':' + version, registry + '/website:' + version] ,
            )
################################################################################
k8s_yaml('./K8S_Files/deploy.yaml')