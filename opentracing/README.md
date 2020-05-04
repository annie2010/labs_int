<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# OpenTracing G.O.T Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Game Of Thrones (G.O.T) Reloaded!

> In this lab, we are going to decorate a web server using **OpenTracing**.
> There are 2 services involved: **Castle** and **Knight**. The
> Knights want to melt Castles, but if you're a G.O.T fan, you already
> know that only the NightKing can melt a Castle using his undead dragon 🙀🐉...
> The Knight service accepts post requests on */api/v1/melt* and issues a
> post */api/v1/melt* on the Castle service with a given Knight name.
> The Castle service returns either a 200 with a castle melted message if the
> knight is the `NightKing` 😵 or a 417 error with *only NightKing can melt* otherwise.

1. Instrument the Castle service by tracing incoming *melt* requests
   1. Edit your Castle trace and add the following tags to the trace:
      1. http.method
      2. http.url
      3. knight
2. If the given Knight is *NightKing* add a log to the castle span to indicate `the castle is melted`.
3. All other knights should produce a span error (internal/http.go).
4. Span errors are indicated as follows:
   1. Setting a span tag error=true
   2. Adding a structured log on the span using
      1. event=error
      2. message=only the NightKing can melt the castle
5. Using the provided command start the Jaeger servicer
6. In a separate terminal start your Castle and Knight services.
7. Using the Jaeger Dashboard (see command below) validate that your traces are correctly tracking the workload by using different knights.

## Commands

1. Download and Install Docker on your machine [SKIP IF ALREADY INSTALLED!]
   1. Please follow the [Docker For MY_PLATFORM install](https://www.docker.com/products/docker-desktop) instructions
1. Start a Jaeger server

   ```shell
   docker run --name jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
   ```

1. In one terminal launch prometheus with your custom scraper config file

      ```shell
      prometheus --config.file=config/prom_scraper.yml
      ```

1. Prometheus Dashboard

   ```shell
   open http://localhost:9090/graph
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
