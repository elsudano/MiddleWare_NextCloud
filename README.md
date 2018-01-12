# MiddleWare for NextCloud Platform

[![Build Status](https://travis-ci.org/elsudano/MiddleWare_NextCloud.svg?branch=master)](https://travis-ci.org/elsudano/MiddleWare_NextCloud)
[![Maintainability](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud/badges/maintainability.svg)](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud/maintainability)
[![Test Coverage](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud/badges/coverage.svg)](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud/coverage)
[![Issue Count](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud/badges/issue_count.svg)](https://codeclimate.com/github/elsudano/MiddleWare_NextCloud)

<p>Aplicación que pretende ser un paso intermedio para futuras aplicaciones de consulta rápida a la estructura ya montado de NextCloud</p>

<p>Lo que se pretende en este proecto es poder desplegar de manera automática toda la infraestructura necesaria para poder hacer funcionar este middleware y así de esta forma hacer posible la conexión de otras aplicaciones a owncloud de manera mas fácil y trasparente </p>

## Objetivos principales

<p>Los objetivos principales de este proyecto es que se pueda, deplegar lo mas rapidamente una insfraestructura desentralizada on todos los componentes necesarios para que tengamos un asistente personal en nuestro móvil</p>

## Herramientas


## Tests, Intalación y ejecución

## Despliegue en PaaS
Para realizar el despliegue de la aplicación simplemente se ha unido la cuenta de Heroku con la cuenta de Github para que se realicen los despliegues automáticos.

Esto simplemente hay que habilitarlo desde el panel de control del usuario de Heroku, en las settings de la propia aplicación que se ha creado para el despliegue

Despliegue https://middleware-nextcloud.herokuapp.com/

## Despliegue en Docker
Para poder realizar el despliegue de la aplicación se realiza en el proveedor [Zeit.com](Zeit.com)

Contenedor https://middlewarenextcloud-onimeosalb.now.sh

El fichero dockerfile https://hub.docker.com/r/elsudano/middleware_nextcloud/ se sube al repositorio publico de Docker Hub para que sea accesible en cualquier momento.

## Despliegue con Vagrant y Ansible
Antes de poder realizar cualquier despliegue con Vagrant hace falta que tengamos en nuestro repositorio de boxes la maquina base con la cual vamos a empezar a trabajar, para ello tenemos que ejecutar el siguiente comando: `vagrant box add <PATH, URL>`.

En este comando podemos especificar una ruta completa de nuestro ordenador para que Vagrant importe dicha *imagen* o bien podemos usar una ya creada, de las muchas que se encuentran en nuestro proveedor de azure.

Para poder realizar el despliegue de la aplicación es necesario instalar los providers necesarios, para ello utilizaremos el comando: `vagrant plugin install <provider>` en nuestro caso tenemos que usar como prividers, **azure** y **vmware_workstation**

Después de realizar esto nos aseguraremos que tenemos todas las variables de entorno correctas para poder desplegar en Azure.

- AZURE_APPID = Hash de ID de subscripción
- AZURE_PASSWORD = Hash de contraseña de subscripción
- AZURE_SUBSCRIPTION = Variable de Subscripción de Azure
- AZURE_TENANT = Variable del servidor de Dominio para Azure
- AZURE_VM_PASS = Contraseña de la MV
- AZURE_DISPLAYNAME = Nombre Visible de la Conexión
- AZURE_NAME = Nombre de la conexión

Para conseguir todas estas o casi todos estos valores de las variables tenemos que crear nuestro Active Directory en Azure para eso ejecutamos `az ad sp create-for-rbac -o json`, y eso nos dará algunas respuestas.

Después de esto tenemos que listar nuestra subscripción para eso ejecutaremos el comando `az account list --query "[?isDefault].id" -o json` el cual nos dará el ID de la subscripción.

Ahora ya podemos generar nuestro fichero de configuración de Vagrant o bien con el comando `vagrant init` o bien generandolo a mano tal y como se ha realizado en el fichero [Vagrant](https://github.com/elsudano/MiddleWare_NextCloud/blob/master/Vagrantfile)

Una vez que le despliegue de la maquina virtual se lleva a cabo se utiliza ansible para poder configurar la aplicación que queremos ejecutar en dicha maquina, osea tendriamos que configurar las diferentes variables de entorno, programas y dependencias de nuestra aplicación, por ejemplo si necesitamos que nuestra aplicación funcione en un apache se tiene que configurar desde aquí, para ello usamos el [fichero Ansible](https://github.com/elsudano/MiddleWare_NextCloud/blob/master/provision/tasks.yml) para describir nuestro entorno de despliegue.

Y por ultimo utilizaremos Fabric para controlar la aplicación final tal y como se muestra en el [fichero Fabfile](https://github.com/elsudano/MiddleWare_NextCloud/blob/master/despliegue/fabfile.py)

Para comprobar todos los comandos que podemos usar con Fabric ejecutamos: `fab --list` en el mismo directorio donde se encuentra el fichero fabfile

Para poder comprobar que la aplicación funciona correctamente se accederá a la siguiente dirección:

Despliegue final: middleware.westeurope.cloudapp.azure.com

## Sistemas de Integración Continua

Se usará el sistema de integración continua TRAVIS.CI para poder testear de manera fácil y con los menos errores posibles el MiddleWare que se encargara de realizar conexiones con un servidor de Owncloud.

Los tests serán tests unitarios en un primer momento, mas adelante se intentara que sean tests de covertura.
