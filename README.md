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

## Sistemas de Integración Continua

Se usará el sistema de integración continua TRAVIS.CI para poder testear de manera fácil y con los menos errores posibles el MiddleWare que se encargara de realizar conexiones con un servidor de Owncloud.

Los tests serán tests unitarios en un primer momento, mas adelante se intentara que sean tests de covertura.
