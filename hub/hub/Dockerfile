FROM maven:3.8.4 as maven
WORKDIR /app
COPY . /app
RUN mvn package
FROM openjdk:17-jdk
ARG JAR_FILE=mastelemetry-hub-0.0.1.jar
WORKDIR /opt/
COPY --from=maven /app/target/${JAR_FILE} /opt/
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "/opt/mastelemetry-hub-0.0.1.jar"]
