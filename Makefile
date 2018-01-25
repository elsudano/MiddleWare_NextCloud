###############################################################################
# Variable que guarda los ficheros que hay que compilar
SRC_DIR =

###############################################################################
# Extraemos los nombres de los ficheros para los binarios
BIN_DIR = bin/

###############################################################################
# Variables que no cambian para cualquier compilación
# Extención de los ficheros fuente
EXT = go # pueden ser go

###############################################################################
# ruta del comando cliente para el despliegue en Zeit.co
NOW_COMMAND = /home/usuario/now-linux
SSH_COMMAND = ssh
SED_COMMAND = sed

###############################################################################
# Cliente web para linea de comandos
WEBCLI = curl

###############################################################################
# Configuración para las configuraciones remotas
WEB_LOCAL_SRV = localhost:# pruebas en local
URL_HEROKU = middleware-nextcloud.herokuapp.com
URL_ZEIT = middlewarenextcloud-ysiglvkfgt.now.sh
URL_AZURE = middleware.westeurope.cloudapp.azure.com
WEB_HEROKU_SRV = https://${URL_HEROKU}#Pruebas en remoto
WEB_ZEIT_SRV = https://${URL_ZEIT}#Pruebas en remoto
WEB_AZURE_SRV = http://${URL_AZURE}#Pruebas en remoto
ID_RSA_FILE = ~/.ssh/id_rsa_deploying
KNOWN_FILE = ~/.ssh/known_hosts

###############################################################################
# Cabeceras y metodos las peticiones de http que son solicitudes
ID = KJ43MS30MEJQIR6ULBI0U9 # se pasa por parametros para sobreescribir esta variable
STRING2 = -d '{"denomination":"Prueba nueva", "description":"esta es una prueba de addición de evento", "reminders":[{"id":"1", "denomination":"Primer recordatorio", "datestart":"2017/12/10","timestart":"20:00"}], "datestart":"2017/12/10", "dateend":"2017/12/10", "timestart":"22:00","timeend":"23:00"}' --header "Content-Type: application/json" -X POST

###############################################################################
# Nombre del programa Principal
NAME = MiddleWare_NextCloud

###############################################################################
# Comando para eliminar los ficheros
RM = rm -Rf

###############################################################################
# Indica cual es el compilador para los fuentes
COMPILER = go

.PHONY: help
.SECONDARY:

all: clean help

makedir:
	@mkdir ./$(BIN_DIR)

run:
	$(COMPILER) run $(SRC_DIR)*.$(EXT)

build: clean makedir
	$(COMPILER) build -o $(BIN_DIR)$(NAME) $(SRC_DIR)*.$(EXT)

deploy: clean
	$(NOW_COMMAND) -e USER_NEXTCLOUD=${USER_NEXTCLOUD} -e PASS_NEXTCLOUD=${PASS_NEXTCLOUD} -e DOMAIN=${DOMAIN} -e URL_BASE=${URL_BASE} -e PORT=${PORT} --public

azure: clean
	@vagrant up

azure_prov: clean
	@vagrant provision

connect:
	@$(SED_COMMAND) '/$(URL_AZURE)/d' $(KNOWN_FILE) > $(KNOWN_FILE)
	@$(SSH_COMMAND) -i $(ID_RSA_FILE) $(URL_AZURE)

clean_azure: clean
	@vagrant destroy --force

clean_deploy: clean
	$(NOW_COMMAND) rm $(NAME) -y

tests: clean
	$(COMPILER) test $(SRC_DIR)*_test.$(EXT)

test_status: clean
	@echo "LOCAL"
	$(WEBCLI) $(WEB_LOCAL_SRV)${PORT}/status
	@echo "HEROKU"
	$(WEBCLI) $(WEB_HEROKU_SRV)/status
	@echo "ZEIT"
	$(WEBCLI) $(WEB_ZEIT_SRV)/status
	@echo "AZURE"
	$(WEBCLI) $(WEB_AZURE_SRV)/status

test_list: clean
	$(WEBCLI) $(WEB_LOCAL_SRV)${PORT}/list | jq

test_show: clean
	$(WEBCLI) $(WEB_LOCAL_SRV)${PORT}/show/$(ID) | jq

test_new: clean
	$(WEBCLI) $(WEB_LOCAL_SRV)${PORT}/new $(STRING2)

test_update: clean
	$(WEBCLI) $(WEB_LOCAL_SRV)${PORT}/update $(ID)

test_delete: clean
	$(WEBCLI) $(WEB_LOCAL_SRV):${PORT}/delete $(STRING1)

docu:
	@doxygen ./doc/doxys/dox_config

clean:
	$(RM) $(BIN_DIR) doc/html/* doc/latex/*

touch:
	@touch $(SRC_DIR)/*

help:
	@echo "Available targets:"
	@echo "- run          Run Application whitout compile"
	@echo "- build        Generate bin file on directory $(BIN_DIR)"
	@echo "- deploy       Execute now for deploy on Zeit.co"
	@echo "- azure        Execute vagrant up for deploy in Azure Cloud with vagrant"
	@echo "- azure_prov   Provisioning VM with ansible in Azure"
	@echo "- connect      Connect by ssh to Azure Virtual Machine"
	@echo "- clean_azure  Execute vagrant destroy for delete VM to Azure"
	@echo "- clean_deploy Clean all instances of Zeit.co"
	@echo "- tests        Run tests"
	@echo "- test_status  Run test for status command in all containers"
	@echo "- test_list    Run test for list command of webservice"
	@echo "- test_show    Run test for show command of webservice"
	@echo "- test_new     Run test for new command of webservice"
	@echo "- test_update  Run test for update command of webservice"
	@echo "- test_delete  Run test for delete command of webservice"
	@echo "- docu         Create a documentation of project"
	@echo "- clean        Clean up the source directory $(SRC_DIR) and bin directory $(BIN_DIR)"
	@echo "- help         This info"
	@echo
