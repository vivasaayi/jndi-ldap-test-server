.PHONY:all

build:
	docker build -t vivasaayi/jndi-ldap-test-server:0.0.1 .

run:
	docker run -i -p 1389:1389 vivasaayi/jndi-ldap-test-server:0.0.1

push:
	docker push vivasaayi/jndi-ldap-test-server:0.0.1