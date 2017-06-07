package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	offset    = 0
	randrange = 10
)

type randHandler struct {
	rand *rand.Rand
}

func (r randHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var output string

	if req.URL.Path == "/data" {
		randNum := rand.Intn(randrange)
		w.Header().Set("Content-Type", "application/json")
		output = fmt.Sprintf(`{"value": "%d"}`, randNum+offset)
	} else {
		output = html
	}
	fmt.Fprintln(w, output)

}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := randHandler{
		rand.New(s1),
	}

	listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(listen, r))
}

var html = `
<!DOCTYPE html>
<html>
<head>
	<title>Random Data Grapher</title>
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.3.7/css/bootstrap.min.css" />
</head>
<body>
	<div class="container-fluid">
		<h1 id="dotCount">0</h1>
	    <canvas id="widgetChart"></canvas>
	</div>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.5.0/Chart.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
    <script type="text/javascript">
    var ctx = document.getElementById("widgetChart").getContext("2d");
    var data = {
      labels: [],
      datasets: [{
        label: "Widgets",
        backgroundColor: "rgba(64,0,144,0.5)",
        borderColor: "rgba(64,0,144,0.8)",
        data: []
      }]
    };
    var options = {
      legend: {
        display: false
      },
      scales: {
        yAxes: [{
          ticks: {
            min: 0,
            max: 50
          }
        }],
        xAxes: [{
          scaleLabel: {
            display: false
          },
          ticks: {
            display: false
          }
        }]
      }
    };
    var lineChart = new Chart.Line(ctx, {
       data: data,
       options: options
    });
    setInterval(function(){
      $.ajax({ url: "data" }).done(function (newWidgets) {
        data.datasets[0].data.push(newWidgets.value);
		$("#dotCount").text(newWidgets.value);
        data.labels.push(Date.now());
        if (data.labels.length > 30) {
          data.labels.shift();
          data.datasets[0].data.shift();
        }
        lineChart.update();
      });
    }, 1000);
    </script>
</body>
</html>
`
