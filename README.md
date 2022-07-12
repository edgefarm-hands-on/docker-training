---
marp: true
title: Basic Containers Training
paginate: true
header: 'Basic Containers Training'
footer: '![height:40px](.assets/ci4rail_logo.png)'

style: |
  header, footer {
    font-size: 20pt;
  }
  h1{
      padding: 0;
      margin: 0;
  }
  h2, h3{
      padding: 0;
      margin: 5px;
  }
---
# Basic Containers Training

Welcome! Here you will find everything that is needed for this course.

In this training we will talk about the very `basics` of containers like:

* what is a `container`?
* how do they work?
* how do i interact with containers?
* what is the difference between `container images` and `containers`?
* does it even `compare` to a `virtual machine`?

---

We will also take a deeper look at some topics on `managing` container images:

* how do i `start and stop` my containers?
* how do i `build` my own containers?
* how do i `store and share` my own container `images`?

---

We will also take a deeper look at some `technical stuff`:

* what are the most popular container runtimes
* how do i let my container `use` or `provide` `resources` like ports or devices?
* i'm running on ARM. How do i cross-compile my images?

Many sections in this training are coupled with examples and hands-on-lessons.
I don't want you to just listen. I want you to explore for yourself.

So let's start then!

---
# Prerequisites

Wait! Before we begin there are some actions to take.

If you want to follow along and try for yourself you can use a local installation of a container runtime e.g. `docker`. Follow the linked instructions on [how to install docker](https://docs.docker.com/get-docker/).
If you don't want or can't install `docker` locally, please make sure that you signed up for a free account on https://github.com and https://gitpod.io. Using `gitpod` you can also follow along all examples in your web browser. The browser plugin for gitpod is highly recommended as you can simply create an gitpod instance from an github repository -> [Gitpod browser extension](https://www.gitpod.io/docs/browser-extension).

---
Prepare a free account at https://docker.io for uploading your custom docker images.

In this training all examples refer to `docker` - however many other popular container runtimes are pretty compatible in usage to `docker` like `crictl`, `podman`, `nerdctl`, ...
If you feel adventurous use whatever container runtime you want. If you want to stay safestick with `docker`.

---

# Basics

We'll start with the absolute basics before jumping in and do cool stuff.

---
## Before we start

Did you have any contact with the world of containers yet?
Have you ever used containers or docker before?

**Tell us!**

---
## What is a `container` and how does it work?

**A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another.**

Container runtimes use a technology called `namespaces` in the Linux kernel to provide the isolated workspace called the container. When you run a container, the runtime creates a set of namespaces for that container.

These namespaces provide a layer of isolation. Each aspect of a container runs in a separate namespace and its access is limited to that namespace.

---
**Wait, what does that mean in the language of simple people?**

![center w:500px](https://www.docker.com/wp-content/uploads/2021/11/container-what-is-container.png.webp)

---

The container runtime (today it's docker) runs on the host sharing the same infrastructure. Docker then runs the applications that are encapsulated in containers `App A` - `App F`. The running containers share the same underlying Linux Kernel. This means the application does not bring it's own kernel. It just brings what the application needs to run in terms of libraries, dependencies and so on. Each application can be configured to access a limited amount of CPU or RAM resources.

---
**In short, what are the advantages containers bring to us?**

1. All the applications dependencies are packed up in one simple thing that can be easily versioned and shared.
2. If you've build a container image and it runs well on your machine, you can be absolutely sure that this will just work fine on another machine (depending on the CPU architecture).
3. There is no runtime overhead. Processes within the container are handeled by the hosts Linux kernel.
4. The container image brings its own root filesystem. Is does not matter which Linux distribution runs the container.
---
## How do i interact with containers?

Assuming you are running docker you can simply start a command shell on your PC. Assuming from now on that you are using linux.

Running the `docker --help` command gives various help what you can do.

```sh
$ docker --help

Usage:  docker [OPTIONS] COMMAND
...
```

---
Just to name a few: 
* ps (list curtently running)
* run (instanciates a new container)
* build (builds a new container image)
* push (pushes a locally present image to a container registry)
* pull (pulls a remote present image from a container registry)
* rm (deletes containers)

There many more options. But we will take a closer look later on in this training.

---

## What is the difference between `container images` and `containers`?

Simple: the container is a running instance of an image

You can instanciate as many containers from the same image as you like.
So you are not distributing a container, but a container image.

---
## Does it even `compare` to a `virtual machine`?

No, not really. Virtual machines bring a bunch of CPU and memory resources, have their own operating system running an provide a own Linux kernel. The container however shares a bunch of resources like CPU, Memory, disk, Linux kernel and so on. 

Head over to [this external link for a deeper comparison](https://www.netapp.com/blog/containers-vs-vms/)  if you like.

---
## What's the big picture in terms of `docker`?

![center](https://docs.docker.com/engine/images/architecture.svg)

---
The container `registry` is a place where you can upload your images and provide access for people to pull them too. Popular ones are [Docker Hub](https://hub.docker.com) or the [github container registry](https://github.com/features/packages). They both offer free plans.

The `client` builds, pushes, pulls and runs images. That's you making the calls what to do.

The `DOCKER_HOST` can be running on the same local machine the client does stuff. However, it can also be some remote machine. The docker daemon performs the action, e.g. it pulls an image from the registry and runs it with parameters the client passed.

---

**Need a break before we are going to try out some cool containers?**
![w:400px center](https://thumbs.dreamstime.com/b/coffee-mug-17527982.jpg)

---

# Let's 


Now we are are ready to try out running our first container. Run this in your shell:

```sh
$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
2db29710123e: Pull complete 
Digest: sha256:13e367d31ae85359f42d637adf6da428f76d75dc9afeb3c21faea0d976f5c651
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash
...
```

---

Excersise:

1. Write a `Dockerfile` and `build` a container image that contains the application `exercise/webserver-volume-access`. Follow the `exercise/webserver-volume-access/README.md` on how to build the application.
Make sure the resulting image is as `small as possible`!
2. `Run` the container so, that the container has the following characteristics:

* let the internal container port `8080` be accesible at the hosts port `9090`.
* give the container access to the hosts file `exercise/webserver-volume-access/input` in containers filesystem at `/app/input`.

3. Test your setup by updating the input file. After requesting the application the new values shall be printed.

```sh
$ echo "hello world" > examples/webserver-volume-access/input
$ curl localhost:9090
hello World
$ echo "I am standing on the moon" > examples/webserver-volume-access/input
$ curl localhost:9090
I am standing on the moon
```
