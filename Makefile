#APP_COMPOSE := $("./Docker/app.yaml")
#KAFKA_COMPOSE := $("./Docker/broker.yaml")
#NEBULA_COMPOSE := $("./Docker/nebula.yaml")

run:
	docker-compose -f ./Docker/app.yaml -f ./Docker/broker.yaml -p health-sentinel up