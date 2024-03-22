# Description 

A small system for collecting multiple distributed IBM MAS Manage instance logs for administrative purposes and optimal organization.

...Note, this is currently in development not production ready...but is close for prime time..

# Architecture

![Image Description](./arch.png)

# Installation 

#### @ build hub. -> ``` cd hub/ ```

1. change .env values if needed. 
2. you should be able to just run ```docker compose up --build -d ```

#### @ build client collectors.

1. make sure to change any appropriate .env values to match the OC4 cluster destination.
2. build the node/Dockerfile image. 
3. push the image to a client's OpenShift repo.
4. update the k8s configs in the ``` ./kubernetes ``` folder.
5. push those configs to the OC4 cluster.



