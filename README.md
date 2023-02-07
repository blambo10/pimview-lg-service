# Pimview Backend

This project is the back end to "pimview".

The core is a broker of whichever flavour is desired,
then the front end should have components built to suit and the sub layer should process the queue 
and translate to any iot device that isnt natively compatible with mqtt.

![image](https://user-images.githubusercontent.com/37164299/211139130-5237f1b5-c4cc-4841-b9d3-90b03fce408b.png)

### `run the mqtt broker`

Current implementation uses rabitmq with the mqtt and websocket plugins enabled

### `go run main.go run sub`

Ensure the broker variables are configured appropriately
