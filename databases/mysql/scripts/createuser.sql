CREATE USER 'docker-example-app'@'%' IDENTIFIED BY 'docker-example-2024#';

GRANT SELECT, INSERT, UPDATE, DELETE ON exampledb.* TO 'docker-example-app'@'%';