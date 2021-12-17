.PHONY:all

build:
	docker build -t jndi-ldap-test-server .

run:
	docker run -i -p 1389:1389 jndi-ldap-test-server