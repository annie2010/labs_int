<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Picker Service Container Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission...

> Build and Dockerize a Word Picker Web Service!

1. Clone the [Labs Repo](https://github.com/gopherland/labs_int)
2. cd picker_svc
3. Make sure Docker is installed (see installation below)!
4. Create your module file for your repository using your OWN github user handle!
5. Ensure all the tests are passing!
6. Build a picker service executable by hand
   1. Make sure you can set the service version via your build command by leveraging build flags
   2. Run the picker service locally
   3. Ensure it prints your custom version upon starting!
   4. Exercise the picker service endpoints and make sure all is cool!
7. Modify the provided Dockerfile and .dockerignore files to build a Docker image
   1. Make sure you can set the application version via your build command
   2. Make sure your image is not loading extraneous files!
8. Run your Picker service Docker image locally (see commands below)
9. Ensure it prints your custom version upon starting
10. Verify your service endpoints are working correctly on the container
11. Note the size of your current Docker image and your local executable.
12. Next change the Docker base image to use a scratch image instead
13. Rebuild your Docker image
14. Repeat the step above to launch your new Docker container and validate the
    service API
15. Now the surprise!
    Compare the size of this new image you've just build vs the old image?
16. BONUS! Publish your Picker Servie Docker image
    1. Publish your image on a public registry (DockerHub, Quay,...)
    2. Share your image url with your classmates so they can check out your work!

## Installation

1. Download and Install Docker on your machine [SKIP IF ALREADY INSTALLED!]
   1. See [Docker install](https://www.docker.com/products/docker-desktop) instructions

## Commands

1. Build a Docker images

   ```shell
   docker build --rm -t picker:0.0.1 .
   ```

2. Run a Docker image exposing 4500 locally

   ```shell
   docker run -it -p 4500:4500 picker:0.0.1
   ```

3. List your Docker image

   ```shell
   docker images | grep picker
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)