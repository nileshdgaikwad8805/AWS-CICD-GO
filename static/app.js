(function () {
  const $ = (id) => document.getElementById(id);
  let baseUptime = 0;
  let baseAt = Date.now();

  function fmtUptime(total) {
    const d = Math.floor(total / 86400), h = Math.floor((total % 86400) / 3600),
          m = Math.floor((total % 3600) / 60), s = total % 60;
    return (d ? d + "d " : "") + (h || d ? h + "h " : "") + m + "m " + s + "s";
  }

  function setStatus(up) {
    $("statusDot").className = "dot " + (up ? "up" : "down");
    $("statusText").textContent = up ? "All systems operational" : "Service unreachable";
    $("statusBig").textContent = up ? "UP" : "DOWN";
    $("statusBig").className = "big " + (up ? "ok" : "bad");
  }

  async function load() {
    const t0 = performance.now();
    try {
      const res = await fetch("/api/info", { cache: "no-store" });
      if (!res.ok) throw new Error("HTTP " + res.status);
      const d = await res.json();
      $("latency").textContent = "latency: " + Math.round(performance.now() - t0) + " ms";
      $("version").textContent = "v" + d.version;
      $("goVersion").textContent = d.goVersion + " · " + d.os;
      $("serverTime").textContent = d.serverTime.replace("T", " ").replace("Z", "");
      $("hostname").textContent = "host: " + d.hostname;
      $("startedAt").textContent = "since " + d.startedAt.replace("T", " ").replace("Z", " UTC");
      $("envPill").textContent = d.environment;
      baseUptime = d.uptimeSeconds;
      baseAt = Date.now();
      setStatus(true);
    } catch (e) {
      setStatus(false);
    }
  }

  async function healthCheck() {
    const btn = $("checkBtn");
    btn.textContent = "Checking…";
    try {
      const res = await fetch("/health", { cache: "no-store" });
      const d = await res.json();
      setStatus(d.status === "UP");
      btn.textContent = d.status === "UP" ? "✔ Healthy" : "✖ Unhealthy";
    } catch (e) {
      setStatus(false);
      btn.textContent = "✖ Unreachable";
    }
    setTimeout(() => (btn.textContent = "Run health check"), 2500);
  }

  setInterval(() => {
    const secs = baseUptime + Math.floor((Date.now() - baseAt) / 1000);
    $("uptime").textContent = fmtUptime(secs);
  }, 1000);

  $("year").textContent = new Date().getFullYear();
  $("checkBtn").addEventListener("click", healthCheck);
  load();
  setInterval(load, 5000);
})();
