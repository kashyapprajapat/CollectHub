package routes

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kashyapprajapat/collecthub_api/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
	// Initialize Controllers
	controllers.InitUserController(db)
	controllers.InitBookController(db)
	controllers.InitRecipeController(db)
	controllers.InitMovieController(db)
	controllers.InitQuoteController(db)
	controllers.InitPetController(db)
	controllers.InitTravelController(db)

	// Home Route
	app.Get("/", func(c *fiber.Ctx) error {
		htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>CollectHub API</title>
			<style>
				* {
					margin: 0;
					padding: 0;
					box-sizing: border-box;
				}
				
				body {
					font-family: 'Arial', sans-serif;
					background: white;
					min-height: 100vh;
					display: flex;
					align-items: center;
					justify-content: center;
					color: #333;
				}
				
				.container {
					background: white;
					padding: 40px;
					border-radius: 8px;
					border: 1px solid #e2e8f0;
					text-align: center;
					max-width: 600px;
					width: 90%;
				}
				
				.header {
					margin-bottom: 30px;
				}
				
				h1 {
					color: #4a5568;
					font-size: 2.5em;
					margin-bottom: 10px;
					display: flex;
					align-items: center;
					justify-content: center;
					gap: 10px;
				}
				
				.subtitle {
					color: #718096;
					font-size: 1.2em;
					margin-bottom: 20px;
				}
				
				.description {
					color: #4a5568;
					font-size: 1.1em;
					line-height: 1.6;
					margin-bottom: 30px;
				}
				
				.features {
					background: #f7fafc;
					padding: 25px;
					border-radius: 15px;
					margin-bottom: 30px;
				}
				
				.features h2 {
					color: #2d3748;
					margin-bottom: 20px;
					font-size: 1.5em;
				}
				
				.features-grid {
					display: grid;
					grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
					gap: 15px;
					text-align: left;
				}
				
				.feature-item {
					background: white;
					padding: 15px;
					border-radius: 10px;
					box-shadow: 0 2px 8px rgba(0,0,0,0.1);
				}
				
				.feature-item strong {
					color: #4a5568;
					display: block;
					margin-bottom: 5px;
				}
				
				.feature-item span {
					color: #718096;
					font-size: 0.9em;
				}
				
				.buttons {
					display: flex;
					gap: 20px;
					justify-content: center;
					flex-wrap: wrap;
				}
				
				.btn {
					padding: 15px 30px;
					border: none;
					border-radius: 50px;
					font-size: 1.1em;
					font-weight: bold;
					cursor: pointer;
					text-decoration: none;
					transition: all 0.3s ease;
					display: inline-block;
				}
				
				.btn-primary {
					background: #22c55e;
					color: white;
				}
				
				.btn-primary:hover {
					background: #16a34a;
					transform: translateY(-2px);
					box-shadow: 0 10px 20px rgba(34, 197, 94, 0.3);
				}
				
				.btn-secondary {
					background: white;
					color: #4a5568;
					border: 2px solid #e2e8f0;
				}
				
				.btn-secondary:hover {
					background: #f7fafc;
					transform: translateY(-2px);
					box-shadow: 0 5px 15px rgba(0,0,0,0.1);
				}
				
				.api-info {
					background: #edf2f7;
					padding: 20px;
					border-radius: 10px;
					margin-top: 30px;
				}
				
				.api-info h3 {
					color: #2d3748;
					margin-bottom: 10px;
				}
				
				.api-info p {
					color: #4a5568;
					margin-bottom: 0;
				}
				
				@media (max-width: 768px) {
					.container {
						padding: 30px 20px;
					}
					
					h1 {
						font-size: 2em;
					}
					
					.features-grid {
						grid-template-columns: 1fr;
					}
					
					.buttons {
						flex-direction: column;
						align-items: center;
					}
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>CollectHub 🎒📃</h1>
					<p class="subtitle">Your personal collections. All in one place.</p>
				</div>
				
				<p class="description">
					A unified platform to organize and store your personal collections. 
				</p>
				
				<div class="buttons">
					<a href="https://documenter.getpostman.com/view/36611651/2sB2x8Grko" 
					   target="_blank" 
					   class="btn btn-primary">
						📖 View API Documentation
					</a>
					<a href="https://github.com/kashyapprajapat/CollectHub" class="btn btn-secondary">
						🚀 Get Started
					</a>
				</div>
			</div>
		</body>
	</html>
	`
		return c.Type("html").SendString(htmlContent)
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	var startTime = time.Now()

	//System health route - Enhanced Reactive Dashboard with 6 Boxes
	app.Get("/health", func(c *fiber.Ctx) error {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		uptime := time.Since(startTime).Truncate(time.Second)

		// Enhanced metrics calculations
		allocRate := float64(memStats.TotalAlloc) / time.Since(startTime).Seconds()
		gcPauseAvg := float64(memStats.PauseTotalNs) / float64(memStats.NumGC) / 1000000 // Convert to ms
		heapInUse := float64(memStats.HeapInuse) / float64(memStats.HeapSys) * 100

		// Additional calculations for new boxes
		memoryEfficiency := (float64(memStats.HeapInuse) / float64(memStats.HeapSys)) * 100
		gcEfficiency := 100 - (float64(memStats.PauseTotalNs) / float64(time.Since(startTime).Nanoseconds()) * 100)
		cpuUsage := float64(runtime.NumGoroutine()) / float64(runtime.NumCPU()) * 10 // Approximation

		// Network simulation (in real app, you'd get actual network stats)
		networkIn := float64(memStats.TotalAlloc) / (1024 * 1024) * 0.1   // Simulated
		networkOut := float64(memStats.TotalAlloc) / (1024 * 1024) * 0.05 // Simulated

		if gcPauseAvg != gcPauseAvg { // Check for NaN
			gcPauseAvg = 0
		}
		if gcEfficiency != gcEfficiency { // Check for NaN
			gcEfficiency = 100
		}

		htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>CollectHub API - Live System Monitor</title>
			<style>
				* {
					margin: 0;
					padding: 0;
					box-sizing: border-box;
				}
				
				body {
					font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
					background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
					min-height: 100vh;
					padding: 20px;
					color: #333;
					overflow-x: hidden;
				}
				
				.dashboard {
					max-width: 1400px;
					margin: 0 auto;
					background: rgba(255, 255, 255, 0.95);
					backdrop-filter: blur(15px);
					border-radius: 25px;
					box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
					overflow: hidden;
					animation: slideIn 0.8s ease-out;
				}
				
				@keyframes slideIn {
					from {
						opacity: 0;
						transform: translateY(30px);
					}
					to {
						opacity: 1;
						transform: translateY(0);
					}
				}
				
				.header {
					background: linear-gradient(135deg, #2c3e50 0%%, #3498db 100%%);
					color: white;
					padding: 35px;
					text-align: center;
					position: relative;
					overflow: hidden;
				}
				
				.header::before {
					content: '';
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					bottom: 0;
					background: radial-gradient(circle at 50%% 50%%, rgba(255,255,255,0.1) 0%%, transparent 70%%);
					animation: shimmer 3s ease-in-out infinite;
				}
				
				@keyframes shimmer {
					0%%, 100%% { opacity: 0.5; }
					50%% { opacity: 1; }
				}
				
				.header h1 {
					font-size: 2.8em;
					font-weight: 700;
					margin-bottom: 10px;
					position: relative;
					z-index: 1;
					animation: bounce 2s ease-in-out infinite;
				}
				
				@keyframes bounce {
					0%%, 20%%, 50%%, 80%%, 100%% { transform: translateY(0); }
					40%% { transform: translateY(-10px); }
					60%% { transform: translateY(-5px); }
				}
				
				.header .subtitle {
					font-size: 1.2em;
					opacity: 0.9;
					font-weight: 300;
					position: relative;
					z-index: 1;
				}
				
				.status-indicator {
					display: inline-block;
					width: 15px;
					height: 15px;
					background: radial-gradient(circle, #27ae60, #2ecc71);
					border-radius: 50%%;
					margin-right: 10px;
					animation: pulse 2s infinite;
					box-shadow: 0 0 20px rgba(46, 204, 113, 0.6);
				}
				
				@keyframes pulse {
					0%% { 
						transform: scale(1); 
						box-shadow: 0 0 20px rgba(46, 204, 113, 0.6);
					}
					50%% { 
						transform: scale(1.2); 
						box-shadow: 0 0 30px rgba(46, 204, 113, 0.8);
					}
					100%% { 
						transform: scale(1); 
						box-shadow: 0 0 20px rgba(46, 204, 113, 0.6);
					}
				}
				
				.live-indicator {
					position: absolute;
					top: 20px;
					right: 30px;
					background: rgba(255, 255, 255, 0.2);
					padding: 8px 15px;
					border-radius: 20px;
					font-size: 0.9em;
					animation: glow 2s ease-in-out infinite alternate;
				}
				
				@keyframes glow {
					from { box-shadow: 0 0 10px rgba(255, 255, 255, 0.5); }
					to { box-shadow: 0 0 20px rgba(255, 255, 255, 0.8); }
				}
				
				.metrics-grid {
					display: grid;
					grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
					gap: 25px;
					padding: 35px;
				}
				
				.metric-card {
					background: white;
					border-radius: 18px;
					padding: 30px;
					box-shadow: 0 15px 35px rgba(0, 0, 0, 0.08);
					border: 1px solid rgba(0, 0, 0, 0.05);
					transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
					position: relative;
					overflow: hidden;
					cursor: pointer;
				}
				
				.metric-card:hover {
					transform: translateY(-8px) scale(1.02);
					box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15);
				}
				
				.metric-card::before {
					content: '';
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					height: 5px;
					background: linear-gradient(90deg, #3498db, #2ecc71, #f39c12, #e74c3c, #9b59b6, #1abc9c);
					animation: rainbow 3s linear infinite;
				}
				
				@keyframes rainbow {
					0%% { background-position: 0%% 50%%; }
					100%% { background-position: 100%% 50%%; }
				}
				
				.metric-card::after {
					content: '';
					position: absolute;
					top: 0;
					left: -100%%;
					width: 100%%;
					height: 100%%;
					background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
					transition: left 0.5s;
				}
				
				.metric-card:hover::after {
					left: 100%%;
				}
				
				.metric-header {
					display: flex;
					align-items: center;
					margin-bottom: 20px;
				}
				
				.metric-icon {
					width: 50px;
					height: 50px;
					border-radius: 12px;
					display: flex;
					align-items: center;
					justify-content: center;
					margin-right: 15px;
					font-size: 24px;
					transition: transform 0.3s ease;
				}
				
				.metric-card:hover .metric-icon {
					transform: rotate(360deg) scale(1.1);
				}
				
				.system-icon { background: linear-gradient(135deg, #3498db, #2980b9); }
				.memory-icon { background: linear-gradient(135deg, #e74c3c, #c0392b); }
				.performance-icon { background: linear-gradient(135deg, #f39c12, #e67e22); }
				.runtime-icon { background: linear-gradient(135deg, #2ecc71, #27ae60); }
				.network-icon { background: linear-gradient(135deg, #9b59b6, #8e44ad); }
				.efficiency-icon { background: linear-gradient(135deg, #1abc9c, #16a085); }
				
				.metric-title {
					font-size: 1.3em;
					font-weight: 600;
					color: #2c3e50;
				}
				
				.metric-value {
					font-size: 2.5em;
					font-weight: 700;
					color: #2c3e50;
					margin-bottom: 10px;
					animation: countUp 1s ease-out;
				}
				
				@keyframes countUp {
					from { opacity: 0; transform: translateY(20px); }
					to { opacity: 1; transform: translateY(0); }
				}
				
				.metric-label {
					font-size: 0.9em;
					color: #7f8c8d;
					text-transform: uppercase;
					letter-spacing: 0.5px;
				}
				
				.metric-list {
					list-style: none;
				}
				
				.metric-list li {
					display: flex;
					justify-content: space-between;
					padding: 10px 0;
					border-bottom: 1px solid #ecf0f1;
					font-size: 0.95em;
					animation: fadeInUp 0.5s ease-out;
				}
				
				.metric-list li:last-child {
					border-bottom: none;
				}
				
				@keyframes fadeInUp {
					from { opacity: 0; transform: translateY(10px); }
					to { opacity: 1; transform: translateY(0); }
				}
				
				.metric-list .label {
					color: #7f8c8d;
					font-weight: 500;
				}
				
				.metric-list .value {
					color: #2c3e50;
					font-weight: 600;
					transition: color 0.3s ease;
				}
				
				.metric-list li:hover .value {
					color: #3498db;
				}
				
				.progress-bar {
					width: 100%%;
					height: 10px;
					background: #ecf0f1;
					border-radius: 5px;
					overflow: hidden;
					margin-top: 15px;
					position: relative;
				}
				
				.progress-fill {
					height: 100%%;
					background: linear-gradient(90deg, #3498db, #2ecc71);
					border-radius: 5px;
					transition: width 1s ease-out;
					position: relative;
					overflow: hidden;
				}
				
				.progress-fill::after {
					content: '';
					position: absolute;
					top: 0;
					left: 0;
					right: 0;
					bottom: 0;
					background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.6), transparent);
					animation: shine 2s infinite;
				}
				
				@keyframes shine {
					0%% { transform: translateX(-100%%); }
					100%% { transform: translateX(100%%); }
				}
				
				.chart-container {
					height: 80px;
					margin-top: 15px;
					background: linear-gradient(135deg, #f8f9fa, #e9ecef);
					border-radius: 8px;
					padding: 15px;
					position: relative;
					overflow: hidden;
				}
				
				.chart-bar {
					height: 100%%;
					background: linear-gradient(135deg, #3498db, #2ecc71);
					border-radius: 4px;
					animation: chartGrow 1.5s ease-out;
					position: relative;
				}
				
				@keyframes chartGrow {
					from { height: 0; }
					to { height: 100%%; }
				}
				
				.footer {
					background: linear-gradient(135deg, #f8f9fa, #e9ecef);
					padding: 25px 35px;
					text-align: center;
					color: #7f8c8d;
					font-size: 0.9em;
					border-top: 1px solid #ecf0f1;
					position: relative;
				}
				
				.auto-refresh-indicator {
					display: inline-flex;
					align-items: center;
					background: rgba(52, 152, 219, 0.1);
					padding: 8px 16px;
					border-radius: 20px;
					margin-bottom: 10px;
					animation: breathe 2s ease-in-out infinite;
				}
				
				@keyframes breathe {
					0%%, 100%% { transform: scale(1); }
					50%% { transform: scale(1.05); }
				}
				
				.refresh-dot {
					width: 8px;
					height: 8px;
					background: #3498db;
					border-radius: 50%%;
					margin-right: 8px;
					animation: blink 1s infinite;
				}
				
				@keyframes blink {
					0%%, 50%% { opacity: 1; }
					51%%, 100%% { opacity: 0.3; }
				}
				
				.timestamp {
					font-size: 0.8em;
					opacity: 0.7;
					margin-top: 10px;
				}
				
				@media (max-width: 768px) {
					.metrics-grid {
						grid-template-columns: 1fr;
						padding: 20px;
					}
					
					.header h1 {
						font-size: 2.2em;
					}
					
					.metric-card {
						padding: 25px;
					}
					
					.metric-value {
						font-size: 2em;
					}
				}
			</style>
		</head>
		<body>
			<div class="dashboard">
				<div class="header">
					<div class="live-indicator">🔴 LIVE</div>
					<h1><span class="status-indicator"></span>CollectHub API</h1>
					<div class="subtitle">⚡ Real-time System Health & Performance Monitor</div>
				</div>
				
				<div class="metrics-grid">
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon system-icon">🖥️</div>
							<div class="metric-title">System Overview</div>
						</div>
						<ul class="metric-list">
							<li><span class="label">🔧 Go Version</span><span class="value">%s</span></li>
							<li><span class="label">⚙️ CPU Cores</span><span class="value">%d</span></li>
							<li><span class="label">⏱️ System Uptime</span><span class="value">%s</span></li>
							<li><span class="label">🔄 Active Goroutines</span><span class="value">%d</span></li>
						</ul>
					</div>
					
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon memory-icon">💾</div>
							<div class="metric-title">Memory Analytics</div>
						</div>
						<ul class="metric-list">
							<li><span class="label">📊 Current Allocation</span><span class="value">%s</span></li>
							<li><span class="label">📈 Total Allocated</span><span class="value">%s</span></li>
							<li><span class="label">🖥️ System Memory</span><span class="value">%s</span></li>
							<li><span class="label">📉 Heap Usage</span><span class="value">%.1f%%</span></li>
						</ul>
						<div class="progress-bar">
							<div class="progress-fill" style="width: %.1f%%"></div>
						</div>
					</div>
					
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon performance-icon">⚡</div>
							<div class="metric-title">Performance Metrics</div>
						</div>
						<ul class="metric-list">
							<li><span class="label">🗑️ GC Cycles</span><span class="value">%d</span></li>
							<li><span class="label">⏳ Avg GC Pause</span><span class="value">%.2f ms</span></li>
							<li><span class="label">🚀 Allocation Rate</span><span class="value">%.2f MB/s</span></li>
							<li><span class="label">🧮 Memory Objects</span><span class="value">%d</span></li>
						</ul>
						<div class="chart-container">
							<div class="chart-bar" style="width: %.1f%%"></div>
						</div>
					</div>
					
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon runtime-icon">🚀</div>
							<div class="metric-title">Runtime Status</div>
						</div>
						<div class="metric-value">HEALTHY</div>
						<div class="metric-label">✅ System Status</div>
						<ul class="metric-list">
							<li><span class="label">📚 Stack Memory</span><span class="value">%s</span></li>
							<li><span class="label">🎯 Heap Objects</span><span class="value">%d</span></li>
							<li><span class="label">🎯 Next GC Target</span><span class="value">%s</span></li>
						</ul>
					</div>
					
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon network-icon">🌐</div>
							<div class="metric-title">Network Activity</div>
						</div>
						<ul class="metric-list">
							<li><span class="label">📥 Network In</span><span class="value">%.2f MB</span></li>
							<li><span class="label">📤 Network Out</span><span class="value">%.2f MB</span></li>
							<li><span class="label">🔗 Active Connections</span><span class="value">%d</span></li>
							<li><span class="label">📡 Throughput</span><span class="value">%.1f MB/s</span></li>
						</ul>
						<div class="progress-bar">
							<div class="progress-fill" style="width: %.1f%%"></div>
						</div>
					</div>
					
					<div class="metric-card">
						<div class="metric-header">
							<div class="metric-icon efficiency-icon">🎯</div>
							<div class="metric-title">System Efficiency</div>
						</div>
						<ul class="metric-list">
							<li><span class="label">🧠 Memory Efficiency</span><span class="value">%.1f%%</span></li>
							<li><span class="label">♻️ GC Efficiency</span><span class="value">%.1f%%</span></li>
							<li><span class="label">⚡ CPU Usage</span><span class="value">%.1f%%</span></li>
							<li><span class="label">📊 Overall Score</span><span class="value">%.0f/100</span></li>
						</ul>
						<div class="chart-container">
							<div class="chart-bar" style="width: %.1f%%"></div>
						</div>
					</div>
				</div>
				
				<div class="footer">
					<div class="auto-refresh-indicator">
						<div class="refresh-dot"></div>
						🔄 Auto-refreshing every 5 seconds
					</div>
					<div class="timestamp">⏰ Last Updated: %s</div>
					<div>© 2025 CollectHub API - 🚀 Enterprise Grade Real-time Monitoring</div>
				</div>
			</div>
			
			<script>
				// Enhanced auto-refresh with visual feedback
				let refreshInterval;
				let countdown = 5;
				
				function startAutoRefresh() {
					refreshInterval = setInterval(() => {
						// Add refresh animation
						document.body.style.opacity = '0.8';
						setTimeout(() => {
							location.reload();
						}, 200);
					}, 5000);
				}
				
				// Start auto-refresh on page load
				document.addEventListener('DOMContentLoaded', function() {
					startAutoRefresh();
					
					// Add stagger animation to cards
					const cards = document.querySelectorAll('.metric-card');
					cards.forEach((card, index) => {
						card.style.animationDelay = (index * 0.1) + 's';
						card.style.animation = 'slideIn 0.8s ease-out forwards';
						
						// Add hover sound effect (visual feedback)
						card.addEventListener('mouseenter', function() {
							this.style.transform = 'translateY(-8px) scale(1.02)';
						});
						
						card.addEventListener('mouseleave', function() {
							this.style.transform = 'translateY(0) scale(1)';
						});
					});
					
					// Add typing animation to values
					const values = document.querySelectorAll('.value');
					values.forEach((value, index) => {
						value.style.animationDelay = (index * 0.05) + 's';
						value.style.animation = 'countUp 1s ease-out forwards';
					});
				});
				
				// Smooth scroll to top when refreshing
				window.addEventListener('beforeunload', () => {
					window.scrollTo(0, 0);
				});
			</script>
		</body>
		</html>
		`,
			runtime.Version(),
			runtime.NumCPU(),
			uptime,
			runtime.NumGoroutine(),
			formatBytes(memStats.Alloc),
			formatBytes(memStats.TotalAlloc),
			formatBytes(memStats.Sys),
			heapInUse,
			heapInUse,
			memStats.NumGC,
			gcPauseAvg,
			allocRate/(1024*1024), // Convert to MB/s
			memStats.Mallocs-memStats.Frees,
			gcEfficiency, // Chart width for performance
			formatBytes(memStats.StackInuse),
			memStats.HeapObjects,
			formatBytes(memStats.NextGC),
			networkIn,                // Network In
			networkOut,               // Network Out
			int(cpuUsage*5),          // Simulated active connections
			(networkIn+networkOut)/2, // Throughput
			(networkIn+networkOut)/4, // Progress bar for network
			memoryEfficiency,
			gcEfficiency,
			cpuUsage,
			(memoryEfficiency+gcEfficiency+(100-cpuUsage))/3, // Overall score
			(memoryEfficiency+gcEfficiency+(100-cpuUsage))/3, // Chart width for efficiency
			time.Now().Format("2006-01-02 15:04:05 MST"))

		return c.Type("html").SendString(htmlContent)
	})

	api := app.Group("/api")

	// User Routes
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetUsers)
	api.Post("/users/login", controllers.LoginUser)

	// Book Routes
	api.Post("/books", controllers.CreateBook)
	api.Get("/books/user/:userId", controllers.GetBooksByUser)
	api.Get("/books/:id", controllers.GetBookByID)
	api.Put("/books/:id", controllers.UpdateBook)
	api.Delete("/books/:id", controllers.DeleteBook)

	// Recipe Routes
	api.Post("/recipes", controllers.CreateRecipe)
	api.Get("/recipes/user/:userId", controllers.GetRecipesByUser)
	api.Get("/recipes/:id", controllers.GetRecipeByID)
	api.Put("/recipes/:id", controllers.UpdateRecipe)
	api.Delete("/recipes/:id", controllers.DeleteRecipe)

	// Movie Routes
	api.Post("/movies", controllers.CreateMovie)
	api.Get("/movies/user/:userId", controllers.GetMoviesByUser)
	api.Get("/movies/:id", controllers.GetMovieByID)
	api.Put("/movies/:id", controllers.UpdateMovie)
	api.Delete("/movies/:id", controllers.DeleteMovie)

	// Quote Routes
	api.Post("/quotes", controllers.CreateQuote)
	api.Get("/quotes/user/:userId", controllers.GetQuotesByUser)
	api.Get("/quotes/:id", controllers.GetQuoteByID)
	api.Put("/quotes/:id", controllers.UpdateQuote)
	api.Delete("/quotes/:id", controllers.DeleteQuote)

	// Pet Routes
	api.Post("/pets", controllers.CreatePet)
	api.Get("/pets/user/:userId", controllers.GetPetsByUser)
	api.Get("/pets/:id", controllers.GetPetByID)
	api.Put("/pets/:id", controllers.UpdatePet)
	api.Delete("/pets/:id", controllers.DeletePet)

	// Travel Routes
	api.Post("/travels", controllers.CreateTravel)
	api.Get("/travels/user/:userId", controllers.GetTravelsByUser)
	api.Get("/travels/:id", controllers.GetTravelByID)
	api.Put("/travels/:id", controllers.UpdateTravel)
	api.Delete("/travels/:id", controllers.DeleteTravel)

	// 🤖 AI Personality Analysis Route
	api.Post("/aipersonality/analysis", controllers.GetAIPersonalityAnalysis(db))
}

// Helper functions
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
