# Calculator CLI

Made for SE lab submission.
Goal is to get hands on with docker and jenkins.

Examples:
calculator add 2 3 -> 5
calculator div 10 2 -> 5

Jenkins is running at localhost:8080

```
cat ~/.jenkins/secrets/initialAdminPassword
```

```
https://hub.docker.com/repository/docker/deepya123/calculator-imt2023006/general
```

```
docker pull deepya123/calculator-imt2023006:latest
docker run --rm deepya123/calculator-imt2023006 add 5 7
docker run --rm deepya123/calculator-imt2023006 sub 10 4
docker run --rm deepya123/calculator-imt2023006 mul 6 3
docker run --rm deepya123/calculator-imt2023006 div 10 2
```
