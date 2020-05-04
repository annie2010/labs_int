<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Prometheus Hangman Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> We're going to play a hangman game. The game consists of a
> hangman service hangman and a CLI to submit guesses. The hangman
> service queries a given dictionary to get a list of words for the guess
> word. To play the game, we are going to leverage Prometheus metrics to
> track good/bad guess counts as well as tracking win rates. Sounds cool?

1. Instrument the hangman code base and add 2 prometheus counters to track your
   good and bad guesses (see game.go).
2. Next define a prometheus gauge to track your game results:
   ie +1 for wins and -1 for loss (see tally.go)
3. Install prometheus and configure the scraper to scrape your hangman service on a given port.
4. Start your hangman service
5. You can now enjoy the fruits of your labor by firing off the provided hangman CLI (cmd/cli/main.go) and try out your guessing skills while watching your game performance in the prometheus dashboard...

## Commands

1. Download install Prometheus
   1. For OSX use the following command

      ```shell
      cd /tmp
      wget https://github.com/prometheus/prometheus/releases/download/v2.18.0-rc.1/prometheus-2.18.0-rc.1.darwin-amd64.tar.gz
      tar -xvzf /tmp/prometheus-2.18.0-rc.1.darwin-amd64.tar.gz
      IMPORTANT!! Make sure to copy the prometheus binary so that it is in your $PATH
      ```

   2. For other platforms please see [Prometheus Install](https://prometheus.io/download)

1. In one terminal launch prometheus with your custom scraper config file

      ```shell
      prometheus --config.file=~/gopherland/labs/prometheus/config/prom_scraper.yml
      ```

1. Prometheus Dashboard

   ```shell
   open http://localhost:9090/graph
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
