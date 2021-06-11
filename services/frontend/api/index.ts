import express from 'express'
import bodyParser from 'body-parser'

import capes from './routes/router'
const app = express()

app.use(bodyParser.json())
app.use(capes)

module.exports = app
