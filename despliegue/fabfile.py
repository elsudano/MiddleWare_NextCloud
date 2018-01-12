import os
from fabric.api import env, local, run, sudo

env.user = 'usuario'
env.hosts = ['middleware.westeurope.cloudapp.azure.com']
env.key_filename = '~/.ssh/id_rsa_deploying'

def levantar_maquina():
    local('vagrant up --no-provision')

def destruir_maquina():
    local('vagrant destroy --force')

def configurar_maquina():
    local('vagrant provision')
    local('sed "/middleware.westeurope.cloudapp.azure.com/d" ~/.ssh/known_hosts > ~/.ssh/known_hosts')
    if run('echo $DOMAIN') == '':
        sudo('echo DOMAIN="' + os.environ['DOMAIN'] + '" >> /etc/environment')
    if run('echo $URL_BASE') == '':
        sudo('echo URL_BASE="' + os.environ['URL_BASE'] + '" >> /etc/environment')
    if run('echo $USER_NEXTCLOUD') == '':
        sudo('echo USER_NEXTCLOUD="' + os.environ['USER_NEXTCLOUD'] + '" >> /etc/environment')
    if run('echo $PASS_NEXTCLOUD') == '':
        sudo('echo PASS_NEXTCLOUD="' + os.environ['PASS_NEXTCLOUD'] + '" >> /etc/environment')
    if run('echo $PORT') == '':
        sudo('echo PORT="80"' + ' >> /etc/environment')

def ejecutar_aplicacion():
    sudo('~/go/bin/MiddleWare_NextCloud &')

def detener_aplicacion():
    local('curl http://middleware.westeurope.cloudapp.azure.com/exit &')

def deploy():
    levantar_maquina()
    configurar_maquina()
    ejecutar_aplicacion()

def undeploy():
    detener_aplicacion()
    destruir_maquina()
