# Using KF
This guide will walk you through how to push a spring app 
to kf using the the Cloud Native App Bundle (CNAB) buildpack. 

To start you'll need to have completed the [install](./install.md)
and have the mini broker running. It's also handy to configure
your registry in a global environment variable. 
```
export KF_REGISTRY=gcr.io/<PROJECT_ID>
```

## Sample Apps & Buildpacks
This repo includes a few sample apps that integrate with
the provided buildpack. This is intended to give you a
baseline framework for developing your own apps and
buildpacks as they are needed. 

### Spring Pet Clinic
The Pet Clinic app included the third_party directory
is a jaava Spring Boot app built using Maven. It is a 
fork of the [Spring Pet Clinic](pet-clinic) app.

To use this app you'll need to install the CNAB buildpack
included in the samples. 

```
kubectl apply -f ./samples/buildpacks/cnab.yaml
```

Now you'll create a my-sql instance from the Mini broker
that we'll be using for our app. You can do that with the 
follow create service command. 
```
kf create-service mysql 5-7-14 petdb
```

The fastest workflow for creating apps bound to services
is to create the binding before pushing the app. To do this
you can run the following bind-service command. 
```
kf bind-service petclinic petdb --binding-name petclinic-mypets
```

Finally you need to push the application with the necessary 
settings for the environment variables.  
```
kf push petclinic \
 --container-registry $KF_REGISTRY \
 -e VCAP_APPLICATION='{"name":"petclinic"}' \
 -e VCAP_SERVICES="$(kf vcap-services petclinic)" \
 -e DATABASE=mysql \
 -e SPRING.DATASOURCE.INITIALIZATION_MODE=always \
 -e SPRING.DATASOURCE.URL='jdbc:mysql://${vcap.services.kf-binding-petclinic-petdb.credentials.host}:${vcap.services.kf-binding-petclinic-petdb.credentials.port}/petclinic' \
 -e SPRING.DATASOURCE.USERNAME=root \
 -e SPRING.DATASOURCE.PASSWORD='${vcap.services.kf-binding-petclinic-petdb.credentials.mysql-root-password}'
```

Once the app has completed it's build (usually about 120 seconds)
you can use the proxy command to proxy the app and browse using
[local host](http://localhost:8080/). 

```
kf proxy petclinic
```

[pet-clinic]: https://github.com/spring-projects/spring-petclinic