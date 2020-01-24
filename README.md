# GoLang(Beego) + Docker Compose

Horeca API Report
======

NOTE
----
The master branch will always contain the latest stable version. If you wish to check older versions or newer ones currently under development, please switch to the relevant branch.

Get Started
-----------

#### Requirements

To run this application on your machine, you need at least:

* docker-compose = 3.7
* GO = 1.12.8-r0


Application flow pattern:
---------------------
http://git365.eggdigital.com/horeca/horeca-api-report

Run the docker for development:
---------------------
You can now build, create, start, and attach to containers to the environment for your application. To build the containers use following command inside the project root:

```bash
docker-compose up -d --build
```

Map the domain
------------------------------------
Open the hosts file on your local machine `/etc/hosts`.
```bash
127.0.0.1  api-report.eggsmartpos.local
```

Running Application
------------------------------------
Open the browser
```bash
http://api-report.eggsmartpos.local:8304
```

