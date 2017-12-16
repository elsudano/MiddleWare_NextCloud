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

###############################################################################
# Cliente web para linea de comandos
WEBCLI = curl

###############################################################################
# Cliente web para linea de comandos
WEBSRV = localhost:# pruebas en local
#WEBSRV = https://${DOMAIN}${URL_BASE}:#Pruebas en remoto

###############################################################################
# Cabeceras y metodos las peticiones de http que son solicitudes
ID = KJ43MS30MEJQIR6ULBI0U9 # se pasa por parametros para sobreescribir esta variable
STRING2 = -d '{"id":"", "denomination":"Prueba nueva", "description":"esta es una prueba de addición de evento", "reminders":[{"id":"1", "denomination":"Primer recordatorio", "datestart":"2017/12/10","timestart":"20:00"}], "datestart":"2017/12/10", "dateend":"2017/12/10", "timestart":"22:00","timeend":"23:00"}' --header "Content-Type: application/json" -X POST

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

all: clean build

makedir:
	mkdir ./$(BIN_DIR)

run:
	$(COMPILER) run $(SRC_DIR)*.$(EXT)

build: clean makedir
	$(COMPILER) build -o $(BIN_DIR)$(NAME) $(SRC_DIR)*.$(EXT)

deploy: clean
	$(NOW_COMMAND) -e USER_NEXTCLOUD=${USER_NEXTCLOUD} -e PASS_NEXTCLOUD=${PASS_NEXTCLOUD} -e DOMAIN=${DOMAIN} -e URL_BASE=${URL_BASE} -e PORT=${PORT} --public

cleandeploy: clean
	$(NOW_COMMAND) rm $(NAME) -y

tests: clean
	$(COMPILER) test $(SRC_DIR)*_test.$(EXT)

test_list: clean
	$(WEBCLI) $(WEBSRV)${PORT}/list | jq

test_show: clean
	$(WEBCLI) $(WEBSRV)${PORT}/show/$(ID) | jq

test_new: clean
	$(WEBCLI) $(WEBSRV)${PORT}/new $(STRING2)

test_update: clean
	$(WEBCLI) $(WEBSRV)${PORT}/update $(ID)

test_delete: clean
	$(WEBCLI) $(WEBSRV):${PORT}/delete $(STRING1)

docu:
	doxygen ./doc/doxys/dox_config

clean:
	$(RM) $(BIN_DIR) doc/html/* doc/latex/*

touch:
	touch $(SRC_DIR)/*

help:
	@echo "Available targets:"
	@echo "- run         Run Application whitout compile"
	@echo "- build       Generate bin file on directory $(BIN_DIR)"
	@echo "- deploy       Execute now for deploy on Zeit.co"
	@echo "- clean       Clean up the source directory $(SRC_DIR) and bin directory $(BIN_DIR)"
	@echo "- test        Run tests"
	@echo "- help        This info"
	@echo
	@echo "Available variables:"
	@echo "- SRC_DIR      default: $(SRC_DIR)"
	@echo
