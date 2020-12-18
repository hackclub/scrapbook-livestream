const http = require("http")
const express = require("express")
const ws = require("ws")

const app = express()
const server = http.createServer(app)

const wss = new ws.Server({ server })

app.use(express.json())

app.get("/", (req, res) => {
  res.send("howdy")
})

app.post("/webhook", (req, res) => {
  res.end();

  const {name, username} = req.body

  wss.clients.forEach(client => {
    client.send(`@${name} posted to their scrapbook! scrapbook.hackclub.com/${username}`)
  })
})

server.listen(process.env.PORT || 3000, () => {
  console.log("app started")
})
