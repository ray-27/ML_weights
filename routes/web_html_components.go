package handler

var Chart_component string = `
<div class="chart-wrapper">
					<canvas id="%s" width="400" height="200"></canvas>
					<script>
						(function initializeChart() {
							const ctx = document.getElementById("%s").getContext("2d");

							const chart = new Chart(ctx, {
								type: "line",
								data: {
									labels: [],
									datasets: [
										{
											label: "Chart ",
											data: [],
											borderColor: "blue",
											backgroundColor: "rgba(0, 0, 255, 0.1)",
											fill: true,
										},
									],
								},
								options: {
									responsive: true,
									scales: {
										x: {
											type: "time",
											time: {
												unit: "second",
												displayFormats: {
													second: "h:mm:ss a",
												},
											},
										},
										y: {
											beginAtZero: true,
										},
									},
								},
							});

							const socket = new WebSocket("ws://localhost:8080/ws");

							socket.onmessage = function (event) {
								const data = JSON.parse(event.data);
								const timestamp = new Date(data.timestamp * 1000);
								const value = data.value;

								chart.data.labels.push(timestamp);
								chart.data.datasets[0].data.push(value);

								chart.update();
							};

							socket.onerror = function (error) {
								console.error("WebSocket error:", error);
							};
						})();
					</script>
				</div>
`
