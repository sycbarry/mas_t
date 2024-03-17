FROM maven:3.6.3 AS maven
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN mvn package
FROM openjdk:8-jdk-alpine
ARG JAR_FILE=backend-0.0.1.jar
WORKDIR /opt/app
COPY --from=maven /usr/src/app/target/${JAR_FILE} /opt/app/
#COPY ./src/main/resources/static/images/* /opt/app/
EXPOSE 8080
ENTRYPOINT ["java","-jar","${JAR_FILE}.jar"]
